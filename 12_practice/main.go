// 1: intx float 64y type conversion sample
// package main

// import "fmt"

// func main() {
// 	x := 75
// 	var y float64
// 	y = float64(x) // type (value)
// 	fmt.Println(y)
// }

// 2: multiple assing sample x,y = y,x
// package main

// import "fmt"

// func main() {
// 	x := 5
// 	y := 10
// 	fmt.Println("x:", x, "y:", y)

// 	x, y = y, x // x=y,y=x
// 	fmt.Println("x:", x, "y:", y)

// }

// 3: non Englihs variable names
// package main

// import "fmt"

// func main() {
// 	yas := 40
// 	fmt.Println(yas)
// }

// 4: Shadowing kavramı -> gölgeleme
// package main

// import "fmt"

// func main() {
// 	x := 5
// 	if true {
// 		x := 10
// 		x++
// 		fmt.Println(x)
// 	}
// 	fmt.Println(x)
// }

// 5: 40 as a string
package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 40
	s := string(65)
	w := string("65") // parantez içine almadan 65 yazarsam ASCII karşılığını 'A' verir.
	fmt.Printf("%v %T \n", x, x)
	fmt.Printf("%v %T \n", s, s)
	fmt.Printf("%v %T \n", w, w)

	y := strconv.Itoa(x)
	fmt.Printf("%v %T \n", y, y) // "40"
}
