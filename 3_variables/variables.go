package main

import "fmt"

func main() {
	// var name string = "Hello Go"

	// type inference
	var name = "Hello Go"
	fmt.Println(name)

	var isAdult = true
	fmt.Println(isAdult)
	// fmt.Println("Hello, World!")

	var age int = 24
	fmt.Println(age)

	// shorthand syntax
	names := "Hello Go" // only works inside functions and not outside and when we are declaring and initializing at the same time
	fmt.Println(names)

	// when you need to declare but want to assing later so you need to do like
	var name2 string
	name2 = "hello go"
	fmt.Println(name2)

	var price = 4.5
	fmt.Println(price)

	price3 := 6.9
	fmt.Println(price3)
}
