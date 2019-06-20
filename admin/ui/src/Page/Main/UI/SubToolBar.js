import React from 'react'
import { withStyles } from '@material-ui/core/styles'
import PropTypes from 'prop-types'
import AppBar from '@material-ui/core/AppBar'
import Tabs from '@material-ui/core/Tabs'
import Tab from '@material-ui/core/Tab'
import AdapterLink from '../../../Components/AdapterLink'

const styles = {}

function SubToolBar({ classes, match }) {
  return (
    <AppBar position="static" color="default">
      <Tabs
        indicatorColor="primary"
        textColor="primary"
        variant="scrollable"
        scrollButtons="auto"
        value={match.params.service}
      >
        <Tab
          component={AdapterLink}
          label="view"
          value="view"
          to={`/fsm/${match.params.id}/view`}
          activeClassName={classes.active}
        />
        <Tab
          component={AdapterLink}
          label="view-editor"
          value="view-editor"
          to={`/fsm/${match.params.id}/view-editor`}
          activeClassName={classes.active}
        />
        <Tab
          component={AdapterLink}
          label="json-editor"
          value="json-editor"
          to={`/fsm/${match.params.id}/json-editor`}
          activeClassName={classes.active}
        />
      </Tabs>
    </AppBar>
  )
}

SubToolBar.propTypes = {
  classes: PropTypes.object.isRequired, // eslint-disable-line
  match: PropTypes.shape({
    params: PropTypes.object.isRequired,
  }).isRequired,
}

export default withStyles(styles)(SubToolBar)
