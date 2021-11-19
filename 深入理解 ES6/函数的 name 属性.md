ECMAScript 6 程序中所有的函数的 name 属性都有一个合适的值。

```js
function funcName1() {}
let funcName2 = function () {}
console.log(funcName1.name, funcName2.name) // funcName1 funcName2
```

```js
// 具名函数的名称优先级更高
let funcName3 = function funcName4() {}
console.log(funcName3.name) // funcName4
```

```js
// 使用 bind 创建的函数，带有 bound 标识
function funcName5() {
  console.log(this.name)
}
let obj = {
  name: 'Object',
}
let bindFunc = funcName5.bind(obj)
console.log(bindFunc.name) // bound funcName5
```

函数 name 属性的值不一定引用同名变量，它只是协助调试用的额外信息，所以不能使用 name 属性的值来获取对于函数的引用。
