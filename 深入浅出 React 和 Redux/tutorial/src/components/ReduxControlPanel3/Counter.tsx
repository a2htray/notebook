import React, {Component} from 'react'

interface CounterProps {
  caption: string
  value: number
  onIncrement: (caption: string) => void
  onDecrement: (caption: string) => void
}

class Counter extends Component<CounterProps> {
  render = () => {
    const {caption, value, onDecrement, onIncrement} = this.props

    return (<div>
      <button onClick={() => onIncrement(caption)}>+</button>
      <button onClick={() => onDecrement(caption)}>-</button>
      <span>{caption} count: {value}</span>
    </div>)
  }
}

export default Counter
