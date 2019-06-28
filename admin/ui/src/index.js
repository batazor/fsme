import React from 'react'
import ReactDOM from 'react-dom'
import { Provider } from 'react-redux'
import { ConnectedRouter } from 'connected-react-router'
import { PersistGate } from 'redux-persist/es/integration/react'
import store, { history, persistor } from './store/configureStore'
import App from './Page/App'
import * as serviceWorker from './serviceWorker'

ReactDOM.render(
  (
    <Provider store={store}>
      <PersistGate
        // loading={<Loading />}
        persistor={persistor}
      >
        <ConnectedRouter history={history}>
          <App />
        </ConnectedRouter>
      </PersistGate>
    </Provider>
  ), document.getElementById('root'), // eslint-disable-line
)

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: http://bit.ly/CRA-PWA
serviceWorker.unregister()
