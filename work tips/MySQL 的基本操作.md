MySQL 的基本操作
======================

### 用户

#### 新增用户及权限

> requirement: 提前建好数据库

```sql
CREATE USER 'your_username'@'host_you_need' IDENTIFIED BY 'your_password';
GRANT ALL PRIVILEGES ON 'database_you_want'.* TO 'your_username'@'host_you_need';
```

授权格式：`GRANT ALL PRIVILEGES ON #1.#2 TO '#3'@'#4';`

