import React, { Component } from 'react'
import PropTypes from 'prop-types'
import { bindActionCreators } from 'redux'
import { withRouter, Route } from 'react-router-dom'
import { connect } from 'react-redux'
import { withStyles } from '@material-ui/core/styles'
import Paper from '@material-ui/core/Paper'

import Home from '../Home'
import GraphEditor from '../GraphEditor'
import ViewEditorPage from '../ViewEditor'
import JSONEditorPage from '../JSONEditor'
import Toolbar from './UI/ToolBar'
import Menu from './UI/Menu'
import SubToolBar from './UI/SubToolBar'
import SpeedDial from '../../Containers/SpeedDial'
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

  componentDidMount() {
    this.props.listAction() // eslint-disable-line
  }

  onEvent(args) {
    const { id } = this.props.match.params  // eslint-disable-line

    args.shift()
    this.props.sendEventAction(id, args.join(' ')) // eslint-disable-line
  }

  onRedirect = url => this.props.history.push(url) // eslint-disable-line

  onDelete(args, print, runCommand) { // eslint-disable-line
    const { id } = this.props.match.params // eslint-disable-line

    if (id && id !== 'new') {
      this.props.deleteFsmAction(id) // eslint-disable-line
        .then(() => this.onRedirect('/fsm/new/json-editor'))
    }
  }

  onChangeOpenDrawer() {
    this.setState(state => ({ isOpenMenu: !state.isOpenMenu }))
  }

  render() {
    const { classes, fsm } = this.props
    const { isOpenMenu } = this.state

    return (
      <div className={classes.root}>
        <Toolbar
          onChangeOpenDrawer={this.onChangeOpenDrawer}
        />

        <Route path="/fsm/:id" component={SubToolBar} />

        <main className={classes.main}>
          <Paper className={classes.rootPaper} elevation={1}>
            <Route path="/" exact component={Home} />
            <Route path="/fsm/:id/view" exact component={GraphEditor} />
            <Route path="/fsm/:id/view-editor" exact component={ViewEditorPage} />
            <Route path="/fsm/:id/json-editor" exact component={JSONEditorPage} />

            <Terminal
              onEvent={this.onEvent}
              onDelete={this.onDelete}
            />
          </Paper>
        </main>

        <Menu
          fsm={fsm}
          isOpenMenu={isOpenMenu}

          onChangeOpenDrawer={this.onChangeOpenDrawer}
        />

        <Route path="/" exect component={SpeedDial} />
        <Route path="/:service/:id/:type" component={SpeedDial} />
      </div>
    )
  }
}

MainPage.propTypes = {
  classes: PropTypes.object.isRequired, // eslint-disable-line
  fsm: PropTypes.shape({
    Name: PropTypes.string.isRequired,
  }).isRequired,
  match: PropTypes.shape({
    params: PropTypes.object.isRequired,
  }).isRequired,
  history: PropTypes.arrayOf(PropTypes.string).isRequired,

  listAction: PropTypes.func.isRequired,
  sendEventAction: PropTypes.func.isRequired,
  deleteFsmAction: PropTypes.func.isRequired,
}

function mapStateToProps(state) {
  return {
    fsm: state.fsm.list,
  }
}

function mapDispatchToProps(dispatch) {
  return {
    listAction: bindActionCreators(list, dispatch),
    sendEventAction: bindActionCreators(sendEvent, dispatch),
    deleteFsmAction: bindActionCreators(remove, dispatch),
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(withStyles(styles)(withRouter(MainPage)))
