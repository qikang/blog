---
title: k8s 安装
date: 2025-10-10
author: 管理员
tags: k8s,容器
summary: 本文主要介绍k8s的安装部署，详细结交安装过程中需要注意的点。文档详细介绍在 Ubuntu 24.04 LTS 系统上安装 Kubernetes 1.32.10 集群的完整步骤。Kubernetes（简称 k8s）是一个开源的容器编排平台，用于自动化容器化应用的部署、扩展和管理。
---
# Kubernetes 1.32.10 安装指南 (Ubuntu 24.04 LTS)

## 文档概述

本文档详细介绍在 Ubuntu 24.04 LTS 系统上安装 Kubernetes 1.32.10 集群的完整步骤。Kubernetes（简称 k8s）是一个开源的容器编排平台，用于自动化容器化应用的部署、扩展和管理。

---

## 目录

1. [环境准备](#1-环境准备)
2. [系统配置](#2-系统配置)
3. [安装容器运行时](#3-安装容器运行时)
4. [安装 Kubernetes 组件](#4-安装-kubernetes-组件)
5. [初始化集群](#5-初始化集群)
6. [添加工作节点](#6-添加工作节点)
7. [验证集群状态](#7-验证集群状态)
8. [部署示例应用](#8-部署示例应用)

---

## 1. 环境准备

### 1.1 硬件要求

| 角色 | CPU | 内存 | 磁盘 |
|------|-----|------|------|
| Master 节点 | 2 核+ | 2GB+ | 20GB+ |
| Worker 节点 | 1 核+ | 1GB+ | 20GB+ |

### 1.2 网络要求

- 所有节点之间网络互通
- 关闭防火墙或开放必要端口
- 支持 Internet 访问（用于拉取镜像）

### 1.3 节点规划

本示例使用 1 个 Master 节点和 2 个 Worker 节点：

| 主机名 | IP 地址 | 角色 |
|---------|----------|------|
| k8s-master | 192.168.1.10 | Master |
| k8s-worker1 | 192.168.1.11 | Worker |
| k8s-worker2 | 192.168.1.12 | Worker |

### 1.4 更新系统

首先更新系统软件包到最新版本：

```bash
# 更新软件包列表
sudo apt update

# 升级已安装的软件包
sudo apt upgrade -y

# 重启系统确保内核更新生效
sudo reboot
```

**命令解释**：
- `apt update`：从配置的软件源下载最新的软件包列表信息
- `apt upgrade`：根据更新后的列表，升级所有可升级的软件包
- `-y`：自动确认所有提示，避免交互式询问
- `reboot`：重启系统以加载新的内核

---

## 2. 系统配置

### 2.1 设置主机名

在每个节点上设置对应的主机名：

```bash
# 在 Master 节点执行
sudo hostnamectl set-hostname k8s-master

# 在 Worker1 节点执行
sudo hostnamectl set-hostname k8s-worker1

# 在 Worker2 节点执行
sudo hostnamectl set-hostname k8s-worker2
```

**命令解释**：
- `hostnamectl`：用于查询和更改系统主机名
- `set-hostname`：设置静态主机名

### 2.2 配置 hosts 文件

在所有节点上编辑 hosts 文件，添加节点 IP 和主机名映射：

```bash
sudo vim /etc/hosts
```

添加以下内容：

```
192.168.1.10 k8s-master
192.168.1.11 k8s-worker1
192.168.1.12 k8s-worker2
```

**命令解释**：
- `vim`：文本编辑器，用于编辑文件
- `/etc/hosts`：本地 DNS 解析文件，将主机名解析为 IP 地址

### 2.3 关闭 Swap 分区

Kubernetes 要求关闭 Swap 分区，否则会影响调度：

```bash
# 立即关闭 Swap
sudo swapoff -a

# 永久关闭 Swap（编辑 fstab 文件）
sudo sed -i '/swap/s/^\/dev/#\/dev/' /etc/fstab
```

**命令解释**：
- `swapoff -a`：立即关闭所有 Swap 分区
- `sed -i`：使用 sed 进行文本替换，将 swap 行的开头注释掉
- `/etc/fstab`：文件系统表，定义分区挂载点

### 2.4 加载内核模块

加载 Kubernetes 所需的内核模块：

```bash
# 加载 required 模块
sudo modprobe overlay
sudo modprobe br_netfilter

# 配置模块开机自加载
sudo tee /etc/modules-load.k8s.conf <<EOF
overlay
br_netfilter
EOF
```

**命令解释**：
- `modprobe`：加载内核模块
- `overlay`：容器叠加文件系统，用于 Docker 的存储驱动
- `br_netfilter`：网桥过滤模块，用于 iptables 规则
- `/etc/modules-load.k8s.conf`：系统启动时自动加载的模块列表

### 2.5 配置网络参数

配置 Kubernetes 所需的网络参数：

```bash
# 创建 sysctl 配置文件
sudo tee /etc/sysctl.d/k8s.conf <<EOF
net.bridge.bridge-nf-call-iptables  = 1
net.bridge.bridge-nf-call-ip6tables = 1
net.ipv4.ip_forward                 = 1
EOF

# 重新加载 sysctl 配置
sudo sysctl --system
```

**命令解释**：
- `net.bridge.bridge-nf-call-iptables`：让 iptables 处理桥接网络流量
- `net.ipv4.ip_forward`：启用 IP 转发，允许数据包路由
- `sysctl --system`：重新加载所有 sysctl 配置

### 2.6 关闭防火墙（可选）

如果测试环境可以关闭防火墙：

```bash
sudo ufw disable
```

**或者开放必要端口**：

```bash
# Master 节点需要开放的端口
sudo ufw allow 6443  # Kubernetes API Server
sudo ufw allow 2379 # etcd client
sudo ufw allow 2380 # etcd peer
sudo ufw allow 10250 # Kubelet API
sudo ufw allow 10259 # kube-scheduler
sudo ufw allow 10257 # kube-controller-manager

# 重新加载防火墙
sudo ufw reload
```

**端口说明**：
- 6443：Kubernetes API Server 监听端口
- 2379/2380：etcd 数据库通信端口
- 10250：Kubelet API 端口
- 10259/10257：Kubernetes 控制平面组件端口

---

## 3. 安装容器运行时

Kubernetes 1.27+ 推荐使用 containerd 作为容器运行时。

### 3.1 安装 containerd

```bash
# 安装 containerd
sudo apt install -y containerd

# 创建 containerd 配置目录
sudo mkdir -p /etc/containerd
```

**命令解释**：
- `containerd`：Docker 容器运行时的核心组件

### 3.2 配置 containerd

生成默认配置文件并进行修改：

```bash
# 生成默认配置
sudo containerd config default > /etc/containerd/config.toml
```

编辑配置文件，修改 SystemdCgroup 为 true：

```bash
sudo vim /etc/containerd/config.toml
```

找到以下部分：

```toml
[plugins."io.containerd.grpc.v1.cri".containerd.runtimes]
  [plugins."io.containerd.grpc.v1.cri.containerd.runtimes.runc]
    ...
    [plugins."io.containerd.grpc.v1.cri.containerd.runtimes.runc.options]
      SystemdCgroup = false
```

修改为：

```toml
[plugins."io.containerd.grpc.v1.cri".containerd.runtimes]
  [plugins."io.containerd.grpc.v1.cri.containerd.runtimes.runc]
    ...
    [plugins."io.containerd.grpc.v1.cri.containerd.runtimes.runc.options]
      SystemdCgroup = true
```

**配置解释**：
- `SystemdCgroup`：使用 systemd cgroup 管理容器资源，实现更精细的资源控制

### 3.3 启动 containerd

```bash
# 启动 containerd
sudo systemctl start containerd

# 设置开机自启动
sudo systemctl enable containerd

# 检查运行状态
sudo systemctl status containerd
```

**命令解释**：
- `systemctl start`：启动服务
- `systemctl enable`：设置服务开机自启动
- `systemctl status`：查看服务运行状态

---

## 4. 安装 Kubernetes 组件

### 4.1 添加 Kubernetes APT 源

```bash
# 安装依赖包
sudo apt install -y apt-transport-https ca-certificates curl gnupg

# 添加 Kubernetes GPG 密钥
curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.32/deb/Release.key | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg

# 添加 Kubernetes APT 源
echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.32/deb/ /' | sudo tee /etc/apt/sources.list.d/kubernetes.list
```

**命令解释**：
- `apt-transport-https`：支持通过 HTTPS 协议访问 APT 源
- `ca-certificates`：CA 证书，用于验证 HTTPS 连接
- `gnupg`：GNU 隐私工具，用于处理 GPG 密钥
- `curl -fsSL`：下载文件（-f 失败时返回错误，-s 静默模式，-L 跟随重定向）

### 4.2 安装 Kubernetes 组件

```bash
# 更新 APT 源
sudo apt update

# 安装 kubelet、kubeadm、kubectl
sudo apt install -y kubelet=1.32.10-1.1 kubeadm=1.32.10-1.1 kubectl=1.32.10-1.1

# 锁定版本，防止意外升级
sudo apt-mark hold kubelet kubeadm kubectl
```

**组件说明**：
- `kubelet`：运行在每个节点上的代理，负责管理容器生命周期
- `kubeadm`：集群初始化工具，用于快速部署 Kubernetes 集群
- `kubectl`：命令行工具，用于与集群交互

**版本说明**：
- 如果 1.32.10 版本不可用，请使用可用版本：
  ```bash
  # 查看可用版本
  apt-cache madison kubeadm
  ```

---

## 5. 初始化集群

### 5.1 在 Master 节点初始化

```bash
sudo kubeadm init \
  --pod-network-cidr=10.244.0.0/16 \
  --service-cidr=10.96.0.0/12 \
  --kubernetes-version=1.32.10
```

**参数说明**：
- `--pod-network-cidr`：Pod 网络使用的 IP 段（Flannel 网络插件使用）
- `--service-cidr`：Service 网络使用的 IP 段
- `--kubernetes-version`：指定 Kubernetes 版本

**初始化输出**：
初始化成功后，会看到类似以下的输出：

```
Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 192.168.1.10:6443 --token xxxxxx \
    --discovery-token-ca-cert-hash sha256:xxxxxx
```

### 5.2 配置 kubectl

```bash
# 创建 .kube 目录
mkdir -p $HOME/.kube

# 复制配置文件
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config

# 设置权限
sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

**命令解释**：
- `admin.conf`：包含管理员凭证的配置文件
- `.kube/config`：kubectl 默认读取的配置文件路径

### 5.3 安装网络插件

Kubernetes 需要网络插件来实现 Pod 网络功能。这里使用 Flannel：

```bash
# 下载并安装 Flannel
kubectl apply -f https://raw.githubusercontent.com/flannel-io/flannel/master/Documentation/kube-flannel.yml
```

**网络插件说明**：
- Flannel 为每个 Pod 分配唯一的 IP，实现跨主机容器通信
- 10.244.0.0/16 是 Flannel 使用的默认 Pod 网络

如果下载失败，可以手动创建配置文件：

```yaml
# kube-flannel.yml
apiVersion: v1
kind: Namespace
metadata:
  name: flannel
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: flannel
  namespace: flannel
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: flannel
rules:
  - apiGroups: [""]
    resources: ["pods"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["nodes"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["configmaps"]
    verbs: ["get", "list", "watch", "create", "update", "patch"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: flannel
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: flannel
subjects:
  - kind: ServiceAccount
    name: flannel
    namespace: flannel
---
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: flannel
  namespace: flannel
spec:
  podSelector: {}
  policyTypes:
    - Ingress
    - Egress
---
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: kube-flannel-ds
  namespace: flannel
  labels:
    app: flannel
spec:
  selector:
    matchLabels:
      app: flannel
  template:
    metadata:
      labels:
        app: flannel
    spec:
      serviceAccountName: flannel
      hostNetwork: true
      containers:
        - name: kube-flannel
          image: docker.io/flannel/flannel:v0.25.1
          command:
            - /opt/bin/flanneld
          args:
            - --ip-masq
            - --kube-subnet-mgr
          securityContext:
            privileged: true
          env:
            - name: POD_NAME
              valueFrom:
                fieldRef:
                  fieldPath: metadata.name
            - name: POD_NAMESPACE
              valueFrom:
                fieldRef:
                  fieldPath: metadata.namespace
          volumeMounts:
            - name: run
              mountPath: /run/flannel
            - name: cni
              mountPath: /etc/cni/net.d
            - name: flannel-cfg
              mountPath: /etc/kube-flannel/
      volumes:
        - name: run
          hostPath:
            path: /run/flannel
        - name: cni
          hostPath:
            path: /etc/cni/net.d
        - name: flannel-cfg
          configMap:
            name: kube-flannel-cfg
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: kube-flannel-cfg
  namespace: flannel
  labels:
    app: flannel
data:
  cni-conf.json: |
    {
      "name": "cbr0",
      "cniVersion": "0.3.1",
      "plugins": [
        {
          "type": "flannel",
          "delegate": {
            "hairpinMode": true,
            "isDefaultGateway": true
          }
        },
        {
          "type": "portmap",
          "capabilities": {
            "portMappings": true
          }
        }
      ]
    }
  net-conf.json: |
    {
      "Network": "10.244.0.0/16",
      "Backend": {
        "Type": "vxlan"
      }
    }
```

保存为 `kube-flannel.yml` 并执行：

```bash
kubectl apply -f kube-flannel.yml
```

---

## 6. 添加工作节点

### 6.1 获取 join 命令

在 Master 节点初始化时已经生成了 join 命令。如果需要重新生成：

```bash
# 在 Master 节点执行
kubeadm token list

# 如果 token 已过期，重新生成
kubeadm token create --print-join-command
```

### 6.2 获取 CA cert hash

```bash
openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | openssl rsa -pubin -outform der 2>/dev/null | openssl dgst -sha256 -hex | sed 's/^.* //'
```

### 6.3 在 Worker 节点执行 join

```bash
sudo kubeadm join 192.168.1.10:6443 \
  --token xxxxxx \
  --discovery-token-ca-cert-hash sha256:xxxxxx
```

**参数说明**：
- `192.168.1.10:6443`：Master 节点的 API Server 地址
- `--token`：用于身份验证的令牌
- `--discovery-token-ca-cert-hash`：CA 证书指纹，用于验证 Master 身份

### 6.4 验证节点加入

在 Master 节点执行：

```bash
kubectl get nodes
```

应该能看到所有节点，状态为 NotReady（等待网络插件安装完成）

---

## 7. 验证集群状态

### 7.1 查看节点状态

```bash
kubectl get nodes -o wide
```

### 7.2 查看 Pod 状态

```bash
# 查看所有 Pod
kubectl get pods -A

# 查看系统 Pod
kubectl get pods -n kube-system
```

### 7.3 常用故障排查命令

```bash
# 查看节点详细信息
kubectl describe node <node-name>

# 查看 Pod 日志
kubectl logs -n kube-system <pod-name>

# 查看组件状态
kubectl get cs

# 重置集群（如果需要重新安装）
# kubeadm reset
```

---

## 8. 部署示例应用

### 8.1 部署 Nginx 应用

```bash
# 创建 deployment
kubectl create deployment nginx --image=nginx:latest

# 查看 deployment
kubectl get deployment

# 查看 pod
kubectl get pods
```

### 8.2 暴露服务

```bash
# 使用 NodePort 暴露服务
kubectl expose deployment nginx --type=NodePort --port=80

# 查看服务
kubectl get svc
```

### 8.3 扩缩容

```bash
# 扩容到 3 个副本
kubectl scale deployment nginx --replicas=3

# 缩容到 1 个副本
kubectl scale deployment nginx --replicas=1
```

---

## 附录：常用命令速查

### 资源管理

```bash
# 查看资源
kubectl get <resource> [name] [-n namespace]

# 创建资源
kubectl create -f <file.yaml>

# 应用资源
kubectl apply -f <file.yaml>

# 删除资源
kubectl delete <resource> name [-n namespace]

# 查看资源详情
kubectl describe <resource> name [-n namespace]

# 查看资源日志
kubectl logs <pod-name> [-n namespace]

# 进入容器
kubectl exec -it <pod-name> -- /bin/bash
```

### 常用资源类型

- `pods` / `po`：Pod
- `services` / `svc`：Service
- `deployments` / `deploy`：Deployment
- `replicasets` / `rs`：ReplicaSet
- `configmaps` / `cm`：ConfigMap
- `secrets`：Secret
- `nodes` / `no`：节点
- `namespaces` / `ns`：命名空间

---

## 常见问题

### Q1: kubeadm init 失败

**解决**：检查系统要求，确保满足以下条件：
- 关闭 Swap
- 加载 br_netfilter 模块
- 配置 net.ipv4.ip_forward=1

### Q2: 节点加入失败

**解决**：
- 检查 Master 节点防火墙是否开放 6443 端口
- 确认 token 未过期
- 检查节点时间是否同步

### Q3: Pod 一直处于 Pending

**解决**：
- 检查网络插件是否正确安装
- 检查节点资源是否充足

### Q4: 容器镜像拉取失败

**解决**：
- 配置国内镜像加速器
- 或使用私有镜像仓库

---

## 参考文档

- [Kubernetes 官方文档](https://kubernetes.io/zh/docs/)
- [kubeadm 安装指南](https://kubernetes.io/zh/docs/setup/production-environment/tools/kubeadm/)
- [Flannel 网络插件](https://github.com/flannel-io/flannel)

---

*文档创建日期：2026-03-17*
