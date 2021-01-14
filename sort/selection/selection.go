package selection

import (
	"fmt"
	"time"
)

/**
根据最大值排序
*/
func SelectionMaxSort(array []int) time.Duration {
	start := time.Now()
	count := 0 //记录比较次数
	for end := len(array) - 1; end > 0; end-- {
		maxIndex := 0
		for begin := 1; begin <= end; begin++ {
			count++
			//这里等于的可以建立稳定排序
			if array[maxIndex] <= array[begin] {
				maxIndex = begin
			}
		}
		array[maxIndex], array[end] = array[end], array[maxIndex]
	}
	fmt.Printf("[SelectionMaxSort]数组长度：%d,比较次数：%d \n", len(array), count)
	fmt.Println("[SelectionMaxSort]", array)
	return time.Now().Sub(start)
}
