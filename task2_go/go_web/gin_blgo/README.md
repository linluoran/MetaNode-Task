# Gin Blog - Go Gin 博客系统

![Go Version](https://img.shields.io/badge/go-1.24%2B-blue)
![Gin Framework](https://img.shields.io/badge/gin-1.10.1-green)

Gin Blog 是一个基于 Go 语言和 Gin 框架构建的高性能博客系统，提供完整的文章管理、用户认证和内容展示功能。

## 功能特性

- 🚀 高性能 Gin 框架驱动
- 🔐 JWT 用户认证与权限管理
- 📝 文章编辑与展示
- 🔍 文章标题搜索功能
- 📊 分页与数据统计
- 📱 RESTful API 设计
- 📈 Zap 高性能日志记录

## 运行环境

- Go 1.24+ ([下载地址](https://golang.org/dl/))
- MySQL 5.7+ 

## 项目结构

~~~bash
.
├── README.md
├── etc																		# 配置文件
│   └── config.yaml
├── internal
│   ├── config														# 加载配置
│   │   └── config.go
│   ├── handler													 	# 视图
│   │   ├── comment_handler.go
│   │   ├── post_handler.go
│   │   └── user_handler.go
│   ├── middleware											  # 中间件
│   │   └── middleware.go
│   ├── model														  # 数据库
│   │   ├── comment.go
│   │   ├── post.go
│   │   └── user.go
│   ├── pkg
│   │   ├── dao														# mysql
│   │   │   └── mysql.go
│   │   └── logger												# zap日志
│   │       └── logger.go									
│   ├── routers														# 路由
│   │   ├── comment_router.go
│   │   ├── post_router.go
│   │   ├── router.go											# 主路由
│   │   └── user_router.go
│   └── types															# 请求体
│       └── types.go											
├── logs																	# 日志存储目录
└── main.go																# 主程序入口
~~~

## 快速开始

### 1. 克隆项目

```bash
git clone https://github.com/linluoran/MetaNode-Task.git
cd MetaNode-Task/task2_go/go_web/gin_blgo
```

### 2. 配置环境

复制示例配置文件并修改：

```bash
cp config/config.example.yaml config/config.yaml
```

编辑 `config/config.yaml` 文件，配置数据库等信息：

```yaml
Name: "gin blog"
Env: "dev"
Host: 127.0.0.1
Port: 8080


Mysql:
  Username: "gin"
  Password: "gin123456"
  Host: "127.0.0.1"
  Port: "3306"
  DBname: "gozero"
  Timeout: "10s"
  MaxOpenConns: 100    # 最大连接数
  MaxIdleConns: 10     # 空闲连接数
  ConnMaxLifetime: 30m # 连接最大存活时间

Log:
  LogPath: "logs"  # 日志文件路径
  MaxSize: 20             # 单个文件最大大小(MB)
  MaxBackups: 5           # 保留的旧日志文件数
  MaxAge: 7              # 日志保留天数
  Compress: true           # 是否压缩旧日志

Jwt:
  Secret: "jGudRl2qx0zAbckdK1unsq8vxSy1riQ1HkVTn59qois="
  TokenExpire: 24
  Issuer: "gin blog"
```

### 3. 安装依赖

```bash
go mod tidy
```

### 4. 初始化数据库

```sql
-- MySQL 示例
CREATE DATABASE IF NOT EXISTS gozero CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 5. 启动应用

```bash
# 直接运行
go run main.go
```

应用将在 `http://localhost:8080` 启动

## API 文档

### https://doc.apipost.net/docs/499205ad9cbe000?locale=zh-cn

| 方法 | 路径                     | 描述         | 处理器函数             |
| :--- | :----------------------- | :----------- | :--------------------- |
| POST | `/api/v1/user/register`  | 用户注册     | `UserRegisterHandler`  |
| POST | `/api/v1/user/login`     | 用户登录     | `UserLoginHandler`     |
| POST | `/api/v1/post/create`    | 创建文章     | `PostCreateHandler`    |
| POST | `/api/v1/post/list`      | 获取文章列表 | `PostListHandler`      |
| POST | `/api/v1/post/detail`    | 获取文章详情 | `PostDetailHandler`    |
| POST | `/api/v1/post/update`    | 更新文章     | `PostUpdateHandler`    |
| POST | `/api/v1/post/delete`    | 删除文章     | `PostDeleteHandler`    |
| POST | `/api/v1/comment/create` | 创建评论     | `CommentCreateHandler` |
| POST | `/api/v1/comment/list`   | 获取评论列表 | `CommentListHandler`   |

## 使用 Makefile 命令

```bash
# 查看所有可用命令
make help

# 开发模式运行 (带热重载)
make dev

# 编译项目
make build

# 运行测试
make test

# 清理构建文件
make clean

# 格式化代码
make fmt

# 静态检查
make lint
```

## 贡献指南

欢迎贡献代码！请遵循以下步骤：

1. Fork 项目仓库
2. 创建特性分支 (`git checkout -b feature/your-feature`)
3. 提交更改 (`git commit -am 'Add some feature'`)
4. 推送到分支 (`git push origin feature/your-feature`)
5. 创建 Pull Request

## 技术支持

如有任何问题，请提交 issue 或联系：  
your.email@example.com