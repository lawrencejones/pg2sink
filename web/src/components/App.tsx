import React from 'react';
import { BrowserRouter as Router, Switch, Route, Redirect, Link } from 'react-router-dom';

import './App.css';
import Tables from './Tables';

function App(): JSX.Element {
  return (
    <Router>
      <div>
        <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
          <div className="container-fluid">
            <Link to="/" className="navbar-brand">
              pgsink admin
            </Link>
            <div className="navbar-collapse">
              <ul className="navbar-nav me-auto mb-2 mb-lg-0">
                <li className="nav-item">
                  <Link to="/tables" className="nav-link">
                    Tables
                  </Link>
                </li>
                <li className="nav-item">
                  <Link to="/imports" className="nav-link">
                    Imports
                  </Link>
                </li>
              </ul>
            </div>
          </div>
        </nav>

        <Switch>
          <Route path="/tables">
            <Tables schema="public"></Tables>
          </Route>
          // Catch-all route, bring us to tables if we don't recognise the path
          <Route>
            <Redirect to="/tables" />
          </Route>
        </Switch>
      </div>
    </Router>
  );
}

export default App;
