package main

import "fmt"

func main() {
	// var nama = "golang"
	// var age = 5
	// var address = "google"
	// fmt.Printf("Halo namaku %s, umurku %d, dan aku berasal dari %s", nama, age, address)
	// fmt.Println("ini adalah bahasa pemograman ==>", nama)
	// fmt.Println("saya sudah berumur ==>", age)
	// fmt.Println("saya berasal dari ==>", address)

	// var name = "Golang"
	// var age = 5

	// fmt.Printf("%T, %T", name, age)

	var student1, student2, student3 string = "satu", "dua", "tiga"
	var first, second, third int
	first, second, third = 1, 2, 3

	fmt.Println(student1, student2, student3)
	fmt.Println(first, second, third)

	var firstVariable string

	_ = firstVariable

	name := "android"
	age := 20

	fmt.Printf("%T, %T", name, age)

}
