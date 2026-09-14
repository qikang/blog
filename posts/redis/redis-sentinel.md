---
title: Redis集群与哨兵模式配置
date: 2026-03-04
author: 西康
tags: [Redis, 数据库, 集群]
---

# Redis集群与哨兵模式配置

构建高可用的Redis架构。

## 哨兵模式配置

```conf
sentinel monitor mymaster 127.0.0.1 6379 2
sentinel down-after-milliseconds mymaster 5000
sentinel failover-timeout mymaster 60000
```

## 集群模式

```bash
redis-cli --cluster create 127.0.0.1:7001 127.0.0.1:7002 127.0.0.1:7003 --cluster-replicas 1
```

## Python客户端

```python
import redis

r = redis.Redis(host='localhost', port=6379, decode_responses=True)
r.set('key', 'value')
print(r.get('key'))
```

## 应用场景

- 会话存储
- 缓存层
- 消息队列
- 实时排行榜
