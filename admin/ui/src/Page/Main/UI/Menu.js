import React from 'react'
import { makeStyles } from '@material-ui/styles';
import { Link } from "react-router-dom";
import IconButton from '@material-ui/core/IconButton'
import Drawer from '@material-ui/core/Drawer';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import Divider from '@material-ui/core/Divider';
import Typography from '@material-ui/core/Typography'

const useStyles = makeStyles({
  drawer: {
    width: "20em",
    flexShrink: 0,
  },
  drawerPaper: {
    width: "20em",
  },
  drawerHeader: {
    display: 'flex',
    alignItems: 'center',
    padding: '0 8px',
    justifyContent: 'space-between',
  },
});

export default props => {
  const classes = useStyles();

  return (
    <Drawer
      className={classes.drawer}
      variant="persistent"
      anchor="left"
      open={props.isOpenMenu}
      classes={{
        paper: classes.drawerPaper,
      }}
    >
      <div className={classes.drawerHeader}>
        <Typography variant="h3" color="inherit">
          FSME-UI
        </Typography>

        <IconButton onClick={props.onChangeOpenDrawer}>
          <ChevronLeftIcon />
        </IconButton>
      </div>

      <Divider />

      <List>
        {
          Object.keys(props.fsm)
            .filter(key => key !== "new")
            .map((key, index) => (
              <Link key={`/fsm/${props.fsm[key]._id}`} to={`/fsm/${props.fsm[key]._id}/view`}>
                <ListItem button key={key}>
                  <ListItemText primary={props.fsm[key].Name} />
                </ListItem>
              </Link>
            ))
        }
      </List>
    </Drawer>
  )
}
