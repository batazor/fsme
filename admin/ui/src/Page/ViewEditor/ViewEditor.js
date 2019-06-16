import React, { Component } from 'react'
import PropTypes from 'prop-types'
import { connect } from 'react-redux'

import ViewEditor from '../../Components/ViewEditor/ViewEditor'

class ViewEditorPage extends Component {
  constructor(props) {
    super(props)

    this.state = {}
  }

  render() {
    const fsm = this.props.fsm[this.props.match.params.id]

    return (
      <ViewEditor fsm={fsm} />
    )
  }
}

ViewEditorPage.propTypes = {
  classes: PropTypes.object.isRequired,
};

function mapStateToProps(state) {
  return {
    fsm: state.fsm.list,
  }
}

function mapDispatchToProps(dispatch) {
  return {}
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(ViewEditorPage);
