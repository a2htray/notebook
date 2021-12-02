对象字面量隐式含 any 类型的读取
===========================

### 问题

以下是我的部分代码：

```typescript
const myObj: object = {}
const propname = 'propname'

myObj[propname] = 'string'
```

得到如下的错：

```bash
ERROR in path/to/file.ts(4,1)
TS7053: Element implicitly has an 'any' type because expression of type '"propname"' can't be used to index type '{}'.
  Property 'propname' does not exist on type '{}'.
```

哪里错了，该如何解决？

### 回答

你不得不定义一个以字符串为 key 的对象：

```typescript
const myObj: {[index: string]:any} = {}
```

### 总结