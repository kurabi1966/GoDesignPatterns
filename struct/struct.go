package main

import "fmt"

type rectangle struct {
	length  int
	breadth int
	color   string
	x       []int
}

func (r *rectangle) print() {
	fmt.Println(*r)
}

func main() {
	// var rect1 = &rectangle{10, 20, "Green"} // Can't skip any value
	// fmt.Println(rect1)

	// var rect2 = &rectangle{}
	// rect2.length = 10
	// rect2.color = "Red"
	// fmt.Println(rect2) // breadth skipped

	// var rect3 = &rectangle{}
	// (*rect3).breadth = 10
	// (*rect3).color = "Blue"
	// // fmt.Println(rect3) // length skipped

	// rect3.print()

	rect4 := rectangle{breadth: 20, color: "Yellow", length: 55, x: []int{1, 2, 3}}

	// fmt.Println(rect4)
	rect4.print()
}
