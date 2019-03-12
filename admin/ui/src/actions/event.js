import * as FSM from '../constants/fsm'
import { ApiClient, EventApi } from '../swagger/src'

export function sendEvent(nameEveent = "") {
  const client = new ApiClient()
  client.basePath = "http://localhost:44555"
  const api = new EventApi(client)
  api.sendEventFSM(nameEveent)
    .then(
      resp => console.warn('sendEvent', resp),
      err => console.warn('sendEvent err', err)
    )

  return d => d({
    type: "",
  })
}
