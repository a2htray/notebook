import React, {Component} from 'react'

const counterStyle = {
  margin: '16px',
}

// Component<{}, {count: number}>
//   {} 表示 Props 实现的接口
//   {count: number} 表示 State 实现的接口
class ClickCounter extends Component<{}, {count: number}> {
  constructor(props: {}) {
    super(props)
    // 使用 bind 方法，将方法内的 this 指向 ClickCounter 实例
    this.onClickButton = this.onClickButton.bind(this)
    this.state = {
      count: 1,
    }
  }

  onClickButton() {
    this.setState({
      count: this.state.count + 1
    })
  }

  render() {
    return (<div style={counterStyle}>
      <button onClick={this.onClickButton}>Click me</button>
      <div>
        Click Count: {this.state.count}
      </div>
    </div>)
  }
}

export default ClickCounter
