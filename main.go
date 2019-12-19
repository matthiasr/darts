package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	runs = 100
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < runs; i++ {
		fmt.Printf("%d ", NewGame().Run())
	}
}
