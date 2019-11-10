import React from 'react'
import PropTypes from 'prop-types'
import { compose, withApollo } from 'react-apollo'

import AppBar from '@material-ui/core/AppBar'
import { withStyles } from '@material-ui/core/styles'
import IconButton from '@material-ui/core/IconButton'
import ArrowHome from '@material-ui/icons/Close'
import ArrowBack from '@material-ui/icons/ArrowBackIos'

import { router } from '../../utils'
import { H5 } from '../Text'
import { queries } from '../../constants'

const styles = theme => ({
  navbar: {
    borderBottom: `1px solid ${theme.palette.divider}`,
    alignItems: 'center',
    margin: 0,
  },
  title: {
    color: theme.palette.grey[700],
    fontSize: '110%',
  },
})

class NavBar extends React.Component {
  state = {
    dropdownAnchor: null,
    drawer: false,
  }

  render() {
    const { title, children, classes, backHref = '/', ...props } = this.props
    return (
      <AppBar position="static" {...props} className={classes.navbar}>
        <IconButton onClick={router.to(backHref)}>
          <ArrowBack />
        </IconButton>
        <div>
          <H5 className={classes.title}>{title}</H5>
        </div>
        <IconButton onClick={router.to('/app')}>
          <ArrowHome />
        </IconButton>
      </AppBar>
    )
  }

  logout = () => {
    this.props.client.mutate({ mutation: queries.LOGOUT })
  }
}

NavBar.propTypes = {
  title: PropTypes.string,
  backHref: PropTypes.string,
  client: PropTypes.object.isRequired,
}

export default compose(
  withStyles(styles),
  withApollo
)(NavBar)
