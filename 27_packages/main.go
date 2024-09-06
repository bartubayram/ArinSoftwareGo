/* package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Lütfen Sınav Griniz:")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString(('\n'))
	fmt.Println(value)
}
*/

/* package main

import "fmt"

func main() {
	fmt.Println("Arin") // fmt paketini import ediyor önce.Burda yazdığımız Println() metodunu almak iin fmt paketini kullanırız. Bu paketş kullanmka için de öncelikle import etmemiz lazım.
} */

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	fmt.Println(strings.Contains("seafood", "foo"))
// 	/* fmt.Println(strings.Contains("seafood", "bar"))
// 	fmt.Println(strings.Contains("seafood", ""))
// 	fmt.Println(strings.Contains("", "")) */
// 	// Contains -> sağında yazılan ifadenin solundaki kelime içerisinde var mı yok mu diye göstermektedir.
// 	// fmt.Println(strings.Count("animatrix", "a"))
// }
//trytrytrytry

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ToUpper("Gopher"))
}
