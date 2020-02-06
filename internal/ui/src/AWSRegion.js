import React from 'react';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import MenuItem from '@material-ui/core/MenuItem';
import Menu from '@material-ui/core/Menu';

export default class AWSRegion extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      selectedProfile: 0,
      selectedRegion: 0,
      regionMenuOpen: true,
      anchorEl: null,
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
    'Asia Pacific (Singapore) ap-southeast-1',
    'Asia Pacific (Sydney) ap-southeast-2',
    'Asia Pacific (Tokyo) ap-northeast-1',
    'Canada (Central) ca-central-1',
    'China (Beijing) cn-north-1',
    'China (Ningxia) cn-northwest-1',
    'EU (Frankfurt) eu-central-1 ',
    'EU (Ireland) eu-west-1',
    'EU (London) eu-west-2',
    'EU (Paris) eu-west-3',
    'EU (Stockholm) eu-north-1',
    'Middle East (Bahrain) me-south-1',
    'South America (Sao Paulo) sa-east-1'
  ];

  toggleRegion = (index) => {
    this.setState({ selectedRegion: index })
  }

  setAnchorEl = element => {
    this.setState({ anchorEl: element })
  }

  handleClickListItem = event => {
    this.setAnchorEl(event.currentTarget);
  };

  handleMenuItemClick = (event, index) => {
    this.toggleRegion(index);
    this.setAnchorEl(null);

    let words = this.regions[index].split(" ");
    let code = words[words.length-1];
    this.props.setRegion(code);
  };

  handleClose = () => {
    this.setAnchorEl(null);
  };

  render() {
    return (
      <div>
        <List component="nav" aria-label="Region">
          <ListItem
            button
            aria-haspopup="true"
            aria-controls="lock-menu"
            aria-label="Region"
            onClick={this.handleClickListItem}
          >
            <ListItemText primary="Region" secondary={this.regions[this.state.selectedRegion]} />
          </ListItem>
        </List>
        <Menu
          id="lock-menu"
          anchorEl={this.state.anchorEl}
          keepMounted
          open={Boolean(this.state.anchorEl)}
          onClose={this.handleClose}
        >
          {this.regions.map((region, index) => (
            <MenuItem
              key={region}
              selected={index === this.state.selectedRegion}
              onClick={event => this.handleMenuItemClick(event, index)}
            >
              {region}
            </MenuItem>
          ))}
        </Menu>
      </div>
    )
  }
}