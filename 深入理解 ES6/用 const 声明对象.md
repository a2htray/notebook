const 声明不允许修改绑定，但允许修改值

```js
const person = {
    age: 18,
}

// 允许
person.age = 19

// 语法错误
person = {
    age: 19,
}
```


