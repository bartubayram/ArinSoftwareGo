package main

import "fmt"

func main() {

	// if <boolean expression> {code}else {code}

	/*x := 27
	if x%2 == 0 {
		fmt.Println("x çift sayidir")
	} else {
		fmt.Println("not")
	}
	if x%2 != 0 {
		fmt.Println("x tek sayidir")
	}
	*/
	// if !false {
	// 	fmt.Println("!false ")
	// }

	/*
		x := 27
		if x%2 == 0 {
			fmt.Println("Cift sayidir.")
		} else {
			fmt.Println("tek sayi")
		}
	*/
	x := 2
	if x < 0 {
		fmt.Println(x, "negatif sayidir")
	} else if x%2 == 0 {

		fmt.Println(x, "cift sayidir")

	} else {
		fmt.Println(x, "tek sayidir.")
	}
	fmt.Println(x)
	// if <boolean expression>{code } else if <boolean expression> else {code}
}
