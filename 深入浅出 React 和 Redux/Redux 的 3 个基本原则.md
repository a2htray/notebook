Redux 的 3 个基本原则
==========================

* 唯一数据源（Single Source of Truth）

唯一数据源指的是应用的状态数据应该只存储在唯一的一个Store上。

* 保持状态只读（State is read-only）

保持状态只读，就是说不能去直接修改状态，要修改Store的状态，必须要通过派发一个 action 对象完成。

* 数据改变只能通过纯函数完成（Changes are made with pure functions）

所说的纯函数就是 Reducer。Redux 的 reducer 只负责计算状态，却并不负责存储状态。



