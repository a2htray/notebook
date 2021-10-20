import React, {Component} from 'react'
import Counter from "./Counter";

interface ControlPanelProps {
}

interface ControlPanelState {
  total: number
}

const initValues = [0, 10, 20]

class ControlPanel extends Component<ControlPanelProps, ControlPanelState> {
  constructor(props: ControlPanelProps) {
    super(props)
    this.onUpdateCounter = this.onUpdateCounter.bind(this)

    this.state = {
      total: initValues.reduce((a, b) => a + b),
    }
  }

  render = () => {
    return (<div>
      {['first', 'second', 'third'].map((caption, idx) => {
        return <Counter key={idx} caption={caption} initValue={initValues[idx]} onUpdate={this.onUpdateCounter}/>
      })}
      <span>total: {this.state.total}</span>
    </div>)
  }

  onUpdateCounter(newValue: number, prevValue: number) {
    const total = this.state.total - prevValue + newValue
    this.setState({
      total,
    })
  }
}

export default ControlPanel
