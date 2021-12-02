TS - 对象可以为 null
=========================

[原问题](https://stackoverflow.com/questions/49431880/ts2531-object-is-possibly-null/58774464#58774464)

### 问题

我有下面一个函数：

```js
uploadPhoto() {
    var nativeElement: HTMLInputElement = this.fileInput.nativeElement;

    this.photoService.upload(this.vehicleId, nativeElement.files[0])
        .subscribe(x => console.log(x));
}
```

然而，在 `nativeElement.files[0]` 上，我得到一个 typescript 类型错误，`Object is possibly 'null'`,有人能解决吗？

### 回答

TypeScript 3.7 中引入了 `Optional Chaining`，针对有可能为 `null` 的值，可以这样操作：

```typescript
nativeElement?.file?.name
```



### 总结
