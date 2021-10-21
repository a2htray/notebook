import Action, {DECREMENT, INCREMENT} from "./actions";

export interface ComponentState {
  [K: string]: number
}

// reducer 处理接收到动作的逻辑
const reducer = (state: ComponentState={}, action: Action): ComponentState => {
  switch (action.type) {
    case INCREMENT:
      return {...state, [action.caption]: state[action.caption]+1}
    case DECREMENT:
      return {...state, [action.caption]: state[action.caption]-1}
    default:
      return state
  }
}

export default reducer
