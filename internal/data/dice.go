package data

import (
	"math/rand"
	"time"
)

const (
	lowerBound = 1
	D4         = 4
	D6         = 6
	D8         = 8
	D20        = 20
	D100       = 100
)

type Die interface {
	Roll(upperBound uint)
}

type Dice struct {
	NumberOfSides int
}

func generateRandomNumber(upper int) int {
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(upper-lowerBound+1) + lowerBound
	return result
}

func (d Dice) Roll() int {
	result := generateRandomNumber(d.NumberOfSides)
	return result
}
