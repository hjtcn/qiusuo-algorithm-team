/*
120. 三角形最小路径和
给定一个三角形 triangle ，找出自顶向下的最小路径和。

每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。

 

示例 1：

输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
输出：11
解释：如下面简图所示：
   2
  3 4
 6 5 7
4 1 8 3
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
示例 2：

输入：triangle = [[-10]]
输出：-10
 

提示：

1 <= triangle.length <= 200
triangle[0].length == 1
triangle[i].length == triangle[i - 1].length + 1
-104 <= triangle[i][j] <= 104
 

进阶：

你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题吗？

*/

/*
    思路：从上到小将原数替换为最小和
    注意数组溢出和特殊边界  

    题解：发现自底向上的特殊处理要更少一些
        并且直接返回triangle[0][0]即可
*/
var minimumTotal = function(triangle) {
    let len1=triangle.length
    let min=Infinity
    if(len1==1){
        return triangle[0][0]
    }
    for(let i=0;i<len1;i++){
        for(let j=0;j<triangle[i].length;j++){
            if(i==0){
                break;
            }
            if(j==0){
                triangle[i][j]=triangle[i-1][j]+triangle[i][j]
            }
            else if(j==triangle[i].length-1){
                triangle[i][j]=triangle[i-1][j-1]+triangle[i][j]
            }
            else{
                triangle[i][j]=Math.min(triangle[i-1][j],triangle[i-1][j-1])+triangle[i][j]
            }
            if(i==len1-1){
                min=Math.min(triangle[i][j],min)
            }
        }
    }
    return min
};


var minimumTotal = function(triangle) {
    let len1=triangle.length
    for(let i=len1-2;i>=0;i--){
        for(let j=0;j<triangle[i].length;j++){
            triangle[i][j]=Math.min(triangle[i+1][j+1],triangle[i+1][j])+triangle[i][j]
        }
    }
    return triangle[0][0]
}

/*
    时间复杂度：O(N^2)
    空间复杂度：O(1)
*/