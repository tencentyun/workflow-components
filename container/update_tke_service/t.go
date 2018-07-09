package main

import "fmt"

type A struct{}

func main() {
	var containers map[string]string
	fmt.Println(containers)
	fmt.Println(containers == nil)
	fmt.Printf("%p", &containers)

	var a []int64

	fmt.Printf("%p", &a)
	fmt.Println(a == nil)

	var x *A

	fmt.Printf("%p", x)
	fmt.Println(x == nil)

	var y *A

	fmt.Printf("%p", &y)
	fmt.Println(&y == nil)
}
