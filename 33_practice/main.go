// 1-) 5 string elemandan oluşan bir array ve 5 int elemandan oluşan slice oluşturup index numaralarıyla bilrikte gösterelim
/* package main

import "fmt"

func main() {
	myArr := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(myArr)
	for index, value := range myArr {
		fmt.Println(index, "-->", value)
	}
	fmt.Println()
	mySlc := []int{6, 7, 8, 9, 10}
	fmt.Println(mySlc)
	for i, v := range mySlc {
		fmt.Println(i, "-->", v)
	}

} */

// 2-) [0,1,2,3,4,5,6,7,8] slice'dan  [0,1,2,3],[4,5,6],[6,7,8],[2,3,6,7] slice'larını oluşturun
/* package main

import "fmt"

func main() {
	mySlc := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(mySlc)

	fmt.Println(mySlc[:4])
	fmt.Println(mySlc[4:7])
	fmt.Println(mySlc[6:])

	mySlc1 := mySlc[2:4]
	mySlc2 := mySlc[6:8]
	mySlc3 := append(mySlc1, mySlc2...)
	fmt.Println(mySlc3)
} */

// 3-) slice'lar için copy metodunu ve assign(=) ile farkını açıklayınız

/*  package main

import "fmt"

func main() { */
/* mySlc := []int{1, 2, 3}

mySlc2 := make([]int, 2)

fmt.Println(mySlc)
fmt.Println(mySlc2)
copy(mySlc2, mySlc) // ---> mySlc2'ye  mySlc'yi kopyalar..
fmt.Println(mySlc)
fmt.Println(mySlc2)
mySlc[0] = 100 // ---> mySlc değerine etki etmez çünkğ copy yaptık
fmt.Println(mySlc)
fmt.Println(mySlc2) */
/* 	OUTPUT
   	[1 2 3]
   	[0 0]
   	[1 2 3]
   	[1 2]
   	[100 2 3]
   	[1 2] */

/* 	mySlc := []int{1, 2, 3}

mySlc2 := make([]int, 2)

fmt.Println(mySlc)
fmt.Println(mySlc2)
mySlc2 = mySlc      // ---> assign (=) -> PASS BY REFERANCE
fmt.Println(mySlc2) // --> burda mySlc2'nin 2 indeksli olması önemli değil direkt mySlc'yi mySlc2'ye eşitler
mySlc[0] = 100      //
fmt.Println(mySlc)
fmt.Println(mySlc2) */

/*
		 	OUTPUT
			[1 2 3]
			[0 0]
			[1 2 3] // ***important
			[100 2 3]
			[100 2 3]

}*/

// 4-) map gösterimi ile yazar ve yazarlara ait kitapların isimlerini for range ile gösteriniz.
package main

import "fmt"

func main() {
	myMap := map[string][]string{ // önemli bi kullanım dikkat et key değeri string value'lar ise   slice !!!!!

		"yasar kemal":     []string{"ince memed1,", "Yer demir gök bakır"},
		"sabahattin ali":  []string{"kuyucaklı yusuf", "kürk montolu madonna"},
		"Haruki Murakami": []string{"1084", "Dans dans dans", "Kumandanı öldürmek"},
	}
	fmt.Println(myMap)
	fmt.Println()
	fmt.Println(myMap["sabahattin ali"])
	fmt.Println(myMap["Haruki Murakami"][0])
	fmt.Println()
	fmt.Println("for range kullanımı--:")
	fmt.Println()
	for key, value := range myMap {
		fmt.Println("Yazarimiz :", key)
		fmt.Println("bazi kitaplari asagidadir:")
		for i, v := range value {
			fmt.Println("\t", i+1, ".", v)
		}
		fmt.Println()
	}
}
