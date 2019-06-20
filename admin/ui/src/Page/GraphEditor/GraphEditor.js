import React from 'react'
import PropTypes from 'prop-types'
import { connect } from 'react-redux'

import GraphEditor from '../../Components/GraphEditor'

function GraphEditorPage(props) {
  const { fsm, match } = props

  return (
    <GraphEditor
      fsm={fsm[match.params.id]}
    />
  )
}

GraphEditorPage.propTypes = {
  classes: PropTypes.object.isRequired, // eslint-disable-line
  fsm: PropTypes.shape({
    Name: PropTypes.string.isRequired,
  }).isRequired,
  match: PropTypes.shape({
    params: PropTypes.object.isRequired,
  }).isRequired,
}

function mapStateToProps(state) {
  return {
    fsm: state.fsm.list,
  }
}

function mapDispatchToProps() {
  return {}
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(GraphEditorPage)
