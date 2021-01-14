package bubble

import (
	"fmt"
	"time"
)

//ViolenceBubbleSort:暴力排序
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

//BubbleSort:排序过的数字就过滤，所以长度是一个动态的
func BubbleSort(array []int) time.Duration {
	start := time.Now()
	count := 0 //记录比较次数
	for end := len(array) - 1; end > 0; end-- {
		for begin := 1; begin <= end; begin++ {
			count++
			if array[begin] < array[begin-1] {
				array[begin], array[begin-1] = array[begin-1], array[begin]
			}
		}
	}
	fmt.Printf("[BubbleSort1]数组长度：%d,比较次数：%d \n", len(array), count)
	fmt.Println("[BubbleSort1]", array)
	return time.Now().Sub(start)
}

//BubbleSort2:假如给你的数据经过某次已经有序了
func BubbleSort2(array []int) time.Duration {
	start := time.Now()
	count := 0 //记录比较次数
	for end := len(array) - 1; end > 0; end-- {
		//标记是否排序过
		sorted := true
		for begin := 1; begin <= end; begin++ {
			count++
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
	fmt.Printf("[BubbleSort2]数组长度：%d,比较次数：%d \n", len(array), count)
	fmt.Println("[BubbleSort2]", array)
	return time.Now().Sub(start)
}

//BubbleSort3:假如给你的数据已经有序就不用找排序了
//数组后面有序的个数越多比较次数越少
func BubbleSort3(array []int) time.Duration {
	start := time.Now()
	count := 0 //记录比较次数
	for end := len(array) - 1; end > 0; end-- {
		index := 1
		for begin := 1; begin <= end; begin++ {
			count++
			if array[begin] < array[begin-1] {
				array[begin], array[begin-1] = array[begin-1], array[begin]
				index = begin
			}
		}
		end = index
	}
	fmt.Printf("[BubbleSort3]数组长度：%d,比较次数：%d \n", len(array), count)
	fmt.Println("[BubbleSort3]", array)
	return time.Now().Sub(start)
}

//MaxDepth: sort排序深度算法
func MaxDepth(size int) int {
	fmt.Println("size=", size)
	var depth int
	for i := size; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}
