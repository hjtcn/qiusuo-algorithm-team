/*
62. 不同路径
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？

 

示例 1：


输入：m = 3, n = 7
输出：28
示例 2：

输入：m = 3, n = 2
输出：3
解释：
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右
3. 向下 -> 向右 -> 向下
示例 3：

输入：m = 7, n = 3
输出：28
示例 4：

输入：m = 3, n = 3
输出：6
 

提示：

1 <= m, n <= 100
题目数据保证答案小于等于 2 * 109

*/


/*
思路:看到示例1的图，一瞬间脑子里想的满多的。
    想到了递归，想到了扫雷，想到了最短路径
    最后还是想着先乖乖的把动态规划模拟出来吧。
    因此，第一步，先写状态转移。
    其实这种网格的是最容易画动态转移过程的。
    依然以m=3,n=7为例
    0，1，1， 1， 1， 1，1，
    1，2，3， 4， 5， 6，7
    1，3，6，10，15，21，28
    这样，一下是不是清晰多了。dp[i][j]=dp[i-1][j]+dp[i][j-1]

    难点在哪：
    1).m=1,n=1的特殊情况要确定。经过执行测试dp[1][1]=1
    2).对于js来说比较困难的dp初始化。一开始我初始化的为0，后来发现为0不仅没有什么作用，而且在第一行和第一列的时候还要特殊赋值。因此最终将dp全部初始化为1.哒哒哒
*/

var uniquePaths = function(m, n) {
    let dp = Array.from(Array(m),()=>Array(n).fill(1))
    for(let i=1;i<m;i++){
        for(let j=1;j<n;j++){
          dp[i][j]=dp[i-1][j]+dp[i][j-1]
        }
    }
    return dp[m-1][n-1]
};

/*
写出这道题还挺开心的，然后去看了看题解
https://leetcode-cn.com/problems/unique-paths/solution/di-gui-dong-tai-gui-hua-jiang-wei-zu-he-xlwrx/
这个题解不停优化，我一看，有递归，那我不能放过呀，原来递归这想法还是可以的，再抠一会递归吧，试一试，记得更加深刻
*/

var uniquePaths = function(m, n) {
    let count=1
    let dfs=(x,y)=>{
        if(x>=m||y>=n){
            return ;
        }
        count+=1
        dfs(x+1,y)
        dfs(x,y+1)
    }
    dfs(1,1)
    return count
};
//这个代码在m=51,n=6的时候超时了。
//想想怎么缓存一下呢。
var uniquePaths = function(m, n) {
    let cache=Array.from(Array(m+1),()=>Array(n+1).fill(0))
    let dfs=(x,y)=>{
        let count=0
        if(cache[x][y]) return cache[x][y]
        if(x<=0||y<=0){
            return 0;
        }
        if(x==1||y==1) return 1 
        count+=dfs(x-1,y)
        count+=dfs(x,y-1)
        cache[x][y]=count
        return count
    }
    return dfs(m,n)
};

//嘿嘿。还是抠出来了，哒哒哒.上面题解代码可读性太差了，不想降维进一步优化了。
/*
    时间复杂度：O(N*M)
    空间复杂度：O(N*M)
*/
