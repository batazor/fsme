import React, { Component } from 'react'
import PropTypes from 'prop-types'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { withStyles } from '@material-ui/core/styles'
import AppBar from '@material-ui/core/AppBar'
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import Toolbar from '@material-ui/core/Toolbar'
import Typography from '@material-ui/core/Typography'
import IconButton from '@material-ui/core/IconButton'
import MenuIcon from '@material-ui/icons/Menu'
import Paper from '@material-ui/core/Paper'

import GraphEditor from '../../Containers/GraphEditor'
import JSONEditor from '../../Containers/JSONEditor'
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

    this.state = {
      typeEditor: 'view',
    }

    this.onEvent = this.onEvent.bind(this)
    this.onChangeTypeEditor = this.onChangeTypeEditor.bind(this)
  }

  onEvent(args, print, runCommand) {
    args.shift()
    this.props.sendEvent(args.join(" "))
  }

  onChangeTypeEditor(event, typeEditor) {
    this.setState({ typeEditor })
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
          <AppBar position="static" color="default">
            <Tabs
              value={this.state.typeEditor}
              onChange={this.onChangeTypeEditor}
              indicatorColor="primary"
              textColor="primary"
              variant="scrollable"
              scrollButtons="auto"
            >
              <Tab value="view" label="view" />
              <Tab value="json-editor" label="json-editor" />
            </Tabs>
          </AppBar>

          <Paper className={classes.rootPaper} elevation={1}>
            {
              this.state.typeEditor === "view" && (
                  <GraphEditor fsm={this.props.fsm['ID_1']} />
              )
            }
            {
              this.state.typeEditor === "json-editor" && (
                <Paper className={classes.rootPaper} elevation={1}>
                  <JSONEditor code={this.props.fsm['ID_1']} />
                </Paper>
              )
            }

            <Terminal onEvent={this.onEvent} />
          </Paper>
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
