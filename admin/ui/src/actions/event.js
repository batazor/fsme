import * as FSM from '../constants/fsm'

export function sendEvent(id, nameEveent = '') {
  return dispatch => fetch(`${process.env.REACT_APP_API_URL}/${id}/event/${nameEveent}`, {
    method: 'POST',
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json',
    },
    credentials: 'include',
  })
    .then(response => response.json())
    .then(response => {
      dispatch({
        type: FSM.UPDATE,
        payload: response,
      })
    })
    .catch(error => console.error(error))
}
