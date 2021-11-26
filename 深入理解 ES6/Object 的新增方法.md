当你想在 JavaScript 中比较两个值时，可能习惯于使用相等运算符（==）或全等运算符（===），许多开发者更喜欢后者，从而避免在比较
时触发强制类型转换的行为。但即使全等运算符也不完全准确，如：

```js
console.log(+0 == -0) // true
console.log(+0 === -0) // true
console.log(Object.is(+0, -0)) // false
console.log(NaN == NaN) // false
console.log(NaN === NaN) // false
console.log(Object.is(NaN, NaN)) // true
console.log(5 == 5) // true
console.log(5 == '5') // true
console.log(5 === 5) // true
console.log(5 === '5') // false
console.log(Object.is(5, 5)) // true
console.log(Object.is(5, '5')) // false
```