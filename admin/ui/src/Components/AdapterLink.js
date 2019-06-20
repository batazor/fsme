import React from 'react'
import { Link } from 'react-router-dom'

// required for react-router-dom < 6.0.0
// see https://github.com/ReactTraining/react-router/issues/6056#issuecomment-435524678
const AdapterLink = React.forwardRef((props, ref) => <Link innerRef={ref} {...props} />)

export default AdapterLink
