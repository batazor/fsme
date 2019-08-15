import React from 'react'
import { NavLink } from 'react-router-dom'

// required for react-router-dom < 6.0.0
// see https://github.com/ReactTraining/react-router/issues/6056#issuecomment-435524678
const AdapterLink = React.forwardRef((props, ref) => <NavLink innerRef={ref} {...props} />) // eslint-disable-line

export default AdapterLink
