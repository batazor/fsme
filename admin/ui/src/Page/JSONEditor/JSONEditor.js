import React, { Component } from 'react'
import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import { updateLocale } from '../../actions/fsm'

import JSONEditor from '../../Components/JSONEditor'

class JSONEditorPage extends Component {
  constructor(props) {
    super(props)

    this.state = {}

    this.onChangeFSM = this.onChangeFSM.bind(this)
  }

  onChangeFSM(code) {
    this.props.updateLocaleActions(code)
  }

  render() {
    const fsm = this.props.fsm[this.props.match.params.id]

    return (
      <JSONEditor
        code={fsm}
        onChange={this.onChangeFSM}
      />
    )
  }
}

function mapStateToProps(state) {
  return {
    fsm: state.fsm.list,
  }
}

function mapDispatchToProps(dispatch) {
  return {
    updateLocaleActions: bindActionCreators(updateLocale, dispatch),
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(JSONEditorPage);
