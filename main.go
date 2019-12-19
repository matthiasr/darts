package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	runs = 1e6
)

func main() {
	rand.Seed(time.Now().UnixNano())
	s := summary{}
	start := time.Now()
	for i := 0; i < runs; i++ {
		score := NewGame().Run()
		s.Observe(score)
	}
	end := time.Now()

	fmt.Println("Summary:", s)
	fmt.Println("Average:", s.Average())
	fmt.Println("Duration:", end.Sub(start))
}
