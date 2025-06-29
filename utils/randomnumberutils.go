package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomNumber(length int) string {
	// make random numbers using time seeds
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	result := ""
	for i := 0; i < length; i++ {
		result += fmt.Sprintf("%d", r.Intn(10))
	}
	return result
}
