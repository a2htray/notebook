export const INCREMENT = 'INCREMENT'
export const DECREMENT = 'DECREMENT'

export default interface Action {
  type: string
  caption: string
}

export const increment = (caption: string): Action => {
  return {
    type: INCREMENT,
    caption: caption,
  }
}

export const decrement = (caption: string): Action => {
  return {
    type: DECREMENT,
    caption: caption,
  }
}
