package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
	d = iota
)

const (
	e = d + iota
	f = iota
	g = b
)

func main() {

	fmt.Println(a, b, c, d, e, f, g)
}
