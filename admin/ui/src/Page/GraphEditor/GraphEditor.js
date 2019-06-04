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
      <GraphEditor fsm={this.props.fsm} />
    )
  }
}

GraphEditorPage.propTypes = {
  classes: PropTypes.object.isRequired,
};

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
)(GraphEditorPage);
