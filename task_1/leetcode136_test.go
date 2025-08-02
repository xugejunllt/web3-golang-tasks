/**
136. 只出现一次的数字：给定一个非空整数数组，
除了某个元素只出现一次以外，其余每个元素均出现两次。
找出那个只出现了一次的元素。可以使用 for 循环遍历数组，
结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，
然后再遍历 map 找到出现次数为1的元素。
**/

package task1

import (
	"fmt"
	"testing"
)

/*
*
方法一，异或
1.​任何数和 0 异或，仍然是它本身​
a⊕0=a
例如：5 ^ 0 = 5（二进制 101 ^ 000 = 101）

2.​任何数和自身异或，结果为 0​
a⊕a=0
例如：5 ^ 5 = 0（二进制 101 ^ 101 = 000）

3.​异或运算满足交换律和结合律​
a⊕b⊕a=(a⊕a)⊕b=0⊕b=b
这意味着异或的顺序不影响最终结果

复杂度分析：
​时间复杂度​：O(n)，只需遍历一次数组。
​空间复杂度​：O(1)，仅使用常数级别的额外空间。
*
*/
func singleNumber1(nums []int) int {
	result := 0
	for _, v := range nums {
		result ^= v
	}
	fmt.Println("异或结果：", result)
	return result
}

// go test -v .\task_1\leetcode136_test.go -run ^TestSingleNumber1$
func TestSingleNumber1(t *testing.T) {
	nums := []int{22, 22, 11, 33, 33}
	// nums := []int{22, 22, 11, 33, 44}
	singleNumber1(nums)
	// t.Log(result)

}

/*
*
方法二:hash统计
将切片里每个元素的个数都存入hashmap，k=元素值，v=元素个数，元素个数为1的那个就是结果

复杂度分析：
​时间复杂度​：O(n)，遍历两次数组。
​空间复杂度​：O(n)，存储唯一元素集合。
*
*/
func singleNumber2(nums []int) int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	for num, count := range freq {
		if count == 1 {
			fmt.Println("哈希表的计算结果：", num)
			return num
		}
	}
	return -1
}

// go test -v .\task_1\leetcode136_test.go -run ^TestSingleNumber2$
func TestSingleNumber2(t *testing.T) {
	nums := []int{22, 22, 11, 33, 33}
	singleNumber2(nums)
	// t.Log(result)

}
