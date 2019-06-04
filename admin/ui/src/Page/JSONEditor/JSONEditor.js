import React, { Component } from 'react'
import { connect } from 'react-redux'

import JSONEditor from '../../Components/JSONEditor'

class JSONEditorPage extends Component {
  constructor(props) {
    super(props)

    this.state = {}

    this.onChangeFSM = this.onChangeFSM.bind(this)
  }

  onChangeFSM(code) {
    // this.setState({ newFSM: code })
  }

  render() {
    return (
      <JSONEditor
        code={this.props.fsm}
        onChange={this.onChangeFSM}
      />
    )
  }
}

function mapStateToProps(state) {
  // Get id current FSM
  const idFSM = state.router.location.pathname.split('/')[2]

  return {
    fsm: state.fsm.fsm[idFSM],
  }
}

function mapDispatchToProps(dispatch) {
  return {}
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(JSONEditorPage);
