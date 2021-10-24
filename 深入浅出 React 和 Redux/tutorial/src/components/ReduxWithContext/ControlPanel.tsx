import React, {Component} from 'react'
import Counter from "./Counter";

interface ControlPanelProps {
}

interface ControlPanelState {
}

class ControlPanel extends Component<ControlPanelProps, ControlPanelState> {
  constructor(props: ControlPanelProps) {
    super(props)
  }

  render = () => {
    return (<div>
      {['first', 'second', 'third'].map((caption, i) => {
        return <Counter key={i} caption={caption} />
      })}
    </div>)
  }
}

export default ControlPanel
