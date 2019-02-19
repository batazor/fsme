import * as FSM from '../constants/fsm'
import { ApiClient, FsmApi } from '../swagger/src'

export function list(opts = {}) {
  const client = new ApiClient()
  client.basePath = "http://localhost:34801"
  const api = new FsmApi(client)
  api.getFSMList(opts)
    .then(
      resp => console.warn('getFSMList', resp),
      err => console.warn('getFSMList err', err)
    )

  return d => d({
    type: "",
  })
}

export function get(id = 1) {
  // const resp = FsmApi.getFSM(id)
  // console.warn('getFSM', resp)
}

export function add(opts = {}) {
  // const resp = FsmApi.addFSM(opts)
  // console.warn('addFSM', resp)
}

export function update(id = 1, opts = {}) {
  // const resp = FsmApi.updateFSM(id, opts)
  // console.warn('updateFSM', resp)
}

export function remove(id = 1) {
  // const resp = FsmApi.destroyFSM(id)
  // console.warn('destroyFSM', resp)
}
