package main

import (
	"github.com/zhaoyunxing92/algorithms/src/sort/bubble"
	"github.com/zhaoyunxing92/algorithms/src/tools"
)

func main() {
	seed := tools.Random(4, 400, 50000)
	var a, b = seed[:], seed[:]

	bubble.BubbleSort(a)
	bubble.ViolenceBubbleSort(b)
}
