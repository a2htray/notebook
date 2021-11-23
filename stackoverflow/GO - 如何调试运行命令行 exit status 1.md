如何调试运行命令行 exit status 1
================================


[原问题](https://stackoverflow.com/questions/18159704/how-to-debug-exit-status-1-error-when-running-exec-command-in-golang/18159705)


### 问题

当我运行下面的代码时：

```go
cmd := exec.Command("find", "/", "-maxdepth", "1", "-exec", "wc", "-c", "{}", "\\")
var out bytes.Buffer
cmd.Stdout = &out
err := cmd.Run()
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println("Result: " + out.String())
```

我得到了如下的这个错误

```bash
exit status 1
```

然而，这对于 debug 没有帮助。我应该如何获得更详细的信息。

### 回答

解决方案是使用 `Command` 对象的 `Stderr` 属性，你可以这么做：

```go
cmd := exec.Command("find", "/", "-maxdepth", "1", "-exec", "wc", "-c", "{}", "\\")
var out bytes.Buffer
var stderr bytes.Buffer
cmd.Stdout = &out
cmd.Stderr = &stderr
err := cmd.Run()
if err != nil {
    fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
    return
}
fmt.Println("Result: " + out.String())
```

运行上述的代码，你会得到下面的错误信息：

```bash
exit status 1: find: -exec: no terminating ";" or "+"
```

### 总结

使用 `Stderr` 属性。

有时命令是把错误直接输出在 `标准输出`，需要区别对待。