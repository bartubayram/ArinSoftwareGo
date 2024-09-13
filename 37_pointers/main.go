package main

import "fmt"

func main() {
	/* name := "Arin"
	fmt.Println(name)
	fmt.Println(&name) // name değişkeninin bulunduğu hafızadaki adresi... */
	//(& ---> address operator)

	/* 	fmt.Println(x)
	   	fmt.Println(&x)
	   	fmt.Println()

	   	fmt.Printf("%T,%v \n", x, x)
	   	fmt.Printf("%T,%v \n", &x, &x) */

	/* 	y := &x
	   	fmt.Println(y)
	   	fmt.Printf("%T,%v \n", y, y)

	   	z := &name
	   	fmt.Println(z)
	   	fmt.Printf("%T,%v \n", z, z) */
	/* 	x := 22

	   	fmt.Println(x)           // 22
	   	fmt.Println(&x)          // 22 adresi
	   	fmt.Println(*(&x))       // 22 (* operatörü adresteki değeri gösteri --> dereferencing)
	   	fmt.Println(&(*(&x)))    // 22 adresi
	   	fmt.Println(*(&(*(&x)))) // 22
	   	fmt.Println(3 * 5) */

	/* 	x1 := 10
	   	x2 := x1
	   	fmt.Println(x1, x2)
	   	fmt.Println(&x1, &x2)
	   	x1 = 5
	   	fmt.Println(x1, x2) */
	/* x1 := 10
	x2 := &x1
	fmt.Println(x1, x2)
	fmt.Println(x1, *x2)

	*x2 = 3
	fmt.Println(x1, x2)
	fmt.Println(x1, *x2)

	x3 := &x1
	 *x3 = 5
	fmt.Println(x1, *x2, *x3)
	*/
	x1 := [4]int{1, 10, 100, 1000} // array pass by value
	// x1 := []int{1, 10, 100, 1000} // pass by referance
	x2 := x1
	fmt.Println(x1, x2)
	fmt.Println(&x1[0], &x2[0])
	x2[0] = 3
	fmt.Println(x2)
	fmt.Println(x1)

}
