import { createStore } from 'redux'
import reducer from './reducer'

const initValues = {
  first: 0,
  second: 10, 
  third: 20,
}
// store 结合 reducer 和初始值，创造一个 store
const store = createStore(reducer, initValues)

export default store
