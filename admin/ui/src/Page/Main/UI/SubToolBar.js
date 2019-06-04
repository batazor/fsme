import React from 'react'
import AppBar from '@material-ui/core/AppBar'
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import { Link } from "react-router-dom";

export default ({ match }) => {
  return (
    <AppBar position="static" color="default">
      <Tabs
        indicatorColor="primary"
        textColor="primary"
        variant="scrollable"
        scrollButtons="auto"
      >
        <Link to={`/fsm/${match.params.id}/view`}>
          <Tab label="view" />
        </Link>
        <Link to={`/fsm/${match.params.id}/json-editor`}>
          <Tab label="json-editor" />
        </Link>
      </Tabs>
    </AppBar>
  )
}
