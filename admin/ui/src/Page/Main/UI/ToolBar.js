import React from 'react'
import { makeStyles } from '@material-ui/styles'
import { Link } from 'react-router-dom'
import AppBar from '@material-ui/core/AppBar'
import Toolbar from '@material-ui/core/Toolbar'
import Typography from '@material-ui/core/Typography'
import IconButton from '@material-ui/core/IconButton'
import MenuIcon from '@material-ui/icons/Menu'

const useStyles = makeStyles({
  menuButton: {
    marginLeft: -12,
    marginRight: 20,
  },
  grow: {
    flexGrow: 1,
  },
})

export default props => {
  const classes = useStyles()

  return (
    <AppBar position="static">
      <Toolbar>
        <IconButton
          className={classes.menuButton}
          color="inherit"
          aria-label="Menu"

          onClick={props.onChangeOpenDrawer}
        >
          <MenuIcon />
        </IconButton>

        <Typography variant="h6" color="inherit" className={classes.grow}>
          <Link to="/">FSME-UI</Link>
        </Typography>
      </Toolbar>
    </AppBar>
  )
}
