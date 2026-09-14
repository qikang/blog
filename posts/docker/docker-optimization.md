---
title: Docker镜像优化技巧
date: 2026-03-08
author: 管理员
tags: Docker,DevOps,优化
summary: 分享Docker镜像构建和运行时的优化技巧，帮助创建更小、更快的镜像。
---

# Docker镜像优化技巧

优化Docker镜像可以显著提升构建速度和运行效率。

## 多阶段构建

使用多阶段构建减小镜像体积：

```dockerfile
# 构建阶段
FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# 运行阶段
FROM alpine:3.22
COPY --from=builder /app .
CMD ["./app"]
```

## 减少层数

合并RUN指令：

```dockerfile
# 不推荐
RUN apt-get update
RUN apt-get install -y nginx
RUN rm -rf /var/lib/apt/lists/*

# 推荐
RUN apt-get update && \
    apt-get install -y nginx && \
    rm -rf /var/lib/apt/lists/*
```

## .dockerignore

排除不必要的文件：

```
node_modules
.git
*.md
.env
```

## 使用轻量基础镜像

```dockerfile
# 不推荐
FROM ubuntu:22.04

# 推荐
FROM alpine:3.22
# 或
FROM scratch
```

## 按需安装依赖

```dockerfile
# 只在需要时安装
FROM alpine
RUN apk add --no-cache python3 && \
    rm -rf /var/cache/apk/*
```

## 使用构建缓存

合理安排指令顺序以充分利用缓存：

```dockerfile
# 先复制依赖文件
COPY go.mod go.sum ./
RUN go mod download

# 再复制源代码
COPY . .
RUN go build .
```

## 总结

通过这些优化技巧，可以创建更小、更快、更安全的Docker镜像。
