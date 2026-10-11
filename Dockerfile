# 仅支持 linux/amd64 平台
FROM --platform=linux/amd64 golang:1.27.3-alpine AS builder

# 国内网络拉依赖更稳(可按需删除)
ENV GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /src

# 先拷贝 go.mod/go.sum 单独拉依赖,利用层缓存
COPY go.mod go.sum ./
RUN go mod download

# 再拷贝源码
COPY . .

# 编译产物命名 blog(和 build-run.sh 一致)
RUN go build -trimpath -ldflags="-s -w" -o /out/blog .

# ============ 运行阶段 ============
FROM --platform=linux/amd64 alpine:3.20

# 添加 CA 证书和时区数据(HTTP 服务最小必要依赖)
RUN apk add --no-cache ca-certificates tzdata

# 拷贝构建产物
COPY --from=builder /out/blog /usr/local/bin/blog

# 暴露端口(与 main.go 的 PORT 默认值对齐)
EXPOSE 8083

# 工作目录:挂载 posts 用
WORKDIR /app

# 环境变量基线(可被 compose 的 environment 覆盖)
ENV PORT=":8083" \
    POSTS_DIR="./posts" \
    SITE_NAME="西康的博客" \
    PAGE_SIZE=5

ENTRYPOINT ["/usr/local/bin/blog"]
