/* package main

import "fmt"

func main() {
	merhaba("Bartu", 7) // Argument Function Run
}

func merhaba(name string, age int) { // parametre Function write
	fmt.Printf("Adım %s, yaşım %d", name, age)
}
*/

/* package main

import "fmt"

func main() {
	z := result(50)
	fmt.Println(z)
}
func result(grade int) string {
	if grade >= 50 {
		return "geçtiniz"
	} else {
		return "kaldiniz"
	}
}
*/

/*
	package main

import "fmt"

	func main() {
		z := result(40)
		fmt.Println(z)
	}

	func result(grade int) string {
		if grade >= 50 {
			return "geçtiniz"
		}
		return "kaldiniz"
		fmt.Println("Hatali yazim ")

}
*/
/* package main

import "fmt"

func main() {
	merhaba("Arin", 6)

	name := "Elis"
	age := 4
	fmt.Printf("Adım %s, yaşım %d", name, age)

}
func merhaba(name string, age int) {
	fmt.Printf("Adım %s, yaşım %d\n", name, age)
} */

// Programlama dillerinde ya kendi fonksiyonlarımızı yazarız ya da zaten bizim adımıza dilin çekirdeğinde daha önce gerkli olan birçok fonksiyon yazılmıştır bunları kulanırız.Belli işlemleri yapmak için farklı paketlerdeki farklı işlemleri kullanırız. Farklı paketlerdeki fonksiyonlara Method diyoruz
// Methodlar için farklı veri tiplerine ait olan fonksiyonlar da diyebiliriz.
/* package main

import (
	"fmt"
	"time"
)

func main() {

	var x int = 10
	fmt.Println(x)
	var moment time.Time = time.Now() // Now --->method
	fmt.Println(moment)

} */
/* package main

import (
	"bufio"
	"fmt"
	"os"
)

// tüm işlmeler değişkenismi.Paketismi diye kalıplandırılır.
func main() {
	fmt.Print("Lütfen sınav sonucunuzu girin: ")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n') // _ blank identifier
	fmt.Println(value)
	//_ blank identifier ---> Oluşturulan değişkenler kesinlikle kullanılmalıdır.Kullanmak istemediğimiz değişkenler de olabilir bazı durumlarda. Bu tarz durumlarda blank identifier kullanırız.
} */
package main

import "fmt"

func main() {
	z, y := bölme(104, 5)
	fmt.Println(z, "-", y)

}
func bölme(bölünen, bölen int) (bölüm int, kalan int) {

	return bölünen / bölen, bölünen % bölen
}
