package main

import (
	"fmt"
	"github.com/zhaoyunxing92/algorithms/src/tools"
	"time"
)

func main() {
	array := tools.Random(40, 500, 100000)
	duration := bubbleSort(array)
	fmt.Println(duration)
}

func bubbleSort(array []int) time.Duration {
	start := time.Now()
	//排序过的数字就过滤，所有长度是一个动态的
	for end := len(array); end > 0; end-- {
		for begin := 1; begin < end; begin++ {
			//如果前面数小于后面数就交换
			if array[begin] < array[begin-1] {
				// set sort/sort.go:273
				array[begin], array[begin-1] = array[begin-1], array[begin]
			}
		}
	}
	ent := time.Now().Sub(start)
	return ent
}

