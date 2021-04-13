// 64. 最小路径和
// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

// 说明：每次只能向下或者向右移动一步。

 

// 示例 1：


// 输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
// 输出：7
// 解释：因为路径 1→3→1→1→1 的总和最小。
// 示例 2：

// 输入：grid = [[1,2,3],[4,5,6]]
// 输出：12
 

// 提示：

// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 200
// 0 <= grid[i][j] <= 100


/*
    思路：这个题和昨天的题查询过程差不多。但是今天的需要加上当前val,保持最小和。
        首先画状态转移过程
        1，3，1            1,   (3+1),(4+1)
        1，5，1    -----> (1+1),(2+5),(5+1)
        4，2，1           (2+4),(6+2),(6+1)

        dp[i][j]=Math.min(dp[i-1][j],dp[i][j-1])+dp[i][j]
        难点在哪？
        二维数组的溢出问题。。。
        首先我写的时候，想要直接修改输入数组。
        但是dp[i-1][j]和dp[i][j-1],当i=0或者j=0时溢出，需要特殊处理
*/

var minPathSum = function(grid) {
    let m=grid.length,n=grid[0].length
    for(let i=0;i<grid.length;i++){
        for(let j=0;j<grid[0].length;j++){
            if(i==0&&j==0){
                continue;
            }
            else if(i==0){
                 grid[i][j]=grid[i][j-1]+grid[i][j]
            }
            else if(j==0){
                  grid[i][j]=grid[i-1][j]+grid[i][j]
            }
            else{
                 grid[i][j]=Math.min(grid[i-1][j],grid[i][j-1])+grid[i][j]
            }
         
        }
    }
    return grid[m-1][n-1]
};

/*
    时间复杂度：O(M*N)
    空间复杂度：O(1).直接改输入数组，不需要额外的空间复杂度
*/

//降维处理

var minPathSum = function(grid) {
    let dp=[]
    for(let i=0;i<grid.length;i++){
        for(let j=0;j<grid[0].length;j++){
            if(i==0&&j==0){
                dp[j]=grid[i][j]
            }
            else if(i==0){
                dp[j]=dp[j-1]+grid[i][j]
            }
            else if(j==0){
                dp[j]=dp[j]+grid[i][j]
            }
            else{
                dp[j]=Math.min(dp[j-1],dp[j])+grid[i][j]
            }
        }
    }
    return dp[grid[0].length-1]
};

/*
降维的状态转移不太好理解。为什么状态转移方程可以变为：
dp[j]=Math.min(dp[j-1],dp[j])+grid[i][j]了呢
此时dp[j-1]代表的是上次计算的左边的位置。而dp[j]代表的是上次计算的上边的位置。
展开空间想象。剪枝不好弄呀.

时间复杂度：O(M*N)
空间复杂度：O(max(m,n))
*/