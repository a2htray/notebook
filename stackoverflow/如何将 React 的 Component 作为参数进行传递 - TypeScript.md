如何将 React 的 Component 作为参数进行传递 - TypeScript
========================================================

### 问题

使用场景非常简单，在一个 Component 仓库中，我希望能够动态地调整不同的 React Component 的显示：

```js
import { Component } from 'react';
import DummyComponent from './DummyComponent';

class Registry {
  TestComponent: Component = DummyComponent;
}
```

`DummyComponent`

```js
import { Component } from 'react';

export default class DummyComponent extends Component {
  constructor(props: any) {
    super(props);
    throw new Error('Cannot use this dummy component');
  }
}
```

但在 TS 下报错，如下：

`ts(2740): Type 'typeof DummyComponent' is missing the following properties from type 'Component<{}, {}, any>': context, setState, forceUpdate, render, and 3 more.`

想请问哪里错了，我仅仅希望 `TestComponent` 是一个 React Component 就行，它不需要太多的属性。

线上复现的地址 [StackBlitz example](https://stackblitz.com/edit/react-ts-aa73j8)，报错信息可能略有不同，但我相信应该是同一个问题。

`tsconfig.json`

```json
{
  "compilerOptions": {
    "target": "es2015",
    "jsx": "react",
    "esModuleInterop": true,
    "forceConsistentCasingInFileNames": true,
    "lib": ["esnext"],
    "noEmit": true,
    "noFallthroughCasesInSwitch": true,
    "noImplicitReturns": true,
    "noUnusedLocals": true,
    "noUnusedParameters": true,
    "strict": true
  },
  "include": ["src/**/*"]
}
```

### 回答

`Component` 是组件类的一个**实例**，但是 `DummyCompoent` 是一个类。

应该是这样的：

```js
class Registry {
  TestComponent: ComponentClass = DummyComponent;
}
```

在这个问题中，`TestComponent` 不能只限定为类组件，可以定义如下：

```js
class Registry {
  TestComponent: ComponentType = DummyComponent;
}
```

### 总结

`React.ComponentClass` `React.ComponentType`

使用场景：在一个文件中定义多个路由（react-dom-router），实现 URL 与 Component 的一一对应。同时，编写一个处理多个路由的方法，输出类似以下的代码：

```js
<Route path={route.path} render={props => (
    <route.component {...props} />
)}/>
...
```

关键在于如何声明 `component` 属性的类型。

组件分类组件 `ComponentClass` 和 `FunctionComponent`，而 `ComponentType` 的定义是两者其一：

```js
type ComponentType<P = {}> = ComponentClass<P> | FunctionComponent<P>;
```