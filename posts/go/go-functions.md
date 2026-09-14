---
title: Go语言函数详解
date: 2026-03-14
author: 管理员
tags: Go,后端,函数
summary: 详细介绍Go语言中函数的定义、使用方法以及高级特性。
---

# Go语言函数详解

函数是Go语言中组织代码的基本单元，本文将详细介绍函数的定义和使用。

## 函数定义

```go
func greet(name string) string {
    return "Hello, " + name
}
```

## 多返回值

Go语言支持函数返回多个值：

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

## 命名返回值

Go允许为返回值命名：

```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}
```

## 变长参数

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

## 函数作为值

Go语言中函数是一等公民，可以作为参数传递：

```go
func apply(fn func(int) int, n int) int {
    return fn(n)
}

result := apply(func(x int) int {
    return x * 2
}, 5)  // result = 10
```

## 闭包

闭包是一个函数及其引用环境的组合：

```go
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}
```

## 总结

函数是Go语言的核心概念，熟练掌握函数的各种用法对于编写高质量Go代码至关重要。
