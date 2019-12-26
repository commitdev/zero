import React from 'react';
import Grid from '@material-ui/core/Grid';
import Button from '@material-ui/core/Button';

import Service from './Service.js';

export default class Services extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      services: []
    }
  }

  add = () => {
    let services = this.state.services;
    services.push({ name: "", description: ""});
    this.setState({services: services}, () => {
      this.props.setServices(this.state.services);
    });
  }

  update = (index, data) => {
    let services = this.state.services;
    services[index] = data;
    this.setState({services: services}, () => {
      this.props.setServices(this.state.services);
    });
  }
  render() {
    return (
      <div className="services">
        <h2>Services</h2>
        <Grid container spacing={2} direction="column" alignItems="center">
          <Button variant="contained" color="default" onClick={this.add}>+ Add Service</Button>
          {this.props.services.map((s, i) => <Service key={i} serviceID={i} update={this.update} />)}
        </Grid>
      </div>
    );
  }
}
