---
title: Go语言Gin框架实战
date: 2026-03-05
author: 西康
tags: [Go, Gin, 后端]
---

# Go语言Gin框架实战

Gin是Go语言中最流行的Web框架。

## 快速开始

```go
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })
    r.Run()
}
```

## 路由组

```go
v1 := r.Group("/api/v1")
{
    v1.GET("/users", getUsers)
    v1.POST("/users", createUser)
}
```

## 中间件

```go
func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 处理请求前
        c.Next()
        // 处理请求后
    }
}
```
