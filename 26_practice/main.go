//
// 	1-) ** İKİ RAKAM ARASINDA TOPLAMA ÇIKARMA VE ÇARPMA İŞLEMİNİN
// YAPILDIĞI BİR FONKSİYON YAZINIZ.

/* package main

import "fmt"

func main() {
	sum, dif, prof := calculation(10, 5)
	fmt.Println(sum, ",", dif, ",", prof)

}
func calculation(num1, num2 int) (int, int, int) {
	sum := num1 + num2
	dif := num1 - num2
	prod := num1 * num2

	return sum, dif, prod

}  */
// ----------------------------------------------------------------

/*
	2-) ** KULLANICI TARAFINDAN GİRİLEN NOTA GÖRE GEÇTİNİZ YA DA KALDINIZ

GERİ DÖÜNŞÜNÜ YAZDIRINIZ
*/
/* package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Lütfen aldiginiz notu giriniz..:")
	grade, _ := getGrade()

	var result string

	if grade >= 50 {
		result = "gectin"
	} else {
		result = "kaldin"
	}
	fmt.Println(result)

}
func getGrade() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	input = strings.TrimSpace(input) // girilen değerin sağ ve soolundaki boşluklardan kurtul.
	num, err := strconv.Atoi(input)  //string’i integer’a çevirmeye çalışır
	if err != nil {
		fmt.Println(err)
	}
	return num, nil
} */
// ----------------------------------------------------------------

/*
	3-) ** 1 İLE 100 ARASINDAKİ BİR SAYIYI TAHMİN ETME UYGULAMASI YAZINIZ.

TOPLAM TAHMİN HAKKINIZ 10 OLSUN
*/
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	target := numRand(1, 100)
	fmt.Println("1 ile arasındaki sayiyi bulmaya calisin")

	reader := bufio.NewReader(os.Stdin)
	for attempts := 0; attempts < 10; attempts++ {
		fmt.Println(10-attempts, "hakkiniz kaldi")
		fmt.Println("lütfen tahmin yapiniz:")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}

		if num > target {
			fmt.Println("tahmininiz daha büyük daha kucuk bir sayi girin .")
		} else if num < target {
			fmt.Println("tahmininiz daha kucuk daha buyuk bir sayi girin.")
		} else {
			fmt.Println("Dogru Tahmin, hedef sayi", target, "idi..", attempts, " seferde bulundu.")
			break
		}
	}

}
func numRand(min, max int) int {
	rand.Seed(time.Now().Unix()) // her çalıştırdığımızda yeni random sayı verir.
	return rand.Intn(max-min) + min
}
