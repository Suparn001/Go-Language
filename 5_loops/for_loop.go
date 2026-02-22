package main

import "fmt"

// there is only one type of loop in go and that is for loop and it can be used in three different ways

// for -> only construct in go for looping
func main() {
	// while loop
	// i := 1
	// for i <= 5 {
	// 	fmt.Println(i)
	// 	i++ // or i=i+1
	// }

	// infinte loop
	// for {
	// 	println("1")
	// }

	// for i := 1; i <= 5; i++ {
	// break // this will break the loop when i is 1 and it will not print anythings
	// fmt.Println(i)

	// continue // this will skip the current iteration when i is 1 and it will print 2,3,4,5
	// if i == 3 {
	// 	continue
	// }
	// fmt.Println(i)

	// }

	// range
	for i := range 3 {
		fmt.Println(i) // this will not work because range can only be used with arrays, slices, maps, strings and channels

		// rnage means from 0 you have to 3 but it will not include 3 so it will print 0,1,2
	}
}
