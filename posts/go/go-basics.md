---
title: Go语言基础语法详解
date: 2026-03-15
author: 管理员
tags: Go,后端,编程
summary: 详细介绍Go语言的基础语法，包括变量声明、数据类型、控制结构等核心概念。
---

# Go语言基础语法详解

Go语言是一门简洁而强大的编程语言，本文将详细介绍其基础语法。

## 变量声明

Go语言有多种变量声明方式：

```go
// 方式1：var 关键字
var name string = "Tom"

// 方式2：类型推断
var age = 25

// 方式3：简短声明
city := "Beijing"
```

## 基本数据类型

Go语言包含以下基本数据类型：

- **整数**: int, int8, int16, int32, int64
- **浮点数**: float32, float64
- **复数**: complex64, complex128
- **字符串**: string
- **布尔值**: bool

## 控制结构

### if语句

```go
if age >= 18 {
    fmt.Println("成年人")
} else {
    fmt.Println("未成年人")
}
```

### for循环

```go
// 经典for循环
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// range遍历
for index, value := range slice {
    fmt.Println(index, value)
}
```

### switch语句

```go
switch day {
case "Monday":
    fmt.Println("周一")
case "Tuesday":
    fmt.Println("周二")
default:
    fmt.Println("其他")
}
```

## 总结

掌握这些基础语法后，你就可以开始编写简单的Go程序了。下一篇文章我们将介绍Go语言的函数。
