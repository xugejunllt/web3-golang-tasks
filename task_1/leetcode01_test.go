package task1

/**
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数

**/
import (
	"reflect"
	"testing"
)

// 构建hash表 {数值:索引}
func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, exists := hashMap[complement]; exists {
			return []int{i, j}
		}
		hashMap[num] = i

	}

	return nil
}

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	t.Logf("符合条件的两个数组下标: %v,%v", result[0], result[1])
	// 添加一个实际断言
	expected := []int{1, 0}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("期望得到 %v,但实际得到 %v", expected, result)
	}
}
