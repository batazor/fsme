import React, { Component } from 'react'
import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import { updateLocale } from '../../actions/fsm'

import ViewEditor from '../../Components/ViewEditor/ViewEditor'

class ViewEditorPage extends Component {
  render() {
    const fsm = this.props.fsm[this.props.match.params.id]

    return (
      <ViewEditor
        fsm={fsm}
        onChange={this.props.updateLocaleActions}
      />
    )
  }
}

ViewEditorPage.propTypes = {}

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
)(ViewEditorPage)
