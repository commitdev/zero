import React from 'react';
import './App.css';
import Divider from '@material-ui/core/Divider';

import Project from './Project.js';
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
      projectName: '',
      projectDescription: '',
      provider: 'aws',
      frontend: 'react',
      region: 'us-east-1',
      profile: 'default',
      services: []
    };
  }

  setProjectName = (v) => {
    this.setState({ projectName: v });
  }

  setProjectDescription = (v) => {
    this.setState({ projectDescription: v });
  }

  setProvider = (v) => {
    this.setState({ provider: v });
  }

  setFrontend = (v) => {
    this.setState({ frontend: v });
  }

  setRegion = (v) => {
    this.setState({ region: v });
  }

  setProfile = (v) => {
    this.setState({ profile: v });
  }

  setServices = (s) => {
    this.setState({services: s});
  }

  generate = () => {
    let self = this;
    axios.post('http://localhost:8080/v1/generate', {
      "projectName": self.state.projectName,
      "frontendFramework": self.state.frontend,
      "services": self.state.services,
      "infrastructure": {
        "aws": {
          "region": self.state.region,
          "profile": self.state.profile
        }
      }
    })
      .then(function (response) {
        // handle success
        self.setState({ success: true });
      })
      .catch(function (error) {
        // handle error
        console.log(error);
        self.setState({ success: false });
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
        <Project project={this.project} setProjectName={this.setProjectName} setProjectDescription={this.setProjectDescription} />
        <Providers
          provider={this.state.provider}
          setProvider={this.setProvider}
          setRegion={this.setRegion}
          setProfile={this.setProfile} />
        <Frontend frontend={this.state.frontend} setFrontend={this.setFrontend} />
        <Services services={this.state.services} setServices={this.setServices} />
        <Divider variant="middle" />
        <GenerateButton generate={this.generate} />
        {this.state.success !== null && <Complete success={this.state.success} />}
      </div>
    );
  }
}
