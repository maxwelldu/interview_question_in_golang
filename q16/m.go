package main
import "fmt"
func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	// 参数类型不对
	s1 = append(s1, s2)
	//s1 = append(s1, s2...)
	fmt.Println(s1)
}