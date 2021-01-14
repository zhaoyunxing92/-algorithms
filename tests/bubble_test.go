package tests

import (
	"fmt"
	"github.com/zhaoyunxing92/algorithms/sort/bubble"
	"github.com/zhaoyunxing92/algorithms/sort/selection"
	"github.com/zhaoyunxing92/algorithms/tools"
	"testing"
)

//TestViolenceBubbleSort:暴力排序
func TestViolenceBubbleSort(t *testing.T) {
	var seed = tools.Random(4, 40, 5)

	fmt.Println("暴力排序前：", seed)
	bubble.ViolenceBubbleSort(seed)
	fmt.Println("暴力排序后：", seed)
}

func TestBubbleSort(t *testing.T) {
	var seed = tools.Random(4, 40, 5)

	fmt.Println("冒泡排序前：", seed)
	bubble.BubbleSort(seed)
	fmt.Println("冒泡排序后：", seed)
}

func TestBubbleSort2(t *testing.T) {
	var seed = tools.Random(4, 40, 5)
	fmt.Println("冒泡排序前：", seed)
	bubble.BubbleSort2(seed)
	fmt.Println("冒泡排序后：", seed)
}

func TestBubbleSort3(t *testing.T) {
	var seed = tools.Random(4, 40, 5)
	fmt.Println("冒泡排序前：", seed)
	bubble.BubbleSort3(seed)
	fmt.Println("冒泡排序后：", seed)
}

//TestBubbleSortCount:比较几个算法的次数
func TestBubbleSortCount(t *testing.T) {
	//seed := tools.Random(4, 40, 10)
	seed := []int{5, 3, 6, 7, 8, 9, 15, 30, 32, 50}

	size := len(seed)

	fmt.Println("冒泡排序前：", seed)

	seed1 := make([]int, size, size)
	seed2 := make([]int, size, size)
	seed3 := make([]int, size, size)
	seed4 := make([]int, size, size)

	copy(seed1, seed)
	copy(seed2, seed)
	copy(seed3, seed)
	copy(seed4, seed)

	bubble.BubbleSort(seed1)  //比较10次
	bubble.BubbleSort2(seed2) //比较7次
	bubble.BubbleSort3(seed3) //比较4次
	selection.SelectionMaxSort(seed4)
}

//TestMaxDepth:计算深度
func TestMaxDepth(t *testing.T) {
	depth := bubble.MaxDepth(1 << 5)
	fmt.Println(depth)
}
