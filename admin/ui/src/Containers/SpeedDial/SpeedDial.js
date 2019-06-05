import React, { Component } from 'react'
import { withStyles } from '@material-ui/core/styles'
import { withRouter } from 'react-router-dom'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import SpeedDial from '@material-ui/lab/SpeedDial';
import SpeedDialIcon from '@material-ui/lab/SpeedDialIcon';
import SpeedDialAction from '@material-ui/lab/SpeedDialAction';
import FileCopyIcon from '@material-ui/icons/FileCopyOutlined';
import SaveIcon from '@material-ui/icons/Save';
import ShareIcon from '@material-ui/icons/Share';
import DeleteIcon from '@material-ui/icons/Delete';
import AddIcon from '@material-ui/icons/Add';
import { add, update } from '../../actions/fsm'

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
    }));
  };

  handleClick = () => {
    this.setState(state => ({
      open: !state.open,
    }));
  };

  handleOpen = () => {
    if (!this.state.hidden) {
      this.setState({
        open: true,
      });
    }
  };

  handleClose = () => {
    this.setState({
      open: false,
    });
  };

  onRedirect = url => this.props.history.push(url)

  onSave = () => {
    if (this.props.match.params.id === "new") {
      this.props.addActions(this.props.match.params.id)
    } else {
      this.props.updateAction(this.props.match.params.id)
    }
  }

  render() {
    const { classes, match } = this.props;

    return (
      <SpeedDial
        ariaLabel="SpeedDial tooltip example"
        className={classes.speedDial}
        hidden={this.state.hidden}
        icon={<SpeedDialIcon />}
        onBlur={this.handleClose}
        onClick={this.handleClick}
        onClose={this.handleClose}
        onFocus={this.handleOpen}
        onMouseEnter={this.handleOpen}
        onMouseLeave={this.handleClose}
        open={this.state.open}
        direction="left"
      >
        <SpeedDialAction
          key="New"
          icon={<AddIcon />}
          tooltipTitle="New"
          tooltipPlacement="top-start"
          tooltipOpen
          onClick={() => this.onRedirect('/fsm/new/json-editor')}
        />
        {
          match.params.type && [
            <SpeedDialAction
              key="Save"
              icon={<SaveIcon />}
              tooltipTitle="Save"
              tooltipPlacement="top-start"
              tooltipOpen
              onClick={this.onSave}
            />,
            <SpeedDialAction
              key="Copy"
              icon={<FileCopyIcon />}
              tooltipTitle="Copy"
              tooltipPlacement="top-start"
              tooltipOpen
              onClick={this.handleClick}
            />,
            <SpeedDialAction
              key="Delete"
              icon={<DeleteIcon />}
              tooltipTitle="Delete"
              tooltipPlacement="top-start"
              tooltipOpen
              onClick={this.handleClick}
            />,
          ]
        }
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

function mapStateToProps(state) {
  return {
    fsm: state.fsm.list,
  }
}

function mapDispatchToProps(dispatch) {
  return {
    addActions: bindActionCreators(add, dispatch),
    updateActions: bindActionCreators(update, dispatch),
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(withStyles(styles)(withRouter(SpeedDialButton)))
