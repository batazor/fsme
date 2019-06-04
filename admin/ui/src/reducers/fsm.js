import * as FSM from '../constants/fsm'

const initialState = {
  fsm: {
    new: {
      FSM: {
        State: null,
        Transitions: {},
        Events: {},
        Callbacks: {}
      },
      _id: "new",
      description: "Description",
      title: "Title",
    },
  },
}

export default function update(state: Object = initialState, action: Object): Object {
  switch (action.type) {
    case FSM.LIST: {
      return push(state, action.payload)
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
    fsm: {
      ...state.fsm,
      [payload._id]: payload,
    },
  }
}
