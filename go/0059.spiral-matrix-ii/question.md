# [59. 螺旋矩阵 II][link] (Medium)

[link]: https://leetcode.cn/problems/spiral-matrix-ii/

给你一个正整数 `n` ，生成一个包含 `1` 到 `n²` 所有元素，且元素按顺时针顺序螺旋排列的 `n x n` 正方形
矩阵 `matrix` 。

**示例 1：**

![](https://assets.leetcode.com/uploads/2020/11/13/spiraln.jpg)

```
输入：n = 3
输出：[[1,2,3],[8,9,4],[7,6,5]]

```

**示例 2：**

```
输入：n = 1
输出：[[1]]

```

**提示：**

- `1 <= n <= 20`

## 解答
大模拟，如果n为奇数，循环奇数次，最后一次直接给中间变量赋值
如果n为偶数，循环n/2次
每次循环，一条边走过的空格数都比上一次少1，所以offset要一直加1
循环起点每次都是最左上角，直接每个都加1就行了

```go
func generateMatrix(n int) (ans [][]int) {
	if n == 0 {
		return nil
	}
	num := 1
	x, y := 0, 0
	center := n / 2
	loop := n / 2
	offset := 1
	ans = make([][]int, n)

	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}
	for loop > 0 {
		i, j := x, y
		for j = y; j < n-offset; j++ {
			ans[i][j] = num
			num++
		}

		for i = x; i < n-offset; i++ {
			ans[i][j] = num
			num++
		}
		for ; j > y; j-- {
			ans[i][j] = num
			num++
		}
		for ; i > x; i-- {
			ans[i][j] = num
			num++
		}

		loop--
		offset++
		x++
		y++

	}

	if n%2 == 1 {
		ans[center][center] = n * n
	}
	return
}
```