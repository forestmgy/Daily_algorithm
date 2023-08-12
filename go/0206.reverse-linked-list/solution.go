// Created by maguangyu at 2023/08/12 13:54
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
	if head == nil {
		return nil
	}
	ans = new(ListNode)
	cur := head
	for cur != nil {
		//创建一个零时节点存储当前节点的值
		tmp := &ListNode{Val: cur.Val}
		next := ans.Next
		ans.Next = tmp
		tmp.Next = next
		cur = cur.Next
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
