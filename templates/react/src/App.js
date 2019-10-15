import React from 'react';
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import Layout from 'components/layout';

export default function App() {
  return (
    <Layout>
      <Router>
        <Switch>
          <Route path="/a">
            <span>a</span>
          </Route>
          <Route path="/b">
            <span>b</span>
          </Route>
          <Route path="/">
            <span>c</span>
          </Route>
        </Switch>
      </Router>
    </Layout>
  );
}
