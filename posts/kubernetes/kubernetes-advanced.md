---
title: Kubernetes高级特性与最佳实践
date: 2026-03-06
author: 西康
tags: [Kubernetes, DevOps, 容器]
---

# Kubernetes高级特性与最佳实践

深入理解K8s的核心概念和最佳实践。

## ResourceQuota

```yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: my-quota
spec:
  hard:
    requests.cpu: "2"
    requests.memory: 4Gi
    pods: "10"
```

## HPA自动扩缩容

```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: my-app-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: my-app
  minReplicas: 2
  maxReplicas: 10
  metrics:
    - type: Resource
      resource:
        name: cpu
        target:
          type: Utilization
          averageUtilization: 70
```

## 最佳实践

- 使用命名空间隔离资源
- 设置资源限制
- 使用ConfigMap和Secret
- 实施健康检查
