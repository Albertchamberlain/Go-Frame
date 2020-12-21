### os/exec
我第一眼看到这个包的介绍的时候就觉得它是写小工具的好手
本章代码建议在wsl下执行
`exec包执行外部命令，它将os.StartProcess进行包装使得它更容易映射到stdin和stdout，并且利用pipe连接i/o`

LookPath在环境变量中查找科执行二进制文件，如果file中包含一个斜杠，则直接根据绝对路径或者相对本目录的相对路径去查找
```go
func LookPath(file string) (string, error) 
``` 



