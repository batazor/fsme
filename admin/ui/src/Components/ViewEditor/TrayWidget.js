import * as React from 'react'
import { withStyles } from '@material-ui/core/styles'

const styles = {
  tray: {
    background: '#141414',
    flexGrow: 0,
    flexShrink: 0,
  },
}

/**
 * @author Dylan Vorster
 */
class TrayWidget extends React.Component<TrayWidgetProps, TrayWidgetState> {
	static defaultProps: TrayWidgetProps = {};

	constructor(props: TrayWidgetProps) {
	  super(props)
	  this.state = {}
	}

	render() {
	  const { classes } = this.props

	  return <div className={classes.tray}>{this.props.children}</div>
	}
}

export default withStyles(styles)(TrayWidget)
