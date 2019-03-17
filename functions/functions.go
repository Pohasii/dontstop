package functions

import (
	"math/rand"
	"time"
)

// Rund - случайное int число
// принимает int, предел рандомного числа
func Rand(r int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(r)
}
