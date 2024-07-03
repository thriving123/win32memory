# win32lib使用说明

通过一下命令安装win32lib

> go get github.com/thriving123/win32memory


使用方式
```go
import "github.com/thriving123/win32memory"
pid,_ := win32.GetPidByName("你的进程名.exe")
// 读取内存整数型
date,_ = win32.ReadInt(pid, 内存地址)
// ... 其他自测
```