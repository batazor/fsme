import React from 'react'
import PropTypes from 'prop-types'
import { makeStyles } from '@material-ui/styles'
import AppBar from '@material-ui/core/AppBar'
import Toolbar from '@material-ui/core/Toolbar'
import Typography from '@material-ui/core/Typography'
import IconButton from '@material-ui/core/IconButton'
import MenuIcon from '@material-ui/icons/Menu'
import AdapterLink from '../../../Components/AdapterLink'

const useStyles = makeStyles({
  menuButton: {
    marginLeft: -12,
    marginRight: 20,
  },
  grow: {
    flexGrow: 1,
    textDecorationLine: 'none',
  },
})

function ToolBar({ onChangeOpenDrawer }) {
  const classes = useStyles()

  return (
    <AppBar position="static">
      <Toolbar>
        <IconButton
          className={classes.menuButton}
          color="inherit"
          aria-label="Menu"

          onClick={onChangeOpenDrawer}
        >
          <MenuIcon />
        </IconButton>

        <Typography variant="h6" color="inherit" className={classes.grow} component={AdapterLink} to="/">
          FSME-UI
        </Typography>
      </Toolbar>
    </AppBar>
  )
}

ToolBar.propTypes = {
  onChangeOpenDrawer: PropTypes.func.isRequired,
}

export default ToolBar
