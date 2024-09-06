/* // main.go
package main

import (
	"28_own_packages/merhaba"
)

func main() {
	merhaba.Hola()
	merhaba.Hello()
	merhaba.Merhaba()
}
*/

// GETCELCIUS İÇİN
/*
	package main

import (

	"ArinSoftware/getcelcius"
	"fmt"

)

	func main() {
		fmt.Println("Celsius derecesini girin:")
		celsius, err := getcelcius.GetCelcius()
		if err != nil {
			fmt.Println("Hata:", err)

		}

		kelvin := celsius + 273
		fmt.Println(celsius, " Celsius derecesi,", kelvin, " Kelvin derecesine eşittir.")
	}
*/

package main

import (
	"28_own_packages/getcelcius"
	"fmt"
)

func main() {
	fmt.Println("Celsius derecesini girin:")
	celsius, err := getcelcius.GetCelcius()
	if err != nil {
		fmt.Println("Hata:", err)

	}

	fahrenheit := (celsius * 9 / 5) + 32
	fmt.Println(celsius, " Celsius derecesi,", fahrenheit, " Fahrenheit derecesine eşittir.")
}
