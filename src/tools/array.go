package tools

import (
	"math/rand"
)

/**
min :最小值
max :最大值
size :数组长度
*/
func Random(min, max, size int) []int {
	if min > max {
		panic("max mast greater than min")
	}
	if size < 1 {
		size = 1
	}

	array := make([]int, size)
	for i := 0; i < size; i++ {
		numb := rand.Intn(max)
		if numb <= min {
			numb += min
		}
		array[i] = numb
	}
	return array
}
