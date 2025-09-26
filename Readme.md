## 用户登录
/user/register
Service 层 RegisterUser 检查数据库中是否已有该用户名。 已存在 -> 返回错误。 不存在 -> 对密码进行 bcrypt 哈希加密，存入数据库。
/user/login
Service 层 Authenticate 检查用户名和密码是否正确。 正确 -> 生成一个 JWT token。 把 token 保存到 Redis并设置过期时间（24h）
/user/admin(保护接口)
通过 AuthMiddleware 中间件进行鉴权。 成功返回admin登录成功 失败返回403

## 动态
{ 登录后可发 
    /Moment/create

    /Moment/delete

    /Moment/list

    /Moment/update
}# QQzone
