// Created by maguangyu at 2023/07/11 15:54
// leetgo: 1.3.3
// https://leetcode.cn/problems/spiral-matrix-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func generateMatrix(n int) (ans [][]int) {
	if n == 0 {
		return nil
	}
	num := 1
	x, y := 0, 0
	center := n / 2
	loop := n / 2
	offset := 1
	ans = make([][]int, n)

	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}
	for loop > 0 {
		i, j := x, y
		for j = y; j < n-offset; j++ {
			ans[i][j] = num
			num++
		}

		for i = x; i < n-offset; i++ {
			ans[i][j] = num
			num++
		}
		for ; j > y; j-- {
			ans[i][j] = num
			num++
		}
		for ; i > x; i-- {
			ans[i][j] = num
			num++
		}

		loop--
		offset++
		x++
		y++

	}

	if n%2 == 1 {
		ans[center][center] = n * n
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	ans := generateMatrix(n)

	fmt.Println("\noutput:", Serialize(ans))
}
