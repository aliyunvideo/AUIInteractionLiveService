# Live
* 本项目代码仅作范例演示使用。

## 构建执行文件
```bash
go build -o main ./cmd/main.go
```

## 本地运行
```bash
export ADMIN_PASSWORD=your-secret
./main
```
所有API及对应文档:
浏览器打开：http://localhost:7001/swagger/index.html
用户名为 admin，密码为您自行设置

* 注意：目前数据库表为自动创建，线上请关注数据库变更，相关索引及查询优化需自行关注。


## 生成`swagger`文档

```bash
# 安装swagger工具 请具体参考：https://github.com/swaggo/swag
go install github.com/swaggo/swag/cmd/swag@latest

# 动态生成文档
swag init -g pkg/handler/*.go
```


## 基于Swagger生成多语言sdk
如果需要通过服务端应用调用本应用API，可以通过辅助该工具使用：
https://github.com/swagger-api/swagger-codegen

```bash
swagger-codegen generate -i ./docs/swagger.json -l java -o /tmp/test/
```

## 函数使用
请参考：https://docs.serverless-devs.com/serverless-devs/quick_start
