package main

import (
	"fmt"
	"github.com/zhaoyunxing92/algorithms/src/sort/bubble"
	"github.com/zhaoyunxing92/algorithms/src/tools"
	"sort"
)

func main() {
	seed := tools.Random(4, 400, 50000)
	var a, b, c = seed[:], seed[:], seed[:]

	fmt.Println(bubble.BubbleSort(a))
	//提前给c排序
	sort.Ints(c)
	fmt.Println(bubble.BubbleSort2(c))
	fmt.Println(bubble.ViolenceBubbleSort(b))
}
