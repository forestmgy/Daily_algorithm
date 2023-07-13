// Created by maguangyu at 2023/07/13 10:26
// leetgo: 1.3.3
// https://leetcode.cn/problems/design-linked-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type MyLinkedList struct {

}

func Constructor() MyLinkedList {

	return MyLinkedList{}
}

func (m *MyLinkedList) Get(index int) (ans int) {

	return
}

func (m *MyLinkedList) AddAtHead(val int)  {

}

func (m *MyLinkedList) AddAtTail(val int)  {

}

func (m *MyLinkedList) AddAtIndex(index int, val int)  {

}

func (m *MyLinkedList) DeleteAtIndex(index int)  {

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
