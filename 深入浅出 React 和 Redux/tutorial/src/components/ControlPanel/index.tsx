import React, {Component} from 'react'
import Counter from "./Counter";

interface ControlPanelProps {
}

interface ControlPanelState {
  initValues: number[]
}

const initialState: ControlPanelState = {
  initValues: [0, 10, 20],
}

class ControlPanel extends Component<ControlPanelProps, ControlPanelState> {
  constructor(props: ControlPanelProps) {
    super(props)
    this.onClickRerender = this.onClickRerender.bind(this)

    this.state = initialState
  }

  render = () => {
    const {initValues} = this.state
    console.log('render', 'initValues', initValues)

    return (
      <div>
        <h6>Props 和 State 组件属性的使用</h6>
        {['first', 'second', 'third'].map((caption, idx) => {
          console.log('array map', 'initValue', initValues[idx])
          return <Counter key={idx} caption={caption} initValue={initValues[idx]}/>
        })}
        <button onClick={this.onClickRerender}>Click me to re-render</button>
      </div>
    )
  }

  onClickRerender() {
    console.log('onClickRerender')
    this.setState(initialState)
  }
}

export default ControlPanel
