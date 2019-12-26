import React from 'react';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import MenuItem from '@material-ui/core/MenuItem';
import Menu from '@material-ui/core/Menu';

export default class AWSProfile extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      selectedProfile: 0,
      anchorEl: null,
    };
  }

  profiles = [
    'default'
  ];

  toggle = (index) => {
    this.setState({ selectedRegion: index })
  }

  setAnchorEl = element => {
    this.setState({ anchorEl: element })
  }

  handleClickListItem = event => {
    this.setAnchorEl(event.currentTarget);
  };

  handleMenuItemClick = (event, index) => {
    this.toggle(index);
    this.setAnchorEl(null);
    this.props.setProfile(this.profiles[index]);
  };

  handleClose = () => {
    this.setAnchorEl(null);
  };

  render() {
    return (
      <div>
        <List component="nav" aria-label="Profile">
          <ListItem
            button
            aria-haspopup="true"
            aria-controls="lock-menu"
            aria-label="Profile"
            onClick={this.handleClickListItem}
          >
            <ListItemText primary="Profile" secondary={this.profiles[this.state.selectedProfile]} />
          </ListItem>
        </List>
        <Menu
          id="lock-menu"
          anchorEl={this.state.anchorEl}
          keepMounted
          open={Boolean(this.state.anchorEl)}
          onClose={this.handleClose}
        >
          {this.profiles.map((profile, index) => (
            <MenuItem
              key={profile}
              selected={index === this.state.selectedProfile}
              onClick={event => this.handleMenuItemClick(event, index)}
            >
              {profile}
            </MenuItem>
          ))}
        </Menu>
      </div>
    )
  }
}