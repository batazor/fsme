import React, { Component } from 'react'
import PropTypes from 'prop-types'
import { connect } from 'react-redux'

import GraphEditor from '../../Components/GraphEditor'

class GraphEditorPage extends Component {
  constructor(props) {
    super(props)

    this.state = {}
  }

  render() {
    return (
      <GraphEditor
        fsm={this.state.newFSM || this.props.fsm['ID_1']}
      />
    )
  }
}

GraphEditorPage.propTypes = {
  classes: PropTypes.object.isRequired,
};

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
)(GraphEditorPage);
