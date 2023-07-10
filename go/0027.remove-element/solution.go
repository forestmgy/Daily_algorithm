// Created by maguangyu at 2023/07/10 15:59
// leetgo: 1.3.3
// https://leetcode.cn/problems/remove-element/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func removeElement(nums []int, val int) (ans int) {
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[ans] = nums[fast]
			// 和fast一起走
			ans++
		}
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	val := Deserialize[int](ReadLine(stdin))
	ans := removeElement(nums, val)

	fmt.Println("\noutput:", Serialize(ans))
}
