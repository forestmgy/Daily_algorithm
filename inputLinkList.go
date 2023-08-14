package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println("请输入一系列整数，用空格分隔:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	nums := parseInput(input)
	head := buildLinkedList(nums)
	printLinkedList(head)
}

// 解析输入字符串为整数切片
func parseInput(input string) []int {
	input = strings.TrimSpace(input)
	strNums := strings.Split(input, " ")
	nums := make([]int, len(strNums))
	for i, str := range strNums {
		nums[i], _ = strconv.Atoi(str)
	}
	return nums
}

// 构建链表
func buildLinkedList(nums []int) *ListNode {
	var head, prev *ListNode
	for _, num := range nums {
		node := &ListNode{Val: num}
		if head == nil {
			head = node
		} else {
			prev.Next = node
		}
		prev = node
	}
	return head
}

// 打印链表
func printLinkedList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Printf("%d -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Println("nil")
}
