import React from 'react';
import Grid from '@material-ui/core/Grid';

import AWSRegion from './AWSRegion.js';
import AWSProfile from './AWSProfile.js';

export default class AWSProvider extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      selectedProfile: 0,
      selectedRegion: 0,
      regionMenuOpen: false,
      profileMenuOpen: false
    };
  }

  profiles = [
    'default',
    'testing',
    'fake-data'
  ];

  render() {
    return (
      <div><h3>AWS</h3>
        <Grid container spacing={4} direction="row" alignItems="center" justify="center">
          <Grid item>
            <AWSRegion setRegion={this.props.setRegion} />
          </Grid>
          <Grid item>
            <AWSProfile setProfile={this.props.setProfile} />
          </Grid>
        </Grid>
      </div>
    )
  }
}