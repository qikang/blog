---
title: Redis缓存实战指南
date: 2026-03-06
author: 管理员
tags: Redis,缓存,数据库
summary: 详细介绍Redis作为缓存的使用场景、数据结构和最佳实践。Redis是一个高性能的内存数据库，常用作缓存层。
---

# Redis缓存实战指南

Redis是一个高性能的内存数据库，常用作缓存层。

## Redis数据结构

### 字符串（String）

```bash
SET user:1 "Tom"
GET user:1
```

### 哈希（Hash）

```bash
HSET user:1 name "Tom" age "25"
HGET user:1 name
HGETALL user:1
```

### 列表（List）

```bash
LPUSH tasks "task1"
LPUSH tasks "task2"
LRANGE tasks 0 -1
```

### 集合（Set）

```bash
SADD tags "golang"
SADD tags "docker"
SADD tags "golang"
SMEMBERS tags
```

### 有序集合（ZSet）

```bash
ZADD leaderboard 100 "user1"
ZADD leaderboard 200 "user2"
ZREVRANGE leaderboard 0 -1 WITHSCORES
```

## 缓存策略

### Cache-Aside模式

```go
// 查询缓存
func GetUser(id int) User {
    // 先查缓存
    cached, err := redis.Get(fmt.Sprintf("user:%d", id))
    if err == nil {
        return unmarshal(cached)
    }

    // 缓存未命中，查询数据库
    user := db.GetUser(id)

    // 存入缓存
    redis.Set(fmt.Sprintf("user:%d", id), marshal(user), time.Hour)

    return user
}
```

### 缓存过期策略

```bash
# 设置过期时间
SET key value EX 3600  # 1小时后过期

# 惰性删除
```

## 应用场景

- 会话缓存
- 热点数据缓存
- 分布式锁
- 消息队列
- 限流控制

## 总结

Redis是现代应用架构中不可或缺的组件，合理使用可以大幅提升系统性能。
