import React, { Component } from 'react'
import { withStyles } from '@material-ui/core/styles'
import SpeedDial from '@material-ui/lab/SpeedDial';
import SpeedDialIcon from '@material-ui/lab/SpeedDialIcon';
import SpeedDialAction from '@material-ui/lab/SpeedDialAction';
import FileCopyIcon from '@material-ui/icons/FileCopyOutlined';
import SaveIcon from '@material-ui/icons/Save';
import ShareIcon from '@material-ui/icons/Share';
import DeleteIcon from '@material-ui/icons/Delete';
import AddIcon from '@material-ui/icons/Add';

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
          icon={<AddIcon />}
          tooltipTitle="New"
          tooltipPlacement="top-start"
          tooltipOpen
          onClick={this.handleClick}
        />
        {
          match.params.type && [
            <SpeedDialAction
              icon={<SaveIcon />}
              tooltipTitle="Save"
              tooltipPlacement="top-start"
              tooltipOpen
              onClick={this.handleClick}
            />,
            <SpeedDialAction
              icon={<FileCopyIcon />}
              tooltipTitle="Copy"
              tooltipPlacement="top-start"
              tooltipOpen
              onClick={this.handleClick}
            />,
            <SpeedDialAction
              icon={<DeleteIcon />}
              tooltipTitle="Delete"
              tooltipPlacement="top-start"
              tooltipOpen
              onClick={this.handleClick}
            />,
          ]
        }
        <SpeedDialAction
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

export default withStyles(styles)(SpeedDialButton);
