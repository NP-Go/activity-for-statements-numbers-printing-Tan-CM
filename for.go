package main

import (
	"fmt"
)

func main() {
	// print 0 to 1000
	for i := 0; i <= 1000; i++ {
		fmt.Printf("%d \n", i)
	}

	fmt.Println("Odd Numbers")
	// print only odd numbers
	for i := 0; i <= 1000; i++ {
		if i%2 == 1 {
			fmt.Printf("%d \n", i)
		}
	}

	// print only even numbers
	fmt.Println("Even Numbers")
	for i := 0; i <= 1000; i++ {
		if i%2 == 0 {
			fmt.Printf("%d \n", i)
		}
	}
}
