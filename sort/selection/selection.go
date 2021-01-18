package selection

import (
	"fmt"
	"time"
)

/**
根据最大值排序
*/
//SelectionMaxSort: 选择最大的值交换
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

//SelectionMinSort:选择最小
func SelectionMinSort(array []int) {
	count := 0 //记录比较次数
	for begin := 0; begin < len(array); begin++ {
		minIdx := begin
		for end := begin + 1; end < len(array); end++ {
			count++
			if array[end] < array[minIdx] {
				minIdx = end
			}
			fmt.Println("最小指针：", minIdx)
		}
		array[minIdx], array[begin] = array[begin], array[minIdx]
	}
	fmt.Printf("[SelectionMinSort]数组长度：%d,比较次数：%d \n", len(array), count)
	fmt.Println("[SelectionMinSort]", array)
}

func Max(array []int) {

	for start := len(array) - 1; start > 0; start-- {
		max := 0
		for end := 1; end <= start; end++ {
			if array[max] < array[end] {
				max = end
			}
		}
		array[start], array[max] = array[max], array[start]
	}
	fmt.Println("[max]", array)
}
