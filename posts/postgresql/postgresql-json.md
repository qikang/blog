---
title: PostgreSQL JSON类型高级应用
date: 2026-03-10
author: 西康
tags: [PostgreSQL, 数据库, JSON]
---

# PostgreSQL JSON类型高级应用

PostgreSQL强大的JSON支持使其成为处理半结构化数据的理想选择。

## JSON类型

```sql
CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    data JSONB
);

INSERT INTO events (data) VALUES
    ('{"type": "click", "page": "/home"}');
```

## 查询JSON

```sql
-- 提取字段
SELECT data->>'type' FROM events;

-- 使用JSONB路径
SELECT data#>>'{user,name}' FROM events;

-- 索引查询
SELECT * FROM events WHERE data @> '{"type":"click"}';
```

## 聚合操作

```sql
SELECT
    data->>'type' as event_type,
    COUNT(*)
FROM events
GROUP BY data->>'type';
```
