# go-study

## Workspaces

[Tutorial: Getting started with multi-module workspaces](https://go.dev/doc/tutorial/workspaces)

### Introduction

- `workspace` 在 `Go 1.18` 版本后引入；
- `workspace` 即工作空间，用于统一管理多个软件模块，如可以快速测试、构建和运行多个模块中的代码；
- `workspace` 使用 `go work` 工具进行管理，配置在 `go.work` 文件中，用法和 `go mod` 相似；
- `workspace` 可以不使用 `replace` 指令就实现替换模块的功能（见示例）。

### Usage

#### go help work [subcommand]

#### go work init [moddirs]

在当前目录下创建一个 `go.work` 文件，即将当前目录标识为一个工作空间：

- 当忽略 `moddirs` 时，仅创建一个不包含任何 `use` 指令的 `go.work` 文件；
- 当添加 `moddirs` 时，需要每个 `moddir` 下都存在 `go.mod` 文件，然后会在 `go.work` 文件中使用 `use` 指令添加这些 `moddir`；

```Shell
# Usage 1: go mod init without any moddir
go mod init

# Usage 1: go mod init with one moddir
go mod init ./module1

# Usage 2: go mod init with multiple moddirs
go mod init ./module1 ./module2
```

当进行上述配置后，可以在工作空间中运行任意模块（即使我们没有在模块内）：

```Shell
go run ./module1
```

#### go work use

将指定模块添加到当前工作空间中。

```Shell
# 将指定模块加入到当前工作空间
go work use ./module2

# 将指定目录下的所有模块都加入到当前工作空间
go work use -r ./modules
```