package main

import "fmt"

func main() {

	num1 := 106
	str1 := string(num1)

	fmt.Printf("%v %T \n", num1, num1)

	fmt.Printf("%v %T", str1, str1)

	// x := 10
	// y := "10"
	// fmt.Println(x + y)

	// var x int16 = 128
	// var y int8

	// y = int8(x)

	// fmt.Println(y)

	// var x int8 = 10
	// var y int16 = 10
	// fmt.Println(x + int8(y))

	// x := 10
	// y := 10.0

	// fmt.Printf("%T %v\n", x, x)

	// fmt.Printf("%T %v \n", y, y) // %v -> değişken type'ını yazdırır // %T -> değişken değerini yazdırır
	// // Type Conversion -> type ()
	// fmt.Println(x + int(y))

	// // bakalım y değişkeninin type değişmiş mi ? -> hayır
	// fmt.Printf("%T %v \n", y, y)

	// // özetle ->
	// /* Type Conversion ile biz yeni bir değer oluşturuyoruz.Ancak değişkeninin asıl veri tipini ve değerini değiştirmiyoruz.

}
