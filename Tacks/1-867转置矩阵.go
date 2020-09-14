/*
 * @lc app=leetcode.cn id=867 lang=golang
 *
 * [867] 转置矩阵
 *
 * https://leetcode-cn.com/problems/transpose-matrix/description/
 *
 * algorithms
 * Easy (67.80%)
 * Likes:    104
 * Dislikes: 0
 * Total Accepted:    25K
 * Total Submissions: 36.9K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给定一个矩阵 A， 返回 A 的转置矩阵。
 *
 * 矩阵的转置是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。
 *
 *
 *
 * 示例 1：
 *
 * 输入：[[1,2,3],[4,5,6],[7,8,9]]
 * 输出：[[1,4,7],[2,5,8],[3,6,9]]
 *
 *
 * 示例 2：
 *
 * 输入：[[1,2,3],[4,5,6]]
 * 输出：[[1,4],[2,5],[3,6]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= A.length <= 1000
 * 1 <= A[0].length <= 1000
 *
 *
 */

// @lc code=start
func transpose1(A [][]int) [][]int {

	// 获取矩阵 行列
	var row int = len(A)
	var col int = len(A[0])
	// 初始化
	res := make([][]int, col)
	// 遍历初始化
	for i, _ := range res {
		res[i] = make([]int, row) //  分配二维矩阵空间
	}

	// 双层for循环
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			res[j][i] = A[i][j]
		}
	}

	return res
}

func transpose(A [][]int) [][]int {

	// 获取矩阵 行列
	row := len(A)
	col := len(A[0])

	// 初始化
	var res [][]int

	// 双层for循环
	for i := 0; i < row; i++ {
		var temp []int // 申请切片
		for j := 0; j < col; j++ {
			temp = append(temp, A[j][i]) // 矩阵转置
		}
		res = append(res, temp) // 采用append进行追加到切片元素中
	}

	return res
}

// @lc code=end

/*

Go 语言切片是对数组的抽象。

思路：
    1、采用双重循环，拿到二维切片也就是矩阵的某个元素。
    2、然后重新赋值给一个新的切片，只是在赋值的时候，复制到斜对角的位置。
	3、那么斜对角其实就是 i、j对换位置。
	4、同时注意go语言的初始化赋值，切片用make申请

时间复杂度：
    跟矩阵的m、n有关，如m行n列的矩阵，那么时间复杂度就是 O(m*n)

语言特性：

- 变量
	声明变量的一般形式是使用 var 关键字：
		例如
			var a int = 10

		变量的初始化时省略变量的类型而由系统自动推断，声明语句写上 var 关键字其实是显得有些多余了

	简写
		简短形式，使用 := 赋值操作符，例如a := 10
		它只能被用在函数体内，而不可以用于全局变量的声明与赋值。使用操作符 := 可以高效地创建一个新的变量，称之为初始化声明。


- 切片
	Go 数组的长度不可改变，在特定场景中这样的集合就不太适用。
	Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

	例如：声明一个slice切片的初始长度为len，并且存储int类型元素
		声明：var slice1 []int = make([]int, len)
		简写：slice2 := make([]int, len)

	切片是可索引的，并且可以由 len() 方法获取长度。所以我们先根据len方法获取到矩阵的行列，同时再遍历赋值的时候，我们采用切片索引跟数组类似进行赋值。

【1】
	提前采用make准确的进行初始化切片。
【2】
	声明切片可以不需要知道长度。
	使用append()去追加切片元素，因为你可能不知道需要申请多大的切片。


*/