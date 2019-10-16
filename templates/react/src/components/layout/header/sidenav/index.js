import React, { Fragment, useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';
import Drawer from '@material-ui/core/Drawer';
import Content from 'components/layout/header/sidenav/content';

const useStyles = makeStyles(theme => ({
  menuButton: {
    marginRight: theme.spacing(2),
  },
}));

export default function Sidenav() {
  const classes = useStyles();
  const [open, setOpen] = useState(false);

  const onClose = () => setOpen(false);
  const onOpen = () => setOpen(true);

  return (
    <Fragment>
      <IconButton
        edge="start"
        className={classes.menuButton}
        color="inherit"
        aria-label="open drawer"
        onClick={onOpen}
      >
        <MenuIcon />
      </IconButton>
      <Drawer open={open} onClose={onClose}>
        <Content />
      </Drawer>
    </Fragment>
  );
}
