/* package main

import "fmt"

func main() { // GO pass by value
	x := 5
	fmt.Println(x)
	double(x)
	fmt.Println(x)
}
func double(num int) {
	num *= 2
	fmt.Println(num)
}
*/

/*
	package main // pass by referance

import "fmt"

	func main() {
		mySlc := []int{1, 10, 100}
		fmt.Println(mySlc)
		double(mySlc)
		fmt.Println(mySlc)

}

	func double(slc []int) {
		for i := 0; i < len(slc); i++ {
			slc[i] *= 2
		}
		fmt.Println(slc)
	}
*/
/* package main // pass by value

import "fmt"

func main() {
	myArr := [3]int{1, 10, 100}
	fmt.Println(myArr)
	double(myArr)
	fmt.Println(myArr)

}
func double(Arr [3]int) {
	for i := 0; i < len(Arr); i++ {
		Arr[i] *= 2
	}
	fmt.Println(Arr)
}
*/

package main

import "fmt"

func main() { // GO pass by value
	x := 5
	fmt.Println(x)
	double(&x)
	fmt.Println(x)
}
func double(num *int) { // double --> pointer Method
	*num *= 2
	fmt.Println(*num)
	fmt.Println(num)
}
