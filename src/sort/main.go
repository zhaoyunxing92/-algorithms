package main

import (
	"fmt"
	"github.com/zhaoyunxing92/algorithms/src/sort/bubble"
	"github.com/zhaoyunxing92/algorithms/src/sort/selection"
	"github.com/zhaoyunxing92/algorithms/src/tools"
	"sort"
)

func main() {
	seed := tools.Random(4, 40, 10)
	var a, b, c = seed[:], seed[:], seed[:]
    //冒泡排序
	fmt.Println(bubble.BubbleSort(a))
	//提前给c排序
	sort.Ints(c)
	fmt.Println(bubble.BubbleSort2(c))
	fmt.Println(bubble.ViolenceBubbleSort(b))

    //选择排序
	ss:=[4]int{10,10,7,9}
	selection.SelectionMaxSort(ss[:])
	fmt.Println(ss)
}
