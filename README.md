# Omiweb Web 微服务框架
**作者**: stormi-li  
**Email**: 2785782829@qq.com  
## 简介

**Omiweb** 是一个轻量级的 Web 微服务框架，旨在简化 Web 服务的构建，并使其能够作为微服务接入注册中心。通过 **Omiweb**，服务能够进行注册、发现，并支持相互调用，从而实现更高效的分布式系统架构。


## 功能

- **支持构建 Web 服务**：通过简单的配置，快速启动一个 Web 服务。
- **支持接入注册中心**：服务可以注册到指定的注册中心，支持服务的发现和互调。

## 教程
### 安装
```shell
go get github.com/stormi-li/omiweb-v1
```
### 使用
```go
import (
	"embed"

	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omiweb-v1"
)

func main() {
	// 初始化 omiweb 客户端，连接到 Redis 服务
	omiweb := omiweb.NewClient(&redis.Options{Addr: "localhost:6379"})

	// 生成默认的 Web 模板文件
	omiweb.GenerateTemplate()

	// "web_demo"：服务器名称，用于标识 Web 服务， "localhost:8080"：服务器地址和端口
	ws := omiweb.NewWebServer("web_demo", "localhost:8080")

	// - 1：权重为 1
	ws.Start(1)
}
```