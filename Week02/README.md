学习笔记

```go
├── main.go
└── src
    ├── dao
		│   └── user_dao.go // 使用 gorm 操作数据库，使用 errors.Wrap()
    ├── dbs
    │   └── gorm.go // 使用 gorm 连接数据库
    ├── handlers
    │   └── user_handler.go // 用户 http 请求的处理函数
    ├── logic
    │   └── user_logic.go // http 请求的逻辑处理，日志打印
    └── models
        └── user_model.go // 结构体
```

访问 http://localhost:9090/v1/user?uid=1 有数据返回的 `json`

```json
{
  "data": {
    "id": 2,
    "name": "lisi",
    "user_pwd": "",
    "addtime": ""
  },
	"message": "success"
}
```

访问 http://localhost:9090/v1/user?uid=4 没有数据返回的 `json`

```json
{
	"message": "根据uid查找用户信息失败"
}
```

记录的日志为 

```bash
[GIN] 2020/12/02 - 17:13:02 | 404 |     611.698µs |             ::1 | GET      "/v1/user?uid=4"
2020/12/02 17:13:02 根据uid: 4,未找到用户信息
Week02/src/dao.UserDAO.FindOneByUid
	/Go-000/Week02/src/dao/user_dao.go:21
Week02/src/logic.userLogic.FindOne
	/Go-000/Week02/src/logic/user_logic.go:15
Week02/src/handlers.UserController.GetUserInfo
	/Go-000/Week02/src/handlers/user_handler.go:17
github.com/gin-gonic/gin.(*Context).Next
	/Users/go/pkg/mod/github.com/gin-gonic/gin@v1.6.3/context.go:161
github.com/gin-gonic/gin.RecoveryWithWriter.func1
	/Users/go/pkg/mod/github.com/gin-gonic/gin@v1.6.3/recovery.go:83
github.com/gin-gonic/gin.(*Context).Next
	/Users/go/pkg/mod/github.com/gin-gonic/gin@v1.6.3/context.go:161
github.com/gin-gonic/gin.LoggerWithConfig.func1
	/Users/go/pkg/mod/github.com/gin-gonic/gin@v1.6.3/logger.go:241
github.com/gin-gonic/gin.(*Context).Next
	/Users/go/pkg/mod/github.com/gin-gonic/gin@v1.6.3/context.go:161
github.com/gin-gonic/gin.(*Engine).handleHTTPRequest
	/Users/go/pkg/mod/github.com/gin-gonic/gin@v1.6.3/gin.go:409
github.com/gin-gonic/gin.(*Engine).ServeHTTP
	/Users/go/pkg/mod/github.com/gin-gonic/gin@v1.6.3/gin.go:367
net/http.serverHandler.ServeHTTP
	/usr/local/go/src/net/http/server.go:2843
net/http.(*conn).serve
	/usr/local/go/src/net/http/server.go:1925
runtime.goexit
	/usr/local/go/src/runtime/asm_amd64.s:1374
```

访问 http://localhost:9090/v1/user?uid=wo 查询出错返回的 `json`

```json
{
	"message": "根据uid查找用户信息失败"
}
```

记录的日志为 

```bash
[0.266ms] [rows:0] SELECT * FROM `users` WHERE wo
[GIN] 2020/12/02 - 17:15:20 | 404 |     441.806µs |             ::1 | GET      "/v1/user?uid=wo"
2020/12/02 17:15:20 Error 1054: Unknown column 'wo' in 'where clause'
根据uid: wo,查找用户信息发生错误
Week02/src/dao.UserDAO.FindOneByUid
	/Users/Go-000/Week02/src/dao/user_dao.go:18
Week02/src/logic.userLogic.FindOne
	/Users/Go-000/Week02/src/logic/user_logic.go:15
Week02/src/handlers.UserController.GetUserInfo
	/Users/Go-000/Week02/src/handlers/user_handler.go:17
github.com/gin-gonic/gin.(*Context).Next
	/Users/go/pkg/mod/github.com/gin-gonic/gin@v1.6.3/context.go:161
github.com/gin-gonic/gin.RecoveryWithWriter.func1
	/Users/go/pkg/mod/github.com/gin-gonic/gin@v1.6.3/recovery.go:83
github.com/gin-gonic/gin.(*Context).Next
	/Users/go/pkg/mod/github.com/gin-gonic/gin@v1.6.3/context.go:161
github.com/gin-gonic/gin.LoggerWithConfig.func1
	/Users/go/pkg/mod/github.com/gin-gonic/gin@v1.6.3/logger.go:241
github.com/gin-gonic/gin.(*Context).Next
	/Users/go/pkg/mod/github.com/gin-gonic/gin@v1.6.3/context.go:161
github.com/gin-gonic/gin.(*Engine).handleHTTPRequest
	/Users/go/pkg/mod/github.com/gin-gonic/gin@v1.6.3/gin.go:409
github.com/gin-gonic/gin.(*Engine).ServeHTTP
	/Users/go/pkg/mod/github.com/gin-gonic/gin@v1.6.3/gin.go:367
net/http.serverHandler.ServeHTTP
	/usr/local/go/src/net/http/server.go:2843
net/http.(*conn).serve
	/usr/local/go/src/net/http/server.go:1925
runtime.goexit
	/usr/local/go/src/runtime/asm_amd64.s:1374
```

