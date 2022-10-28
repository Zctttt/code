package cmd

import (
	"crypto/rand"
	"math/big"
)

func ShackDict() int64 {
	return Rand(1, 6)
}
func Rand(head, end int64) int64 {
	result, _ := rand.Int(rand.Reader, big.NewInt(end-head))
	return result.Int64() + head
}
