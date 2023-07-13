# [209. 长度最小的子数组][link] (Medium)

[link]: https://leetcode.cn/problems/minimum-size-subarray-sum/

给定一个含有 `n` 个正整数的数组和一个正整数 `target` **。**

找出该数组中满足其和 `≥ target` 的长度最小的 **连续子数组** `[numsₗ, numsₗ₊₁, ..., numsᵣ₋₁, numsᵣ]` 
，并返回其长度 **。** 如果不存在符合条件的子数组，返回 `0` 。

**示例 1：**

```
输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。

```

**示例 2：**

```
输入：target = 4, nums = [1,4,4]
输出：1

```

**示例 3：**

```
输入：target = 11, nums = [1,1,1,1,1,1,1,1]
输出：0

```

**提示：**

- `1 <= target <= 10⁹`
- `1 <= nums.length <= 10⁵`
- `1 <= nums[i] <= 10⁵`

**进阶：**

- 如果你已经实现 `O(n)` 时间复杂度的解法, 请尝试设计一个 `O(n log(n))` 时间复杂度的解法。

## 解答
我感觉这个代码还是挺好懂的，代码功力还需要增强
```go
func minSubArrayLen(target int, nums []int) (ans int) {
	var sum, left, right int
	ans = math.MaxInt
	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			ans = min(ans, right-left+1)
			sum -= nums[left]
			left++
		}
		right++

	}
	if ans == math.MaxInt {
		ans = 0
	}
	return
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
```