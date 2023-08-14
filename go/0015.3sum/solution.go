// Created by maguangyu at 2023/08/13 19:19
// leetgo: 1.3.3
// https://leetcode.cn/problems/3sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func threeSum(nums []int) (ans [][]int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := threeSum(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
