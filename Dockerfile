FROM swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/library/golang:1.23.8-alpine AS builder

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 安装 git 与 CA 证书:Go module 在校验/解析依赖时需要 git,https 代理依赖 CA 证书
RUN apk add --no-cache git ca-certificates

WORKDIR /app

# 复制源代码
COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o blog .

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


# 环境变量
ENV PORT=:8083
ENV SITE_NAME="技术博客"
ENV PAGE_SIZE=10

EXPOSE 8083

CMD ["./blog"]
