import React from 'react';
import TextField from '@material-ui/core/TextField';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import MenuItem from '@material-ui/core/MenuItem';
import Menu from '@material-ui/core/Menu';

export default class Service extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: "",
      description: "",
      language: "",
      anchorEl: null,
    }
  }

  updateName = (event) => {
    this.setState({
      name: event.target.value
    }, () => {
      this.props.update(this.props.serviceID, this.state);
    })
  }

  updateDesc = (event) => {
    this.setState({
      description: event.target.value
    }, () => {
      this.props.update(this.props.serviceID, this.state);
    })
  }

  languages = [
    'go',
    'scala',
    'nodejs'
  ];

  setAnchorEl = element => {
    this.setState({ anchorEl: element })
  }

  handleClickListItem = event => {
    this.setAnchorEl(event.currentTarget);
  };

  handleMenuItemClick = (event, index) => {
    this.setAnchorEl(null);
    this.setState({ language: this.languages[index] }, () => {
      this.props.update(this.props.serviceID, this.state);
    });
  };

  handleClose = () => {
    this.setAnchorEl(null);
  };

  render() {
    return (
      <div className="service">
        <TextField
          id="standard-basic"
          label="Name"
          margin="normal"
          onChange={this.updateName}
        />
        <TextField
          id="standard-multiline-flexible"
          label="Description"
          multiline
          rowsMax="4"
          margin="normal"
          onChange={this.updateDesc}
        />
        <List component="nav" aria-label="language">
          <ListItem
            button
            aria-haspopup="true"
            aria-controls="lock-menu"
            aria-label="language"
            onClick={this.handleClickListItem}
          >
            <ListItemText primary="Language" secondary={this.state.language} />
          </ListItem>
        </List>
        <Menu
          id="lock-menu"
          anchorEl={this.state.anchorEl}
          keepMounted
          open={Boolean(this.state.anchorEl)}
          onClose={this.handleClose}
        >
          {this.languages.map((language, index) => (
            <MenuItem
              key={language}
              selected={index === this.state.selectedlanguage}
              onClick={event => this.handleMenuItemClick(event, index)}
            >
              {language}
            </MenuItem>
          ))}
        </Menu>
      </div>
    );
  }
}
