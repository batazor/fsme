import * as FSM from '../constants/fsm'

const initialState = {
  list: {
    new: {
      FSM: {
        State: null,
        Transitions: {},
        Events: {},
        Callbacks: {}
      },
      _id: "new",
      description: "Description",
      name: "Title",
    },
  },
}

export default function update(state: Object = initialState, action: Object): Object {
  switch (action.type) {
    case FSM.LIST: {
      action.payload.forEach(fsm => state = push(state, fsm))
      return state
    }
    case FSM.GET: {
      return push(state, action.payload)
    }
    case FSM.UPDATE: {
      return push(state, action.payload)
    }
    case FSM.ADD: {
      return push(state, action.payload)
    }
    case FSM.REMOVE: {
      const STATE = state
      delete STATE.fsm[action.payload.id]
      return STATE
    }
    default:
      return state
  }
}

function push(state, payload) {
  return {
    ...state,
    list: {
      ...state.list,
      [payload._id]: payload,
    },
  }
}
