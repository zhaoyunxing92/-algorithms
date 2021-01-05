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
