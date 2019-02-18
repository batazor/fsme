import React, { Component } from 'react';
import { withStyles } from '@material-ui/core/styles';
import Terminal from 'terminal-in-react'

const styles = {
  root: {
    display: "flex",
    justifyContent: "center",
    overflow: 'auto',
  },
};

class Command extends Component {

  constructor(props) {
    super(props);

    this.state = {}
  }

  render() {
    const { classes } = this.props;

    return (
      <div className={classes.root}>
        <Terminal
          color='green'
          backgroundColor='black'
          barColor='black'
          style={{ fontWeight: "bold", fontSize: "1em" }}
          commands={{}}
          descriptions={{}}
          showActions={false}
          hideTopBar
          msg='You can write anything here. Example - Hello! My name is Foo and I like Bar.'
        />
      </div>
    );
  }
}

export default withStyles(styles)(Command);
