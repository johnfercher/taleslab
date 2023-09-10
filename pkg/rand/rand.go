package rand

import (
	"math/rand"
	"time"
)

var lastRotation = make(map[string]int)

func DifferentRotation(verticalX bool, ticksOfFreedom int, key string) int {
	value90 := 384
	value270 := 1152
	minTick := 64
	value := 0

	if !verticalX {
		value += value90
	}

	rand.Seed(time.Now().UnixNano())

	if rand.Intn(100)%2 == 0 {
		value += value90
	} else {
		value += value270
	}

	randomValue := rand.Intn(ticksOfFreedom) * minTick

	rand.Seed(time.Now().UnixNano())
	if rand.Intn(100)%2 == 0 {
		value += randomValue
	} else {
		value -= randomValue
	}

	if value == lastRotation[key] {
		return DifferentRotation(verticalX, ticksOfFreedom, key)
	}

	lastRotation[key] = value
	return value
}

var lastDistance = make(map[string]int)

func DifferentIntn(maxRand int, key string) int {
	return getRandom(maxRand, key, 0)
}

func Int() int {
	return rand.Int()
}

func Intn(n int) int {
	return rand.Intn(n)
}

func getRandom(maxRand int, key string, depth uint) int {
	randomValue := rand.Intn(maxRand)
	depth++

	if randomValue == lastDistance[key] && depth < 10 {
		return getRandom(maxRand, key, depth)
	}

	lastDistance[key] = randomValue
	return randomValue
}

func Option(currentValue, maxValue int, divisor float64) bool {
	xPercent := float64(currentValue) / float64(maxValue)
	value := (rand.NormFloat64() / divisor) + xPercent
	if value < 0.5 {
		return false
	} else {
		return true
	}
}
