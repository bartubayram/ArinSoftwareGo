// if grade == 5 {
// 	fmt.Println("Pek iyi")

// } else if grade == 4 {
// 	fmt.Println(" iyi")

// } else if grade == 3 {
// 	fmt.Println("orta ")
// } else if grade == 2 {
// 	fmt.Println("geçer ")

// } else if grade == 1 {
// 	fmt.Println("Başarısız ")

//	} else {
//		fmt.Println("Geçersiz değer")
//	}
package main

import "fmt"

func main() {

	/*
		switch grade := -4; grade {
		case 5: // if grade == 5 {	fmt.Println("Pek iyi") }
			fmt.Println("Pek iyi")
		case 4:
			fmt.Println(" iyi")
			y := 100
			fmt.Println(y)
		case 3:
			fmt.Println("orta ")
		default:
			fmt.Println("Geçersiz Değer")
		case 2:
			fmt.Println("geçer ")
		case 1:
			fmt.Println("başarısız ")

		}*/
	switch {
	case false:
		fmt.Println("Bu yazdığımız konsolda görünmez")
	case true:
		{
			fmt.Println("Ekrana yazdırılacaktır.")
		}
	}
}

/* switch grade := BURDAKİ DEĞİŞKEN TİPİYLE; grade {
	case BURALARDAKİ DEĞİŞKEN TİPİ AYNI OLMALIDIR.:
}*/

// Program case doğrulandıktan itibaren diğer case'lere uğramaz.
