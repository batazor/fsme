import React, { Component } from 'react'
import { BrowserRouter as Router, Route } from 'react-router-dom'

import Main from '../Main'

class App extends Component {
  render() {
    return (
      <div className="App">
        <Router>
          <Route path="/" exact component={Main} />
          <Route path="/:service/:id/:type" exact component={Main} />
        </Router>
      </div>
    )
  }
}

export default App
