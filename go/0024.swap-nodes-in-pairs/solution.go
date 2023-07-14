// Created by maguangyu at 2023/07/13 18:50
// leetgo: 1.3.3
// https://leetcode.cn/problems/swap-nodes-in-pairs/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func swapPairs(head *ListNode) (ans *ListNode) {
	if head == nil || head.Next == nil {
		return head
	}
	ans = new(ListNode)
	ans.Next = head
	pre := ans

	for head != nil && head.Next != nil {
		pre.Next = head.Next
		next := head.Next.Next
		head.Next.Next = head
		head.Next = next
		pre = head
		head = next
	}
	return ans.Next
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	ans := swapPairs(head)

	fmt.Println("\noutput:", Serialize(ans))
}
