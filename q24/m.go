package main

import "fmt"

func main() {
	type MyInt1 int //重新申明一个类型
	type MyInt2 = int //别名，一模一样
	var i int = 9
	var i1 MyInt1 = i
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
}
