package main

import "fmt"

func main() {

	for i, v := range []string{"foo", "bar", "baz"} {
		fmt.Println(i, v)
	}
}
