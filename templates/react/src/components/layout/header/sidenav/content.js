import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import Icon from '@material-ui/core/Icon';
import config from 'config';

const useStyles = makeStyles({
  list: {
    width: 250,
  },
});

export default function Content() {
  const classes = useStyles();
  return (
    <div
      className={classes.list}
      role="presentation"
    >
      <List>
        {
          config.sidenav && config.sidenav.items && config.sidenav.items.map((item, index) => (
            <ListItem button component="a" href={item.path} key={index}>
              <ListItemIcon>
                <Icon>{item.icon}</Icon>
              </ListItemIcon>
              <ListItemText primary={item.label} />
            </ListItem>
          ))
        }
      </List>
    </div>
  );
}
