package task1

/**

编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。
**/

import (
	"fmt"
	"testing"
)

func longestCommonPrefix(strs []string) string {
	//设置数组里第一个元素为基准字符串
	baseStr := strs[0]
	//循环基准字符串的每个字符元素
	for i, baseChar := range baseStr {
		//循环其他字符串
		for _, str := range strs[1:] {
			//1.其他字符串长度不足
			//2.基础字符串的i位置字符与其他字符串的对应位置不一样
			if i > len(str) || baseChar != rune(str[i]) {
				return baseStr[:i]
			}
		}
	}

	return baseStr
}

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	result := longestCommonPrefix(strs)
	fmt.Print("最长前缀:", result)
}
