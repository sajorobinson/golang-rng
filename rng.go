package main

import (
	"fmt"
	"math/rand"
	"time"
)

var max int

func numberGen(n int) (int error) {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(n-1) + 1)
	return
}

func main() {
	max = 1000
	numberGen(max)
}
