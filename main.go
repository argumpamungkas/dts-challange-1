package main

import "fmt"

// CHALLANGE 1

func main() {
	i := 21
	fmt.Printf("%d\n", i)  // Nilai i
	fmt.Printf("%T \n", i) // Tipe data i

	fmt.Println("%") // Menampilkan %

	var j = true
	fmt.Println(j) // j bernilai true

	j = false
	fmt.Println(j) // j bernilai false

	fmt.Println("\u042F") // menampilkan unicode russia : Я

	num := 21
	fmt.Printf("%d \n", num) // menampilkan base 10 : 21
	fmt.Printf("%o \n", num) // menampilkan base 8 : 25

	numBaseSixteen := 15
	fmt.Printf("%x \n", numBaseSixteen) // menampilkan base 16 : f
	fmt.Printf("%X \n", numBaseSixteen) // menampilkan base 16 : F

	unicodeRussia := 'Я'
	fmt.Printf("%U\n", unicodeRussia) // menampilkan unicode karakter Я : U+042F

	var k float64 = 123.456
	fmt.Printf("%f\n", k) // menampilkan float 123.456000
	fmt.Printf("%E\n", k) // menampilkan float scientific 1.1234560E+02
}
