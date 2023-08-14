// Created by maguangyu at 2023/08/13 16:39
// leetgo: 1.3.3
// https://leetcode.cn/problems/valid-anagram/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isAnagram(s string, t string) bool {
	if s == "" || t == "" {
		return false
	}
	if len(s) != len(t) {
		return false
	}
	map1 := make(map[byte]int)
	map2 := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		map1[s[i]] += 1
	}
	for i := 0; i < len(t); i++ {
		map2[t[i]] += 1
	}
	for i := 0; i < len(s); i++ {
		if map1[s[i]] != map2[s[i]] {
			return false
		}
	}
	return true
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	t := Deserialize[string](ReadLine(stdin))
	ans := isAnagram(s, t)

	fmt.Println("\noutput:", Serialize(ans))
}
