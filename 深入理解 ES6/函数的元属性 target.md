在 ECMAScript 5 中，如果想确定一个函数是否通过 new 关键字被调用（或者说，判断该函数是否作为构造函数被调用），最流行的方式是使用 instanceof：

```js
function Person(name) {
  if (this instanceof Person) {
    this.name = name
  } else {
    throw new Error('必须通过 new 关键字调用 Person')
  }
}

var person = new Person('a2htray')
```

```js
var notAPerson = Person('a2htray')
```

```bash
VM2873:5 Uncaught Error: 必须通过 new 关键字调用 Person
    at Person (<anonymous>:5:11)
    at <anonymous>:2:18
```

但可以使用 call 来绕过此限制：

```js
var notAPerson2 = Person.call(person, 'a2htray2')
```

**使用函数的元属性 target 作检查**

元属性是指非对象的属性，其可以提供非对象目标的补充信息（例如 new ）。当调用函数的 `[[Construct]]` 方法时，new.target 被赋值
为 new 操作符的目标，通常是新创建对象实例，也就是函数体内 this 的构造函数；如果调用 [[Call]] 方法，则 new.target 的值为 undefined。

```js
function Person(name) {
  if (typeof new.target !== 'undefined') {
    this.name = name
  } else {
    throw new Error('必须通过 new 关键字来调用 Person。')
  }
}
let person = new Person('a2htray')
let notAPerson = Person.call(person, 'a2htray')
```

```bash
Uncaught Error: 必须通过 new 关键字来调用 Person。
    at Person (<anonymous>:5:11)
    at <anonymous>:1:25
```