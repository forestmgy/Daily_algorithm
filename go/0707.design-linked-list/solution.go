// Created by maguangyu at 2023/07/13 10:26
// leetgo: 1.3.3
// https://leetcode.cn/problems/design-linked-MyLinkedList/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type MyNode struct {
	val  int
	next *MyNode
}

type MyLinkedList struct {
	fakeHead *MyNode
	size     int
}

func Constructor() MyLinkedList {
	return MyLinkedList{fakeHead: new(MyNode), size: 0}
}

func (m *MyLinkedList) Get(index int) (ans int) {
	if m.fakeHead.next == nil || index < 0 || index > m.size || m.size == 0 {
		return -1
	}
	tmp := m.fakeHead.next
	for i := 0; i < index; i++ {
		if tmp.next == nil {
			return -1
		}
		tmp = tmp.next
	}

	return tmp.val
}

func (m *MyLinkedList) AddAtHead(val int) {
	tmp := &MyNode{val: val}
	tmp.next = m.fakeHead.next
	m.fakeHead.next = tmp
	m.size++
}

func (m *MyLinkedList) AddAtTail(val int) {
	cur := m.fakeHead
	for cur.next != nil {
		cur = cur.next
	}
	tail := MyNode{}
	tail.val = val
	cur.next = &tail
	m.size++
}

func (m *MyLinkedList) AddAtIndex(index int, val int) {
	if index > m.size {
		return
	} else if index < 0 {
		index = 0
	}
	cur := m.fakeHead
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	tmp := MyNode{}
	tmp.val = val
	tmp.next = cur.next
	cur.next = &tmp
	m.size++
}

func (m *MyLinkedList) DeleteAtIndex(index int) {
	if index >= m.size {
		return
	}
	cur := m.fakeHead.next
	if index == 0 {
		m.fakeHead.next = m.fakeHead.next.next
		return
	}
	for i := 0; i < index-1; i++ {
		cur = cur.next
	}
	if index == m.size-1 {
		cur.next = nil
		return
	}
	cur.next = cur.next.next
	m.size--
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	ops := Deserialize[[]string](ReadLine(stdin))
	params := MustSplitArray(ReadLine(stdin))
	output := make([]string, 0, len(ops))
	output = append(output, "null")

	obj := Constructor()

	for i := 1; i < len(ops); i++ {
		switch ops[i] {
		case "get":
			methodParams := MustSplitArray(params[i])
			index := Deserialize[int](methodParams[0])
			ans := Serialize(obj.Get(index))
			output = append(output, ans)
		case "addAtHead":
			methodParams := MustSplitArray(params[i])
			val := Deserialize[int](methodParams[0])
			obj.AddAtHead(val)
			output = append(output, "null")
		case "addAtTail":
			methodParams := MustSplitArray(params[i])
			val := Deserialize[int](methodParams[0])
			obj.AddAtTail(val)
			output = append(output, "null")
		case "addAtIndex":
			methodParams := MustSplitArray(params[i])
			index := Deserialize[int](methodParams[0])
			val := Deserialize[int](methodParams[1])
			obj.AddAtIndex(index, val)
			output = append(output, "null")
		case "deleteAtIndex":
			methodParams := MustSplitArray(params[i])
			index := Deserialize[int](methodParams[0])
			obj.DeleteAtIndex(index)
			output = append(output, "null")
		}
	}
	fmt.Println("\noutput:", JoinArray(output))
}
