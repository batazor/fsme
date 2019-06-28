import * as React from 'react'
import PropTypes from 'prop-types'
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
class TrayWidget extends React.Component {
  static defaultProps = {};

  constructor(props) {
    super(props)
    this.state = {}
  }

  render() {
    const { classes, children } = this.props

    return <div className={classes.tray}>{children}</div>
  }
}

TrayWidget.propTypes = {
  classes: PropTypes.object.isRequired, // eslint-disable-line
  children: PropTypes.node.isRequired,
}

export default withStyles(styles)(TrayWidget)
