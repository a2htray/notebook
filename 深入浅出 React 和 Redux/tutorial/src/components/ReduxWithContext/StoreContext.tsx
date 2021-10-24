import store from './store'
import React from 'react'
import {Store} from 'redux'
import Action from "./actions"
import {ComponentState} from "./reducer";

const StoreContext = React.createContext<{store: Store<ComponentState, Action>}>({
  store: store,
})

export default StoreContext
