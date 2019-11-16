import React from 'react';
import Grid from '@material-ui/core/Grid';
import Button from '@material-ui/core/Button';

export default function Providers() {
  const [services, addService] = React.useState([]);

  // const addService() {

  // }
  return (
    <div className="providers">
      <h2>Services</h2>
      <Grid container spacing={2} direction="column" alignItems="center">
        <Button variant="contained" color="primary">+ Add Service</Button>
        {/* {this.props.children} */}
      </Grid>
    </div>
  );
}
