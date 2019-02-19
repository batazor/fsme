import * as FSM from '../constants/fsm'
import { FsmApi } from '../swagger/src'

export function list(opts = {}) {
  FsmApi.getFSMList(opts, resp => console.warn('getFSMList', resp))
}

export function get(id = 1) {
  FsmApi.getFSM(id, resp => console.warn('getFSM', resp))
}

export function add(opts = {}) {
  FsmApi.addFSM(opts, resp => console.warn('addFSM', resp))
}

export function update(id = 1, opts = {}) {
  FsmApi.updateFSM(id, opts, resp => console.warn('updateFSM', resp))
}

export function remove(id = 1) {
  FsmApi.destroyFSM(id, resp => console.warn('destroyFSM', resp))
}
