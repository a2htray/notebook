import React, {Component} from 'react'
import {Store} from 'redux'
import {ComponentState} from './reducer'
import Action, {increment, decrement} from './actions'
import StoreContext from "./StoreContext";

interface CounterProps {
  caption: string
}

interface Context {
  store: Store<ComponentState, Action>
}

class Counter extends Component<CounterProps, unknown> {
  constructor(props: CounterProps, ctx: Context) {
    super(props, ctx)
    console.log('Component With Context', ctx)
  }

  render = () => {
    return (
      <StoreContext.Consumer>
        {({store}) => {
          return <div>
            <button onClick={() => store.dispatch(increment(this.props.caption))}>+</button>
            <button onClick={() => store.dispatch(decrement(this.props.caption))}>-</button>
            <span>value: {store.getState()[this.props.caption]}</span>
          </div>
        }}

      </StoreContext.Consumer>
      )
  }
}

export default Counter
