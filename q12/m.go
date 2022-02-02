package main

func GetValue() int {
	return 1
}

func main() {
	i := GetValue()
	//Invalid type switch guard: i.(type) (non-interface type int on left)
	//类型断言需要是接口类型才可以
	switch i.(type) {
		case int:
			println("int")
		case string:
			println("string")
		case interface{}:
			println("interface")
		default:
			println("unknown")
	}
}
