// Created by maguangyu at 2023/07/10 17:14
// leetgo: 1.3.3
// https://leetcode.cn/problems/squares-of-a-sorted-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func sortedSquares(nums []int) (ans []int) {
	left, right, k := 0, len(nums)-1, len(nums)-1
	ans = make([]int, len(nums))
	for left <= right {
		lm, rm := nums[left]*nums[left], nums[right]*nums[right]
		if lm < rm {
			ans[k] = rm
			right--
		} else {
			ans[k] = lm
			left++
		}
		k--
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := sortedSquares(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
