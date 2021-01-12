package tests

import (
	"fmt"
	"github.com/zhaoyunxing92/algorithms/sort/bubble"
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
	seed := tools.Random(4, 40, 10)
	fmt.Println("冒泡排序前：", seed)
	seed1 := seed[:]
	seed2 := seed[:]
	seed3 := seed[:]
	bubble.BubbleSort(seed1)
	bubble.BubbleSort2(seed2)
	bubble.BubbleSort3(seed3)
	fmt.Println("冒泡排序后：", seed1)
}

//TestMaxDepth:计算深度
func TestMaxDepth(t *testing.T) {
	depth := bubble.MaxDepth(1 << 5)
	fmt.Println(depth)
}
