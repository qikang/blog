---
title: Go语言并发编程
date: 2026-03-11
author: 管理员
tags: Go,后端,并发
summary: 详细介绍Go语言的并发编程模型，包括goroutine、channel和同步机制。
---

# Go语言并发编程

Go语言原生支持并发编程，这是其最强大的特性之一。

## Goroutine

goroutine是由Go运行时管理的轻量级线程：

```go
// 启动一个goroutine
go func() {
    fmt.Println("Hello from goroutine")
}()

// 调用函数
go doSomething()
```

## Channel

channel是goroutine之间通信的管道：

```go
// 创建channel
ch := make(chan int)

// 发送数据
ch <- 10

// 接收数据
value := <-ch
```

### 带缓冲的channel

```go
ch := make(chan int, 10)  // 缓冲大小为10
```

### 关闭channel

```go
close(ch)
```

## Select

select用于监听多个channel：

```go
select {
case msg := <-ch1:
    fmt.Println("Received:", msg)
case msg := <-ch2:
    fmt.Println("Received:", msg)
case <-time.After(time.Second):
    fmt.Println("Timeout")
}
```

## 同步机制

### WaitGroup

```go
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(n int) {
        defer wg.Done()
        fmt.Println("Task", n)
    }(i)
}

wg.Wait()  // 等待所有goroutine完成
```

### Mutex

```go
var mu sync.Mutex
var counter int

func increment() {
    mu.Lock()
    defer mu.Unlock()
    counter++
}
```

## 总结

Go的并发模型简单而强大，合理使用goroutine和channel可以编写高效并发程序。
