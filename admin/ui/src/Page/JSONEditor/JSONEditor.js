import React from 'react'
import PropTypes from 'prop-types'
import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import { updateLocale } from '../../actions/fsm'

import JSONEditor from '../../Components/JSONEditor'

function JSONEditorPage(props) {
  const { fsm, match, updateLocaleActions } = props

  return (
    <JSONEditor
      code={fsm[match.params.id]}
      onChange={updateLocaleActions}
    />
  )
}

JSONEditorPage.propTypes = {
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
)(JSONEditorPage)
