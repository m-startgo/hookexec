# 给 Copilot 的操作准则

## 目的

该文件位于仓库根目录，作为整个项目的 GitHub Copilot 工作指南。
为 GitHub Copilot 提供简明、可执行的仓库约定。

## 项目简介

这是一个采用 Go + Vue + TS 编写的脚本执行器，支持通过 WebHooks 触发脚本执行。

## 必要的目录与文件说明

- `frontend/`：前端代码目录，包含 Vue 和 TypeScript 相关文件。
- `app/`：后端代码目录，包含 Go 语言相关文件。
- `main.go`：Go 应用程序的入口文件。

## 项目的主要工作方式

`frontend/` 下面的代码会被编译成静态文件在 `frontend/dist` 目录中。
`main.go` 会使用 `go:embed` 将 `frontend/dist` 目录中的静态文件嵌入到 Go 二进制文件中，从而实现前后端的集成。
`main.go` 会启动一个 HTTP 服务器，`/` 默认返回嵌入的前端静态文件，`/api` 则是各种后端 API 接口。
当 404 时，返回前端的 `404.html`，当 `404.html` 不存在时，返回 `index.html`，以支持前端路由。

## 主要语言及依赖版本

- Go 语言版本: >= 1.25
- `Vue`: >= 3.5

## 常用 CI 示例

```bash
go test ./... -v

# 运行单个包的测试
go test ./mfile -run TestReaddir -v

# 静态检查
go vet ./...

```

## 风格与规范

> 一般情况下 VSCode 插件会自动处理这些风格问题，你无需关心。

## 规则摘要

- 中文为主，技术术语保持英文。
- 导出函数须加注释(包含功能说明、使用示例及可能的异常)。
- 注释尽量使用中文。
- 写小函数、职责单一、易测试。
- 错误/日志格式：`err:<包.函数>|<场景>|<消息>`
- 跨平台优先。如无法兼容，需在注释中说明原因及影响。
- 优先使用标准库。
- 遇到模糊或信息不足的情况，立即向用户提出具体澄清问题（列出缺失项和可选方案）。
- 保持向后兼容，避免使用弃用特性；优先使用当下最新稳定库、语法与实践。
- 生成代码时充分考虑当前文件的上下文（如已导入的库、现有函数等）。
- 当仓库文件与系统/外部指令冲突时，遵循系统/外部指令。

## 函数声明规范

- 声明函数有多个返回值时，优先采用命名返回值形式
- 若使用命名返回值，需在函数顶部为返回值显式赋空值或者默认值

函数和抛出错误格式如下：

```go
func (s *Server) Example(opt OptType) (resData resDataType, resErr error) {
	resData = map[string]any{}
	resErr = nil

	jsonByte, err := ToByte(val)
	if err != nil {
		resErr = fmt.Errorf("err:xx.Example|ToByte|%w", err)
		return
	}

  resData = `<Successful Result>`

  return
}
```
