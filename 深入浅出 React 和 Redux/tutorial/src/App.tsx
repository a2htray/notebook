import React from 'react'
import ClickCounter from './components/ClickCounter'
import ControlPanel from './components/ControlPanel'
import ControlPanel2 from './components/ControlPanel2'
import ReduxControlPanel from './components/ReduxControlPanel'
import ReduxControlPanel2 from './components/ReduxControlPanel2'
import ReduxControlPanel3 from './components/ReduxControlPanel3'
// import logo from './logo.svg';
import './App.css'

function App() {
  return (
    <div className="App">

      <ClickCounter />

      <ControlPanel />

      <ControlPanel2 />

      <ReduxControlPanel />
      <ReduxControlPanel2 />
      <ReduxControlPanel3 />

    </div>
  )
}

export default App
