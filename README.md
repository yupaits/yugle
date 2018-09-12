# yugle
万千资讯，一钩钓之。

## 运行步骤

1. 搭建 Golang 和 node.js 运行环境

1. `git clone https://github.com/YupaiTS/yugle.git` 下载代码

1. 执行前端构建命令：

    ```
    cd yugle/ui
    npm install
    npm run build
    ```

1. 使用 `go get xxx` 命令下载依赖包

1. 在本地 MySQL 数据库中新建 `yugle` 数据库

1. 使用 `go test yugle_test.go` 命令运行单元测试用例，插入测试数据

1. 使用 `go run main.go` 命令运行项目，使用测试用例中的用户名密码进行登录