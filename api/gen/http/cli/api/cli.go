// Code generated by goa v3.2.6, DO NOT EDIT.
//
// api HTTP client CLI support package
//
// Command:
// $ goa gen github.com/lawrencejones/pgsink/api/design -o api

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	healthc "github.com/lawrencejones/pgsink/api/gen/http/health/client"
	importsc "github.com/lawrencejones/pgsink/api/gen/http/imports/client"
	subscriptionsc "github.com/lawrencejones/pgsink/api/gen/http/subscriptions/client"
	tablesc "github.com/lawrencejones/pgsink/api/gen/http/tables/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `health check
tables list
subscriptions (get|add-table|stop-table)
imports list
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` health check` + "\n" +
		os.Args[0] + ` tables list --schema "public,payments"` + "\n" +
		os.Args[0] + ` subscriptions get` + "\n" +
		os.Args[0] + ` imports list` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		healthFlags = flag.NewFlagSet("health", flag.ContinueOnError)

		healthCheckFlags = flag.NewFlagSet("check", flag.ExitOnError)

		tablesFlags = flag.NewFlagSet("tables", flag.ContinueOnError)

		tablesListFlags      = flag.NewFlagSet("list", flag.ExitOnError)
		tablesListSchemaFlag = tablesListFlags.String("schema", "public", "")

		subscriptionsFlags = flag.NewFlagSet("subscriptions", flag.ContinueOnError)

		subscriptionsGetFlags = flag.NewFlagSet("get", flag.ExitOnError)

		subscriptionsAddTableFlags    = flag.NewFlagSet("add-table", flag.ExitOnError)
		subscriptionsAddTableBodyFlag = subscriptionsAddTableFlags.String("body", "REQUIRED", "")

		subscriptionsStopTableFlags    = flag.NewFlagSet("stop-table", flag.ExitOnError)
		subscriptionsStopTableBodyFlag = subscriptionsStopTableFlags.String("body", "REQUIRED", "")

		importsFlags = flag.NewFlagSet("imports", flag.ContinueOnError)

		importsListFlags = flag.NewFlagSet("list", flag.ExitOnError)
	)
	healthFlags.Usage = healthUsage
	healthCheckFlags.Usage = healthCheckUsage

	tablesFlags.Usage = tablesUsage
	tablesListFlags.Usage = tablesListUsage

	subscriptionsFlags.Usage = subscriptionsUsage
	subscriptionsGetFlags.Usage = subscriptionsGetUsage
	subscriptionsAddTableFlags.Usage = subscriptionsAddTableUsage
	subscriptionsStopTableFlags.Usage = subscriptionsStopTableUsage

	importsFlags.Usage = importsUsage
	importsListFlags.Usage = importsListUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "health":
			svcf = healthFlags
		case "tables":
			svcf = tablesFlags
		case "subscriptions":
			svcf = subscriptionsFlags
		case "imports":
			svcf = importsFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "health":
			switch epn {
			case "check":
				epf = healthCheckFlags

			}

		case "tables":
			switch epn {
			case "list":
				epf = tablesListFlags

			}

		case "subscriptions":
			switch epn {
			case "get":
				epf = subscriptionsGetFlags

			case "add-table":
				epf = subscriptionsAddTableFlags

			case "stop-table":
				epf = subscriptionsStopTableFlags

			}

		case "imports":
			switch epn {
			case "list":
				epf = importsListFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "health":
			c := healthc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "check":
				endpoint = c.Check()
				data = nil
			}
		case "tables":
			c := tablesc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data, err = tablesc.BuildListPayload(*tablesListSchemaFlag)
			}
		case "subscriptions":
			c := subscriptionsc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "get":
				endpoint = c.Get()
				data = nil
			case "add-table":
				endpoint = c.AddTable()
				data, err = subscriptionsc.BuildAddTablePayload(*subscriptionsAddTableBodyFlag)
			case "stop-table":
				endpoint = c.StopTable()
				data, err = subscriptionsc.BuildStopTablePayload(*subscriptionsStopTableBodyFlag)
			}
		case "imports":
			c := importsc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data = nil
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// healthUsage displays the usage of the health command and its subcommands.
func healthUsage() {
	fmt.Fprintf(os.Stderr, `Provide service health information
Usage:
    %s [globalflags] health COMMAND [flags]

COMMAND:
    check: Health check for probes

Additional help:
    %s health COMMAND --help
`, os.Args[0], os.Args[0])
}
func healthCheckUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] health check

Health check for probes

Example:
    `+os.Args[0]+` health check
`, os.Args[0])
}

// tablesUsage displays the usage of the tables command and its subcommands.
func tablesUsage() {
	fmt.Fprintf(os.Stderr, `Expose Postgres tables, and their import/sync status
Usage:
    %s [globalflags] tables COMMAND [flags]

COMMAND:
    list: List all tables

Additional help:
    %s tables COMMAND --help
`, os.Args[0], os.Args[0])
}
func tablesListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] tables list -schema STRING

List all tables
    -schema STRING: 

Example:
    `+os.Args[0]+` tables list --schema "public,payments"
`, os.Args[0])
}

// subscriptionsUsage displays the usage of the subscriptions command and its
// subcommands.
func subscriptionsUsage() {
	fmt.Fprintf(os.Stderr, `Manage a pgsink subscription
Usage:
    %s [globalflags] subscriptions COMMAND [flags]

COMMAND:
    get: Get current subscription data
    add-table: Add table to publication, relying on an import manager to schedule the job
    stop-table: Stop a table by removing it from the publication, and expiring any import jobs

Additional help:
    %s subscriptions COMMAND --help
`, os.Args[0], os.Args[0])
}
func subscriptionsGetUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] subscriptions get

Get current subscription data

Example:
    `+os.Args[0]+` subscriptions get
`, os.Args[0])
}

func subscriptionsAddTableUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] subscriptions add-table -body JSON

Add table to publication, relying on an import manager to schedule the job
    -body JSON: 

Example:
    `+os.Args[0]+` subscriptions add-table --body '{
      "name": "payments",
      "schema": "public"
   }'
`, os.Args[0])
}

func subscriptionsStopTableUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] subscriptions stop-table -body JSON

Stop a table by removing it from the publication, and expiring any import jobs
    -body JSON: 

Example:
    `+os.Args[0]+` subscriptions stop-table --body '{
      "name": "payments",
      "schema": "public"
   }'
`, os.Args[0])
}

// importsUsage displays the usage of the imports command and its subcommands.
func importsUsage() {
	fmt.Fprintf(os.Stderr, `Manage table imports, scoped to the server subscription ID
Usage:
    %s [globalflags] imports COMMAND [flags]

COMMAND:
    list: List all imports

Additional help:
    %s imports COMMAND --help
`, os.Args[0], os.Args[0])
}
func importsListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] imports list

List all imports

Example:
    `+os.Args[0]+` imports list
`, os.Args[0])
}