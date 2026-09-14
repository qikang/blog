---
title: Go语言接口详解
date: 2026-03-12
author: 管理员
tags: Go,后端,接口
summary: 详细介绍Go语言中接口的定义、实现以及多态特性。
---

# Go语言接口详解

接口是Go语言实现多态的关键特性。

## 接口定义

```go
type Writer interface {
    Write([]byte) (int, error)
}

type Reader interface {
    Read([]byte) (int, error)
}
```

## 接口实现

Go语言使用隐式实现，不需要显式声明：

```go
type File struct {
    name string
}

func (f *File) Write(data []byte) (int, error) {
    // 实现Write方法
    return len(data), nil
}

// File自动实现Writer接口
```

## 空接口

空接口interface{}可以存储任意类型：

```go
func printValue(v interface{}) {
    fmt.Println(v)
}

printValue(42)        // int
printValue("hello")  // string
printValue(3.14)     // float64
```

## 类型断言

```go
var i interface{} = "hello"

s, ok := i.(string)
if ok {
    fmt.Println(s)
}
```

## 类型switch

```go
func printType(v interface{}) {
    switch v.(type) {
    case int:
        fmt.Println("整数")
    case string:
        fmt.Println("字符串")
    case bool:
        fmt.Println("布尔值")
    default:
        fmt.Println("未知类型")
    }
}
```

## 接口组合

```go
type ReadWriter interface {
    Reader
    Writer
}
```

## 总结

接口是Go语言的核心特性，熟练使用接口可以写出更加灵活和可维护的代码。
