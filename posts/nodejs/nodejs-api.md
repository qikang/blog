---
title: Node.js RESTful API开发指南
date: 2026-03-14
author: 西康
tags: [Node.js, API, 后端]
---

# Node.js RESTful API开发指南

使用Express框架快速构建RESTful API。

## 初始化项目

```bash
npm init -y
npm install express
```

## 创建服务器

```javascript
const express = require('express');
const app = express();

app.get('/api/users', (req, res) => {
    res.json([{id: 1, name: '张三'}]);
});

app.listen(3000);
```

## 中间件

```javascript
app.use(express.json());
app.use(authMiddleware);
```

## 路由参数

```javascript
app.get('/api/users/:id', (req, res) => {
    const user = getUserById(req.params.id);
    res.json(user);
});
```
