// 给定一个矩阵 A， 返回 A 的转置矩阵。

// 矩阵的转置是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。

 

// 示例 1：

// 输入：[[1,2,3],[4,5,6],[7,8,9]]
// 输出：[[1,4,7],[2,5,8],[3,6,9]]
// 示例 2：

// 输入：[[1,2,3],[4,5,6]]
// 输出：[[1,4],[2,5],[3,6]]
 

// 提示：

// 1 <= A.length <= 1000
// 1 <= A[0].length <= 1000

/**
 * @param {number[][]} A
 * @return {number[][]}
 */
var transpose = function(A) {
    let row=A.length;
    let col=A[0].length
    let B=[]
    for(let i=0;i<col;i++){
        let C=[]
        for(let j=0;j<row;j++){
          C.push(A[j][i])
        }
        B.push(C)
    }
    return B
};

/** 题解
反转矩阵
定义空数组，拆解原数组，进行填充赋值A[j][i]
复杂度分析：
	时间复杂度：O(R*C)
	与双层数组的ROW和COL有关

	空间复杂度：O(R*C)
	与双层数组的ROW和COL有关
 */


