package main

import "fmt"

func main() {
	/* underArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(underArray)

	mySlc := underArray[2:5] // -> underArray'deki indeks numarlarını(aralık olarak-> 2-3-4. indekleri yazdırır.) göstermektedir.
	fmt.Println(mySlc)
	mySlc2 := underArray[:6]
	fmt.Println(mySlc2)
	mySlc3 := underArray[3:]
	fmt.Println(mySlc3)
	mySlc4 := underArray[:]
	fmt.Println(mySlc4)

	fmt.Println(mySlc)
	mySlc[0] = 100
	fmt.Println(mySlc) */

	/* 	mySlc := []int{1, 2, 3} // mySliceUnderArray
	   	fmt.Println(mySlc) */

	/* 	mySlc = append(mySlc, 4, 5)
	   	fmt.Println(mySlc) */

	/* mySlc2 := append(mySlc, 4, 5, 6) // mySlice2UnderArray -> aynı underarray üzerinden işlem yapmıyor çünkü ekleme yapınca aynı array'e ait olmus olmuyor(array'lere ekleme çıkarma yapılamaz.)
	fmt.Println(mySlc2)
	mySlc[0] = 100
	fmt.Println(mySlc)
	fmt.Println(mySlc2) */

	/* mySlc := []int{1, 2, 3}
	mySlc2 := []int{4, 5}
	fmt.Println(mySlc)
	mySlc = append(mySlc, mySlc2...)
	fmt.Println(mySlc) */

	// mySlc := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// mySlc = mySlc[2:] // ilk n elamanı silme için [n:]
	/* 	fmt.Println(mySlc)
	 */
	/* mySlc = mySlc[:len(mySlc)-3] // son n elemanı silme için [:len(mySlc)-(n)]
	fmt.Println(mySlc) */

	/* mySlc = mySlc[2:]
	mySlc = mySlc[:len(mySlc)-3]
	fmt.Println(mySlc)
	*/

	/* 	var myArr [4]int
	   	fmt.Println(myArr)

	   	var mySlc []int
	   	fmt.Println(mySlc)

	   	var mySlc2 []int
	   	mySlc2 = make([]int, 4) // Zero değerler Slice Elemanlarına ait olan
	   	fmt.Println(mySlc2)

	   	var mySlc3 []bool
	   	mySlc3 = make([]bool, 2) // Zero değerler Slice Elemanlarına ait olan
	   	fmt.Println(mySlc3) */

	var mySlc4 []int
	fmt.Printf("%#v", mySlc4) // declare ettik ama oluşturmadık yani bu slice yok
	fmt.Println()             // for line
	mySlc5 := make([]int, 0)  // declare ettik ve oluşturduk bu slice var ama elemanı yok yani boş
	fmt.Printf("%#v", mySlc5)
}
