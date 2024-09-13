/* // 1-) underlying type 'int' olacak şekilde kendi veri tipinizi
// oluşturun veri tipi ve değerini 10 yazdır
package main

import "fmt"

type myType int

func main() {
	x := 10
	fmt.Printf("%T,%v \n", x, x)

	y := myType(10)
	fmt.Printf("%T,%v", y, y)
	// output --> int,10
	//main.myType,10%
}
*/

// 2-) Başlangıç değeri 10 olan bir X değişkeni oluşturun sonrasında
// bulunduğu adres üzerindne y değişkenini tanımlayıp x değişkenini 20 yapınız..
/* package main

import "fmt"

type myType int

func main() {
	x := 10
	fmt.Println(x)
	y := &x
	fmt.Println(y)
	*y = 20
	fmt.Println(*y)
	fmt.Println(x)
} */

// 3-) Underlying type struct olan Rectangle type oluşturunuz.
// display, area, circumference metodlarını yazınız.
/* package main

import "fmt"

type rectangle struct {
	a, b int
}

func (r rectangle) display() {
	fmt.Println("Kenar Uzunlukları ..:", r.a, "ve ", r.b)
}
func (r rectangle) area() int {
	return (r.a * r.b)
}
func (r rectangle) circumference() int {
	return 2 * (r.a + r.b)
}
func main() {
*/
/* 	r1 := rectangle{
	a: 10,
	b: 20,
} */
/* r1 := rectangle{5, 20}
	fmt.Println(r1)
	r1.display()
	fmt.Println("Alanı ..:", r1.area())
	fmt.Println("Çeversi ..:", r1.circumference())
	r1.circumference()

} */

// 4-) family aile bireyleri şeklinde veri tipi oluşturalım.,underlying veri tipi struct. Aile bireylerinin hepsini farklı
// şekilde tanımlayınız . Sonrasında for döngüsünde yazdırınız . field age,name, isMarried field.

package main

import "fmt"

type family struct {
	age       int
	name      string
	isMarried bool
}

func showFamily() []family {
	family1 := family{
		name:      "x",
		age:       4,
		isMarried: false,
	}
	family2 := family{
		name: "y",
		age:  3,
	}
	family3 := family{
		21,
		"bartu",
		false,
	}

	var family4 family
	family4.name = "gamze"
	family4.age = 12
	family4.isMarried = false

	return []family{family1, family2, family3, family4}
}
func main() {
	families := showFamily()
	for i := 0; i < len(families); i++ {
		fmt.Printf("%v,%T\n", families[i], families[i])

	}
}
