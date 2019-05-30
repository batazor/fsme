import * as FSM from '../constants/fsm'

export function sendEvent(nameEveent = "") {
  console.log(nameEveent)

  return (dispatch, getState) => fetch(`${process.env.REACT_APP_API_URL}/1/event/${nameEveent}`, {
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
        payload: response[0],
      })
    })
    .catch(error => console.error(error))
}
