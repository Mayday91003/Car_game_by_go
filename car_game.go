package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	FinishLine = 50
	CarCount   = 5
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cars := make([]int, CarCount)
	for i := range cars {
		cars[i] = 0
	}

	for {
		for i, car := range cars {
			move := rand.Intn(10)
			cars[i] += move
			fmt.Printf("Car %d moved %d spaces, now at %d\n", i+1, move, car)
			if car >= FinishLine {
				fmt.Printf("Car %d has won the race!", i+1)
				return
			}
		}
	}
}
