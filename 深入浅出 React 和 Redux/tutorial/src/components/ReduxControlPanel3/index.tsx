import React, {Component} from 'react'
import Counter from "./Counter";
import store from "./store";
import {decrement, increment} from "./actions";

interface ControlPanelState {
  values: number[]
}

class ControlPanel extends Component<{}, ControlPanelState> {
  constructor(props: {}) {
    super(props)
    this.state = {
      values: this.getOwnState(),
    }
  }

  componentDidMount = () => {
    store.subscribe(() => this.setState({
      values: this.getOwnState(),
    }))
  }

  getOwnState = (): number[] => {
    const {first, second, third} = store.getState()
    return [first, second, third]
  }

  render = () => {
    console.log('聪明组件与傻瓜组件2', 'this.state', this.state)
    const {values} = this.state
    return (<div>
      <h6>聪明组件与傻瓜组件2</h6>
      {['first', 'second', 'third'].map((caption, i) => {
        return <Counter key={caption} caption={caption} value={values[i]} onIncrement={this.onIncrement} onDecrement={this.onDecrement} />
      })}
      <span>total: {values.reduce((a, b) => a+b)}</span>
    </div>)
  }

  onIncrement = (caption: string) => {
    console.log('聪明组件与傻瓜组件2', 'onIncrement', 'caption', caption)
    store.dispatch(increment(caption))
  }

  onDecrement = (caption: string) => {
    store.dispatch(decrement(caption))
  }
}

export default ControlPanel
