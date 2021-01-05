# 冒泡排序

## 定义

重复地走访要排序的元素，依次比较两个相邻元素的大小依次完成排序。如图

![bubble sort](../../img/bubble_sort.gif)

## 先看总结

* 最好时间复杂度:O(n) 也就是数组提前有序
* 最坏时间复杂度:O(n<sup>2</sup>)
* 是一种稳定排序（能保证相对位置不改变）
* 原地算法（没有依赖额外的资源，仅依赖输出覆盖输入），也就是空间复杂度低

## 开始推断

开始你想直接暴力点两个for就完事是代码如下

```go
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
	ent := time.Now().Sub(start)
	return ent
}
```

但是聪明的你立马想到了我每次遍历一遍后就已经确定了最大的值了是不是下次比较的时候就可以排除掉它，于是有了下面的代码。动态的长度

```go
func BubbleSort(array []int) time.Duration {
	start := time.Now()
	for end := len(array)-1; end > 0; end-- {
		for begin := 1; begin <= end; begin++ {
			if array[begin] < array[begin-1] {
				array[begin], array[begin-1] = array[begin-1], array[begin]
			}
		}
	}
	ent := time.Now().Sub(start)
	return ent
}
```

写完代码后你在感叹自己的聪明时又感觉那里怪怪的，你在想如果给的数组本来就有序的那么是不是就不用继续了，激动的你立马想到了添加一个`flag`记录下是否有序，于是有了下面的代码

```go
func BubbleSort2(array []int) time.Duration {
	start := time.Now()
	for end := len(array)-1; end > 0; end-- {
		sorted:=true
		for begin := 1; begin <= end; begin++ {
			//能进来说明需要排序了
			if array[begin] < array[begin-1] {
				array[begin], array[begin-1] = array[begin-1], array[begin]
				sorted=false
			}
		}
		if sorted {
			break
		}
	}
	return time.Now().Sub(start)
}
```

就在你暗暗窃喜的时候发现数组排好序的概率太低了，然是如果是部分有序的概率还是有的于是你又修改了代码.记录下最后一次交换的位置

```GO
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
```