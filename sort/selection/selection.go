package selection

import "time"

/**
根据最大值排序
*/
func SelectionMaxSort(array []int) time.Duration {
	start := time.Now()
	for end := len(array) - 1; end > 0; end-- {
		maxIndex := 0
		for begin := 1; begin <= end; begin++ {
			//这里等于的可以建立稳定排序
			if array[maxIndex] <= array[begin] {
				maxIndex = begin
			}
		}
		array[maxIndex], array[end] = array[end], array[maxIndex]
	}
	return time.Now().Sub(start)
}
