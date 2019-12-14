# go-mod 使用
v1.13 默认开启默认开启go mod

说明：现在不允许在 GOPATH 下使用 gomod


#第一步：在GOPATH下创建包 GOPATH=D:\gopro
```
进入GOPATH：cd D:\gopro
创建包路径： md src/mypkg/calc
初始化mod：go mod init mypkg/calc
在caclc下编辑包内容：

package calc
import (
	"fmt"
)
func init() {
	fmt.Println("this is calc init")
}

//func add(a, b int) int {
func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

```
 #第一步：引用包文件
 在其他非GOPATH路径下创建工程目录 如：D:\go\myPro\main.go
 引用代码
 ```
 package main //必须

import (
	"mypkg/calc"
	"fmt"
)

func init() {
	fmt.Println("this is main init")
}

func main() {
	a := calc.Add(10, 20)
	fmt.Println("a = ", a)

	fmt.Println("r = ", calc.Minus(10, 5))
}

 ```
 