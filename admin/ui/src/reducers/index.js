import { persistCombineReducers } from 'redux-persist'
import { connectRouter } from 'connected-react-router'
import localForage from 'localforage'
import fsm from './fsm'

const config = {
  key: 'fsme-ui',
  storage: localForage,
  blacklist: ['router'],
}

export default history => persistCombineReducers(config, {
  router: connectRouter(history),

  fsm,
})
