// Created by maguangyu at 2023/08/13 18:49
// leetgo: 1.3.3
// https://leetcode.cn/problems/ransom-note/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func canConstruct(ransomNote string, magazine string) bool {
	map1 := make(map[byte]int)
	map2 := make(map[byte]int)
	for i := 0; i < len(ransomNote); i++ {
		map1[ransomNote[i]] += 1
	}
	for i := 0; i < len(magazine); i++ {
		map2[magazine[i]] += 1
	}
	for i := 0; i < len(ransomNote); i++ {
		if v, ok := map2[ransomNote[i]]; ok {
			if v > 0 {
				map2[ransomNote[i]]--
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	ransomNote := Deserialize[string](ReadLine(stdin))
	magazine := Deserialize[string](ReadLine(stdin))
	ans := canConstruct(ransomNote, magazine)

	fmt.Println("\noutput:", Serialize(ans))
}
