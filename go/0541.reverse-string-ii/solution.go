// Created by maguangyu at 2023/08/14 10:12
// leetgo: 1.3.3
// https://leetcode.cn/problems/reverse-string-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverseStr(s string, k int) string {
	count := 0
	length := len(s)
	str := []byte(s)
	for i := 0; i < length; i++ {
		leftCount := length - count
		if leftCount < k {
			reverse(str, count, length-1)
			break
		} else if leftCount < k*2 && leftCount >= k {
			reverse(str, count, count+k-1)
			break
		} else {
			reverse(str, count, count+k-1)
			count += 2 * k
		}
	}
	return string(str)
}

func reverse(s []byte, left, right int) {
	for left < right {
		s[left], s[right] = s[right], s[left]
		left, right = left+1, right-1
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := reverseStr(s, k)

	fmt.Println("\noutput:", Serialize(ans))
}
