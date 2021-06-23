/*
63. 不同路径 II
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？



网格中的障碍物和空位置分别用 1 和 0 来表示。

 

示例 1：


输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
输出：2
解释：
3x3 网格的正中间有一个障碍物。
从左上角到右下角一共有 2 条不同的路径：
1. 向右 -> 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右 -> 向右
示例 2：


输入：obstacleGrid = [[0,1],[0,0]]
输出：1
 

提示：

m == obstacleGrid.length
n == obstacleGrid[i].length
1 <= m, n <= 100
obstacleGrid[i][j] 为 0 或 1

*/

/*
    思路：一开始用的递归，摸索
        最后一个case超时，好吧，需要用动规去记录状态。
        动规：想这种网格，动规的状态转移方程还是比较好看出来的。
        dp[i][j]=dp[i-1][j]+dp[i][j-1]
        此时的重点就是防止数组溢出。
        如果i>0,j>0.则dp[i][j]=dp[i-1][j]+dp[i][j-1]
        如果j!=0,则dp[i][j]=dp[i][j-1]
        如果i!=0,则d[i][j]=dp[i-1][j]
        如果i=0,j=0,则dp[i][j]=1
        
        以及二维数组的初始化也要记住：
        let dp=Array.from(Array(len),()=>Array(len).fill(0))
    题解：降维动态规划，每一行每一行滚动记录。。。
        dp[j]原数为上一行的值
        dp[j-1]为当前行左边的值
        因此dp[j]=dp[j]+dp[j-1]

*/

var uniquePathsWithObstacles = function(obstacleGrid) {
    let column=obstacleGrid.length,row=obstacleGrid[0].length,count=0
    let dfs=(i,j)=>{
        if(i<0||j<0||i>=column||j>=row||obstacleGrid[i][j]==1){
            return;
        }
        if(i==column-1&&j==row-1){
            count++;
            return;
        }
        obstacleGrid[i][j]=1
        dfs(i+1,j)//向下
        dfs(i,j+1)//向右
        obstacleGrid[i][j]=0
    }
    dfs(0,0)
    return count
};
// @lc code=end

//递归超时，来吧，需要动态规划


var uniquePathsWithObstacles = function(obstacleGrid) {
    let column=obstacleGrid.length,row=obstacleGrid[0].length,count=0
    if(obstacleGrid[0][0]==1) return 0
    let dp=Array.from(Array(column),()=>Array(row).fill(0))
    for(let i=0;i<column;i++){
        for(let j=0;j<row;j++){
            if(obstacleGrid[i][j]==1){
                dp[i][j]=0
                continue
            }
            if(i>0&&j>0){
                dp[i][j]=dp[i-1][j]+dp[i][j-1]
            }
            else if(j!==0){
                dp[i][j]=dp[i][j-1]
            }
            else if(i!==0){
                dp[i][j]=dp[i-1][j]
            }
            else{
                dp[i][j]=1
            }
        }
    }
    return dp[column-1][row-1]
};
/*
    时间复杂度:O(M*N)
    空间复杂度:O(M*N)
 */

// 空间复杂度优化:降维
var uniquePathsWithObstacles = function(obstacleGrid) {
    let column=obstacleGrid.length,row=obstacleGrid[0].length,count=0
    if(obstacleGrid[0][0]==1) return 0
    let result=Array(row).fill(0)
    result[0]=1
    for(let i=0;i<column;i++){
        for(let j=0;j<row;j++){
            if(obstacleGrid[i][j]==1){
                result[j]=0
                continue
            }
            if(j>0){
                result[j]+=result[j-1]
            }
            
        }
    }
    return result[row-1]
};

/*
    时间复杂度:O(M*N)
    空间复杂度:O(M)
 */

/*
    算法是个不断优化的过程。
*/