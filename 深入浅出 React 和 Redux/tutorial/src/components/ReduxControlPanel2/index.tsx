import React, {Component} from 'react'
import Counter from './Counter';
import store from './store'
import {ComponentState} from './reducer'

interface ControlPanelState {
  values: number[]
}

class ControlPanel extends Component<{}, ControlPanelState> {
  constructor(props: {}) {
    super(props)
  }

  componentDidMount() {
    store.subscribe(this.onChange)
  }

  onChange = () => {
    console.log('聪明组件与傻瓜组件日志', 'call onChange')
    this.forceUpdate()
  }

  render = () => {
    return (<React.Fragment>
      <h6>聪明组件与傻瓜组件1</h6>
      {['first', 'second', 'third'].map((caption, _) => {
        return <Counter key={caption} caption={caption} count={store.getState()[caption]} />
      })}
      <span>total: {total(store.getState())}</span>
    </React.Fragment>)
  }
}

export default ControlPanel

function total(state: ComponentState): number {
  let sum = 0

  for (let k in state) {
    if (!isNaN(state[k])) {
      sum += state[k]
    }
  }

  return sum
}
