package main
import"fmt"

type T1 struct { }
func (t T1) m1() {
	fmt.Println("T1.m1")
}

type T2 = T1
type MyStruct struct {
	T1
	T2
}

func main() {
	my := MyStruct{}
	my.m1()
	//./m.go:17:4: ambiguous selector my.m1
}
