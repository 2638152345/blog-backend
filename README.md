Blog Backend (Gin + GORM)
一、项目简介

本项目是一个使用 Go 语言 开发的个人博客系统后端，基于 Gin 框架 和 GORM ORM，实现了博客文章的基本管理功能（CRUD），并支持 用户认证（JWT） 和 评论功能。

项目主要功能包括：

用户注册与登录（JWT 鉴权）

博客文章的创建、读取、更新、删除（CRUD）

评论的创建与读取

统一错误处理与日志记录

二、运行环境

操作系统：Windows + WSL / Linux / macOS

Go 版本：Go 1.20+

数据库：SQLite

依赖管理：Go Modules

三、项目结构说明
blog-backend/
├── main.go                # 程序入口
├── go.mod
├── go.sum
├── README.md
├── config/                # 配置相关（数据库、日志）
│   ├── db.go
│   └── logger.go
├── models/                # GORM 模型定义
│   ├── user.go
│   ├── post.go
│   └── comment.go
├── handlers/              # 业务处理（Controller）
│   ├── auth.go
│   ├── post.go
│   └── utils.go
├── middleware/            # 中间件（JWT 鉴权）
│   └── jwt.go
└── scripts/               # 接口测试脚本
    └── curl_test.sh

四、依赖安装
1️⃣ 初始化依赖（如首次运行）
go mod tidy


主要依赖包括：

github.com/gin-gonic/gin

gorm.io/gorm

gorm.io/driver/sqlite

github.com/golang-jwt/jwt/v5

五、启动项目
go run main.go


启动成功后，服务默认监听：

http://localhost:8080

六、数据库说明

使用 SQLite

数据库文件在程序启动时自动创建

表结构由 GORM 自动迁移生成

包含以下表：

users：用户信息

posts：博客文章

comments：文章评论

七、接口说明（简要）
1️⃣ 用户认证
方法	路径	说明
POST	/register	用户注册
POST	/login	用户登录，返回 JWT
2️⃣ 文章管理（需 JWT）
方法	路径	说明
POST	/posts	创建文章
GET	/posts	获取文章列表
GET	/posts/:id	获取文章详情
PUT	/posts/:id	更新文章（作者）
DELETE	/posts/:id	删除文章（作者）
3️⃣ 评论功能（需 JWT）
方法	路径	说明
POST	/comments	创建评论
GET	/comments	获取文章评论
八、接口测试（curl）

本项目使用 curl 在 WSL 环境中对接口进行测试，测试脚本位于：

scripts/curl_test.sh


测试内容包括：

用户注册

用户登录（获取 JWT）

创建文章

获取文章列表

创建评论

获取评论列表

执行测试脚本
chmod +x scripts/curl_test.sh
./scripts/curl_test.sh


⚠️ 脚本中使用了 jq 解析 JSON，如未安装请先执行：

sudo apt update
sudo apt install jq -y

九、错误处理与日志

接口返回统一的 HTTP 状态码与 JSON 错误信息

使用日志记录系统启动与运行信息，方便调试和维护