# 选择排序

## 定义

简单点说选择排序就是在数组总找出最大或最小的值与数组中其他元素比较。如图

![select sort](../../img/select_sort.gif)

## 先看总结

* 最坏时间复杂度:O(n<sup>2</sup>)
* 是一种稳定排序（能保证相对位置不改变）
* 原地算法

## 开始代码

```go
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
```