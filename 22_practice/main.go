// 1 ile 10 arasındaki sayıları if yapıssıyls tek-çift olarak yazdırınız.
/* package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "çift sayi")
		} else {
			fmt.Println(i, "tek sayi")
		}
	}
} */

// 2-) for yapısını kullanarak Go'da olmayan While döngüsüne örnek veriniz.package main

/* package main

import "fmt"

func main() {
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}
}
*/

//3-)  switch fallthrough ifadesini açıklayınız.
// -> Normal switch yapısında doğru case ile karşılaşıldığında diğer case'ler test edilmez ama diğer case'lerin de test edilmesini istediğimiz durumda fallthrough kullanılır
/* package main

import "fmt"

func main() {
	switch x := 25; {
	case x < 20:
		fmt.Printf("%d 20'den küçüktür\n", x)
		fallthrough

	case x < 50:
		fmt.Printf("%d 50'den küçüktür\n", x)
		fallthrough
	case x < 100:
		fmt.Printf("%d 100'den küçüktür\n", x)
		fallthrough

	case x < 200:
		fmt.Printf("%d 200'den küçüktür\n", x)

	}

} */

// 4-) Aşağıdaki if döngüsünü daha idiomatic hale getiriniz. -> (Daha kısa hale getirmek anlamındadır)

/* package main

import "fmt"

func main() {
	if x := 20; x%2 == 0 {
		fmt.Println(x, "cifttir")
	} else {
		fmt.Println(x, "tektir")
	}
} */

/* package main

import "fmt"

func main() {
	x := 20
	if x%2 == 0 {
		fmt.Println(x, "cifttir")
		return
	}
	fmt.Println(x, "tektir")
} */

// 5-) 1- 50 arasındaki asal sayıları yazdır
package main

import "fmt"

func main() {
	var x, y int

	for x = 2; x < 50; x++ {
		for y = 2; y < (x / y); y++ {
			if x%y == 0 {
				break
			}
		}
		if y > (x / y) {
			fmt.Printf("%d asal sayıdır\n ", x)
		}
	}
}
