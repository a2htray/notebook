import EventEmitter = NodeJS.EventEmitter
import {ComponentDispatcher} from './dispatcher'
import {Action} from './actions'
import {DECREMENT, INCREMENT} from "./actionTypes";

const countValues: {
  [key: string]: number
} = {
  first: 0,
  second: 10,
  third: 20,
}

const EVENT_CHANGE = 'EVENT_CHANGE'

class counterStore extends EventEmitter {
  dispatchToken = ''

  getCountValues = () => {
    return countValues
  }

  emitChange = () => {
    this.emit(EVENT_CHANGE)
  }

  addChangeListener = (callback: (...args: any[]) => void) => {
    this.on(EVENT_CHANGE, callback)
  }

  removeChangeListener = (callback: (...args: any[]) => void) => {
    this.removeListener(EVENT_CHANGE, callback)
  }

}

const CounterStore = new counterStore()

CounterStore.dispatchToken = ComponentDispatcher.register((action): void => {
  if (action.type === INCREMENT) {
    countValues[action.caption]++
    CounterStore.emitChange()
  } else if (action.type === DECREMENT) {
    countValues[action.caption]--
    CounterStore.emitChange()
  }
})


class summaryStore extends EventEmitter {
  dispatchToken = ''
  getSummary = (): number => {
    let summary = 0
    const values = CounterStore.getCountValues()
    for (let key in values) {
      summary += values[key]
    }
    return summary
  }

  emitChange = () => {
    this.emit(EVENT_CHANGE)
  }

  addChangeListener = (callback: (...args: any[]) => void) => {
    this.on(EVENT_CHANGE, callback)
  }

  removeChangeListener = (callback: (...args: any[]) => void) => {
    this.removeListener(EVENT_CHANGE, callback)
  }
}

const SummaryStore = new summaryStore()
SummaryStore.dispatchToken = ComponentDispatcher.register((action) => {
  if (action.type === INCREMENT || action.type === DECREMENT) {
    ComponentDispatcher.waitFor([CounterStore.dispatchToken])
    SummaryStore.emitChange()
  }
})

export {
  CounterStore,
  SummaryStore,
}

// export const CounterStore = Object.assign({}, EventEmitter.prototype, {
//   getCountValues: function () {
//     return countValues
//   },
//   emitChange: function () {
//     this.emit()
//   }
// })
