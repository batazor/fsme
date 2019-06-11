import React, { Component } from 'react'
import PropTypes from 'prop-types'
import { bindActionCreators } from 'redux'
import { withRouter } from 'react-router-dom'
import { connect } from 'react-redux'
import { withStyles } from '@material-ui/core/styles'
import Paper from '@material-ui/core/Paper'
import { Route } from "react-router-dom";

import Home from '../Home'
import GraphEditorPage from '../GraphEditor'
import JSONEditorPage from '../JSONEditor'
import Toolbar from './UI/ToolBar'
import Menu from './UI/Menu'
import SubToolBar from './UI/SubToolBar'
import SpeedDial from '../../Containers/SpeedDial';
import Terminal from '../../Containers/Terminal'
import { list, remove } from '../../actions/fsm'
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
    this.onDelete = this.onDelete.bind(this)

    // UI
    this.onChangeOpenDrawer = this.onChangeOpenDrawer.bind(this)
  }

  onEvent(args, print, runCommand) {
    args.shift()
    this.props.sendEvent(args.join(" "))
  }

  onDelete(args, print, runCommand) {
    const { id } = this.props.match.params

    if (id && id !== "new") {
      this.props.deleteFsmAction(id)
        .then(res => this.onRedirect('/fsm/new/json-editor'))
    }
  }

  onChangeOpenDrawer() {
    this.setState({ isOpenMenu: !this.state.isOpenMenu })
  }

  onRedirect = url => this.props.history.push(url)

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

        <Route path="/fsm/:id" component={SubToolBar} />

        <main className={classes.main}>
          <Paper className={classes.rootPaper} elevation={1}>
            <Route path="/" exact component={Home} />
            <Route path="/fsm/:id/view" exact component={GraphEditorPage} />
            <Route path="/fsm/:id/json-editor" exact component={JSONEditorPage} />

            <Terminal
              onEvent={this.onEvent}
              onDelete={this.onDelete}
            />
          </Paper>
        </main>

        <Menu
          fsm={this.props.fsm}
          isOpenMenu={this.state.isOpenMenu}

          onChangeOpenDrawer={this.onChangeOpenDrawer}
        />

        <Route path="/" exect component={SpeedDial} />
        <Route path="/:service/:id/:type" component={SpeedDial} />
      </div>
    )
  }
}

MainPage.propTypes = {
  classes: PropTypes.object.isRequired,
};

function mapStateToProps(state) {
  return {
    fsm: state.fsm.list,
  }
}

function mapDispatchToProps(dispatch) {
  return {
    listActions: bindActionCreators(list, dispatch),
    sendEvent: bindActionCreators(sendEvent, dispatch),
    deleteFsmAction: bindActionCreators(remove, dispatch),
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(withStyles(styles)(withRouter(MainPage)));
