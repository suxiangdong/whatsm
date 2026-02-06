# WhatsApp Web Project

## 项目概述

该项目通过 GoFrame 开发，实现了 Web 版本 WhatsApp 登录、消息发送等功能。支持跨平台编译，能够在 Windows、Linux 和 macOS 系统上运行。

## 应用版本
whatsm v1.0.0

## 依赖
golang 1.24.0

## 功能

1. **Web 登录**
    - 提供登录接口，返回配对码和二维码，用于登录 Web 版本的 WhatsApp。

2. **发送消息**
    - 支持通过 API 发送 WhatsApp 消息。

## 配置文件

- 项目内默认配置文件路径：`manifest/config/config.yaml`

## 已知问题

- 频繁登录退出会导致账号封禁，封禁时间为 6 小时。

## 编译与运行
1. **编译支持**
    - 使用 `gf build` 命令进行编译，支持的操作系统包括：
        - Windows、Linux、Darwin (macOS)
        - 支持的架构：arm64 和 amd64

2. **启动项目**
    - 启动命令：`./whats start`
    - 配置文件可通过 `-c` 参数指定，默认配置文件已打包进二进制文件内。

3. **默认配置文件**
    - 默认配置文件端口为 `8090`。
   
4. **编译**：
   - 安装gfcli `wget -O gf https://github.com/gogf/gf/releases/latest/download/gf_$(go env GOOS)_$(go env GOARCH) && chmod +x gf && ./gf install -y && rm ./gf`
   - 使用 `gf build` 命令进行编译，支持交叉编译： Windows、Linux 和 macOS 系统，arm64和amd64架构。

5. **启动**：使用 `./whats start` 启动项目，可以通过 `-c` 参数指定配置文件。

6. **API 文档**：启动后，访问 [Swagger 文档](http://localhost:8090/swagger) 查看所有支持的 API 接口。