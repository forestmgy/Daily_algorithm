// Created by maguangyu at 2023/08/12 16:40
// leetgo: 1.3.3
// https://leetcode.cn/problems/linked-list-cycle-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func detectCycle(head *ListNode) (ans *ListNode) {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = head
			for slow != fast {
				slow, fast = slow.Next, fast.Next
			}
			return slow
		}
	}
	return nil
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	pos := Deserialize[int](ReadLine(stdin))
	ans := detectCycle(head, pos)

	fmt.Println("\noutput:", Serialize(ans))
}
