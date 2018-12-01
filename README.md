# yugle

[![Build Status](https://travis-ci.com/YupaiTS/yugle.svg?branch=master)](https://travis-ci.com/YupaiTS/yugle)

万千资讯，一钩钓之。

## 使用技术

- gin
- gorm
- go_spider
- robfig/cron

## 运行步骤

1. 搭建 Golang 和 node.js 运行环境

1. `git clone https://github.com/YupaiTS/yugle.git` 下载代码

1. 使用 `go get xxx` 命令下载依赖包

1. 在本地 MySQL 数据库中新建 `yugle` 数据库

1. 使用 `go test yugle_test.go` 命令运行单元测试用例

1. 使用 `go run main.go` 命令运行项目