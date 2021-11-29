React 的 JSX 中 pre 标签格式化代码问题
=========================

[原问题](https://stackoverflow.com/questions/37847885/formatting-code-with-pre-tag-in-react-and-jsx)

### 问题

我在 JSX 中使用 pre 标签，但 pre 标签的格式化效果交没有出来。为了达到目的，我只能通过下面的代码进行实现

```jsx harmony
const someCodeIWantToFormat = "var foo = 1"
const preBlock = { __html: "<pre>" + pythonCode + "</pre>" };
return(
  <div dangerouslySetInnerHTML={ preBlock } />;
)
```

### 回答

可以使用模板字面量的方法，进行换行、留白的效果

```jsx harmony
class SomeComponent extends React.Component {
   render() {
        return (
          <pre>{`
            Hello   ,   
            World   .
          `}</pre>
        )
    }
}
```

### 总结

使用模板字面量
