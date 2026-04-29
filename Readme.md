# QQZone

基于 **Go + Gin + Vue3** 的社交动态平台，支持用户注册登录、好友管理、动态发布（图文/视频）、父子评论等功能。

## 🛠 技术栈

| 层级 | 技术 |
|------|------|
| 后端框架 | Gin (Go) |
| 数据库 | MySQL (GORM) |
| 缓存 / Session | Redis |
| 文件存储 | MinIO |
| 鉴权 | JWT + Redis 黑名单 |
| 前端框架 | Vue 3 (Composition API) |
| UI 组件库 | Element Plus |
| 构建工具 | Vite |
| 类型支持 | TypeScript |

## 📁 项目结构

```
QQZone/
├── main.go                 # 项目入口
├── config.yaml             # 配置文件
├── go.mod / go.sum         # Go 依赖
├── api/                    # 接口层（处理 HTTP 请求）
│   ├── article.go
│   ├── comment.go
│   ├── friend.go
│   └── user.go
├── service/                # 业务逻辑层
│   ├── article.go
│   ├── comment.go
│   ├── friend.go
│   └── user.go
├── model/                  # 数据模型（GORM）
│   ├── article.go
│   └── user.go
├── middleware/              # 中间件
│   ├── auth.go             # JWT 鉴权
│   └── admin.go            # 管理员权限
├── router/                 # 路由注册
│   └── router.go
├── global/                 # 全局变量（DB、Redis、MinIO）
│   └── global.go
├── initialize/             # 初始化模块
│   ├── server.go           # 配置加载
│   ├── mysql.go
│   ├── redis.go
│   └── minio.go
├── utils/                  # 工具函数
│   ├── jwt.go
│   └── minio.go
└── web/                    # Vue3 前端
    ├── src/
    │   ├── api/            # API 请求封装
    │   ├── router/         # 前端路由 + 导航守卫
    │   ├── stores/         # Pinia 状态管理
    │   ├── utils/          # Axios 封装
    │   └── views/          # 页面视图
    ├── vite.config.ts
    └── package.json
```

## 🚀 快速启动

### 后端

```bash
# 1. 确保 MySQL、Redis、MinIO 已启动
# 2. 修改 config.yaml 中的配置

# 3. 安装依赖
go mod tidy

# 4. 运行
go run main.go
# 或 go build -o main.exe . && ./main.exe
```

### 前端

```bash
cd web
npm install
npm run dev        # http://localhost:5173
```

### 生产构建

```bash
cd web
npm run build      # 输出到 web/dist/
```

## 🔌 API 接口

### 用户模块

| 方法 | 路径 | 功能 | 鉴权 |
|------|------|------|------|
| POST | `/user/register` | 用户注册 | ❌ |
| POST | `/user/login` | 用户登录 | ❌ |
| DELETE | `/user/logout` | 退出登录 | ✅ |
| GET | `/user/friends` | 好友列表 | ✅ |
| POST | `/user/friends/add/:id` | 添加好友 | ✅ |
| DELETE | `/user/friends/delete/:id` | 删除好友 | ✅ |
| GET | `/user/admin` | 管理员页面 | ✅ + admin |

### 动态模块

| 方法 | 路径 | 功能 | 鉴权 |
|------|------|------|------|
| GET | `/articles` | 动态列表 | ❌ |
| GET | `/articles/:id` | 动态详情 | ❌ |
| POST | `/articles/create` | 发布动态 | ✅ |
| DELETE | `/articles/:id` | 删除动态 | ✅ |

### 评论模块

| 方法 | 路径 | 功能 | 鉴权 |
|------|------|------|------|
| GET | `/articles/:id/comments` | 获取评论 | ❌ |
| POST | `/articles/:id/comments` | 发表评论 | ✅ |
| DELETE | `/articles/comments/:id` | 删除评论 | ✅ |

## 🔐 鉴权说明

- 登录后返回 JWT Token，存储在 Redis Session 中
- 请求需携带 `Authorization: Bearer <token>`
- 退出登录时 Token 加入 Redis 黑名单
