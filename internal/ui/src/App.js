import React from 'react';
import './App.css';
import Divider from '@material-ui/core/Divider';


import Providers from './Providers.js';
import Frontend from './Frontend.js';
import Services from './Services.js';
import GenerateButton from './GenerateButton.js';
import Complete from './Complete.js';

import axios from 'axios';

export default class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      success: null,
      provider: 'aws',
      frontend: 'react',
      services: [
        {
          name: "test service 1",
          description: "Things.",
          language: "go"
        }
      ]
    };
  }

  setProvider = (p) => {
    this.setState({ provider: p });
  }

  setFrontend = (f) => {
    this.setState({ frontend: f });
  }

  addServices = (s) => {
    this.setState({ services: [1] });
  }

  generate = () => {
    let self = this;
    axios.post('http://localhost:8080/v1/generate', {
      frontendFramework: this.state.frontend,
    })
      .then(function (response) {
        // handle success
        console.log(response);
        this.setState({success: true});
      })
      .catch(function (error) {
        // handle error
        console.log(error);
        self.setState({success: false});
      })
      .finally(function () {
        // always executed
      });
  }

  render() {
    console.log(this.state);
    return (
      <div className="App">
        <h1>Commit0 - Create a new project</h1>
        <Providers provider={this.state.provider} setProvider={this.setProvider} />
        <Frontend frontend={this.state.frontend} setFrontend={this.setFrontend} />
        <Services services={this.state.services} addService={this.addService} />
        <Divider variant="middle" />
        <GenerateButton generate={this.generate} />
        {this.state.success !== null && <Complete succes={this.state.success} />}
      </div>
    );
  }
}
