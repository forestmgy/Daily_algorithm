// Created by maguangyu at 2023/07/13 14:43
// leetgo: 1.3.3
// https://leetcode.cn/problems/reverse-linked-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverseList(head *ListNode) (ans *ListNode) {
	ans = new(ListNode)
	for head != nil {
		tmp := &ListNode{Val: head.Val}
		tmp.Next = ans.Next
		ans.Next = tmp
		head = head.Next
	}
	return ans.Next
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	ans := reverseList(head)

	fmt.Println("\noutput:", Serialize(ans))
}
