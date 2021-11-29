
ECMAScript 6 引入了 Super 引用的特性，使用它可以更便捷地访问对象原型。

在ECMAScript 5中可以这样实现：

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

let friend = {
  getGreeting() {
    return Object.getPrototypeOf(this).getGreeting.call(this) + ', hi!'
  }
}
```

```js
Object.setPrototypeOf(friend, person)
console.log(friend.getGreeting()) // hello, hi!
console.log(Object.getPrototypeOf(friend) === person) // true
```

```js
Object.setPrototypeOf(friend, dog)
console.log(friend.getGreeting()) // woof, hi!
console.log(Object.getPrototypeOf(friend) === dog) // true
```

ECMAScript 6 引入了 super 关键字。简单来说，Super 引用相当于指向对象原型的指针，实际上也就是 Object.getPrototypeOf(this) 的值。

**必须要在使用简写方法的对象中使用 Super 引用**

```js
let friend = {
  getGreeting() {
    return super.getGreeting.call(this) + ', hi!'
  }
}
```

