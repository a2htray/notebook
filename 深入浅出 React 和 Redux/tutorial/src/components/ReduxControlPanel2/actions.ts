import {Action as ReduxAction} from 'redux'

// 动作关键字
export const INCREMENT = 'INCREMENT'
export const DECREMENT = 'DECREMENT'

// Action 动作类型
export default interface Action extends ReduxAction<string> {
  caption: string
}

export const increment = (caption: string): Action => {
  return {
    type: INCREMENT,
    caption: caption
  }
}

export const decrement = (caption: string): Action => {
  return {
    type: DECREMENT,
    caption: caption,
  }
}
