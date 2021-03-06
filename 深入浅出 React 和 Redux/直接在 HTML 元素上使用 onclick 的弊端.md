
* onclick 添加的事件处理函数是在全局环境下执行的，这污染了全局环境，很容易产生意料不到的后果；
* 给很多 DOM 元素添加 onclick 事件，可能会影响网页的性能，毕竟，网页需要的事件处理函数越多，性能就会越低；
* 对于使用 onclick 的 DOM 元素，如果要动态地从 DOM 树中删掉的话，需要把对应的时间处理器注销，假如忘了注销，就可能造成内存泄露，这样的bug很难被发现。

JSX 的 onClick 使用了事件委托 (event delegation) 的方式处理点击事件，无论有多少个 onClick 出现，其实最后都只在 DOM 树上添加了一个事件处理函数，挂在最顶层的 DOM 节点上。因为 React 控制了组件的生命周期，在 unmount 的时候自然能够清除相关的所有事件处理函数，内存泄露也不再是一个问题。