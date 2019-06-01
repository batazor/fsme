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
  grow: {
    flexGrow: 1,
  },
  menuButton: {
    marginLeft: -12,
    marginRight: 20,
  },
})
