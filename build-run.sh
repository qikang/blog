#!/bin/bash
set -ex

IMAGE_NAME="blog:latest"
PROJECT_DIR="$(cd "$(dirname "$0")" && pwd)"

docker system prune --all --force --volumes > /dev/null

docker compose down

cd "$PROJECT_DIR"

docker build -t $IMAGE_NAME .

cd "$PROJECT_DIR"
docker compose up -d

echo "服务地址: http://localhost:8083"

