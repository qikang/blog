---
title: MySQL事务与锁机制
date: 2026-03-11
author: 西康
tags: [MySQL, 数据库, 事务]
---

# MySQL事务与锁机制

事务保证数据库操作的ACID特性。

## 事务基本操作

```sql
START TRANSACTION;
UPDATE accounts SET balance = balance - 100 WHERE user_id = 1;
UPDATE accounts SET balance = balance + 100 WHERE user_id = 2;
COMMIT;
-- 或回滚
ROLLBACK;
```

## 隔离级别

- READ UNCOMMITTED
- READ COMMITTED
- REPEATABLE READ（默认）
- SERIALIZABLE

## 行锁与表锁

```sql
-- 行锁
SELECT * FROM users WHERE id = 1 FOR UPDATE;

-- 表锁
LOCK TABLES users WRITE;
```

## 死锁处理

```sql
-- 查看死锁
SHOW ENGINE INNODB STATUS;
```
