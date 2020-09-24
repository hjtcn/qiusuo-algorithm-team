// 3 x 3 的幻方是一个填充有从 1 到 9 的不同数字的 3 x 3 矩阵，其中每行，每列以及两条对角线上的各数之和都相等。

// 给定一个由整数组成的 grid，其中有多少个 3 × 3 的 “幻方” 子矩阵？（每个子矩阵都是连续的）。

//  

// 示例：

// 输入: [[4,3,8,4],
//       [9,5,1,9],
//       [2,7,6,2]]
// 输出: 1
// 解释: 
// 下面的子矩阵是一个 3 x 3 的幻方：
// 438
// 951
// 276

// 而这一个不是：
// 384
// 519
// 762

/**
 * @param {number[][]} grid
 * @return {number}
 */
var numMagicSquaresInside = function (grid) {
    let res=0
    let row = grid.length, col = grid.length;
    for (let i = 0; i < row - 2; i++) {
        for (let j = 0; j < col - 2; j++) {
            //判断该3X3的矩阵元素是否存在大于9或小于1的数
            if(ifContinue(grid[i][j])||ifContinue(grid[i][j + 1])||ifContinue(grid[i][j + 2])||ifContinue(grid[i + 1][j])||ifContinue(grid[i + 1][j + 1])||ifContinue(grid[i + 1][j + 2])||ifContinue(grid[i + 2][j])||ifContinue(grid[i + 2][j + 1])||ifContinue(grid[i + 2][j + 2])){
               continue;
            }
            let sum1 = grid[i][j] + grid[i][j + 1] + grid[i][j + 2]//第一列
            let sum2 = grid[i + 1][j] + grid[i + 1][j + 1] + grid[i + 1][j + 2]//第二列
            let sum3 = grid[i + 2][j] + grid[i + 2][j + 1] + grid[i + 2][j + 2]//第三列
            let sum4 = grid[i][j] + grid[i+1][j ] + grid[i+2][j]
            let sum5 = grid[i][j+1] + grid[i + 1][j + 1] + grid[i + 2][j +1]
            let sum6 = grid[i][j+2] + grid[i + 1][j + 2] + grid[i + 2][j + 2]
            let sum7 = grid[i][j] + grid[i + 1][j + 1] + grid[i + 2][j + 2]
            let sum8 = grid[i+2][j] + grid[i + 1][j + 1] + grid[i][j + 2] 
            //计算行列斜的八个和，且判断是否为15
            //grid[i][j]!=5是为了去除所有数字都为5的情况
            if(grid[i][j]!==5&&sum1==15&&sum2==15&&sum3==15&&sum4==15&&sum5==15&&sum6==15&&sum7==15&&sum8==15){
                res+=1
            }
        }
    }
    return res
};
function ifContinue(item){//判断数值是否大于9或者小于1
    return item>9||item<1
}

/**题解
 暴力解决。双层遍历行列，注意防止数组溢出，并且计算三行三列两斜线的和，判断和都为15，特殊情况，判断各元素中是否有包含大于9或者小于1的数。以及所有元素为5的情况
  复杂度分析：
    时间复杂度是:O(N*M)
    双层遍历行列的值
    空间复杂度：O(1）
    声明几个变量，存储8个和
 */

