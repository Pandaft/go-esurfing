# go-esurfing

## 简介

基于 Go 实现登入和登出广东天翼校园网的命令行工具。

<br />

## 命令行

```text
> go-esurfing -h

基于 Go 语言实现登入和登出广东天翼校园网的命令行工具
项目 GitHub：https://github.com/Pandaft/go-esurfing

Usage:
  go-esurfing [flags]
  go-esurfing [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  login       登入
  version     输出版本

Flags:
  -h, --help   help for go-esurfing

Use "go-esurfing [command] --help" for more information about a command.
```

### 登入

```text
> go-esurfing login -h

登入广东天翼校园网

对于 nasip 和 clientip 参数：
  - 本机未登入，且在本机登入时，可不填写
  - 本机已登入，或在远程登入时，必须填写

Usage:
  go-esurfing login [flags]

Flags:
  -a, --acc string        账号
  -c, --clientip string   clientIP
  -d, --debug             调试模式 (default true)
  -h, --help              help for login
  -m, --mac string        MAC 地址
  -n, --nasip string      nasIP
  -p, --pwd string        密码
```

### 输出版本

```text
> go-esurfing version -h

输出当前 go-esurfing 具体版本

Usage:
  go-esurfing version [flags]

Flags:
  -h, --help   help for version
```

<br />

## 参考项目

此项目核心实现方法参考自 [Z446C/ESC-Z](https://github.com/Z446C/ESC-Z/) ，可以认为是 Go 复刻版。

<br />

## 其他

敬请期待。
