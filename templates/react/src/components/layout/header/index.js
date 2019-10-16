import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import config from 'config';
import Sidenav from 'components/layout/header/sidenav';
import Account from 'components/layout/header/account';

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  title: {
    flexGrow: 1,
  },
}));

export default function MenuAppBar() {
  const classes = useStyles();
  return (
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          { config && config.sidenav && config.sidenav.enabled && <Sidenav />}
          <Typography variant="h6" className={classes.title}>
            { config && config.app && config.app.name }
          </Typography>
          { config && config.account && config.account.enabled && <Account />}
        </Toolbar>
      </AppBar>
    </div>
  );
}
