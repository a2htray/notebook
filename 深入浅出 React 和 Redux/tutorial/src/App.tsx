import React from 'react'
import ClickCounter from './components/ClickCounter'
import ControlPanel from './components/ControlPanel'
import ControlPanel2 from './components/ControlPanel2'
import ReduxControlPanel from './components/ReduxControlPanel'
import ReduxControlPanel2 from './components/ReduxControlPanel2'
import ReduxControlPanel3 from './components/ReduxControlPanel3'
import './App.css'
import ReduxWithContext from './components/ReduxWithContext'
import ReduxConnect from './components/ReduxConnect'

function App() {
  return (
    <div className="App">
      <ClickCounter />
      <ControlPanel />
      <ControlPanel2 />
      <ReduxControlPanel />
      <ReduxControlPanel2 />
      <ReduxControlPanel3 />
      <ReduxWithContext />

      <ReduxConnect />

    </div>
  )
}

export default App
