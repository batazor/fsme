export default theme => ({
  root: {
    display: 'grid',
    gridTemplateRows: 'auto 1fr',
    flex: 1,
    overflow: 'auto',
  },
  main: {
    overflow: "auto",
    display: "flex",
    flexDirection: "column",
  },
  rootPaper: {
    display: 'grid',
    gridTemplateColumns: '1fr auto',
    flex: 1,
    [theme.breakpoints.down('sm')]: {
      gridTemplateColumns: '1fr',
      gridTemplateRows: 'minmax(20em, auto)',
    },
  },
  drawer: {
    width: "20em",
    flexShrink: 0,
  },
  drawerPaper: {
    width: "20em",
  },
  drawerHeader: {
    display: 'flex',
    alignItems: 'center',
    padding: '0 8px',
    ...theme.mixins.toolbar,
    justifyContent: 'flex-end',
  },
})
