// Created by maguangyu at 2023/08/12 15:47
// leetgo: 1.3.3
// https://leetcode.cn/problems/intersection-of-two-linked-lists/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func getIntersectionNode(headA, headB *ListNode) (ans *ListNode) {
	for headA != nil {
		tmp := headB
		for tmp != nil {
			if headA == tmp {
				return tmp
			}
			tmp = tmp.Next
		}
		headA = headA.Next
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	intersectVal := Deserialize[int](ReadLine(stdin))
	listA := Deserialize[*ListNode](ReadLine(stdin))
	listB := Deserialize[*ListNode](ReadLine(stdin))
	skipA := Deserialize[int](ReadLine(stdin))
	skipB := Deserialize[int](ReadLine(stdin))
	ans := getIntersectionNode(intersectVal, listA, listB, skipA, skipB)

	fmt.Println("\noutput:", Serialize(ans))
}
