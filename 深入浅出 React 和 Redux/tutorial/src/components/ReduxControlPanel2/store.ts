import {createStore} from 'redux'
import reducer, {ComponentState} from "./reducer";

const initialState: ComponentState = {
  first: 0,
  second: 10,
  third: 20,
}

const store = createStore(reducer, initialState)

export default store
