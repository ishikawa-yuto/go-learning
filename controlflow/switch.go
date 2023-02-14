package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main2() {
	sec := time.Now().Unix()
	rand.Seed(sec)
	i := rand.Int31n(10)

	switch i {
	case 0:
		fmt.Print("zero...")
	case 1:
		fmt.Print("one...")
	case 2:
		fmt.Print("two...")
	}

	fmt.Println("ok")
}
