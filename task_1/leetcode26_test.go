package task1

/**
删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，
返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j]
不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位

问题描述：
给定一个 ​非严格递增排列​ 的数组 nums，要求：

1.​原地删除重复元素，使每个元素只出现一次。
2.保持元素的 ​相对顺序​ 不变。
3.返回唯一元素的个数 k，并确保 nums 的前 k 个元素是唯一的。
**/

import (
	"reflect"
	"testing"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

func TestRemoveDuplicates(t *testing.T) {
	input := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// fmt.Print("加1结果:", plusOne(input))
	result := removeDuplicates(input)
	t.Logf("去重后的结果: %v", result)
	// 添加一个实际断言
	expected := 5
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("期望得到 %v,但实际得到 %v", expected, result)
	}
}
