import React, {Component} from 'react'
import store from "./store";
import {decrement, increment} from "./actions";

interface CounterProps {
  caption: string
  count: number
}

// Counter 傻瓜组件不需要维护自己的状态
// 只利用 props 的变化而重新渲染
class Counter extends Component<CounterProps> {
  render = () => {
    const {caption, count} = this.props
    return (
      <div>
        <button onClick={this.onIncrement}>+</button>
        <button onClick={this.onDecrement}>-</button>
        <span>{caption} count: {count}</span>
      </div>
    )
  }

  onIncrement = () => {
    store.dispatch(increment(this.props.caption))
  }

  onDecrement = () => {
    store.dispatch(decrement(this.props.caption))
  }
}

export default Counter
