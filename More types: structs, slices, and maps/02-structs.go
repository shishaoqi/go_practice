package main

import "fmt"

// A struct is a collection of fields.
type Vertex struct {
	X int
	Y int
}

func main() {
	p := Vertex{1, 2}

	fmt.Println(p)

	// 普通点位符
	fmt.Printf("相应值的默认格式：%v\n", p)
	fmt.Printf("打印结构体时，会添加字段名：%+v\n", p)
	fmt.Printf("相应值的Go语法表示：%#v\n", p)
	fmt.Printf("相应值的类型的Go语法表示：%T\n", p)
	fmt.Printf("字面上的百分号，并非值的点位符：%%\n")
}
