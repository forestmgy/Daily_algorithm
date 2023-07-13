// Created by maguangyu at 2023/07/11 11:06
// leetgo: 1.3.3
// https://leetcode.cn/problems/minimum-size-subarray-sum/

package main

import (
	"fmt"
	"math"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minSubArrayLen(target int, nums []int) (ans int) {
	var sum, left, right int
	ans = math.MaxInt
	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			ans = min(ans, right-left+1)
			sum -= nums[left]
			left++
		}
		right++

	}
	if ans == math.MaxInt {
		ans = 0
	}
	return
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// target := Deserialize[int](ReadLine(stdin))
	// nums := Deserialize[[]int](ReadLine(stdin))
	target := 11
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1}
	ans := minSubArrayLen(target, nums)

	fmt.Println("\noutput:", Serialize(ans))
}
