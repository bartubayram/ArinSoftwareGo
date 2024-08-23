// addition  + -> 15 + 10, 10,15 operand, + operator
package main

import "fmt"

func main() {
	// x, y := 15, 10
	// fmt.Printf("%T %v \n", x, x)
	// fmt.Printf("%T %v \n", y, y)

	// fmt.Printf("%T %v \n", (x + y), (x + y))
	// fmt.Printf("%T %v \n", (x - y), (x - y))
	// fmt.Printf("%T %v \n", (x * y), (x * y))
	// fmt.Printf("%T %v \n", (9.0 / 3), (9.0 / 3))
	// aynı veri tiplerinin işlemi sonucunun tipi veri tipleriyle aynı olur.
	//int / int = int
	// fmt.Printf("%T %v \n", (x % y), (x / y))
	// z := 5.0 / 2
	// fmt.Printf("%T %v \n", z, z)

	// Increment ++ ,Decrement  -- , POSTFIX
	x := 10
	fmt.Println(x)
	// x = x + 1
	// fmt.Println(x)
	// x++
	// fmt.Println(x)

	// // bazı programlama dillerinde hem postfix hem de prefix oalrak kullanılmaktadır fakat
	// //golang'ca sadece postfix olarak kullanılır

	// fmt.Println(x++) // bu hatalı bir yazımdır.	sebebi ise go prog dilince icrement dve decrement birere statementtir bi satıda sadece bi statement olacagından bunu yapamayız (sonucta println de bi statementtir)

	// Decrement
	x = x - 1
	fmt.Println(x)

	x--
	fmt.Println(x)

}
