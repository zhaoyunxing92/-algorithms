package array

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

/*
给定一个整数数组，判断是否存在重复元素。

如果存在一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/contains-duplicate
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//使用map
//对于数组中每个元素，我们将它插入到哈希表中。如果插入一个元素时发现该元素已经存在于哈希表中，则说明存在重复的元素。
func containsDuplicate(nums []int) bool {
	if nums == nil || len(nums) < 2 {
		return false
	}
	set := map[int]struct{}{}
	for _, num := range nums {
		if _, ok := set[num]; ok {
			return true
		}
		set[num] = struct{}{}
	}
	return false
}

//排序,先排序然后相邻的两个对比
func containsDuplicate2(nums []int) bool {
	size := len(nums)
	if nums == nil || size < 2 {
		return false
	}
	sort.Ints(nums)
	for fast := 1; fast < size; fast++ {
		if nums[fast] == nums[fast-1] {
			return true
		}
	}
	return false
}

func TestContainsDuplicate(t *testing.T) {
	var (
		nums      []int
		duplicate bool
	)

	nums = []int{1, 2, 3, 3}

	duplicate = containsDuplicate(nums)
	assert.True(t, duplicate)
	duplicate = containsDuplicate2(nums)
	assert.True(t, duplicate)

}
