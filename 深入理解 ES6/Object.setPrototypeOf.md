ECMAScript 5 中添加了 Object.getPrototypeOf() 方法来返回任意指定对象的原型，但仍缺少对象在实例化后改变原型的标准方法。

在 ECMAScript 6 中添加了 Object.setPrototypeOf() 方法来改变这一现状，通过这个方法可以改变任意指定对象的原型，它接受两个参数：
被改变原型的对象及替代第一个参数原型的对象：

```js
let person = {
  getGreeting() {
    return 'hello'
  }
}
let dog = {
  getGreeting() {
    return 'woof'
  }
}

let friend = Object.create(person)
console.log(Object.getPrototypeOf(friend) === person)

Object.setPrototypeOf(friend, dog)
console.log(Object.getPrototypeOf(friend) === dog)
```

```bash
true
true
```

对象原型的真实值被储存在内部专用属性 `[[Prototype]]` 中