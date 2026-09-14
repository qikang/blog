---
title: GitLab CI/CD流水线配置详解
date: 2026-03-08
author: 西康
tags: [CI/CD, DevOps, GitLab]
---

# GitLab CI/CD流水线配置详解

使用GitLab CI/CD实现自动化部署。

## .gitlab-ci.yml基础

```yaml
stages:
  - build
  - test
  - deploy

build:
  stage: build
  script:
    - npm install
    - npm run build

test:
  stage: test
  script:
    - npm test
```

## 多环境部署

```yaml
deploy:
  stage: deploy
  script:
    - ./deploy.sh
  environment:
    name: production
    url: https://example.com
  only:
    - main
```

## Docker镜像构建

```yaml
build-image:
  stage: build
  image: docker:latest
  script:
    - docker build -t myapp:$CI_COMMIT_SHA .
    - docker push myapp:$CI_COMMIT_SHA
```
