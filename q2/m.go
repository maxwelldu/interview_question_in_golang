package main

// 以下代码没问题

type student struct {
	Name string
	Age int
}
func paseStudent() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	for _,stu := range stus {
		m[stu.Name] = &stu
	}
}
func main() {
	paseStudent()
}