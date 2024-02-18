# web-app-server

> 将前端 `web` 模拟成软件的形式启动

## 使用

- 安装

```bash
go install  -ldflags -H=windowsgui github.com/QuietSugar/web-app-server@latest
```

- 启动

```bash
# 指定路径
./web-app-server -w .
# 执行ico路径
./web-app-server -i ./default.ico
```
