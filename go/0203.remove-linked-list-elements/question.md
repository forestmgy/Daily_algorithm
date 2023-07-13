# [203. 移除链表元素][link] (Easy)

[link]: https://leetcode.cn/problems/remove-linked-list-elements/

给你一个链表的头节点 `head` 和一个整数 `val` ，请你删除链表中所有满足 `Node.val == val` 的节点，并返
回 **新的头节点** 。

**示例 1：**

![](https://assets.leetcode.com/uploads/2021/03/06/removelinked-list.jpg)

```
输入：head = [1,2,6,3,4,5,6], val = 6
输出：[1,2,3,4,5]

```

**示例 2：**

```
输入：head = [], val = 1
输出：[]

```

**示例 3：**

```
输入：head = [7,7,7,7], val = 7
输出：[]

```

**提示：**

- 列表中的节点数目在范围 `[0, 10⁴]` 内
- `1 <= Node.val <= 50`
- `0 <= val <= 50`


## 解答
说两点
1. 创建虚拟头节点的目的是：无需额外处理head节点，使其和其他普通节点一样
2. 在`ans = &ListNode{}` 之前，ans的值为`nil` ，对`ans.Next` 进行赋值会触发空指针异常，所以需要一个初始化操作

```go
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
```