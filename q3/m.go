package main

import (
	"fmt"
	"runtime"
	"sync"
)

// A输出的都是A10；
// B输出的是无序的B0-B9
func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
				fmt.Println("A: ", i)
				wg.Done()
			}()
	}
	for i:= 0; i < 10; i++ {
		go func(i int) {
				fmt.Println("B: ", i)
				wg.Done()
			}(i)
	}
	wg.Wait()
}