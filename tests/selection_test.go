package tests

import (
	"fmt"
	"github.com/zhaoyunxing92/algorithms/sort/selection"
	"testing"
)

func TestSelectionMaxSort(t *testing.T) {
	//seed := tools.Random(4, 40, 5)
	seed := []int{2, 3, 1, 4, 5}
	fmt.Println("选择排序排序前：", seed)
	selection.SelectionMaxSort(seed)
}

func TestSelectionMinSort(t *testing.T) {
	//seed := tools.Random(4, 40, 5)
	seed := []int{2, 3, 1, 4, 5}
	fmt.Println("选择排序排序前：", seed)
	selection.SelectionMinSort(seed)
}

func TestMax(t *testing.T) {
	//seed := tools.Random(4, 40, 5)
	seed := []int{2, 3, 1, 5, 4}
	fmt.Println("选择排序排序前：", seed)
	selection.Max(seed)
}
