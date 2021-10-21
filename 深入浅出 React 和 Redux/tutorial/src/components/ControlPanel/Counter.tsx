import React, {Component} from 'react'

interface CounterProps {
  caption: string
  initValue: number
}

interface CounterState {
  count: number
}

class Counter extends Component<CounterProps, CounterState> {
  constructor(props: CounterProps) {
    // 由外部组件传递的值，用 props 接收
    super(props)
    this.onClickDecrementButton = this.onClickDecrementButton.bind(this)
    this.onClickIncrementButton = this.onClickIncrementButton.bind(this);

    // 组件内值要变的，选择放在 state 中
    this.state = {
      count: props.initValue | 0,
    }
  }

  render = () => {
    console.log('lifecycle', 'render', this.props.caption)

    const {caption} = this.props
    const {count} = this.state

    return (
      <div>
        <button onClick={this.onClickIncrementButton}>+</button>
        <button onClick={this.onClickDecrementButton}>-</button>
        <span>{caption} count: {count}</span>
      </div>
    )
  }

  onClickIncrementButton() {
    this.setState({
      count: this.state.count + 1,
    })
  }

  onClickDecrementButton() {
    this.setState({
      count: this.state.count - 1,
    })
  }
  // 使用 componentWillReceiveProps 监听外部组件的传入的 props 值的变化
  componentWillReceiveProps(nextProps: Readonly<CounterProps>, nextContext: any) {
    console.log('componentWillReceiveProps', nextProps)
    this.setState({
      count: nextProps.initValue,
    })
  }

  componentWillMount() {
    console.log('lifecycle', 'componentWillMount', this.props.caption)
  }
  //
  // UNSAFE_componentWillMount() {
  //   console.log('lifecycle', 'UNSAFE_componentWillMount', this.props.caption)
  // }

  componentDidMount() {
    console.log('lifecycle', 'componentDidMount', this.props.caption)
  }
}


export default Counter
