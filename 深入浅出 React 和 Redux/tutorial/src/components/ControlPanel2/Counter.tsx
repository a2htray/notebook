import React, {Component} from 'react'

interface CounterProps {
  caption: string
  initValue: number

  onUpdate(newValue: number, prevValue: number): void
}

interface CounterState {
  count: number
}

class Counter extends Component<CounterProps, CounterState> {
  constructor(props: CounterProps) {
    super(props)
    this.onClickIncrementButton = this.onClickIncrementButton.bind(this)
    this.onClickDecrementButton = this.onClickDecrementButton.bind(this)
    this.updateCount = this.updateCount.bind(this);

    this.state = {
      count: this.props.initValue | 0,
    }
  }

  render = () => {
    return (<div>
      <button onClick={this.onClickIncrementButton}>+</button>
      <button onClick={this.onClickDecrementButton}>-</button>
      <span>{this.props.caption} count: {this.state.count}</span>
    </div>)
  }

  onClickIncrementButton() {
    this.updateCount(true)
  }

  onClickDecrementButton() {
    this.updateCount(false)
  }

  updateCount(isIncrement: boolean) {
    const prevValue = this.state.count
    const newValue = isIncrement ? prevValue + 1 : prevValue - 1

    this.setState({
      count: newValue,
    })

    this.props.onUpdate(newValue, prevValue)
  }
}

export default Counter
