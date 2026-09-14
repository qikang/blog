#!/bin/bash

set -e

# 禁用 BuildKit，使用传统 builder（解决私有仓库 TLS 证书问题）
export DOCKER_BUILDKIT=0

IMAGE_NAME="blog:latest"
PROJECT_DIR="$(cd "$(dirname "$0")" && pwd)"

echo "=== 停止服务 ==="
docker compose down

echo "=== 开始构建镜像 ==="
cd "$PROJECT_DIR"

docker build -t $IMAGE_NAME -f blog-backend/Dockerfile blog-backend/

echo "=== 启动服务 ==="
cd "$PROJECT_DIR"
docker compose up -d

echo "=== 完成! ==="
echo "服务地址: http://localhost:8083"
