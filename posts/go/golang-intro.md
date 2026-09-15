---
title: Go 语言入门指南
date: 2023-01-15
author: Go 爱好者
tags: Go,后端,编程语言
summary: 详细介绍 Go 语言的基础语法和核心概念，适合初学者入门学习。
---

# Go 语言入门指南

Go 语言（又称 Golang）是 Google 开发的一种静态类型、编译型语言。本文将介绍 Go 语言的基础知识。

## 为什么选择 Go？

Go 语言具有以下优点：

- **简洁高效** - 语法简洁，编译速度快
- **并发支持** - 原生支持 goroutine 和 channel
- **标准库丰富** - 强大的标准库
- **部署简单** - 编译成单一可执行文件

## 基础语法

### 变量声明

```go
// 方式一：var 声明
var name string = "Tom"

// 方式二：简短声明
name := "Tom"

// 常量
const PI = 3.14159
```

### 函数

```go
func add(a, b int) int {
    return a + b
}

// 多返回值
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

### 并发

```go
// 启动 goroutine
go func() {
    fmt.Println("Hello from goroutine")
}()

// channel
ch := make(chan int)
ch <- 1    // 发送
value := <-ch  // 接收
```

## 常用标准库

| 包名 | 用途 |
|------|------|
| fmt | 格式化 I/O |
| net/http | HTTP 服务器 |
| encoding/json | JSON 解析 |
| log | 日志记录 |

## 总结

Go 语言是一门非常适合构建高性能服务端应用的编程语言。它的学习曲线平缓，性能优秀，值得推荐！

下一篇文章我们将介绍如何使用 Go 构建 RESTful API。
