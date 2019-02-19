import { createStore, compose, applyMiddleware } from 'redux'
import { routerMiddleware } from 'connected-react-router'
import { persistStore } from 'redux-persist'
import thunk from 'redux-thunk'
import createHistory from 'history/createBrowserHistory'
import rootReducer from '../reducers'

export const history = createHistory()

const enhancers = []

if (process.env.NODE_ENV !== 'production') {
  /* eslint-disable no-underscore-dangle */
  const devTools = window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__()
  enhancers.push(devTools)
  /* eslint-enable */
}

const store = createStore(
  rootReducer(history),
  compose(
    applyMiddleware(thunk, routerMiddleware(history)),
    ...enhancers,
  ),
)

export const persistor = persistStore(store)

export default store
