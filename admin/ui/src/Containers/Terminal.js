import React, { Component } from 'react'
import { withStyles } from '@material-ui/core/styles'
import Terminal from 'terminal-in-react'

const styles = {
  root: {
    display: 'flex',
    justifyContent: 'center',
    overflow: 'auto',
  },
}

class Command extends Component {
  constructor(props) {
    super(props)

    this.state = {
      commands: {
        event: props.onEvent,
        delete: props.onDelete,
      },
      descriptions: {
        event: 'enter `event [nameEvent]` for send event and change state',
        delete: 'enter `delete` for delete current FSM',
      },
    }
  }

  render() {
    const { classes } = this.props
    const { commands, descriptions } = this.state

    return (
      <div className={classes.root}>
        <Terminal
          color="green"
          backgroundColor="black"
          barColor="black"
          style={{ fontWeight: 'bold', fontSize: '1em' }}
          commands={commands}
          descriptions={descriptions}
          showActions={false}
          hideTopBar
          msg="These shell commands are defined internally.  Type `help' to see this list."
        />
      </div>
    )
  }
}

export default withStyles(styles)(Command)
