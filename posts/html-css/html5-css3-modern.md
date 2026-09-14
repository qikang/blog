---
title: 现代HTML5与CSS3布局技术
date: 2026-03-07
author: 西康
tags: [HTML, CSS, 前端]
---

# 现代HTML5与CSS3布局技术

掌握Flexbox和Grid实现现代网页布局。

## Flexbox布局

```css
.container {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.item {
    flex: 1;
}
```

## CSS Grid

```css
.grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 20px;
}
```

## 响应式设计

```css
@media (max-width: 768px) {
    .grid {
        grid-template-columns: 1fr;
    }
}
```

## 新特性

- CSS变量
- calc()函数
- clip-path
- backdrop-filter
