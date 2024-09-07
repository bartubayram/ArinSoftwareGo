package main

import "fmt"

func main() {
	// aynı veri tipine ait birden fazla değişken varsa liste içerisinde tutarız bu liste yapısına array denir
	/* 	city1 := "istanbul"
	   	city2 := "roma"

	   	city3 := "tahran"
	   	city4 := "belgrad"
	   	fmt.Println(city1, city2, city3, city4) */
	// cities := [4]string{"istanbul", "roma", "tahran", "belgrad"}

	/*
		 	cities := [...]string{"istanbul", "roma", "tahran", "belgrad"}

			fmt.Println(cities)
			fmt.Println(cities[0]) // zero base index
			fmt.Println(cities[3])
			fmt.Println(len(cities))
	*/
	// başlangıç değeri kendinden olan array nasıl oluştururuz?
	/* var myArray [5]int
	fmt.Println(myArray)
	myArray[0] = 13
	fmt.Println(myArray)
	myArray[len(myArray)-1] = 200
	fmt.Println(myArray) */

	/* 	var myArr [4]int
	   	var myArr2 [5]int */

	/* 	if myArr == myArr2 {
		fmt.Println("message")
	} */
	// fmt.Println(myArr, "\n", myArr2)

	/* cities := [4]string{"istanbul", "roma", "tahran", "belgrad"}
	for i := 0; i < len(cities); i++ {
		fmt.Println(i, cities[i])
	}
	cities[0] = "Ankara"
	fmt.Println()
	for i := 0; i < len(cities); i++ {
		fmt.Println(i, cities[i])
	} */
	/*
		myArr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

		// Tekrardan aynı array'e atacağız.
		myArr = mySquare(myArr) // First Class Functions
		fmt.Println(myArr) */

	// FOR- RANGE
	// for  index, value := range myArr{}
	cities := [4]string{"istanbul", "roma", "tahran", "belgrad"}
	/* 	for index, city := range cities { // -> range yapsıı burda dizimizin son elemanına kadar olan elemanları yazdırmamızı sağlar

		fmt.Println(index, city)
		fmt.Println(city) ---> Hata Alırız
	} */
	for _, city := range cities { //Blank Identifier ..!
		fmt.Println(city)
	}

}

/* func mySquare(arr [10]int) [10]int {
for i := 0; i < len(arr); i++ {
	arr[i] = arr[i] * arr[i]
}
return arr
}*/
