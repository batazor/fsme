export default theme => ({
  root: {
    flex: 1,
  },
  main: {
    overflow: "auto",
    display: "flex",
    flexDirection: "column",
  },
  rootPaper: {
    display: 'grid',
    gridTemplateColumns: '1fr 1fr',
    flex: 1,
    [theme.breakpoints.down('sm')]: {
      gridTemplateColumns: '1fr',
      gridTemplateRows: 'minmax(20em, auto)',
    },
  },
  speedDial: {
    position: 'absolute',
    top: theme.spacing(4),
    right: theme.spacing(3),
  },
})
