package main

import "fmt"

func main() {

	// Map -> key value çiftlerinden oluşur.
	/* 	myMap := map[string]int{
	   		"Ahmet":   40, //  Ahmet 40 değerine denk gelen bi değerimizdir.
	   		"Ayşe":    17,
	   		"Selim":   14,
	   		"Mustafa": 70,
	   	}

	   	fmt.Println(myMap)
	   	fmt.Println(myMap["Ahmet"])
	   	fmt.Println(myMap["Mustafa"], myMap["Selim"])
	   	fmt.Println(myMap["merhaba"]) */

	/* 	myMap := map[string]int{
	   		"Ahmet":   40, //  Ahmet 40 değerine denk gelen bi değerimizdir.
	   		"Ayşe":    17,
	   		"Selim":   14,
	   		"Mustafa": 70,
	   	}
	   	fmt.Println(myMap["Asli"])
	   	// key ve value aynı veri tipinde olmasına gerek yok !!!!!!
	*/
	/* isMarried := map[string]bool{
		"Ahmet":  true,
		"Ayşe":   false,
		"Mahmut": false,
	}
	fmt.Println(isMarried)
	fmt.Println(isMarried["Ahmet"])
	fmt.Println(isMarried["Mahmut"]) */
	/*
		myMap := make(map[string]int)
		fmt.Println(myMap)
		fmt.Println(myMap["Ali"]) */
	studentGrades := map[string]int{
		"Arin":  80,
		"Ahmet": 29,
		"Selim": 72,
		"Ayşe":  77,
		"Çınar": 0,
	}
	/* fmt.Println(studentGrades)
	fmt.Println(studentGrades["Arin"])
	fmt.Println(studentGrades["Çınar"])
	fmt.Println(studentGrades["Elis"]) */
	// Çınar'ın aldığı not 0 ama ben yanlış isim yazmış olsam da 0 değerini döndürecekti (Defaul value)
	// Key değeri verilen bir elemanın gerçekten map'in içerisinde bulunup bulunmadığını nasıl anlayabileceğiz..?
	/*
		value, ok := studentGrades["Elis"]
		fmt.Println(value, ok)

		_, o := studentGrades["Arin"]
		fmt.Println(o) */

	// fmt.Println(studentGrades)
	// studentGrades["Mahmut"] = 55
	// fmt.Println(studentGrades)
	/* 	fmt.Println(studentGrades)

	   	delete(studentGrades, "Selim")
	   	fmt.Println(studentGrades)
	   	fmt.Println(len(studentGrades)) // key-value çifti olarak ele alıyor. */

	/* sg := studentGrades // maps ---> pass by reference
	sg["bartu"] = 100
	fmt.Println(sg)
	fmt.Println(studentGrades) */

	for k, v := range studentGrades {
		fmt.Println(k, v)

	}
}
