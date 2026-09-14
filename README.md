# Go Blog

一个轻量级、高性能的静态博客系统，使用 Go 语言开发，支持 Markdown 撰写文章。

## 特性

- **Markdown 支持** - 使用 Markdown 格式编写博客文章，天然支持代码高亮
- **分类管理** - 通过文件夹自动管理文章分类
- **标签系统** - 灵活的文章标签管理
- **归档功能** - 按年份自动归档文章
- **中文友好** - URL 支持中文slug，界面完全中文本地化
- **响应式设计** - 适配桌面端和移动端浏览
- **高性能** - Go 语言开发，静态页面渲染，性能卓越
- **Docker 部署** - 一键 Docker 部署

## 技术栈

- **后端**: Go 1.23
- **路由**: Gorilla Mux
- **Markdown**: gomarkdown/markdown
- **部署**: Docker + Docker Compose

## 目录结构

```
blog/
├── main.go                 # 程序入口
├── config/
│   └── config.go           # 配置管理
├── handlers/
│   ├── handler.go          # 基础处理器
│   ├── page_home.go        # 首页
│   ├── page_post.go        # 文章页
│   ├── page_archives.go    # 归档页
│   ├── page_tags.go        # 标签列表页
│   ├── page_tag.go         # 标签详情页
│   └── page_about.go       # 关于页
├── services/
│   ├── markdown.go         # Markdown 解析
│   ├── markdown_post.go    # 文章解析
│   ├── markdown_list.go    # 文章列表
│   ├── markdown_tags.go    # 标签服务
│   └── markdown_archive.go # 归档服务
├── models/
│   └── post.go             # 数据模型
├── router/
│   └── router.go           # 路由配置
├── templates/              # HTML 模板
│   ├── index.html         # 首页模板
│   ├── post.html          # 文章页模板
│   ├── archives.html      # 归档页模板
│   ├── tags.html          # 标签列表模板
│   ├── tag.html           # 标签详情模板
│   ├── about.html         # 关于页模板
│   └── _sidebar.html      # 侧边栏模板
├── static/                 # 静态资源
│   ├── css/               # 样式文件
│   └── images/            # 图片资源
├── posts/                  # 博客文章目录
│   └── welcome.md         # 示例文章
├── Dockerfile             # Docker 镜像构建
└── docker-compose.yml     # Docker Compose 配置
```

## 安装部署

### 方式一：Docker 部署（推荐）

1. 克隆或下载本项目

2. 进入项目目录：
   ```bash
   cd blog
   ```

3. 使用 Docker Compose 启动：
   ```bash
   docker compose up -d
   ```

4. 访问博客：`http://localhost:8083`

### 方式二：本地运行

1. 安装 Go 1.23 或更高版本

2. 克隆项目：
   ```bash
   git clone <repository-url>
   cd blog
   ```

3. 安装依赖：
   ```bash
   go mod download
   ```

4. 运行程序：
   ```bash
   go run main.go
   ```

5. 访问博客：`http://localhost:8083`

### 方式三：手动构建

1. 构建二进制文件：
   ```bash
   go build -o blog main.go
   ```

2. 运行：
   ```bash
   ./blog
   ```

## 配置说明

可以通过环境变量配置博客：

| 环境变量 | 默认值 | 说明 |
|---------|--------|------|
| PORT | :8083 | 服务端口 |
| SITE_NAME | 西康的博客 | 网站名称 |
| PAGE_SIZE | 5 | 首页每页文章数 |
| POSTS_DIR | ./posts | 文章目录 |
| STATIC_DIR | ./static | 静态文件目录 |
| TEMPLATE_DIR | ./templates | 模板目录 |

### Docker 环境变量示例

```yaml
services:
  blog:
    environment:
      - PORT=:8083
      - SITE_NAME=我的技术博客
      - PAGE_SIZE=10
      - POSTS_DIR=/app/posts
```

## 如何使用

### 撰写新文章

1. 在 `posts` 目录下创建 Markdown 文件（建议按分类新建子目录）

2. 文件头部添加 Frontmatter 元信息：
   ```markdown
   ---
   title: 文章标题
   date: 2026-03-18
   author: 作者名
   tags: 标签1, 标签2
   summary: 文章摘要
   ---

   # 正文内容

   开始撰写你的文章...
   ```

3. 文章将自动出现在博客的对应位置

### Frontmatter 字段说明

| 字段 | 必填 | 说明 |
|------|------|------|
| title | 是 | 文章标题 |
| date | 否 | 发布日期，格式：YYYY-MM-DD |
| author | 否 | 作者名，默认值可在 services/markdown_post.go 中修改 |
| tags | 否 | 标签，多个用逗号分隔 |
| summary | 否 | 文章摘要，显示在列表页 |

### 文章分类

- 在 `posts` 目录下创建子目录
- 子目录名称即为分类名称
- 放入子目录的 .md 文件会自动归入该分类

示例：
```
posts/
├── welcome.md              # 无分类
├── go/
│   ├── golang-intro.md    # 分类：go
│   └── go-basics.md
└── docker/
    ├── docker-basics.md   # 分类：docker
    └── docker-compose.md
```

### 访问路径

- 首页：`/`
- 文章：`/post/{slug}`
- 归档：`/archives`
- 标签列表：`/tags`
- 标签详情：`/tag/{tagname}`
- 关于：`/about`

## 更新文档

### 更新文章

直接在 `posts` 目录下编辑或新增 Markdown 文件，然后重启服务即可生效。

### 添加自定义页面

1. 在 `templates` 目录下创建新的 HTML 模板
2. 在 `handlers` 目录下创建对应的处理器
3. 在 `router/router.go` 中添加路由

### 修改样式

- 全局样式：`static/css/style.css`
- 页面特定样式：
  - 首页：`static/css/index.css`
  - 文章页：`static/css/post.css`
  - 归档页：`static/css/archives.css`
  - 标签页：`static/css/tag.css` 和 `static/css/tags.css`
  - 关于页：`static/css/about.css`

### 修改侧边栏

编辑 `templates/_sidebar.html` 模板文件。

## 亮点与优势

### 1. 简单高效

无需数据库，纯静态文件存储，文章就是 Git 仓库中的 Markdown 文件，便于版本管理和备份。

### 2. 性能卓越

Go 语言开发，编译为单一可执行文件，启动快、内存占用低，适合个人博客场景。

### 3. 易于定制

- 清晰的代码结构，便于二次开发
- 模板系统基于 Go 的 html/template，简单易学
- 静态资源分离，可自由定制样式

### 4. 部署便捷

- Docker 一键部署
- 支持热更新（挂载 posts 目录到容器）
- 交叉编译支持多平台

### 5. 中文优化

- 原生支持中文 URL slug
- 完整的中文界面
- 适合国内技术社区

## 常见问题

### Q: 如何修改博客标题？
A: 修改环境变量 `SITE_NAME`，或直接在 `main.go` 中修改默认值。

### Q: 如何修改每页显示文章数？
A: 修改环境变量 `PAGE_SIZE`。

### Q: 如何添加评论功能？
A: 当前版本为静态博客，可集成第三方评论服务（如 Giscus、Disqus 等）。

### Q: 如何实现图片上传？
A: 建议使用图床服务（如 GitHub 图床、sm.ms 等），在 Markdown 中引用图片 URL。

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License
