import React from 'react'
import ClickCounter from './components/ClickCounter'
import ControlPanel from './components/ControlPanel'
import ControlPanel2 from './components/ControlPanel2'
import ReduxControlPanel from './components/ReduxControlPanel'
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
      <h6>使用 Redux 的全局状态示例</h6>
      <ReduxControlPanel />
    </div>
  )
}

export default App
