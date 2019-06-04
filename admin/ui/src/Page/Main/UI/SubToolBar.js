import React from 'react'
import AppBar from '@material-ui/core/AppBar'
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import { Link } from "react-router-dom";

export default () => {
  return (
    <AppBar position="static" color="default">
      <Tabs
        indicatorColor="primary"
        textColor="primary"
        variant="scrollable"
        scrollButtons="auto"
      >
        <Link to="#">
          <Tab label="view" />
        </Link>
        <Link to="#">
          <Tab label="json-editor" />
        </Link>
      </Tabs>
    </AppBar>
  )
}
