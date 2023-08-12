// Created by maguangyu at 2023/08/12 14:13
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
	slow, fast := head, head.Next
	cur := ans
	for fast != nil {
		slow.Next = fast.Next
		fast.Next = slow
		cur.Next = fast
		cur = slow
		if slow.Next == nil {
			break
		}
		slow = slow.Next
		if slow.Next == nil {
			break
		}
		fast = slow.Next

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
