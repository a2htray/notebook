使用 const 声明的是常量，其值一旦被设定后不可更改

```js
const maxIdx = 1024
```

语法错误，修改常量值

```js
maxIdx = 1
```

语法错误，常量未赋值

```js
const name
```

常量同样也不会被提升至作用域顶部

```js
if (1 === 1) {
    const minIdx = 0
}
// 此处不能访问常量 minIdx
```

与 let 相似，在同一作用域用 const 声明已经存在的标识符也会导致语法错误

```js
var message = "hello"
var age = 18
```


语法错误
```js
const message
```

