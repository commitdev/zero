import React from 'react';
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import Layout from 'components/layout';
import config from 'config';

const renderView = (view) => {
  return (
    <Route path={`${view.path}`} component={require(`views/${view.component}`).default} />
  )
}

export default function App() {


  return (
    <Layout>
      <Router>
        <Switch>
          {
            config.views && config.views.map(renderView)
          }
        </Switch>
      </Router>
    </Layout>
  );
}
