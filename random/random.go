package random

import "math/rand"

func Calcule() {
	for i := 0; i < 5; i++ {
		println(rand.Intn(10))
	}
}
