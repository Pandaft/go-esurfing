# go-esurfing <sup>v0.1.2</sup>

基于 Go 实现登入和登出广东天翼校园网的命令行工具。

<br />

## 命令行

```text
> ./go-esurfing -h

基于 Go 实现登入和登出广东天翼校园网的命令行工具 (v0.1.2)
GitHub: https://github.com/Pandaft/go-esurfing

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
> ./go-esurfing login -h

登入广东天翼校园网

必填参数：username, password

对于 nasip 和 clientip 参数：
  - 本机未登入，且在本机登入时，可不填写
  - 本机已登入，或在远程登入时，必须填写

对于 mac 参数：
  - 暂未发现对登入功能有实际影响
  - 不填写默认为 00-00-00-00-00-00

Usage:
  go-esurfing login [flags]

Flags:
  -n, --nasip    string   认证服务器 IP
  -c, --clientip string   登录设备 IP
  -m, --mac      string   MAC 地址
  -u, --username string   账号
  -p, --password string   密码
  -d, --debug             调试模式
  -h, --help              help for login
```

### 输出版本

```text
> ./go-esurfing version -h

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

## 免责声明

**此项目仅供研究、学习和交流，请勿用于商业或非法用途，开发者与协作者不对使用者负任何法律责任，使用者自行承担因不当使用所产生的后果与责任。**

**This project is only for research, learning and exchange. Do not use it for commercial or illegal purposes. Developers and collaborators do not assume any legal responsibility for users. Users bear the consequences and responsibilities arising from improper use.**

<br />

## 其他

广东天翼校园网 QQ 交流群：791455104（[点此加入](http://qm.qq.com/cgi-bin/qm/qr?_wv=1027&k=yTA84KiemCppMD5Y2CDepUsnVRo59dOS&authKey=CH%2Bb2yFiTVPqLOjdwrEGXGVvmhWTURTFX8yM5eRA7ipWh5fOKAIpJRqCKDIWZT7V&noverify=0&group_code=791455104)）
