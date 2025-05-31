package common

import "math/rand"

func RandomChoice[T any](arr []T) T {
	index := rand.Intn(len(arr))
	return arr[index]
}
