import * as React from 'react'
import { withStyles } from '@material-ui/core/styles'

const styles = {
  trayItem: {
    color: 'white',
    padding: '5px',
    margin: '0px 10px',
    border: 'solid 1px',
    borderRadius: '5px',
    marginBottom: '2px',
    cursor: 'pointer',
  },
}

class TrayItemWidget extends React.Component {
  constructor(props: TrayItemWidgetProps) {
    super(props)
    this.state = {}
  }

  render() {
    const { classes, color, name, model } = this.props // eslint-disable-line

    return (
      <div
        style={{ borderColor: color }}
        draggable
        onDragStart={event => {
          event.dataTransfer.setData('storm-diagram-node', JSON.stringify(model))
        }}
        className={classes.trayItem}
      >
        {name}
      </div>
    )
  }
}

export default withStyles(styles)(TrayItemWidget)
