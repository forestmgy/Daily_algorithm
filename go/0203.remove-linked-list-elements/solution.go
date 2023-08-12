// Created by maguangyu at 2023/08/12 13:17
// leetgo: 1.3.3
// https://leetcode.cn/problems/remove-linked-list-elements/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func removeElements(head *ListNode, val int) (ans *ListNode) {
	if head == nil {
		return
	}
	ans = new(ListNode)
	ans.Next = head
	cur := ans
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			if cur.Next.Next != nil {
				cur.Next = cur.Next.Next
			} else {
				cur.Next = nil
			}
			continue
		}
		cur = cur.Next

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
