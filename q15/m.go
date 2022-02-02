package main

import "fmt"

func main() {
	// new出来的是一个切片类型的指针；正确的方式是make出来一个切片才可以操作
	list := new([]int)
	//list := make([]int,0)
	list = append(list, 1)
	fmt.Println(list)
}