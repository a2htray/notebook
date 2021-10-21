import React, {Component} from 'react'
import { Unsubscribe } from 'redux';
import { decrement, increment } from './actions';
import store from './store';

interface CounterProps {
  caption: string
}

interface CounterState {
  count: number
  unsubscribe: Unsubscribe,
}

class Counter extends Component<CounterProps, CounterState> {
  constructor(props: CounterProps) {
    super(props)
    this.onChange = this.onChange.bind(this)
    this.onIncrement = this.onIncrement.bind(this)
    this.onDecrement = this.onDecrement.bind(this)

    this.state = {
      count: store.getState()[this.props.caption],
      unsubscribe: () => {},
    }
  }

  render = () => {
    return (<div>
      <button onClick={this.onIncrement}>+</button>
      <button onClick={this.onDecrement}>-</button>
      <span>{this.props.caption} count: {this.state.count}</span>
    </div>)
  }
  componentDidMount() {
    this.setState({
      unsubscribe: store.subscribe(this.onChange)
    })
  }
  
  componentWillUnmount = () => {
    this.state.unsubscribe()
  }

  onChange() {
    console.log('counter: call onChange')
    this.setState({
      count: store.getState()[this.props.caption],
    })
  }

  onIncrement() {
    store.dispatch(increment(this.props.caption))
  }

  onDecrement() {
    store.dispatch(decrement(this.props.caption))
  }
 
}

export default Counter
