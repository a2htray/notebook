import {createStore} from 'redux'
import reducer from "./reducer";

const store = createStore(reducer, {
  first: 0,
  second: 10,
  third: 20,
})

export default store
