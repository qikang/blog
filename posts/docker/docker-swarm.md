---
title: Docker Swarm容器编排实战
date: 2026-03-03
author: 西康
tags: [Docker, 容器, 编排]
---

# Docker Swarm容器编排实战

使用Docker Swarm管理容器集群。

## 初始化集群

```bash
docker swarm init --advertise-addr 192.168.1.100
```

## 添加节点

```bash
docker swarm join --token SWMTKN-xxx 192.168.1.100:2377
```

## 部署服务

```bash
docker service create \
    --name my-web \
    --replicas 3 \
    -p 80:80 \
    nginx:latest
```

## 滚动更新

```bash
docker service update \
    --image nginx:latest \
    my-web
```
