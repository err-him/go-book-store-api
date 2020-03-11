package utils

import (
	"math/rand"
	"time"
)

const (
	letterString = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func GenerateRandomStr(n int) string {
	//seed everytime otherwiese it will provide the same string every time
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterString[rand.Intn(len(letterString))]
	}
	return string(b)
}
