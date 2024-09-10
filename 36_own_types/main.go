/* package main

// mile float64 üzerine kurulmuş bir veri tipi..!
type mile float64 // ---> underlying veri tipiyle (float64) defined(mile) birbirleriden fakrlıdır ancak birbirlerine dönüştürülebilriler.
type kilometer float64
type mystring string

func main() {
	/* f1 := float64(4.4) // veri tipi float64
	fmt.Printf("%T,%v", f1, f1)

	fmt.Println()
	fmt.Println(3.2 + 4.4) // 7.6
	//fmt.Println(f1 + m1)   // HATA (farklı veri tipleri toplanmaz..)
	// ancak bunları dönüştürebiliriz(type conversion) :
	fmt.Println(m1 + mile(f1))

	fmt.Printf("%T,%v \n", (m1 + mile(f1)), m1+mile(f1))

	fmt.Printf("%T,%v", (f1 + float64(m1)), f1+float64(m1))
*/
/* 	var m1 mile
   	m1 = 3.2 // veri tipi mile'dır. */
/* fmt.Println(m1)
fmt.Printf("%T,%v", m1, m1)
fmt.Println() */

/* m2 := mile(4.6) */

/* k1 := kilometer(7.8)
fmt.Printf("%T,%v", k1, k1)
fmt.Println(f1 + k1) // hata (farklı veri tipleri !!!) */
/* 	fmt.Println(m1 + m2) // her ikisi de mile veri tipincen olduğundan hata almayız!!!
   	fmt.Println(m1 * m2)
   	fmt.Println(m1 + m2 + 2.1)
   	fmt.Println(m1 + 2.1) */

/* 	mystring := "Arin"

   	// m1 := mile(4.2)
   	fmt.Println(strings.ToUpper(mystring)) // mystring string veri tipi üzerine kurulu oolduğundan strings fonksiyonlarını kullanabiliriz.
*/ // fmt.Println(strings.ToUpper(m1))

package main

import "fmt"

type mile float64
type kilometer float64

func main() {
	// m1 =10 k1 = ?
	m1 := mile(10)
	k1 := toKilometer(m1)
	fmt.Println(k1)

	// k1 = 10 ,m1 =?
	k2 := kilometer(10)
	m2 := toMile(k2)
	fmt.Println(m2)
}

func toKilometer(m mile) kilometer {
	return kilometer(m * 1.6)

}
func toMile(k kilometer) mile {
	return mile(k * 0.62)
}
