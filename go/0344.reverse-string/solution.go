// Created by maguangyu at 2023/08/14 10:05
// leetgo: 1.3.3
// https://leetcode.cn/problems/reverse-string/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left, right = left+1, right-1
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[[]byte](ReadLine(stdin))
	reverseString(s)
	ans := s

	fmt.Println("\noutput:", Serialize(ans))
}
