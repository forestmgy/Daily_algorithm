// Created by maguangyu at 2023/08/13 18:33
// leetgo: 1.3.3
// https://leetcode.cn/problems/4sum-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) (ans int) {
	length := len(nums1)
	map1 := make(map[int]int)
	map2 := make(map[int]int)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			map1[nums1[i]+nums2[j]] += 1
			map2[nums3[i]+nums4[j]] += 1
		}
	}
	for k, _ := range map1 {
		if v, ok := map2[0-k]; ok {
			ans += v * map1[k]
		}
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums1 := Deserialize[[]int](ReadLine(stdin))
	nums2 := Deserialize[[]int](ReadLine(stdin))
	nums3 := Deserialize[[]int](ReadLine(stdin))
	nums4 := Deserialize[[]int](ReadLine(stdin))
	ans := fourSumCount(nums1, nums2, nums3, nums4)

	fmt.Println("\noutput:", Serialize(ans))
}
