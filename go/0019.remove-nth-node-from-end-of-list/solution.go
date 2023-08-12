// Created by maguangyu at 2023/08/12 15:27
// leetgo: 1.3.3
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func removeNthFromEnd(head *ListNode, n int) (ans *ListNode) {
	if head == nil || head.Next == nil {
		return nil
	}
	ans = &ListNode{
		Next: head,
	}
	slow, fast := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil { //删除头节点
		ans.Next = ans.Next.Next
		return ans.Next
	}
	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}
	if n == 1 { //删除尾节点
		slow.Next = nil
	} else { //删除常规节点
		slow.Next = slow.Next.Next
	}

	return ans.Next
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	n := Deserialize[int](ReadLine(stdin))
	ans := removeNthFromEnd(head, n)

	fmt.Println("\noutput:", Serialize(ans))
}
