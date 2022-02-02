package main

import (
	"fmt"
	"sync"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	// 使用之前没有初始化，需要加上以下3行对map进行初始化才行
	//if ua.ages == nil {
	//	ua.ages = make(map[string]int)
	//}
	ua.ages[name] = age
}
func ( ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	ua := &UserAges{}
	ua.Add("max", 18)
	fmt.Println(ua.Get("max"))
}
