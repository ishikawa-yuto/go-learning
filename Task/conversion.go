package main

import (
	"fmt"
	"strconv"
)

func main4() {
	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	fmt.Println(i, s)
}
