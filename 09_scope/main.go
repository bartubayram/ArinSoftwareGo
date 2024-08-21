package main

import "fmt"

// var packVar = "Package Scope"
// var funcVar = "Func(Package) Scope"

//-> syntax error: non-declaration statement outside function body

func main() {
	/*if true {
		var blockVar = "Block Scope"
		fmt.Println(blockVar)
	}*/
	var name = "Arin"
	name, surname := "Elis", "yazilim"
	fmt.Println(name, surname)
	// var funcVar = "Func Scope"
	// fmt.Println(funcVar)
	// fmt.Println(packVar)
	// var funcVar = "Func Scope"

	// if true {
	// 	var blockVar = "Block Scope"
	// 	fmt.Println(blockVar)
	// }
	// // fmt.Println(blockVar) -> kendi scope dışında çalışmaz
	// fmt.Println(funcVar)
	// fmt.Println(pacVar)
	// myFunc()
	// myFunc()
}

// func myFunc() {
// 	fmt.Println(funcVar)
// 	// 	// fmt.Println(funcVar) -> undefined: funcVar hatası alırız.
// }
