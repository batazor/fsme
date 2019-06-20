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

export function add(id) {
  return (dispatch, getState) => fetch(process.env.REACT_APP_API_URL, {
    method: 'POST',
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json',
    },
    credentials: 'include',
    body: JSON.stringify({
      ...getState().fsm.list[id],
      _id: undefined, // Delete id for new FSM
    }),
  })
    .then(response => response.json())
    .then(response => {
      dispatch({
        type: FSM.ADD,
        payload: response,
      })
    })
    .catch(error => console.error('error', error))
}

export function update(id) {
  return (dispatch, getState) => fetch(`${process.env.REACT_APP_API_URL}/${id}`, {
    method: 'PATCH',
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json',
    },
    credentials: 'include',
    body: JSON.stringify(getState().fsm.list[id]),
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

export function remove(id) {
  return (dispatch, getState) => fetch(`${process.env.REACT_APP_API_URL}/${id}`, {
    method: 'DELETE',
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json',
    },
    credentials: 'include',
  })
    .then(response => response.json())
    .then(response => {
      dispatch({
        type: FSM.REMOVE,
        payload: { id },
      })
    })
    .catch(error => console.error('error', error))
}

export function updateLocale(fsm) {
  return dispatch => dispatch({
    type: FSM.UPDATE,
    payload: fsm,
  })
}
