import React from 'react'
import PropTypes from 'prop-types'
import Grid from '@material-ui/core/Grid'
import { withStyles } from '@material-ui/core/styles'

import { H5, H6, P, Small } from '../components/Text'
import theme from '../utils/theme'
import { Link } from '../utils/router'

// NOTE: need to make sure react-snap have access to crawl all the static pages
function Footer({ classes }) {
  return (
    <footer className={classes.pageFooter}>
      <Grid container>
        <Grid item xs={6}>
          <div className={classes.logo}>
            <img
              alt="logo"
              src="https://getuikit.com/images/uikit-logo.svg"
              className={classes.img}
            />
            <div className={classes.text}>
              <H5 color="inherit">phrase.home_title</H5>
            </div>
          </div>
          <Small color="inherit">phrase.copyright</Small>
        </Grid>
        <Grid item xs={3}>
          <H6 color="inherit">menu</H6>
          <P color="inherit">item</P>
          <P color="inherit">item</P>
        </Grid>
        <Grid item xs={3}>
          <H6 color="inherit">about</H6>
          <Link href="terms-of-service" className={classes.a}>
            <P color="inherit">tos</P>
          </Link>
        </Grid>
        <Link href="login" />
      </Grid>
    </footer>
  )
}

Footer.propTypes = {
  classes: PropTypes.object.isRequired,
}

export default withStyles({
  pageFooter: {
    flex: 0,
    background: 'black',
    padding: theme.spacing(3),
    color: 'white',
  },
  a: {
    color: 'white',
    textDecoration: 'none',
  },
  logo: {
    marginTop: theme.spacing(1),
    marginBottom: theme.spacing(1),
  },
  img: {
    verticalAlign: 'middle',
    height: theme.spacing(4),
  },
  text: {
    display: 'inline-block',
    marginLeft: theme.spacing(2),
    verticalAlign: 'middle',
  },
})(Footer)
