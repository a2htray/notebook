TypeScript 中 React 如何初始化 State
============================

[问题](https://stackoverflow.com/questions/60209459/how-to-set-initial-react-state-in-typescript)

### 问题

在使用 TypeScript 和 React 过程中，如何有效地初始化 state。因为 `currentVehicle` 不能赋值为一个空的对象，所以下面的代码后报错。
我应该如何进行初始化操作，有最佳实践吗？

```typescript jsx
interface State{
    currentVehicle:Vehicle
}

export default class extends Component<Props, State> {
   state:State={
     currentVehicle:{}
   }
}
```

### 回答

你可以进行如下转换：

```typescript
state:State = {
    currentVehicle:{} as Vehicle
}
```

### 总结