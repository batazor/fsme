import * as FSM from '../constants/fsm'

export function list(opts = {}) {
  return (dispatch, getState) => fetch(`${process.env.REACT_APP_API_URL}`, {
    method: 'GET',
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json',
    },
    credentials: 'include',
  })
    .then(response => response.json())
    .then(response => {
      dispatch({
        type: FSM.LIST,
        payload: response,
      })
    })
    .catch(error => console.error(error))
    // .catch(error => error.then(response => { throw response }))
}

export function get(id = 1) {
  return (dispatch, getState) => fetch(`${process.env.REACT_APP_API_URL}`, {
    method: 'GET',
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json',
    },
    credentials: 'include',
  })
    .then(response => response.json())
    .then(response => {
      dispatch({
        type: FSM.GET,
        payload: response,
      })
    })
    .catch(error => console.error('error', error))
}

export function add(opts = {}) {
  // const resp = FsmApi.addFSM(opts)
  // console.warn('addFSM', resp)
}

export function update(id = "ID_1", fsm) {
  return (dispatch, getState) => fetch(`${process.env.REACT_APP_API_URL}/${id}`, {
    method: 'PATCH',
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json',
    },
    credentials: 'include',
    body: JSON.stringify(fsm),
  })
    .then(response => response.json())
    .then(response => {
      dispatch({
        type: FSM.UPDATE,
        payload: response,
      })
    })
    .catch(error => console.error('error', error))
}

export function remove(id = 1) {
  // const resp = FsmApi.destroyFSM(id)
  // console.warn('destroyFSM', resp)
}
