import React from 'react';
import Grid from '@material-ui/core/Grid';
import ToggleButton from '@material-ui/lab/ToggleButton';
import ToggleButtonGroup from '@material-ui/lab/ToggleButtonGroup';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import MenuItem from '@material-ui/core/MenuItem';
import Menu from '@material-ui/core/Menu';

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

  regions = [
    'US East (Ohio) us-east-2',
    'US East (N. Virginia) us-east-1',
    'US West (N. California) us-west-1',
    'US West (Oregon) us-west-2',
    'Asia Pacific (Hong Kong) ap-east-1',
    'Asia Pacific (Mumbai) ap-south-1',
    'Asia Pacific (Seoul) ap-northeast-2',
    'Asia Pacific (Singapore) ap-southeast-1 ',
    'Asia Pacific (Sydney) ap-southeast-2 ',
    'Asia Pacific (Tokyo) ap-northeast-1 ',
    'Canada (Central) ca-central-1 ',
    'China (Beijing) cn-north-1 ',
    'China (Ningxia) cn-northwest-1',
    'EU (Frankfurt) eu-central-1 ',
    'EU (Ireland) eu-west-1 ',
    'EU (London) eu-west-2 ',
    'EU (Paris) eu-west-3 ',
    'EU (Stockholm) eu-north-1 ',
    'Middle East (Bahrain) me-south-1 ',
    'South America (Sao Paulo) sa-east-1 '
  ];

  profiles = [
    'default',
    'testing',
    'fake-data'
  ];

  handleClickListItemRegion = (event) => {
    this.props.setProvider(this.state);
  }

  handleMenuItemClickRegion = (event, index) => {
    // this.setState({
    //   regionMenuOpen: true
    // })
  };

  handleMenuClick = () => {
    this.setState({ regionMenuOpen: true});
  }

  handleClickListItemProfile = (event) => {
  }

  handleMenuItemClickProfile = (event, index) => {
  };

  handleClose = () => {
    this.setState({ regionMenuOpen: false, profileMenuOpen: false })
  };

  render() {
    return (
      <div><h3>AWS</h3>
        <Grid container spacing={4} direction="row" alignItems="center" justify="center">
          <Grid item>
            <List component="nav" aria-label="Region">
              <ListItem
                button
                aria-haspopup="true"
                aria-controls="lock-menu"
                aria-label="Region"
                onClick={this.handleClickListItemRegion}
              >
                <ListItemText primary="Region" secondary={this.regions[this.state.selectedRegion]} />
              </ListItem>
            </List>
            <Menu
              id="lock-menu"
              anchorEl={this.state.regionMenuOpen}
              keepMounted
              open={this.state.regionMenuOpen}
              onClose={this.handleClose}
            >
              {this.regions.map((region, index) => (
                <MenuItem
                  key={region}
                  selected={index === this.state.selectedRegion}
                  onClick={event => this.handleMenuItemClickRegion(event, index, "region")}
                >
                  {region}
                </MenuItem>
              ))}
            </Menu>
          </Grid>
          <Grid item>
            <List component="nav" aria-label="Profiles">
              <ListItem
                button
                aria-haspopup="true"
                aria-controls="lock-menu"
                aria-label="Profile"
                onClick={this.handleClickListItemProfile}
              >
                <ListItemText primary="Profile" secondary={this.profiles[this.selectedProfile]} />
              </ListItem>
            </List>
            <Menu
              id="lock-menu"
              anchorEl={this.state.currentProfile}
              keepMounted
              open={this.profileMenuOpen}
              onClose={this.handleClose}
            >
              {this.profiles.map((profile, index) => (
                <MenuItem
                  key={profile}
                  selected={index === this.state.selectedProfile}
                  onClick={event => this.handleMenuItemClickProfile(event, index)}
                >
                  {profile}
                </MenuItem>
              ))}
            </Menu>
          </Grid>
        </Grid>
      </div>
    )
  }
}