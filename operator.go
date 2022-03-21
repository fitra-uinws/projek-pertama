package main

import "fmt"

func main() {
	const full_name string = "Golang"
	fmt.Printf("Hello %s", full_name)

	var value = (5 + 2) * 3

	fmt.Println(value)

	var firstcondition bool = 10 == 12

	fmt.Println(firstcondition)

	var right = true
	var wrong = false

	var WrongAndRight = wrong && right

	var WrongOrRight = wrong || right
	var WrongReverse = !wrong

	fmt.Println("wrong and right ", WrongAndRight)
	fmt.Println("wrong or right ", WrongOrRight)
	fmt.Println("wrong reverse ", WrongReverse)
}
