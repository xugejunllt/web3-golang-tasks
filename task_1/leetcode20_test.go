package task1

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。

**/

import (
	"fmt"
	"testing"
)

func isValid(s string) bool {
	stack := make([]rune, 0) //初始化栈
	//map存储标注括号的键值对，用于校验,k-左括号，v-右括号
	pairs := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, char := range s {
		//如果是左括号，压栈
		if _, exists := pairs[char]; exists {
			stack = append(stack, char)
		} else {
			//如果是右括号，但是栈为空，则没有对应括号
			if len(stack) == 0 {
				return false
			}
			//如果是右括号，弹出栈顶元素，跟此右括号比较是否是一对
			//栈顶元素
			top := stack[len(stack)-1]
			//弹出栈顶元素
			stack = stack[:len(stack)-1]
			//比较栈顶元素和右括号是否是一对
			if pairs[top] != char {
				return false
			}
		}

	}
	//全部匹配结束栈应该为空，不为空的话，证明右括号比左括号多
	return len(stack) == 0
}

// go test -v .\task_1\leetcode20_test.go -run ^TestIsValid$
func TestIsValid(t *testing.T) {
	// s := "()[]{}"
	s := "()[]{}"
	fmt.Print("是否符合要求：", isValid(s))
}
