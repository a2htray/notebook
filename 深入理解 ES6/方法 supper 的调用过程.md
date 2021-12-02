在 ECMAScript 6 中正式将方法定义为一个函数，它会有一个内部的 `[[HomeObject]]` 属性来容纳这个方法从属的对象。

```js
let person = {
  getGreeting() {
    return 'hello'
  }
}

let friend = {
  getGreeting() {
    return super.getGreeting() + ' hi'
  }
}

Object.setPrototypeOf(friend, person)

console.log(friend.getGreeting()) // 'hello hi'
```

Super 的所有引用都通过 `[[HomeObject]]` 属性来确定后续的运行过程：

* 第一步是在 `[[HomeObject]]` 属性上调用 `Object.getPrototypeOf()` 方法来检索原型的引用；
* 然后搜寻原型找到同名函数；
* 最后，设置this绑定并且调用相应的方法；


`super.getGreeting() 等价于 person.getGreeting.call(this)`
