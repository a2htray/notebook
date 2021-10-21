import {Dispatcher} from 'flux'
import {Action} from './actions'

const ComponentDispatcher = new Dispatcher<Action>()

export {
  ComponentDispatcher,
}
