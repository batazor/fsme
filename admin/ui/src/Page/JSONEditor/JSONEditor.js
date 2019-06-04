import React, { Component } from 'react'
import { connect } from 'react-redux'

import JSONEditor from '../../Components/JSONEditor'

class JSONEditorPage extends Component {
  constructor(props) {
    super(props)

    this.state = {
      newFSM: {},
    }

    this.onChangeFSM = this.onChangeFSM.bind(this)
  }

  onChangeFSM(code) {
    this.setState({ newFSM: code })
  }

  render() {
    return (
      <JSONEditor
        code={this.state.newFSM || this.props.fsm['ID_1']}
        onChange={this.onChangeFSM}
      />
    )
  }
}

function mapStateToProps(state) {
  return {
    fsm: state.fsm.fsm,
  }
}

function mapDispatchToProps(dispatch) {
  return {}
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(JSONEditorPage);
