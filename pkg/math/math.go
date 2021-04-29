package math

import (
	"math"
	"math/rand"
	"time"
)

func Distance(x1, y1, x2, y2 int) uint16 {
	a := math.Pow(float64(x2-x1), 2.0)
	b := math.Pow(float64(y2-y1), 2.0)
	c := uint16(math.Sqrt(a + b))
	return c
}

var lastRotation = make(map[string]uint16)

func GetRandomSoftlyRotation(verticalX bool, ticksOfFreedom int, key string) uint16 {
	value90 := 384
	value270 := 1152
	minTick := 64
	value := 0

	rand.Seed(time.Now().UnixNano())

	if rand.Intn(100)%2 == 0 {
		value = value90
	} else {
		value = value270
	}

	randomValue := rand.Intn(ticksOfFreedom) * minTick

	rand.Seed(time.Now().UnixNano())
	if rand.Intn(100)%2 == 0 {
		value += randomValue
	} else {
		value -= randomValue
	}

	if uint16(value) == lastRotation[key] {
		return GetRandomSoftlyRotation(verticalX, ticksOfFreedom, key)
	}

	lastRotation[key] = uint16(value)
	return uint16(value)
}

var lastDistance = make(map[string]uint16)

func GetRandomSoftlyDistance(maxDistance int, key string) uint16 {
	randomValue := rand.Intn(maxDistance)

	if uint16(randomValue) == lastDistance[key] {
		return GetRandomSoftlyDistance(maxDistance, key)
	}

	lastDistance[key] = uint16(randomValue)
	return uint16(randomValue)
}
