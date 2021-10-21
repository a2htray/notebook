// 定义返回动作的函数
import { Action as ReduxAction } from 'redux'
import * as ActionTypes from './ActionTypes'

export default interface Action extends ReduxAction<string> {
  caption: string
}

// increment 用于返回增加的 Action
export const increment = (caption: string): Action => {
  return {
    type: ActionTypes.INCREMENT,
    caption,
  }
}

// decrement 用于返回减少的 Action
export const decrement = (caption: string): Action => {
  return {
    type: ActionTypes.DECREMENT,
    caption,
  }
}
