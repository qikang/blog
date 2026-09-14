---
title: DevOps监控与日志系统实践
date: 2026-03-02
author: 西康
tags: [DevOps, 监控, 日志]
---

# DevOps监控与日志系统实践

构建完整的可观测性系统。

## Prometheus监控

```yaml
scrape_configs:
  - job_name: 'node'
    static_configs:
      - targets: ['localhost:9100']
```

## Grafana可视化

Grafana提供丰富的仪表板模板。

## ELK日志系统

```yaml
services:
  elasticsearch:
    image: elasticsearch:8.0
  logstash:
    image: logstash:8.0
  kibana:
    image: kibana:8.0
```

## 告警配置

```yaml
groups:
  - name: alert.rules
    rules:
      - alert: HighMemory
        expr: node_memory_MemAvailable / node_memory_MemTotal * 100 < 10
        for: 5m
        labels:
          severity: warning
```
