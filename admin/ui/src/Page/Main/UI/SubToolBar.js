import React from 'react'
import AppBar from '@material-ui/core/AppBar'
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import { Link } from "react-router-dom";

// required for react-router-dom < 6.0.0
// see https://github.com/ReactTraining/react-router/issues/6056#issuecomment-435524678
const AdapterLink = React.forwardRef((props, ref) => <Link innerRef={ref} {...props} />);

export default ({ match }) => {
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
