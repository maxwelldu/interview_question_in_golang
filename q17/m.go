package main
import "fmt"
func main() {
	sn1 := struct {
		age int
		name string
	}{age: 11, name: "qq"}

	sn2 := struct {
		age int
		name string
	}{age: 11, name: "qq"}
	// 这两个是相等的
	if sn1 == sn2 {
		fmt.Println("sn1== sn2")
	}

	sm1 := struct {
		age int
		m map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	sm2 := struct {
		age int
		m map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	// 结构体中有map类型，无法比较
	//./m.go:28:9: invalid operation: sm1 == sm2 (struct containing map[string]string cannot be compared)
	if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}
}