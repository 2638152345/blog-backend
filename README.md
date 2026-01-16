## 一、项目简介

本项目是一个使用 **Go 语言** 开发的个人博客系统后端，基于 **Gin 框架** 和 **GORM ORM**，实现了博客文章的基本管理功能（CRUD），并支持 **用户认证（JWT）** 和 **评论功能**。

项目主要功能包括：
- 用户注册与登录（JWT 鉴权）  
- 博客文章的创建、读取、更新、删除（CRUD）  
- 评论的创建与读取  
- 统一错误处理与日志记录  

---

## 二、运行环境

- 操作系统：Windows + WSL / Linux / macOS  
- Go 版本：Go 1.20+  
- 数据库：SQLite  
- 依赖管理：Go Modules  

---

## 三、项目结构说明

<pre>
blog-backend/
├── main.go
├── go.mod
├── go.sum
├── README.md
├── config/
│   ├── db.go
│   └── logger.go
├── models/
│   ├── user.go
│   ├── post.go
│   └── comment.go
├── handlers/
│   ├── auth.go
│   ├── post.go
│   └── utils.go
├── middleware/
│   └── jwt.go
└── scripts/
    └── curl_test.sh
</pre>



---

## 四、依赖安装

\`\`\`bash
go mod tidy
\`\`\`

主要依赖：
- github.com/gin-gonic/gin  
- gorm.io/gorm  
- gorm.io/driver/sqlite  
- github.com/golang-jwt/jwt/v5  

---

## 五、启动项目

\`\`\`bash
go run main.go
\`\`\`

默认监听：
http://localhost:8080

---

## 六、数据库说明

- 使用 SQLite  
- 数据库文件在程序启动时自动创建  
- 表结构由 GORM 自动迁移生成  
- 包含表：users、posts、comments  

---

## 七、接口说明（简要）

### 用户认证

| 方法 | 路径 | 说明 |
| ---- | ---- | ---- |
| POST | /register | 用户注册 |
| POST | /login | 用户登录，返回 JWT |

### 文章管理（需 JWT）

| 方法 | 路径 | 说明 |
| ---- | ---- | ---- |
| POST | /posts | 创建文章 |
| GET | /posts | 获取文章列表 |
| GET | /posts/:id | 获取文章详情 |
| PUT | /posts/:id | 更新文章（作者） |
| DELETE | /posts/:id | 删除文章（作者） |

### 评论功能（需 JWT）

| 方法 | 路径 | 说明 |
| ---- | ---- | ---- |
| POST | /comments | 创建评论 |
| GET | /comments | 获取文章评论 |

---

## 八、接口测试（curl）

脚本位置：scripts/curl_test.sh  

执行：
\`\`\`bash
chmod +x scripts/curl_test.sh
./scripts/curl_test.sh
\`\`\`

⚠️ 需安装 jq：
\`\`\`bash
sudo apt update
sudo apt install jq -y
\`\`\`

---

## 九、错误处理与日志

- 接口返回统一 HTTP 状态码和 JSON 错误信息  
- 使用日志记录系统启动和运行信息
