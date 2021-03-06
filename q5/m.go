package main
import (
	"fmt"
	"runtime"
)
func main() {
	runtime.GOMAXPROCS(1)
	var (
		intChan = make(chan int, 1)
		stringChan = make(chan string, 1)
	)
	intChan <- 1
	stringChan <- "hello"
	select {
	case value := <-intChan:
		fmt.Println(value)
	case value := <-stringChan:
		panic(value)
	}
	//有时候1，有时候panic
}