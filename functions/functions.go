package functions

import (
	"math/rand"
	"time"
)

// Rand - случайное int число
// принимает int, предел рандомного числа
func Rand(r int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(r)
}

// MathSign - определяет знак числа
// если отрицательное, то возвращает -1
// если 0, то возвращает 0
// если положительное, то возвращает 1
// если вернуло -2, то это ошибка
func MathSign(x int) int {
	if x > 0 {
		return 1
	}

	if x == 0 {
		return 0
	}

	if x < 0 {
		return -1
	}

	return -2
}
