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
import {
  sendEvent,
} from '../../actions/event'

import styles from './styles'

class MainPage extends Component {
  constructor(props) {
    super(props)

    this.onEvent = this.onEvent.bind(this)
  }

  onEvent(args, print, runCommand) {
    args.shift()
    this.props.sendEvent(args.join(" "))
  }

  componentDidMount() {
    this.props.listActions()
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
            <GraphEditor fsm={this.props.fsm['ID_1']} />
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
    fsm: state.fsm.fsm,
  }
}

function mapDispatchToProps(dispatch) {
  return {
    listActions: bindActionCreators(list, dispatch),
    addActions: bindActionCreators(add, dispatch),
    updateActions: bindActionCreators(update, dispatch),
    removeActions: bindActionCreators(remove, dispatch),
    sendEvent: bindActionCreators(sendEvent, dispatch),
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(withStyles(styles)(MainPage));
