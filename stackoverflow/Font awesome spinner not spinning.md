FontAwesomeIcon 的 spinner 不 spin
============

[原问题](https://stackoverflow.com/questions/52225228/font-awesome-spinner-not-spinning/55944380)

### 问题

我已经完成了 `fontawesome` 的安装，[https://github.com/FortAwesome/vue-fontawesome](https://github.com/FortAwesome/vue-fontawesome)，需要渲染一个 `spinner`，如下：

```js
<font-awesome-icon :icon="['fas','spinner']" />
```

但渲染出的 `spinner` 是静止的，所以又做了如下修改：

```js
<font-awesome-icon :icon="['fas','spinner', 'spin']" />
```

但上述代码又报错了，`Could not find one or more icon(s) undefined`。

有人可以给出正确的方式吗？让我的 spinner 可以动起来。

一些额外的配置如下：

`nuxt.config.js`

```js
    modules: [
        'nuxt-fontawesome'
],


//font-awesome
  fontawesome: {
    imports: [
        {
          set: '@fortawesome/free-solid-svg-icons',
          icons: ['fas']
        },
    ],
  },

build: {
      config.resolve.alias['@fortawesome/fontawesome-free-brands$'] = '@fortawesome/fontawesome-free-brands/shakable.es.js'
      config.resolve.alias['@fortawesome/fontawesome-free-solid$'] = '@fortawesome/fontawesome-free-solid/shakable.es.js'
    }
```

`index.vue`

```js
<template>
  <div>
    <font-awesome-icon :icon="['fas','spinner','spin' ]" />
  </div>
</template>
```

### 回答

需要加个 `spin` 指令：

```js
<font-awesome-icon icon="spinner" spin />
```

### 总结

本人的情况则是基于 React+TypeScript 下，安装 `fontawesome` 如下：

```bash
npm install --save @fortawesome/fontawesome-svg-core @fortawesome/free-solid-svg-icons @fortawesome/react-fontawesome @types/react-fontawesome
```

正确的使用方法如下：

```js
import { faSpinner } from '@fortawesome/free-solid-svg-icons'
import {FontAwesomeIcon} from '@fortawesome/react-fontawesome'

export Example = () => {
    return <FontAwesomeIcon icon={faSpinner} spin/>
}
```
