import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';
import Paper from '@material-ui/core/Paper';
import GraphEditor from '../Containers/GraphEditor'

const styles = {
  root: {
    display: 'grid',
    gridTemplateRows: 'auto 1fr',
    flex: 1,
  },
  rootPaper: {
    display: 'grid',
    gridTemplateRows: '1fr',
    flex: 1,
  },
  grow: {
    flexGrow: 1,
  },
  menuButton: {
    marginLeft: -12,
    marginRight: 20,
  },
};

function ButtonAppBar(props) {
  const { classes } = props;

  return (
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          <IconButton className={classes.menuButton} color="inherit" aria-label="Menu">
            <MenuIcon />
          </IconButton>

          <Typography variant="h6" color="inherit" className={classes.grow}>
            FSME-UI
          </Typography>
        </Toolbar>
      </AppBar>

      <Paper className={classes.rootPaper} elevation={1}>
        <GraphEditor />
      </Paper>
    </div>
  );
}

ButtonAppBar.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(ButtonAppBar);
