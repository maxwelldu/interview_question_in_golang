package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("fatal")
		}
	}()
	defer func() {
		panic("defer panic")
	}()
	// 函数执行完之前会执行defer; panic("panic")； panic之前会先执行defer函数
	panic("panic")
	// defer panic
}