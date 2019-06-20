import React from 'react'
import PropTypes from 'prop-types'
import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import { updateLocale } from '../../actions/fsm'

import ViewEditor from '../../Components/ViewEditor/ViewEditor'

function ViewEditorPage(props) {
  const { fsm, match, updateLocaleActions } = props

  return (
    <ViewEditor
      fsm={fsm[match.params.id]}
      onChange={updateLocaleActions}
    />
  )
}

ViewEditorPage.propTypes = {
  fsm: PropTypes.shape({
    Name: PropTypes.string.isRequired,
  }).isRequired,
  match: PropTypes.shape({
    params: PropTypes.object.isRequired,
  }).isRequired,
  updateLocaleActions: PropTypes.func.isRequired,
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
)(ViewEditorPage)
