
## 简介
这是一个简单的登录系统，实现了用户登录、注册、用户信息的更新、头像上传等功能。

## 技术栈
Angular + Go + postgresSQL + redis。

## 部署
使用docker容器技术 + ngnix反向代理。

## 部分细节
- [ ] 使用邮箱验证实现注册功能。
- [ ] 通过 dgrijalva/jwt 工具，实现 Token 的生成，用以保持用户状态。
- [ ] 使用 AES 对 jwt 进行加密，提高 Token 安全性。
- [ ] 使用加盐哈希的方式，在数据库中存储用户的密码。
- [ ] 使用 beego/logs，实现日志输出。
- [ ] 通过 localstorage 存储前端用户数据。


