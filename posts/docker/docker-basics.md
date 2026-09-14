---
title: Docker入门指南
date: 2026-03-10
author: 管理员
tags: Docker,DevOps,容器
summary: Docker基础知识介绍，包括镜像、容器的基本操作和常用命令。
---

# Docker入门指南

Docker是一个开源的容器化平台，本文将介绍Docker的基础知识。

## 什么是Docker？

Docker是一个用于开发、部署和运行应用程序的容器化平台。容器允许开发者将应用程序及其依赖打包到一个轻量级、可移植的容器中。

## 基本概念

### 镜像（Image）

镜像是只读模板，用于创建容器。可以使用官方镜像或自定义镜像。

```bash
# 拉取镜像
docker pull nginx:latest

# 查看镜像列表
docker images

# 删除镜像
docker rmi nginx:latest
```

### 容器（Container）

容器是镜像的运行实例。

```bash
# 运行容器
docker run -d -p 8080:80 nginx:latest

# 查看运行中的容器
docker ps

# 查看所有容器
docker ps -a

# 停止容器
docker stop container_id

# 删除容器
docker rm container_id
```

## Dockerfile

Dockerfile用于构建自定义镜像：

```dockerfile
FROM golang:1.23

WORKDIR /app

COPY . .

RUN go build -o myapp .

EXPOSE 8080

CMD ["./myapp"]
```

## 数据卷

数据卷用于持久化容器数据：

```bash
# 创建数据卷
docker volume create mydata

# 挂载数据卷
docker run -v mydata:/data myimage
```

## 总结

Docker是现代DevOps的重要工具，掌握基础命令和概念是开发者的必备技能。
