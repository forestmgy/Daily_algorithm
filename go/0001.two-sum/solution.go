// Created by maguangyu at 2023/07/10 12:08
// leetgo: 1.3.3
// https://leetcode.cn/problems/two-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func twoSum(nums []int, target int) (ans []int) {
	hashMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := hashMap[target-nums[i]]; ok {
			return []int{i, v}
		}
		hashMap[nums[i]] = i
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := twoSum(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
