---
title: Vue3 Composition API详解
date: 2026-03-12
author: 西康
tags: [Vue, 前端, Composition API]
---

# Vue3 Composition API详解

Vue3的Composition API提供了更灵活的组件逻辑组织方式。

## setup

```vue
<script setup>
import { ref, computed } from 'vue'

const count = ref(0)
const doubled = computed(() => count.value * 2)
</script>
```

## 响应式引用

```javascript
import { ref, reactive } from 'vue'

// 基础类型用ref
const count = ref(0)

// 对象用reactive
const state = reactive({
    name: '张三',
    age: 25
})
```

## 生命周期

```javascript
import { onMounted, onUnmounted } from 'vue'

onMounted(() => {
    console.log('组件挂载')
})

onUnmounted(() => {
    console.log('组件卸载')
})
```
