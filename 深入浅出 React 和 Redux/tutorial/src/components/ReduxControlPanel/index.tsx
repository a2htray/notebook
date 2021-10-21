import React, {Component} from 'react'
import Counter from "./Counter";
import store from './store';
import {ComponentState} from './reducer'

interface ControlPanelProps {
}

interface ControlPanelState {
  total: number
}

class ControlPanel extends Component<ControlPanelProps, ControlPanelState> {
  constructor(props: ControlPanelProps) {
    super(props)

    this.state = {
      total: total(store.getState()),
    }
  }

  render = () => {
    return (<div>
      <h6>使用 Redux 的全局状态示例</h6>
      {['first', 'second', 'third'].map((caption, idx) => {
        return <Counter key={idx} caption={caption} />
      })}
      <span>total: {this.state.total}</span>
    </div>)
  }

  componentDidMount() {
    store.subscribe(this.onChange)
  }

  onChange = () => {
    console.log('control panel: call onChange')
    
    this.setState({
      total: total(store.getState())
    })
  }
}

function total(state: ComponentState): number {
  let sum = 0

  for (let k in state) {
    if (!isNaN(state[k])) {
      sum += state[k]
    }
  }

  return sum
}

export default ControlPanel
