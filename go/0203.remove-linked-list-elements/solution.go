// Created by maguangyu at 2023/07/11 17:15
// leetgo: 1.3.3
// https://leetcode.cn/problems/remove-linked-list-elements/

package main

import (
	"bufio"
	"fmt"
	. "github.com/j178/leetgo/testutils/go"
	"os"
)

// @lc code=begin

func removeElements(head *ListNode, val int) (ans *ListNode) {
	if head == nil {
		return nil
	}
	ans = &ListNode{}
	ans.Next = head
	var current *ListNode
	current = ans
	for current != nil && current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return ans.Next
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	val := Deserialize[int](ReadLine(stdin))
	ans := removeElements(head, val)

	fmt.Println("\noutput:", Serialize(ans))
}
