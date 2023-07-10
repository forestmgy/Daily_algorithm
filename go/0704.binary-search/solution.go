// Created by maguangyu at 2023/07/10 15:01
// leetgo: 1.3.3
// https://leetcode.cn/problems/binary-search/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func search(nums []int, target int) (ans int) {
	ans = -1
	low := 0
	high := len(nums) - 1
	for low <= high {
		// (low + high) / 2 可能会越界：low+high 太大了
		mid := low + (high-low)/2
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			ans = mid
			return
		}
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := search(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
