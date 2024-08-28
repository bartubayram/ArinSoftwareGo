package main

import "fmt"

func main() {
	//x, y := "a", "b" // 07,98 diye decimal olarak değerlendirir.
	// fmt.Printf("%T,%v \n", x == y, x == y)
	// fmt.Printf("%T,%v \n", x != y, x != y)
	// fmt.Printf("%T,%v \n", x < y, x < y)
	// fmt.Printf("%T,%v \n", x > y, x > y)
	// fmt.Printf("%T,%v \n", x >= y, x >= y)
	// fmt.Printf("%T,%v \n", x <= y, x <= y)

	x, y := 15, 20
	// set1 := (x == y)
	set2 := (x < y)
	set3 := true

	// fmt.Printf("%T,%v \n", set1, set1)
	// fmt.Printf("%T,%v \n", set2, set2)
	// fmt.Printf("%T,%v \n", (set1 && set2), (set1 && set2))
	// fmt.Printf("%v \n", (set3 || set2))
	fmt.Printf("%v \n", (!set3 || set2))

}
