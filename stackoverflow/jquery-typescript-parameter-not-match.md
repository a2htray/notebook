jQuery 中的 与 TypeScript 中的 string 类型不匹配问题
======================

[原地址](https://stackoverflow.com/questions/46007618/jquery-and-typescript-ts-argument-of-type-string-number-string-is-not)

### 问题

我有一个需要改写成 TypeScript 的大型的 Angular 项目。当前的项目中包含着大量 jQuery 的代码，同时在改写的过程中需要实现一个类型的定义。


下面这一行就发现了问题：

```js
if (["Priority", "PM"].indexOf($('option:selected', this).val()) >= 0) {}
```

出现了类型定义的问题。得到的提示如下：

```bash
[TS] Argument of type 'string | number | string[]' is 
not assignable to parameter of type 'string'
```

这是因为 `indexOf` 的函数签名是 `Array.indexOf(searchElement: string, fromIndex?: number): number`，但是 jQuery 的 `val` 的签名是 `JQuery.val(): string | number | string[] (+1 overload)`。

我需要一种方法去覆盖 `val` 的定义，只要一个 `string` 就行。为了不覆盖整个项目单个函数的签名，即只需要在类似的行做个性，我没有想到什么合适的方法。最终的目的就是为了全部脱离 jQuery。有人能指出一个更好的方法吗？

### 回答

我这边实现了你的需求。我使用了类型转换的方法。这需要额外的变量定义。

```ts
let optSel: string = <string>$('option:selected', this).val();

if ($(["Priority", "PM"]).index(optSel) >= 0) {}
```

这么做可以满足 `indexOf` 函数签名的定义。希望这个能帮助到一些人。

### 总结

这里其实是，当一个值具有多种类型时，如何将其转化成其中的一个具体类型。