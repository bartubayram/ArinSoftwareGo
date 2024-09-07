// 3-) Underlying type struct olan Rectangle type oluşturunuz.
// display, area, circumference metodlarını yazınız.
package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	a, b float64
}

func (r rectangle) display() {
	fmt.Println("Kenar Uzunlukları ..:", r.a, "ve ", r.b)
}
func (r rectangle) area() float64 {
	return (r.a * r.b)
}
func (r rectangle) circumference() float64 {
	return 2 * (r.a + r.b)
}

type circle struct {
	r float64
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}
func (c circle) circumference() float64 {
	return math.Pi * 2 * c.r
}
func (c circle) diameter() float64 {
	return 2 * c.r
}

type shape interface {
	area() float64
	circumference() float64
}

func interfaceFunc(i shape) {
	fmt.Println(i)
	fmt.Println(i.circumference())
	fmt.Printf("%T ", i)
	fmt.Println()
}
func main() {

	/* r1 := rectangle{
		a: 10,
		b: 20,
	} */
	/* r1 := rectangle{5, 20}

	fmt.Println("Area ..:", r1.area())
	fmt.Println("Circumference ..:", r1.circumference())
	fmt.Println()
	interfaceFunc(r1) */

	r1 := rectangle{2, 8}
	interfaceFunc(r1)
	fmt.Println()
	c1 := circle{2}
	interfaceFunc(c1)

}
