如何在模板中遍历数组并使用下标
=================

[原地址](https://stackoverflow.com/questions/29762118/range-over-array-index-in-templates)

### 问题

我知道如何在 `range` 中使用下标：

```html
{{range $i, $e := .First}}$e - {{index $.Second $i}}{{end}}
```

那如果数组中包含数组，我又该如何遍历第二个数组呢

```go
type a struct {
   Title []string
   Article [][]string
}
IndexTmpl.ExecuteTemplate(w, "index.html", a)
```

`index.html`

```html
{{range $i, $a := .Title}}
  {{index $.Article $i}}  // Want to range over this.
{{end}}
```

### 回答

你可以使用内嵌的循环：

```html
package main

import (
    "html/template"
    "os"
)

type a struct {
    Title   []string
    Article [][]string
}

var data = &a{
    Title: []string{"One", "Two", "Three"},
    Article: [][]string{
        []string{"a", "b", "c"},
        []string{"d", "e"},
        []string{"f", "g", "h", "i"}},
}

var tmplSrc = `
{{range $i, $a := .Title}}
  Title: {{$a}}
  {{range $article := index $.Article $i}}
    Article: {{$article}}.
  {{end}}
{{end}}`

func main() {
    tmpl := template.Must(template.New("test").Parse(tmplSrc))
    tmpl.Execute(os.Stdout, data)
}
```

### 总结

这个问题的第 1 点在于，使用 range 访问到数组中的下标：

```html
{{range $i, $v := .array}}{{end}}
```

第 2 点在于，使用 index 方法得到数组中的元素。

第 3 点在于，如果元素还是一个数组，则仍然可通过 `{{range $innerV := index $array $i}}` 进行访问