---
title: Go语言结构体与方法
date: 2026-03-13
author: 管理员
tags: Go,后端,结构体
summary: 详细介绍Go语言中结构体的定义、方法绑定以及面向对象编程特性。
---

# Go语言结构体与方法

Go语言没有类，但可以通过结构体和方法来实现面向对象编程。

## 定义结构体

```go
type Person struct {
    Name string
    Age  int
    City string
}
```

## 创建结构体实例

```go
// 方式1：字面量
p1 := Person{Name: "Tom", Age: 25, City: "Beijing"}

// 方式2：使用new
p2 := new(Person)
p2.Name = "Jerry"

// 方式3：指针
p3 := &Person{Name: "Bob", Age: 30}
```

## 方法定义

Go语言中的方法是绑定到特定类型的函数：

```go
// 值接收者方法
func (p Person) SayHello() {
    fmt.Printf("Hello, I'm %s\n", p.Name)
}

// 指针接收者方法
func (p *Person) SetAge(age int) {
    p.Age = age
}
```

## 结构体嵌套

Go语言通过嵌套实现组合：

```go
type Address struct {
    City    string
    Country string
}

type Person struct {
    Name    string
    Address Address  // 嵌套结构体
}
```

## 匿名结构体

```go
person := struct {
    Name string
    Age  int
}{
    Name: "Alice",
    Age:  28,
}
```

## 总结

结构体和方法是Go语言实现面向对象编程的基础，掌握这些概念能够帮助你写出更好的Go代码。
