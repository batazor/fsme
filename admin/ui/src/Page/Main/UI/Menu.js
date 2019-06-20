import React from 'react'
import PropTypes from 'prop-types'
import { makeStyles } from '@material-ui/styles'
import { Link } from 'react-router-dom'
import IconButton from '@material-ui/core/IconButton'
import Drawer from '@material-ui/core/Drawer'
import List from '@material-ui/core/List'
import ListItem from '@material-ui/core/ListItem'
import ListItemText from '@material-ui/core/ListItemText'
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft'
import Divider from '@material-ui/core/Divider'
import Typography from '@material-ui/core/Typography'

const useStyles = makeStyles({
  drawer: {
    width: '20em',
    flexShrink: 0,
  },
  drawerPaper: {
    width: '20em',
  },
  drawerHeader: {
    display: 'flex',
    alignItems: 'center',
    padding: '0 8px',
    justifyContent: 'space-between',
  },
})

function Menu({ fsm, isOpenMenu, onChangeOpenDrawer }) {
  const classes = useStyles()

  return (
    <Drawer
      className={classes.drawer}
      variant="persistent"
      anchor="left"
      open={isOpenMenu}
      classes={{
        paper: classes.drawerPaper,
      }}
    >
      <div className={classes.drawerHeader}>
        <Typography variant="h3" color="inherit">
          FSME-UI
        </Typography>

        <IconButton onClick={onChangeOpenDrawer}>
          <ChevronLeftIcon />
        </IconButton>
      </div>

      <Divider />

      <List>
        {
          Object.keys(fsm)
            .filter(key => key !== 'new')
            .map(key => (
              <Link key={`/fsm/${fsm[key]._id}`} to={`/fsm/${fsm[key]._id}/view`}>
                <ListItem button key={key}>
                  <ListItemText primary={fsm[key].Name} />
                </ListItem>
              </Link>
            ))
        }
      </List>
    </Drawer>
  )
}

Menu.propTypes = {
  isOpenMenu: PropTypes.bool.isRequired,
  fsm: PropTypes.shape({
    Name: PropTypes.string.isRequired,
  }).isRequired,
  onChangeOpenDrawer: PropTypes.func.isRequired,
}

export default Menu
