# gw 网关

---
GW（网关）包含请求转发，请求超时设置，请求限制，该项目对外提供数据增加的接口，暂无页面显示，具体接口文档将跟随项目 README 更新

# 项目依赖
- gin-gonic/gin (所用框架)
- MongoDB (数据存储)
- Redis (缓存及限流等使用)

#功能包含
- 请求转发 (目前只包含:GET/POST)
- 多 IP/域名 配置
- 请求 dns
- 限流
- 数据缓存 

# 如何使用
### 启动项目
```go
go run main.go
```
>* DOC.md 中提供了目前写入数据的接口文档
>* TABLE.md 中提供了MongoDB表设计文档