package task1

/**
给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。
这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
将大整数加 1，并返回结果的数字数组。

示例 1：
输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
加 1 后得到 123 + 1 = 124。
因此，结果应该是 [1,2,4]。
**/

import (
	"reflect"
	"testing"
)

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] < 10 {
			return digits
		}
		digits[i] = 0
	}

	return append([]int{1}, digits...)
}

func TestPlusOne(t *testing.T) {
	input := []int{1, 3, 9}
	// fmt.Print("加1结果:", plusOne(input))
	result := plusOne(input)
	t.Logf("加1结果: %v", result)
	// 添加一个实际断言
	expected := []int{1, 4, 0}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("期望得到 %v,但实际得到 %v", expected, result)
	}
}
