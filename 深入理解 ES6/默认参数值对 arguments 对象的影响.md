> 在 ES5 非严格模式下，函数命名参数的变化会体现在 arguments 对象中

```js
function mixArgs(first, second) {
  console.log(first === arguments[0])
  console.log(second === arguments[1])
  
  first = 'c'
  second = 'd'
  
  console.log(first === arguments[0])
  console.log(second === arguments[1])
}

mixArgs('a', 'b')
```

输出：

```bash
true
true
true
true
```

> 在 ES5 的严格模式下，取消了 arguments 对象的这一行为

```js
function mixArgs(first, second) {
  'use strict'
  
  console.log(first === arguments[0])
  console.log(second === arguments[1])
  
  first = 'c'
  second = 'd'
  
  console.log(first === arguments[0])
  console.log(second === arguments[1])
}

mixArgs('a', 'b')
```

输出：

```bash
true
true
false
false
```

在 ECMAScript 6 中，如果一个函数使用了默认参数值，则无论是否显式定义了严格模式，arguments 对象的行为都将与 ECMAScript 5 严格模式下保持一致。

> 默认参数值，非原始值传参

```js
function getValue() {
  return 5
}

function add(first, second = getValue()) {
  return first + second
}

console.log(add(1, 1)) // 2
console.log(add(1)) // 6，在此时，函数 getValue 才会被调用
```

切记，初次解析函数声明时不会调用 getValue() 方法，只有当调用 add() 函数且不传入第二个参数时才会调用。`默认参数是在函数调用时求值`
