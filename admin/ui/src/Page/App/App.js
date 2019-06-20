import React from 'react'
import { BrowserRouter as Router, Route } from 'react-router-dom'

import Main from '../Main'

function App() {
  return (
    <div className="App">
      <Router>
        <Route path="/" exact component={Main} />
        <Route path="/:service/:id/:type" exact component={Main} />
      </Router>
    </div>
  )
}

export default App
