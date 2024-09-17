/* package main

import (
	"fmt"
	"math"
)

type circle struct {
	r float64
}

func (c circle) display() {
	fmt.Println("A Circle")

}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c1 := circle{5}

	// area1 := go c1.area() // hata almamzıı n nedeni go kelimesinin kondugu fonksiyonda herhangi return değeri alınamaz
	area1 := c1.area()
	fmt.Printf("%.2f\n", area1)

	go c1.display()
} */

/* package main

import "fmt"

func hi() string {
	return "hi"
}

func main() {
	h1 := go hi() // Hata alırız.....!
	fmt.Println(h1)
} */

/* package main

import "fmt"

func merhaba(chan1 chan string) { //step.2 kanalla beraber göndermek istediğimiz fonksiyonumuzu yazdık..
	chan1 <- "hi" // değeri kanal gönderiyor.
} // --> parametresi channel olan fonksiyon ..
func main() {
	// var myChannel chan string
	// myChannel = make(chan string) ---> uzun yolla channel oluşturma

	myChannel := make(chan string) // step.1 channel oluşturduk()

	go merhaba(myChannel)    // go routine oluşturuyoruz buna parametre olarak step1'de tanımlanan channel atadık
	fmt.Println(<-myChannel) // kanaldaki değeri almamızı sağlıyor.
} */

/* package main

import "fmt"

func main() {
	chan1 := make(chan string)
	chan1 <- "ba"
	fmt.Println(<-chan1)
}  */ // deadlock hatası alırızz

//çözüm1
// func main() {

// 	chan1 := make(chan string)
// 	go func() {
// 		chan1 <- "ba"
// 	}()

// 	fmt.Println(<-chan1)

// }

//çözüm2
// func main() {
// 	chan1 := make(chan string, 1)
// 	chan1 <- "ba"
// 	fmt.Println(<-chan1)
// }

package main

import (
	"fmt"
	"math"
)

type circle struct {
	r float64
}

func (c circle) display() {
	fmt.Println("A Circle")

}

func area(c circle, myChannel chan float64) {
	result := math.Pi * c.r * c.r
	myChannel <- result
}

func main() {
	c1 := circle{5}
	chan1 := make(chan float64)

	go area(c1, chan1)
	fmt.Printf("%.2f\n", <-chan1)

	c1.display()
}
