form 的 enctype='multipart/form-data' 是做什么的?
===========================

[原地址](https://stackoverflow.com/questions/4526273/what-does-enctype-multipart-form-data-mean)

### 问题

`enctype="multipart/form-data"` 是什么以及什么时候应该使用

### 回答

当你创建一个 POST 请求，需要对请求中的 Body 体进行编码

HTML 表单提供 3 种编码方式：

* application/x-www-form-urlencoded (默认)
* multipart/form-data
* text/plain

application/json 的编码方式也有实现，但现在已经被弃用了

（相对于表单提交的编码方式，HTTP 请求已经可以在其它方面应用多种不同的编码方式。JSON 是一种 Web 服务中常用的格式，一些也有用在 SOAP中）

大多数开发都不需要特别关心 3 种编码方式的区别，只需要记住下面重要的一点：

* 永远不用使用 text/plain

当你在编写客户端代码时：

* 当 input 的 type 值为 `file` 时，使用 multipart/form-data
* 其它情况下，可以使用 multipart/form-data 或 application/x-www-form-urlencoded，但 application/x-www-form-urlencoded 更为高效

当你在编写服务器端代码时：

* 使用现成的 form 处理库

大多数情况下，你不需要担心处理库的区别。同时，在服务器端处理 form 数据也不要嫌麻烦

有时，你会发现有些库不能同时处理这种编码方式。比如，`body-parser`，NodeJS 中处理表单数据的库，它不能处理 `multipart` 的请求（在文档中有其推荐的其它解决方案）

当你在开发一个解析或生成原始数据 form 库时，就需要关心数据的编码方式了。当然，你也许只是兴趣使然

multipart/form-data 编码方式虽然复杂，但它允许将整个文件作为数据进行发送

HTML 5 引入了 text/plain，但也仅仅在调试时有用

### 总结

* form 的 enctype 值有 3 个
  * application/x-www-form-urlencoded (默认)
  * multipart/form-data
  * text/plain
* application/x-www-form-urlencoded 与 multipart/form-data 区别
  * application/x-www-form-urlencoded 更高效
  * multipart/form-data 可带整个文件数据
* 在 input type 等于 file 时，要使用 multipart/form-data
