// Created by maguangyu at 2023/08/14 14:20
// leetgo: 1.3.3
// https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	reverse(b[0:n])
	reverse(b[n:])
	reverse(b)
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
	n := Deserialize[int](ReadLine(stdin))
	ans := reverseLeftWords(s, n)

	fmt.Println("\noutput:", Serialize(ans))
}
