// Created by maguangyu at 2023/08/13 17:49
// leetgo: 1.3.3
// https://leetcode.cn/problems/intersection-of-two-arrays/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func intersection(nums1 []int, nums2 []int) (ans []int) {
	map2 := make(map[int]bool)
	map3 := make(map[int]bool)

	for i := 0; i < len(nums2); i++ {
		map2[nums2[i]] = true
	}
	for i := 0; i < len(nums1); i++ {
		if map2[nums1[i]] {
			map3[nums1[i]] = true
		}
	}
	for k, _ := range map3 {
		ans = append(ans, k)
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums1 := Deserialize[[]int](ReadLine(stdin))
	nums2 := Deserialize[[]int](ReadLine(stdin))
	ans := intersection(nums1, nums2)

	fmt.Println("\noutput:", Serialize(ans))
}
