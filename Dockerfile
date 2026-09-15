FROM swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/library/golang:1.23.8-alpine AS builder

# 设置代理
RUN go env -w GOPROXY=https://goproxy.cn,direct && \
    go env -w GOPRIVATE=github.com

WORKDIR /app

# 复制源代码
COPY . .

# 使用代理构建
RUN go env -w GOPROXY=https://goproxy.cn,direct && \
    go env -w GOPRIVATE=github.com && \
    CGO_ENABLED=0 GOOS=linux go build -mod=mod -o blog .

# 运行镜像
FROM swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/library/alpine:3.22.2

WORKDIR /app

# 安装 CA 证书
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache ca-certificates tzdata

# 复制构建产物和内容
COPY --from=builder /app/blog .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/data ./data

# 环境变量
ENV PORT=:8083
ENV SITE_NAME="技术博客"
ENV PAGE_SIZE=10

EXPOSE 8083

CMD ["./blog"]
