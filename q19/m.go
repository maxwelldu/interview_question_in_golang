package main
import "fmt"
func GetValue(m map[int]string, id int) (string, bool) {
	if _, exist := m[id]; exist {
		return " 存在数据 " , true
	}
	return nil, false
}
func main() {
	intMap := map[int]string{
		1: "a",
		2: "bb",
		3: "ccc",
	}
	v, err := GetValue(intMap, 3)
	fmt.Println(v, err)
}