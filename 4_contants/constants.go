package main

import "fmt"

// you can declare a constant using the const keyword and you must assign a value to it at the time of declaration and you cannot change the value of a constant once it is assigned
// you can dec;are contant outside the function also

const age = 90

// abc:="" // this is allowed inside the function but not outside the function

const (
	port = 900
	host = "localhost"
)

func main() {
	// const name string = "golang"
	// fmt.Println(age)
	// name = "javascript"

	// constant groupings
	const (
		port = 500
		host = "localhost"
	)
	// fmt.Println(host)
}
fmt.Println(port, host)
