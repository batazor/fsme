import React from 'react'
import PropTypes from 'prop-types'
import AppBar from '@material-ui/core/AppBar'
import Tabs from '@material-ui/core/Tabs'
import Tab from '@material-ui/core/Tab'
import AdapterLink from '../../../Components/AdapterLink'

function SubToolBar({ match }) {
  return (
    <AppBar position="static" color="default">
      <Tabs
        indicatorColor="primary"
        textColor="primary"
        variant="scrollable"
        scrollButtons="auto"
      >
        <Tab
          component={AdapterLink}
          label="view"
          value="view"
          to={`/fsm/${match.params.id}/view`}
        />
        <Tab
          component={AdapterLink}
          label="view-editor"
          value="view-editor"
          to={`/fsm/${match.params.id}/view-editor`}
        />
        <Tab
          component={AdapterLink}
          label="json-editor"
          value="json-editor"
          to={`/fsm/${match.params.id}/json-editor`}
        />
      </Tabs>
    </AppBar>
  )
}

SubToolBar.propTypes = {
  match: PropTypes.shape({
    params: PropTypes.object.isRequired,
  }).isRequired,
}

export default SubToolBar
