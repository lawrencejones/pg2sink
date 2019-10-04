# pg2pubsub

This tool connects to a Postgres database via logical replication, creating the
plumbing required to subscribe to changes on tables within a specific schema.
These changes are then pushed into GCP Pub/Sub.

## Getting started

Boot a Postgres database, then create an example table.

```console
$ createdb pg2pubsub
$ psql pg2pubsub
psql (11.5)
Type "help" for help.

pg2pubsub=# create table example (id bigserial primary key, msg text);
CREATE TABLE

pg2pubsub=# insert into example (msg) values ('hello world');
INSERT 1
```

pg2pubsub will stream these changes from the database and send it to Google's
Pub/Sub service. Two types of message are sent, data and schema. Data messages
contain the inserted/updated/deleted row content in JSON form, with each column
value stored as a string.

Our example would produce the following data message, where `timestamp` is the
time at which the change was committed and `sequence` the operation index within
the transaction:

```json
{
  "timestamp": "2019-10-04T16:05:55.123456+01:00",
  "sequence": 1,
  "namespace": "public",
  "name": "example",
  "before": null,
  "after": {
    "id": "1",
    "msg": "hello world"
  }
}
```

For some, receiving JSON with string keys will be perfectly suitable for their
use cases. For others, it will be important that schemas are available to later
interpret type information such as numeric precision. We publish Avro schema
definitions via a separate Pub/Sub topic that can be used to parse these data
messages:

```json
{
  "timestamp": "2019-10-04T16:05:55.123456+01:00",
  "schema": {
    "namespace": "public.example",
    "type": "record",
    "name": "value",
    "fields": [
      {
        "name": "id",
        "type": ["long", "null"],
        "default": null
      },
      {
        "name": "msg",
        "type": ["string", "null"],
        "default": null
      }
    ]
  }
}
```

Schemas are published whenever we first discover a relation. Use the timestamp
field to order each successive schema event to ensure stale messages don't
override more recent data.
