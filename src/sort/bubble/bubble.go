package bubble

import "time"
/**
暴力排序
*/
func ViolenceBubbleSort(array []int) time.Duration {
	start := time.Now()
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1; j++ {
			//如果前面数小于后面数就交换
			if array[j] > array[j+1] {
				// set sort/sort.go:273
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return time.Now().Sub(start)
}
/**
排序过的数字就过滤，所以长度是一个动态的
 */
func BubbleSort(array []int) time.Duration {
	start := time.Now()
	for end := len(array); end > 0; end-- {
		for begin := 1; begin < end; begin++ {
			if array[begin] < array[begin-1] {
				array[begin], array[begin-1] = array[begin-1], array[begin]
			}
		}
	}
	return time.Now().Sub(start)
}
