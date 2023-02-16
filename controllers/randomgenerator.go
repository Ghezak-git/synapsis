package controllers

import (
	"math/rand"
	"time"
)

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringWithCharset() string {
	charset := "aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ"
	b := make([]byte, 30)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset)-1)]
	}
	return string(b)
}
