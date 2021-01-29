//在经典汉诺塔问题中，有 3 根柱子及 N 个不同大小的穿孔圆盘，盘子可以滑入任意一根柱子。一开始，所有盘子自上而下按升序依次套在第一根柱子上(即每一个盘子只
//能放在更大的盘子上面)。移动圆盘时受到以下限制:
//(1) 每次只能移动一个盘子;
//(2) 盘子只能从柱子顶端滑出移到下一根柱子;
//(3) 盘子只能叠在比它大的盘子上。
//
// 请编写程序，用栈将所有盘子从第一根柱子移到最后一根柱子。
//
// 你需要原地修改栈。
//
// 示例1:
//
//  输入：A = [2, 1, 0], B = [], C = []
// 输出：C = [2, 1, 0]
//
//
// 示例2:
//
//  输入：A = [1, 0], B = [], C = []
// 输出：C = [1, 0]
//
//
// 提示:
//
//
// A中盘子的数目不大于14个。
//
// Related Topics 递归
// 👍 75 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func hanota(A []int, B []int, C []int) []int {
	if A == nil {
		return nil
	}

	move(len(A), &A, &B, &C)
	return C
}
//leetcode submit region end(Prohibit modification and deletion)

func move(n int, A *[]int, B *[]int, C *[]int)  {
	if n == 0 {
		return
	}
	if n == 1 {
		*C = append(*C, (*A)[len(*A) - 1])
		*A = (*A)[:len(*A) - 1]
	} else {
		move(n - 1, A, C, B)
		move(1, A, B, C)
		move(n - 1, B, A, C)
	}
}

/**
递归分治：
n = 1,只有一个盘子，直接从A移动到C
n > 1
	先把最上面 n - 1个盘子从A移动到B（依次分解为子问题，递归）
	再将最大的盘子从A移动到C
	在将B上n - 1个盘子从B移动到C（依次分解为子问题，递归）

尝试去想递归的每一步真的好难，将问题分解子问题，分析子问题就行
指针的运算有点生疏

复杂度分析：
	时间复杂度：O(2^n)
	空间复杂度：O(1)
*/
