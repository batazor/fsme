import React, { Component } from 'react'
import PropTypes from 'prop-types'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { withStyles } from '@material-ui/core/styles'
import AppBar from '@material-ui/core/AppBar'
import Toolbar from '@material-ui/core/Toolbar'
import Typography from '@material-ui/core/Typography'
import IconButton from '@material-ui/core/IconButton'
import MenuIcon from '@material-ui/icons/Menu'
import Paper from '@material-ui/core/Paper'
import GraphEditor from '../../Containers/GraphEditor'
import Terminal from '../../Containers/Terminal'
import {
  list, add, update, remove,
} from '../../actions/fsm'

const styles = {
  root: {
    display: 'grid',
    gridTemplateRows: 'auto 1fr',
    flex: 1,
    overflow: 'auto',
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
  main: {
    display: 'grid',
    flexDirection: 'row',
    gridTemplateColumns: 'auto 25em',
    overflow: 'auto',
  },
};

class MainPage extends Component {
  constructor() {
    super()

    this.onEvent = this.onEvent.bind(this)
  }

  onEvent(args, print, runCommand) {
    console.log(args)
  }

  render() {
    const { classes } = this.props;

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

        <main className={classes.main}>
          <Paper className={classes.rootPaper} elevation={1}>
            <GraphEditor />
          </Paper>

          <Terminal
            onEvent={this.onEvent}
          />
        </main>
      </div>
    )
  }
}

MainPage.propTypes = {
  classes: PropTypes.object.isRequired,
};

function mapStateToProps(state) {
  return {
    fsm: state.fsm,
  }
}

function mapDispatchToProps(dispatch) {
  return {
    listActions: bindActionCreators(list, dispatch),
    addActions: bindActionCreators(add, dispatch),
    updateActions: bindActionCreators(update, dispatch),
    removeActions: bindActionCreators(remove, dispatch),
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(withStyles(styles)(MainPage));
