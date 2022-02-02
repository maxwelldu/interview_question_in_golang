package main

import "fmt"

func Foo(x interface{}) {
	if x== nil {
		fmt.Println("emptyinterface")
		return
	}
	fmt.Println("non-emptyinterface")
}
func main() {
	var x *int = nil
	Foo(x)
	// non-empty interface
}
