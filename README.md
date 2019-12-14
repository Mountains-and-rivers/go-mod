# go-mod 使用
v1.13 默认开启默认开启go mod

说明：1.13不允许在 GOPATH 下使用 go mod


# 第一步：在GOPATH下创建包 GOPATH=D:\gopro

进入GOPATH：```cd D:\gopro  ```  
创建包路径：``` md src/mypkg/calc  ```  
初始化mod：```go mod init mypkg/calc  ```  
在caclc下编辑包内容：
calc.go
```
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
 # 第二步：引用包文件
 在其他非GOPATH路径下创建工程目录 如：D:\go\myPro\main.go  
 引用代码 main.go  
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
 
 # 附件1 自定义包
 ![Image text](https://github.com/Mountains-and-rivers/go-mod/blob/master/image/1.png)
 
 # 附件2 引用自定义包
 ![Image text](https://github.com/Mountains-and-rivers/go-mod/blob/master/image/2.png)
 
 
 
 
 
 
 
 
 
 
 
