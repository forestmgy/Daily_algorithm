// Created by maguangyu at 2023/08/14 13:17
// leetgo: 1.3.3
// https://leetcode.cn/problems/reverse-words-in-a-string/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverseWords(s string) string {
	b := []byte(s)
	//快慢指针去除所有空格，如果遇到空格去除，但是如果slow不是0，说明已经过了第一个单词，需要在前面补一个空格
	slow := 0 // slow 指向下一个即将改变的单词
	for i := 0; i < len(b); i++ {
		if b[i] != ' ' {
			if slow != 0 {
				b[slow] = ' '
				slow++
			}
			for i < len(b) && b[i] != ' ' {
				b[slow] = b[i]
				slow++
				i++
			}
		}
	}
	b = b[:slow]
	reverse(b)
	last := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			reverse(b[last:i])
			last = i + 1
		}
	}
	return string(b)
}

func reverse(b []byte) {
	left, right := 0, len(b)-1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := reverseWords(s)

	fmt.Println("\noutput:", Serialize(ans))
}
