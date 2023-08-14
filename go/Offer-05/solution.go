// Created by maguangyu at 2023/08/14 10:41
// leetgo: 1.3.3
// https://leetcode.cn/problems/ti-huan-kong-ge-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func replaceSpace(s string) string {
	result := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			result = append(result, []byte("%20")...)
		} else {
			result = append(result, s[i])
		}
	}
	return string(result)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := replaceSpace(s)

	fmt.Println("\noutput:", Serialize(ans))
}
