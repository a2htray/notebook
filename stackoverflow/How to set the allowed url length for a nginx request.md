URI 太大，报 414,如何设置
====================

[原地址](https://stackoverflow.com/questions/1067334/how-to-set-the-allowed-url-length-for-a-nginx-request-error-code-414-uri-too)

### 问题

在使用 nginx 的过程中，URI 太大，报如下错：

```bash
error code 414: uri too large
```

有人知道如何设置 nginx 的最大 URI 允许长度吗？

### 回答

[看这里](http://nginx.org/en/docs/http/ngx_http_core_module.html#large_client_header_buffers)

```bash
Syntax:	large_client_header_buffers number size;
Default:	large_client_header_buffers 4 8k;
Context:	http, server
```

> Sets the maximum number and size of buffers used for reading large client request header. A request line cannot exceed the size of one buffer, or the 414 (Request-URI Too Large) error is returned to the client. A request header field cannot exceed the size of one buffer as well, or the 400 (Bad Request) error is returned to the client. Buffers are allocated only on demand. By default, the buffer size is equal to 8K bytes. If after the end of request processing a connection is transitioned into the keep-alive state, these buffers are released.

所以你只需要设置一个比你需要的大一点的值就行了。

### 总结

可以在 http 或 server 上下文中增加！