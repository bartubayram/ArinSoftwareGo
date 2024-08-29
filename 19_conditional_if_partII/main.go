package main

import "fmt"

func main() {
	/* x := 10
	if x := -5; x < 0 {
		fmt.Println(x, "negatifa sayi")
	} else if x%2 == 0 {

		fmt.Println(x, "cift sayidir")

	} else {
		fmt.Println(x, "tek sayidir.")
	}
	fmt.Println(x) */
	if x := 25; x < 0 {
		fmt.Println(x, "negatif tam sayidir.");
		fmt.Println(x, "Bartu");

	} else {
		if x%2 == 0 {
			fmt.Println(x, "cift sayidir")

		} else {
			fmt.Println(x, "tek sayidir")

		}
	}
}
