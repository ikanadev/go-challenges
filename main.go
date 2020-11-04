package main

import "fmt"

type Person struct {
	Age  int
	Name string
}

func main() {
	p := Person{Age: 5}
	fmt.Printf("%+v", p)
	// goChall1()
	// goChall2()
	// goChall3()
	// goChall4()
}
