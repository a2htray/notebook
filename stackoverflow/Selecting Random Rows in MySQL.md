MySQL 随机选择行
==================

[原地址](https://stackoverflow.com/questions/1283640/selecting-random-rows-in-mysql)

### 问题

我在开发一个问答网站，该网站的数据库存着所有的问题。问题分成不同的类型，有数学、科学、历史等。所有的问题都在一个表中。

表结构如下：

```bash
questions ( qno(int)  ,type(int), question, .... ,... )
```

`qno` 是主键，`type` 是类型。

``bash
if type = 1 (math)
  type=2 (science)
``

现在我想针对每一种类型，随机选择不同的问题。比如，我想随机选择 20 个关于数学的问题。

在 MySQL 里有什么方式可以实现？

### 回答

你可以使用 `rand` 函数。使用该函数进行排序，再限制选择的行数。

```sql
select * from table order by rand() limit 10
```

如果你想针对数学的问题，如下：

```sql
select * from table where type = 1 order by rand() limit 10
```

### 总结

`rand` 进行 ORDER BY，再加个 LIMIT