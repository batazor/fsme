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
      return {
        ...state,
        fsm: {
          ...state.fsm,
          [action.payload._id]: action.payload[0],
        },
      }
    }
    case FSM.GET:
    case FSM.UPDATE:
    case FSM.ADD: {
      return {
        ...state,
        fsm: {
          ...state.fsm,
          [action.payload._id]: action.payload,
        },
      }
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
