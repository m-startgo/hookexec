# hookexec

这是一个基于 WebHooks 的脚本执行器

## 项目生成提示词

帮我搭建一个这样的项目：

参考 Wails3 的目录结构，但这不是一个 桌面项目，而是一个网站项目。

1. 存在一个目录为 frontend 的前端目录，使用的可以是 任意 前端框架
2. 项目基地为 go 语言，它内置默认会启动两个服务

服务一： 以 frontend/dist 目录为静态资源目录
使用 go 的 embed 包嵌入 frontend/dist 目录
如果 404 了，则返回 index.html 文件，如果存在 404.html 则返回 404.html 文件

```go
//go:embed all:frontend/dist
var FrontendDist embed.FS
```

服务二： 提供一个 http 接口，用于返回各种数据供前端调用

这两个服务均监听在 9900 上。

例如：

http://localhost:9900/ -> 返回 frontend/dist 目录下的 index.html 文件
http://localhost:9900/api/data -> 返回一些 json 数据

整个服务以 gin 为 web 框架。
