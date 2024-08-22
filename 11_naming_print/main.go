package main

var x = 10

// Bazı değişlenlere burda tanımlanırken diğer paketlere de erişimi olur bunun için değişkeninin adını BÜYÜK HARFLE yazarız

var Y = 100 //  Burdaki Y değişkeninin diğer paketlereden de erişimi vardır anlamındadır

func main() {
	// Print - Println - Printf
	// (Print - Println)  ->raw string = aldıkları gibi yazdırırlar
	// Printf ise formatlanmış şekilde yazdırır.
	// fmt.Println("Merhaba")
	// fmt.Print("Merhaba")
	// fmt.Print("merhaba")
	// name := "Arin"

	// fmt.Print(name)
	// fmt.Println(name)
	// fmt.Printf(name)

	// fmt.Print("My name :", name)
	// fmt.Println()
	// fmt.Println("My name :", name)
	// fmt.Printf("My name : %v", name)  // %v değişeknin içindeki değeri almamıza yarar.
	// fmt.Printf("My type : %T ", name) // %T değişkenin tipini gösterir.
	// fmt.Println()
	// fmt.Printf("My type : %T - My name : %v", name, name) // %T değişkenin tipini gösterir.

	// x := 100
	// y := 20
	// z := 30
	// fmt.Printf("%b %d %o", x, y, z)

	// name, age := "Arin", 5
	// //fmt.Print("Benim Adım: ", name, ", yaşım : ", age)
	// fmt.Println("Benim Adım:", name, ",yaşım :", age)
	// // Println - Printf farkları -> println line atlar ve <<"name:"name>> yazdıgımızda normal boşluk bırakırız güzel dursun diye ama Println değiişken(name)'den önce ve sonrasında kendisi boşluk bırakır.
	// fmt.Printf("Benim Adım: %v ,yaşım : %v", name, age)

	// VISIBILITY(GÖRÜNÜRLÜLÜK) -> Biz herhangi bi şekilde bi değişken tanımladığımız zaman bu değişkenin go programında nereden görünülebilirliği kavramıdı.Tanımlandığı paket dışarısında tnaımlı değildir görünmez
	// fmt.Println(x)
	// my()

	var c string
}

// func my() {
// 	fmt.Println(x)
// }
