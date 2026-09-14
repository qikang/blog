---
title: Docker Compose完全指南
date: 2026-03-09
author: 管理员
tags: Docker,DevOps,Compose
summary: 详细介绍Docker Compose的使用方法，用于定义和运行多容器Docker应用。
---

# Docker Compose完全指南

Docker Compose是一个用于定义和运行多容器Docker应用的工具。

## 基本结构

docker-compose.yml文件定义服务：

```yaml
version: '3.8'

services:
  web:
    image: nginx:latest
    ports:
      - "8080:80"
    depends_on:
      - api

  api:
    build: ./api
    volumes:
      - ./api:/app
    environment:
      - NODE_ENV=production

  db:
    image: postgres:15
    volumes:
      - db-data:/var/lib/postgresql/data
    environment:
      POSTGRES_PASSWORD: secret

volumes:
  db-data:
```

## 常用命令

```bash
# 启动所有服务
docker compose up

# 后台运行
docker compose up -d

# 停止服务
docker compose down

# 查看日志
docker compose logs -f

# 进入容器
docker compose exec web sh

# 构建镜像
docker compose build
```

## 环境变量

可以使用.env文件定义环境变量：

```bash
# .env
POSTGRES_PASSWORD=secret
NODE_ENV=development
```

## 网络配置

```yaml
services:
  web:
    networks:
      - frontend
      - backend

networks:
  frontend:
  backend:
```

## 总结

Docker Compose极大地简化了多容器应用的开发和部署，是现代微服务架构的重要工具。
