package main

import "fmt"

func main() {
	/* var x, y, sum int
	x = 5
	y = 10
	sum = x + y
	fmt.Printf(" %d ve %d toplami = %d", x, y, sum)*/
	// Fonksiyonlar yardımıyla programları daha modüler hale getiririz.(Moduler Programming)
	//Readable code
	// From Complex to Simple
	// fmt.Println(sum(56, 23))
	//func funcName(params) return type { code }
	/* merhaba() */
	// sum2(1, 2)

	// return vs print
	/* z := sum(5, 10)
	fmt.Println(z)
	sum2(6, 11) */
	merhaba2("Bartu", 6)

	// FUNC İSİMLENDİRME KURALLARI:
	/*
		1.İlk karakter harf olmalı.
		2.cameCase -- mySum,myBestFunc
		3.paket dışında kullanacaksak ilk harf büyük olmalı
		
	*/
}
func sum(x, y int) int {
	return x + y
}
func sum2(x, y int) {
	fmt.Println(x + y)
}

func merhaba() {
	fmt.Println("my name is Bartu")
}
func merhaba2(name string, age int) {
	fmt.Printf("Adım %s, yaşım %d ", name, age)
}
