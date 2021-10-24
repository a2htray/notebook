import React from 'react'
import store from './store'
import ControlPanel from "./ControlPanel";
import StoreContext from "./StoreContext";

export default function ReduxWithContext() {

  store.subscribe(() => {
    
  })

  return (
    <StoreContext.Provider value={{
      store: store,
    }}>
      <h6>Redux with Context</h6>
      <ControlPanel />
    </StoreContext.Provider>

  )
}
