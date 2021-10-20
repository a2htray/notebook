Typescript: 如何导出一个变量
======================

[原地址](https://stackoverflow.com/questions/42540745/typescript-how-to-export-a-variable)

### 问题

TS 环境下，`file1.ts` 内容：

```js
export var arr = [1, 2, 3];
```

然后打开另一个文件 `file2.ts`，希望以如下的方式访问 `file1.ts` 中的 `arr`：

```js
import {arr} from './file1';
```

但是，访问 `arr` 不是想要的变量值，而不得不使用 `arr.arr` 的方式得到数组的值。我应该如何通过导出的变量名访问数组值。

### 回答

有两种不同的导出方式：named 和 default。

你可以有多个 named 类型的导出，但只能有一个 default 类型的导出。

```js
// ./file1.ts
const arr = [1,2,3];
export { arr };
```

```js
// ./file2.ts
import { arr } from "./file1";
console.log(arr.length);
```

### 总结

实际的代码中，涉及了 3 个文件：

```bash
routers/
  index.ts
  nav.ts
App.tsx
```

在 `nav.ts` 中写与 Nav 相关的路由，同时 `index.ts` 作为统一导出的出口。

错误的示例如下：

```js
// nav.ts
import {LinkType} from '../types/link';

const navRouters = [
  {
    name: 'Home',
    text: 'Home',
    type: LinkType.Link,
    to: '/',
  }
];

export default navRouters
```

```js
// nav.ts
import navRouters from "./nav";

export default {
  navRouters
}
```

```js
// App.tsx
import {navRouters} from './routers';
```

问题在于 `index.ts` 中的导出多了一个 `default`。

修改：去除 `default`。

```js
// nav.ts
import navRouters from "./nav";

export {
  navRouters
}
```


