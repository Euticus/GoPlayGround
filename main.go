package main

import "fmt"

var x int = 42
var y bool = true

//var z string = "James Bond"

type pineapple bool

var fruit pineapple = true

func main() {

	s := fmt.Sprintf("%v\t%v\t%T\t", x, y, fruit)
	fmt.Println(s)
}
