/*
	package main

import (

	"errors"
	"fmt"

)

	func main() {
		result, err := evenNum(11)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Girilen sayi-->", result)
		}

}

	func evenNum(num int) (int, error) {
		if num%2 != 0 {
			return 0, errors.New("Hata: Cift Sayi girmediniz.")
		}
		return num, nil
	}
*/

// karekök alabilme kodu -> negatif sayıların karekökünü alamayız.
/* package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	result, err := sRoot(4)
	if err != nil {
		fmt.Println("Error Message :", err)

	} else {
		fmt.Println("karekökü : ", result)

	}

}

func sRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("negatif sayilarin karekökü alinamaz.")
	}
	return math.Sqrt(num), nil
}
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Dosyamız:", file)
	}
}
