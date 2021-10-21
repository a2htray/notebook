import Action from './actions'
import * as ActionTypes from './ActionTypes'

export interface ComponentState {
  [K: string]: number
}

const reducer = (state: ComponentState={}, action: Action): ComponentState => {
  const { caption } = action
  if (action.type === ActionTypes.INCREMENT) {
    return { ...state, [caption]: state[caption] + 1 }
  } else {
    return { ...state, [caption]: state[caption] - 1 }
  }
}

// Reducer 函数
export default reducer
