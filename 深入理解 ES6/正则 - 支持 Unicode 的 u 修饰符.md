正则表达式可以完成简单的字符串操作，但默认将字符串中的每一个字符按照 16 位编码单元处理。为解决这个问题，ECMAScript 6 给正则表达式定义了一个支持 Unicode 的 u 修饰符。

```js
'𝌆'.length // 2

/^.$/.test('𝌆') // 2

/^.$/u.test('𝌆') // 2
```

`length` 方法的处理单元仍然是编码单元

`u` 修饰符：以 Unicode 编码处理字符串

判断当前引擎是否支持 `u` 修饰符：

```js
function hasRegExpU() {
  try {
    const _ = new RegExp('.', 'u')
    return true
  } catch (ex) {
    return false
  }
}
```
