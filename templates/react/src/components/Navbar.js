import React from 'react'
import PropTypes from 'prop-types'
import { withApollo } from 'react-apollo'

import AppBar from '@material-ui/core/AppBar'
import Hidden from '@material-ui/core/Hidden'
import Button from '@material-ui/core/Button'
import Menu from '@material-ui/core/Menu'
import MenuItem from '@material-ui/core/MenuItem'
import Divider from '@material-ui/core/Divider'
import Drawer from '@material-ui/core/Drawer'
import IconButton from '@material-ui/core/IconButton'
import MenuIcon from '@material-ui/icons/Menu'
import HomeIcon from '@material-ui/icons/Home'
import UserIcon from '@material-ui/icons/PersonOutline'

import { router } from '../utils'
import { Text } from '../components/Text'
import { queries } from '../constants'

const DropdownContainerStyle = {
  display: 'flex',
  flexDirection: 'row',
  justifyContent: 'center',
  alignItems: 'center',
}

class NavBar extends React.Component {
  state = {
    dropdownAnchor: null,
    drawer: false,
  }

  render() {
    const { user, children, ...props } = this.props
    return (
      <AppBar position="static" {...props}>
        <IconButton onClick={router.to('/')}>
          <HomeIcon />
        </IconButton>

        <div style={DropdownContainerStyle}>
          {children}
          <Hidden xsDown>
            <Button onClick={this.toggleDropdown} style={{ margin: '0 6px' }}>
              {this.renderUserName()}
            </Button>
            <Menu
              anchorEl={this.state.dropdownAnchor}
              open={Boolean(this.state.dropdownAnchor)}
              onClose={this.hideDropdown}
            >
              {this.renderMenuItems()}
            </Menu>
          </Hidden>
          <Hidden smUp>
            <IconButton onClick={this.openDrawer}>
              <MenuIcon />
            </IconButton>
            {this.renderDrawer()}
          </Hidden>
        </div>
      </AppBar>
    )
  }

  renderUserName = () => {
    const { user = '' } = this.props
    return [<UserIcon key={1} />, user && user.replace(/@.*/gi, '')]
  }

  renderMenuItems = () => {
    return [
      <MenuItem key={1} onClick={router.to('/admin/forms')}>
        <Text>forms</Text>
      </MenuItem>,
      <MenuItem key={2} onClick={router.to('/app')}>
        <Text>dashboard</Text>
      </MenuItem>,
      <MenuItem key={3} onClick={this.logout}>
        <Text>logout</Text>
      </MenuItem>,
    ]
  }

  toggleDropdown = e => {
    this.setState({
      dropdownAnchor: this.state.dropdownAnchor ? null : e.currentTarget,
    })
  }

  hideDropdown = () => {
    this.setState({
      dropdownAnchor: null,
    })
  }

  logout = () => {
    this.props.client.mutate({ mutation: queries.LOGOUT })
  }

  renderDrawer = () => {
    const { drawer } = this.state
    return (
      <Drawer anchor="right" open={drawer} onClose={this.closeDrawer}>
        <div
          tabIndex={0}
          role="button"
          onClick={this.closeDrawer}
          onKeyDown={this.closeDrawer}
        >
          <div style={{ width: 250 }}>
            <MenuItem>{this.renderUserName()}</MenuItem>
            <Divider />
            {this.renderMenuItems()}
          </div>
        </div>
      </Drawer>
    )
  }

  openDrawer = () => {
    this.setState({ drawer: true })
  }

  closeDrawer = () => {
    this.setState({ drawer: false })
  }
}

NavBar.propTypes = {
  user: PropTypes.node,
  client: PropTypes.object.isRequired,
}

export default withApollo(NavBar)
