
Function 构造函数，可用于动态创建新的函数，参数为字符串形式的参数，分别为函数的参数及函数体。

```js
let func1 = Function('a', 'b', 'return a+b')
console.log(func1(1, 2)) // 3
```

ES6 增强了 Function 构造函数的功能，使其支持默认参数和不定参数。

```js
let func2 = Function('a', 'b = 1', '...args', 'return a + b + args[0]')
console.log(func2(1, undefined, 2)) // 4
```
