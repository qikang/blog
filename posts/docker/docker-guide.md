---
title: Docker 入门与实践
date: 2026-03-14
author: DevOps 工程师
tags: Docker,DevOps,容器化
summary: 全面介绍 Docker 的基本概念、安装配置和常用命令，帮助快速上手容器技术。
---

# Docker 入门与实践

Docker 是一个开源的容器化平台，可以帮助开发者打包、部署和运行应用程序。本文将介绍 Docker 的基础知识。

## 什么是 Docker？

Docker 是一个用于开发、交付和运行应用程序的开放平台。通过容器化技术，Docker 可以将应用程序及其依赖打包成轻量级、可移植的容器。

## 核心概念

### 镜像（Image）

镜像是一个只读模板，用于创建容器。可以把镜像理解为面向对象中的"类"。

### 容器（Container）

容器是镜像的运行实例。可以把容器理解为面向对象中的"对象"。

### 仓库（Repository）

仓库用于存储镜像，Docker Hub 是最大的公共仓库。

## 常用命令

### 镜像操作

```bash
# 搜索镜像
docker search ubuntu

# 拉取镜像
docker pull ubuntu:20.04

# 查看镜像列表
docker images

# 删除镜像
docker rmi ubuntu:20.04
```

### 容器操作

```bash
# 运行容器
docker run -d -p 8080:80 nginx

# 查看运行中的容器
docker ps

# 查看所有容器
docker ps -a

# 停止容器
docker stop container_id

# 删除容器
docker rm container_id
```

## Dockerfile 示例

```dockerfile
# 基于官方镜像
FROM node:18-alpine

# 设置工作目录
WORKDIR /app

# 复制文件
COPY package*.json ./

# 安装依赖
RUN npm install

# 复制源代码
COPY . .

# 暴露端口
EXPOSE 3000

# 启动命令
CMD ["npm", "start"]
```

## Docker Compose

Docker Compose 用于定义和运行多容器应用：

```yaml
version: '3.8'
services:
  web:
    build: .
    ports:
      - "8080:3000"
  db:
    image: postgres:14
    environment:
      POSTGRES_PASSWORD: secret
```

## 总结

Docker 大大简化了应用的部署和运维工作，是现代 DevOps 不可或缺的技术之一。掌握 Docker 将显著提升开发效率。
