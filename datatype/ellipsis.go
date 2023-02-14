package main

import "fmt"

// func main() {
// 	cities := [...]string{"New York", "Paris", "Berlin", "Madrid"}
// 	fmt.Println("Cities:", cities)
// }

func main() {
	numbers := [...]int{99: -1}
	fmt.Println("First Position:", numbers[0])
	fmt.Println("Last Position:", numbers[99])
	fmt.Println("Length:", len(numbers))
}
