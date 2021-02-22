/*
在经典汉诺塔问题中，有 3 根柱子及 N 个不同大小的穿孔圆盘，盘子可以滑入任意一根柱子。一开始，所有盘子自上而下按升序依次套在第一根柱子上(即每一个盘子只能放在更大的盘子上面)。移动圆盘时受到以下限制:
(1) 每次只能移动一个盘子;
(2) 盘子只能从柱子顶端滑出移到下一根柱子;
(3) 盘子只能叠在比它大的盘子上。
请编写程序，用栈将所有盘子从第一根柱子移到最后一根柱子。
你需要原地修改栈。
示例1:
 输入：A = [2, 1, 0], B = [], C = []
 输出：C = [2, 1, 0]
示例2:
 输入：A = [1, 0], B = [], C = []
 输出：C = [1, 0]
提示:
A中盘子的数目不大于14个。
*/

1. Clearfication:
分析题目意思：
1.1.  每次只能移动一个盘子
1.2. 盘子只能从柱子顶端滑出移到下一根柱子
1.3  盘子只能叠在比它大的盘子上
如果我们手动模拟的话，我们可以将2移动到B，然后将1移动到B，再然后将0移动到C，移动B的1移动到C,最后移动B的2移动到C,这个时候 C的结果就是 [2,1,0] 了
如何将这个过程模拟，然后让计算机去帮我们执行这个过程呢？
A = [2,1,0], B = [],C = []
C = [2,1,0]

func hanota(A []int,B []int,C []int)[]int {
    if A == nil {
        return nil
    }
    move(len(A), &A,&B,&C)
    return C
}
func move(n int,A *[]int,B *[]int,C *[]int) {
    if n == 0 {
        return
    }
    
    if n == 1 {
        // 最小问题处理，只有1个盘子，从A->C
        *C = append(*C, (*A)[len(*A) - 1])
        *A = (*A)[:len(*A)-1]
    }else {
        // 小一级问题处理,把n个盘子从A通过B移动C=1,
        // 1. 先把n-1个盘子从A通过C移动到B，
        // 2. 把A的最后一个盘子移动到C
        // 3. 把n-1个盘子从B通过A移动到C
        move(n - 1,A,C,B)
        move(1, A,B,C)
        move(n - 1,B,A,C)
    }
}

复杂度分析：
时间复杂度：O(2^n) : 递归时间复杂度
空间复杂度：O(2^n)：递归使用的空间复杂度

复杂度分析不一定是对的，后面可以再看下

总结：
1. 通过这个问题看到自己缺少将大问题化为小问题的能力

2. 没有掌握分析递归问题的解决方法，没有从以小见大的思考方法和方法去解决问题

3. 一开始看题解看不懂的时候，去动手，去画图，不要去陷入自己不会的场景中
