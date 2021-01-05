package tests

import (
	"fmt"
	"github.com/zhaoyunxing92/algorithms/sort/selection"
	"github.com/zhaoyunxing92/algorithms/tools"
	"testing"
)

func TestSelectionMaxSort(t *testing.T) {
	var seed = tools.Random(4, 40, 5)
	fmt.Println("选择排序排序前：", seed)
	selection.SelectionMaxSort(seed)
	fmt.Println("选择排序排序后：", seed)
}