如何修改 max_allowed_packet 的大小
=========================

[原问题](https://stackoverflow.com/questions/8062496/how-to-change-max-allowed-packet-size)

### 问题

在使用 mysql 的过程种，我需要处理 blob 字段。但当上传文件大于 1M 时，报了一个错：`Packets larger than max_allowed_packet are not allowed`。

以下是我尝试过的方法：

在 mysql 终端，执行过 `show variables like 'max_allowed_packet'`，显示的是 `1048576`。

然后又执行 `set global max_allowed_packet=33554432`，接着执行 `show variables like 'max_allowed_packet'`，得到了我想要的 `33554432`。

但当我重启时，该值又回到了 `1048576`。哪里错了吗？

### 回答

修改 my.ini 或 my.cnf 文件，在 `[mysqld]` 或 `[client]` 下加入：

`max_allowed_packet=500M`

然后重启。

### 总结

无
