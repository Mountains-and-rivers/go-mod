# go-mod 使用
goang v1.13 默认开启go mod

说明：v1.13不允许在 GOPATH 下使用 go mod (自定义包)


# 第一步：在GOPATH下创建包 （GOPATH=D:\gopro)

进入GOPATH：```cd D:\gopro\src  ```  
创建包路径：``` md mypkg/calc  ```  
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

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

```
 # 第二步：引用包文件
 在其他非GOPATH路径下创建工程目录 如：D:\go\myPro  
 在D:\go\myPro下创建主文件：main.go  
 main.go 引用自定义包：mypkg/calc
 ```
 package main

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
  # 说明：
  在mod模式下 会自动从定义的包路径导入包
 
 # 附件1 自定义包
 ![Image text](https://github.com/Mountains-and-rivers/go-mod/blob/master/image/1.png)
 
 # 附件2 引用自定义包
 ![Image text](https://github.com/Mountains-and-rivers/go-mod/blob/master/image/2.png)
 
 
 
 
 
 
 
 
 
 
 
