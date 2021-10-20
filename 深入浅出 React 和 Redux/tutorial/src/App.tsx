import React from 'react'
import ClickCounter from './components/ClickCounter'
import ControlPanel from './components/ControlPanel'
import ControlPanel2 from './components/ControlPanel2'
// import logo from './logo.svg';
// import './App.css';

function App() {
  return (
    <div className="App">
      <h6>显示点击次数的组件</h6>
      <ClickCounter />
      <h6>Props 和 State 组件属性的使用</h6>
      <ControlPanel />
      <h6>组件向外传递数据</h6>
      <ControlPanel2 />
    </div>
  )
}

export default App
