// 1-) x = x-10 vs x-=10

// package main

// import "fmt"

// func main() {
// 	x := 50
// 	x = x - 10 // assignment statement
// 	// x-10 -> statement, bunu x'e atnaması assignment işlemidir
// 	x -= 10 // assignment operation
// 	// x-=10 burda atmaayla işlem aynı anda yapılmış

// 	fmt.Printf("%v,%T\n", x, x)

// }

// 2-) K = F-32/1.8 + 273  => -40 F kaç K derecedir?
// package main

// import "fmt"

// func main() {
// 	F := -40

// 	K := float64((F-32))/1.8 + 273

// 	fmt.Println(F, "Fahrenant ", K, "Dereceye eşittir")
// 	fmt.Printf("%v,%T \n", K, K)
// 	fmt.Printf("%v,%T \n", F, F)

// }

// 3-) Çıktısı Ne Olur ?

/*
	package main

import "fmt"

	func main() {
		age := 40
		const myAge = age // Hata alırız

		fmt.Printf("%v,%T \n", myAge, myAge)

		// const -> coompile time
		// var -> run time
	}
*/

// 4-) Sabitlerde Shadowing kavramı çalışır mı ?
// package main

// import "fmt"

// var x = 14

// func main() {
// 	var x = 24
// 	fmt.Printf("%v,%T \n", x, x)
// 	write()
// }
// func write() {
// 	fmt.Printf("%v,%T \n", x, x)
// }

// 5-) const c = 4 , cont y = 5.4, x+y ?

// package main

// import "fmt"

// func main() {

// 	const c = 4   // typless
// 	const y = 5.4 // typless

// 	fmt.Printf("%v,%T \n", c+y, c+y) // c+y'nin veir tipi float64
// 	fmt.Printf("%v,%T \n", c, c)     // x'in veri tipi int
// 	fmt.Printf("%v,%T \n", y, y)

// }

// 6-) const x float64 = 6.4, y := 4 + x,  y ? -> 10.4 :D
// package main

// import "fmt"

// func main() {
// 	const x float64 = 6.4
// 	y := 4 + x

// 	fmt.Printf("%v,%T \n", x, x)
// 	fmt.Printf("%v,%T \n", y, y)
// }
