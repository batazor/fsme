import React, { Component } from 'react'
import PropTypes from 'prop-types'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { withStyles } from '@material-ui/core/styles'
import AppBar from '@material-ui/core/AppBar'
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import Paper from '@material-ui/core/Paper'
import { Link, Route } from "react-router-dom";

import Home from '../Home'
import GraphEditorPage from '../GraphEditor'
import JSONEditorPage from '../JSONEditor'
import Toolbar from './UI/ToolBar'
import Menu from './UI/Menu'
import Terminal from '../../Containers/Terminal'
import { list, add, update, remove } from '../../actions/fsm'
import { sendEvent } from '../../actions/event'

import styles from './styles'

class MainPage extends Component {
  constructor(props) {
    super(props)

    this.state = {
      isOpenMenu: false,
    }

    // Terminal
    this.onEvent = this.onEvent.bind(this)
    this.onSave = this.onSave.bind(this)

    // UI
    this.onChangeOpenDrawer = this.onChangeOpenDrawer.bind(this)
  }

  onEvent(args, print, runCommand) {
    args.shift()
    this.props.sendEvent(args.join(" "))
  }

  onSave(args, print, runCommand) {
    const { _id } = this.state.newFSM
    this.props.updateActions(_id, this.state.newFSM)
  }

  onChangeOpenDrawer() {
    this.setState({ isOpenMenu: !this.state.isOpenMenu })
  }

  componentDidMount() {
    this.props.listActions()
  }

  render() {
    const { classes } = this.props;

    return (
      <div className={classes.root}>
        <Toolbar
          onChangeOpenDrawer={this.onChangeOpenDrawer}
        />

        <Route path="/fsm/:id" component={SubAppBar} />

        <main className={classes.main}>
          <Paper className={classes.rootPaper} elevation={1}>
            <Route path="/" exact component={Home} />
            <Route path="/fsm/:id/view" exact component={GraphEditorPage} />
            <Route path="/fsm/:id/json-editor" exact component={JSONEditorPage} />

            <Terminal
              onEvent={this.onEvent}
              onSave={this.onSave}
            />
          </Paper>
        </main>

        <Menu
          fsm={this.props.fsm}
          isOpenMenu={this.state.isOpenMenu}

          onChangeOpenDrawer={this.onChangeOpenDrawer}
        />
      </div>
    )
  }
}

function SubAppBar() {
  return (
    <AppBar position="static" color="default">
      <Tabs
        indicatorColor="primary"
        textColor="primary"
        variant="scrollable"
        scrollButtons="auto"
      >
        <Link to="#">
          <Tab value="view" label="view" />
        </Link>
        <Link to="#">
          <Tab value="json-editor" label="json-editor" />
        </Link>
      </Tabs>
    </AppBar>
  )
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
