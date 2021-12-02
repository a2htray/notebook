如何格式化 JSON
====================

[原问题](https://stackoverflow.com/questions/19038598/how-can-i-pretty-print-json-using-go)

### 问题

请问，有人知道在 Go 中有没有比较简单的 JSON 格式化方式？

### 回答

如果需要像下面格式的打印：

```json
{
    "data": 1234
}
```

最简单的方式是使用 `MarshalIndent`，该方法可以指定缩进的字符串。因此，使用 `json.MarshalIndent(data, "", "    ")` 就可以以 4 个
空格进行缩进。

### 总结