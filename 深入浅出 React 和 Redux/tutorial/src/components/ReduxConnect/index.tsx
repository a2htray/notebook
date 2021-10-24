import React, {Component} from 'react'
import {createStore} from 'redux'
import {connect, Provider} from 'react-redux'

const INCREMENT = 'INCREMENT'
const DECREMENT = 'DECREMENT'

interface Action {
  type: string
}

function increment(): Action {
  return {
    type: INCREMENT,
  }
}

function decrement(): Action {
  return {
    type: DECREMENT,
  }
}

interface AppState {
  value: number
}

const reducer = (state: AppState={value: 0}, action: Action): AppState => {
  switch (action.type) {
    case INCREMENT:
      return {value: state.value+1}
    case DECREMENT:
      return {value: state.value+1}
    default:
      return state
  }
}

const store = createStore(reducer, {value: 10})

interface CounterProps {
  value?: number
  onIncrement(): void
  onDecrement(): void
}

class Counter extends Component<CounterProps, any> {
  render = () => {
    return <div>
      <button onClick={this.props.onIncrement}>+</button>
      <button onClick={this.props.onDecrement}>-</button>
      <span>value: {this.props.value}</span>
    </div>
  }
}

const CounterConnect = connect((state: AppState, _) => {
  return {
    value: state.value,
  }
}, (dispatch, _) => {
  return {
    onIncrement: () => {
      dispatch(increment())
    },
    onDecrement: () => {
      dispatch(decrement())
    }
  }
})(Counter)

interface ReduxConnectProps {
}

interface ReduxConnectState {
}

export default // @ts-ignore
class ReduxConnect extends Component<ReduxConnectProps, ReduxConnectState> {
  render = () => {
    return (<div>
      <Provider store={store}>
        <h6>Redux connect 方法的使用</h6>
        <CounterConnect />
      </Provider>
    </div>)
  }
}
