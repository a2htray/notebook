react-dom-router - Route 中 component 和 render 的区别
===================================================

### 问题

关于 `react-router-dom` (v4.3.1) ，我有两个疑问：

1. 什么时候应该使用 `component` 或 `render`？

```js
<Route exact path='/u/:username/' component={ProfileComponent} />
<Route exact path='/u/:username/' render={() => <ProfileComponent />} />
```

2. 如何访问 `username` 变量？

### 回答

当你将一个组件传递到 `component` 属性时，通过 `props.match.params` 对象可以访问到路径参数。在你的例子中，应该使用 `props.match.params.username`。

```jsx harmony
class ProfileComponent extends React.Component {
  render() {
    return <div>{this.props.match.params.username}</div>;
  }
}
```

当使用 `render` 属性时， 在 `render` 方法中，就可以直接访问到路径参数。

```jsx harmony
<Route
  exact
  path='/u/:username/'
  render={(props) => 
    <ProfileComponent username={props.match.params.username}/>
  }
/>
```

因为 `component` 属性没有办法做到传递额外的值，所以当需要从外界传递值的时候，应该使用 `render`。

### 总结