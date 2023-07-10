# [977. 有序数组的平方][link] (Easy)

[link]: https://leetcode.cn/problems/squares-of-a-sorted-array/

给你一个按 **非递减顺序** 排序的整数数组 `nums`，返回 **每个数字的平方** 组成的新数组，要求也按 **非
递减顺序** 排序。

**示例 1：**

```
输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]
```

**示例 2：**

```
输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]

```

**提示：**

- `1 <= nums.length <= 10⁴`
- `-10⁴ <= nums[i] <= 10⁴`
- `nums` 已按 **非递减顺序** 排序

**进阶：**

- 请你设计时间复杂度为 `O(n)` 的算法解决本问题

## 解答

**常规**
平方，排序

**进阶**
因为不需要去重，所以输入数组和输出数组的长度一样
可以两个指针指向头部和尾部，如果有负数的话，是两边最大，中间最小，从两边向中间遍历就好
```go
func sortedSquares(nums []int) (ans []int) {
	left, right, k := 0, len(nums)-1, len(nums)-1
	ans = make([]int, len(nums))
	for left <= right {
		lm, rm := nums[left]*nums[left], nums[right]*nums[right]
		if lm < rm {
			ans[k] = rm
			right--
		} else {
			ans[k] = lm
			left++
		}
		k--
	}

	return
}
```
