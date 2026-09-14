---
title: React Hooks完全指南
date: 2026-03-13
author: 西康
tags: [React, 前端, Hooks]
---

# React Hooks完全指南

React Hooks让函数组件也能拥有状态。

## useState

```jsx
import { useState } from 'react';

function Counter() {
    const [count, setCount] = useState(0);
    return <button onClick={() => setCount(count + 1)}>{count}</button>;
}
```

## useEffect

```jsx
import { useEffect } from 'react';

function DataFetcher() {
    const [data, setData] = useState(null);

    useEffect(() => {
        fetch('/api/data').then(res => res.json()).then(setData);
    }, []);

    return <div>{data}</div>;
}
```

## useContext

```jsx
const ThemeContext = React.createContext('light');

function App() {
    return (
        <ThemeContext.Provider value="dark">
            <Component />
        </ThemeContext.Provider>
    );
}
```
