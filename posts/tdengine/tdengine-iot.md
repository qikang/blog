---
title: TDengine在物联网数据处理中的应用
date: 2026-03-09
author: 西康
tags: [TDengine, 物联网, 时序数据库]
---

# TDengine在物联网数据处理中的应用

TDengine是专为物联网设计的高性能时序数据库。

## 创建数据库

```sql
CREATE DATABASE demo;
USE demo;
```

## 创建超级表

```sql
CREATE STABLE sensors (
    ts TIMESTAMP,
    temperature FLOAT,
    humidity FLOAT,
    pressure FLOAT
) TAGS (location BINARY(50), device_id BINARY(50));
```

## 插入数据

```sql
INSERT INTO sensor_001 VALUES
    (now, 25.5, 60.0, 101.3)
    (now + 1s, 26.0, 59.5, 101.2);
```

## 查询

```sql
SELECT AVG(temperature) FROM sensors
WHERE location='Beijing'
AND ts > NOW - 1h
INTERVAL(1h);
```
