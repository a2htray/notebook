import * as actionTypes from './actionTypes'
import {ComponentDispatcher} from './dispatcher'

export interface Action {
  type: string
  caption: string
}

export const increment = (caption: string) => {
  ComponentDispatcher.dispatch({
    type: actionTypes.INCREMENT,
    caption: caption,
  })
}

export const decrement = (caption: string) => {
  ComponentDispatcher.dispatch({
    type: actionTypes.DECREMENT,
    caption: caption,
  })
}
