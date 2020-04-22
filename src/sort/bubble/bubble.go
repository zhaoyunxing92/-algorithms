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
	for end := len(array)-1; end > 0; end-- {
		for begin := 1; begin <= end; begin++ {
			if array[begin] < array[begin-1] {
				array[begin], array[begin-1] = array[begin-1], array[begin]
			}
		}
	}
	return time.Now().Sub(start)
}

/**
假如给你的数据已经有序就不用找排序了
*/
func BubbleSort2(array []int) time.Duration {
	start := time.Now()
	for end := len(array)-1; end > 0; end-- {
		sorted := true
		for begin := 1; begin <= end; begin++ {
			//能进来说明需要排序了
			if array[begin] < array[begin-1] {
				array[begin], array[begin-1] = array[begin-1], array[begin]
				sorted = false
			}
		}
		if sorted {
			break
		}
	}
	return time.Now().Sub(start)
}

/**
假如给你的数据已经有序就不用找排序了
*/
func BubbleSort3(array []int) time.Duration {
	start := time.Now()
	for end := len(array)-1; end > 0; end-- {
		index := 1
		for begin := 1; begin <= end; begin++ {
			if array[begin] < array[begin-1] {
				array[begin], array[begin-1] = array[begin-1], array[begin]
				index = begin
			}
		}
		end = index
	}
	return time.Now().Sub(start)
}
