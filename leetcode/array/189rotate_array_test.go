package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
给定一个数组，将数组中的元素向右移动k个位置，其中k是非负数。

进阶：

尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
你可以使用空间复杂度为O(1) 的原地算法解决这个问题吗？


示例 1:

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]


示例2:

输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右旋转 1 步: [99,-1,-100,3]
向右旋转 2 步: [3,99,-1,-100]

提示：

* 1 <= nums.length <= 2 * 104
* -231 <= nums[i] <= 231 - 1
* 0 <= k <= 105

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/problems/rotate-array
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

func rotate(nums []int, k int) {
	size := len(nums)
	//
	k %= size
	rotateAll(nums)
	rotateAll(nums[:k])
	rotateAll(nums[k:])
}

//旋转全部
func rotateAll(nums []int) {
	size := len(nums)
	//边界检查
	mid := len(nums) / 2
	//数组全部翻转
	for i := 0; i < mid; i++ {
		nums[i], nums[size-1-i] = nums[size-1-i], nums[i]
	}
}

/*
### 解题思路
#### 1.使用数组反转
把它比作一个会动的蛇,`k`就是移动几格，这个题的关键点是`k`值的计算，前面几次都没有考虑到`k`超出了数组长度怎么办

例如：[1, 2, 3, 4, 5, 6]  k=11
这个时候蛇的走动步骤如下
```
    //1: 6,1,2,3,4,5
	//2: 5,6,1,2,3,4
	//3: 4,5,6,1,2,3
	//4: 3,4,5,6,1,2
	//5: 2,3,4,5,6,1
	//6: 1,2,3,4,5,6

	//7: 6,1,2,3,4,5
	//8: 5,6,1,2,3,4
	//9: 4,5,6,1,2,3
	//10: 3,4,5,6,1,2
	//11: 2,3,4,5,6,1
```
你就会发现第一步和第七步居然一样走十一步和走五步一样，因此取余就可以了

#### 2.使用额外的数组

使用一个新的数组保存要移动的数组


*/
func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}

	rotate(nums, 11)
	assert.Equal(t, []int{2, 3, 4, 5, 6, 1}, nums)

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	assert.Equal(t, []int{5, 6, 7, 1, 2, 3, 4}, nums)

	nums = []int{-1, -100, 3, 99}
	rotate(nums, 2)
	assert.Equal(t, []int{3, 99, -1, -100}, nums)
}
