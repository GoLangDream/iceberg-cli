# iceberg-cli
iceberg 框架的命令行工具

## install

项目依赖 Git 和 Golang, 请先安装好这两个工具，并配置好 GOPATH, 以及把GOPATH加入 PATH

```shell
sh -c "$(curl -fsSL https://raw.githubusercontent.com/GoLangDream/iceberg-cli/main/install.sh)"
```

## 使用

### 新建一个项目

```shell
iceberg new test_project
```

### 其他一些命令

```shell
iceberg g m create_user                    # 生成一个migration
iceberg g model user name:string age:uint  # 生成一个模型
iceberg g controller user index show       # 生成一个controller 和 对应的 action
```
