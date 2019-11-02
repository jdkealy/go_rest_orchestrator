import React from "react";
import RoutesConfig from "./components/routes_config";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";

export default function Routes() {
  return (
    <Router>
      <div>
        <Switch>
          {RoutesConfig.map((r) => {
            return (
              <Route path={r.route}>{r.cmp}</Route>
            )
          })}
        </Switch>
      </div>
    </Router>
  );
}