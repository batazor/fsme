import React, { Component } from 'react'
import PropTypes from 'prop-types'
import { withStyles } from '@material-ui/core/styles'
import { withRouter } from 'react-router-dom'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import SpeedDial from '@material-ui/lab/SpeedDial'
import SpeedDialIcon from '@material-ui/lab/SpeedDialIcon'
import SpeedDialAction from '@material-ui/lab/SpeedDialAction'
import FileCopyIcon from '@material-ui/icons/FileCopyOutlined'
import SaveIcon from '@material-ui/icons/Save'
import ShareIcon from '@material-ui/icons/Share'
import DeleteIcon from '@material-ui/icons/Delete'
import AddIcon from '@material-ui/icons/Add'
import { add, update, remove } from '../../actions/fsm'

import styles from './styles'

class SpeedDialButton extends Component {
  constructor() {
    super()

    this.state = {
      open: false,
      hidden: false,
    }
  }

  handleVisibility = () => {
    this.setState(state => ({
      open: false,
      hidden: !state.hidden,
    }))
  };

  handleClick = () => {
    this.setState(state => ({
      open: !state.open,
    }))
  };

  handleOpen = () => {
    if (!this.state.hidden) { // eslint-disable-line
      this.setState({
        open: true,
      })
    }
  };

  handleClose = () => {
    this.setState({
      open: false,
    })
  };

  onRedirect = url => this.props.history.push(url) // eslint-disable-line

  onSave = () => {
    /* eslint-disable */
    if (this.props.match.params.id === 'new') {
      this.props.addActions(this.props.match.params.id)
    } else {
      this.props.updateActions(this.props.match.params.id)
    }
    /* eslint-enable */
  }

  onDelete = () => {
    const { id } = this.props.match.params // eslint-disable-line

    if (id && id !== 'new') {
      this.props.deleteFsmAction(id) // eslint-disable-line
        .then(() => {
          this.setState({ open: false })
          this.onRedirect('/fsm/new/json-editor')
        })
    }
  }

  render() {
    const { classes, match } = this.props
    const { hidden, open } = this.state

    const listButton = []
    if (match.params.id !== 'new') {
      listButton.push(
        <SpeedDialAction
          key="New"
          icon={<AddIcon />}
          tooltipTitle="New"
          tooltipPlacement="top-start"
          tooltipOpen
          onClick={() => this.onRedirect('/fsm/new/json-editor')}
        />,
      )
    }
    if (match.params.type) {
      listButton.push(
        <SpeedDialAction
          key="Save"
          icon={<SaveIcon />}
          tooltipTitle="Save"
          tooltipPlacement="top-start"
          tooltipOpen
          onClick={this.onSave}
        />,
      )

      if (match.params.id !== 'new') {
        listButton.push(
          <SpeedDialAction
            key="Copy"
            icon={<FileCopyIcon />}
            tooltipTitle="Copy"
            tooltipPlacement="top-start"
            tooltipOpen
            onClick={this.handleClick}
          />,
        )
        listButton.push(
          <SpeedDialAction
            key="Delete"
            icon={<DeleteIcon />}
            tooltipTitle="Delete"
            tooltipPlacement="top-start"
            tooltipOpen
            onClick={this.onDelete}
          />,
        )
      }
    }

    return (
      <SpeedDial
        ariaLabel="SpeedDial tooltip example"
        className={classes.speedDial}
        hidden={hidden}
        icon={<SpeedDialIcon />}
        onBlur={this.handleClose}
        onClick={this.handleClick}
        onClose={this.handleClose}
        onFocus={this.handleOpen}
        onMouseEnter={this.handleOpen}
        onMouseLeave={this.handleClose}
        open={open}
        direction="left"
      >
        { listButton.map(item => item) }
        <SpeedDialAction
          key="Share"
          icon={<ShareIcon />}
          tooltipTitle="Share"
          tooltipPlacement="top-start"
          tooltipOpen
          onClick={this.handleClick}
        />
      </SpeedDial>
    )
  }
}

SpeedDialButton.propTypes = {
  classes: PropTypes.object.isRequired, // eslint-disable-line
  match: PropTypes.shape({
    params: PropTypes.object.isRequired,
  }).isRequired,

  addActions: PropTypes.func.isRequired,
  updateActions: PropTypes.func.isRequired,
  deleteFsmAction: PropTypes.func.isRequired,
}

function mapStateToProps(state) {
  return {
    fsm: state.fsm.list,
  }
}

function mapDispatchToProps(dispatch) {
  return {
    addActions: bindActionCreators(add, dispatch),
    updateActions: bindActionCreators(update, dispatch),
    deleteFsmAction: bindActionCreators(remove, dispatch),
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(withStyles(styles)(withRouter(SpeedDialButton)))
