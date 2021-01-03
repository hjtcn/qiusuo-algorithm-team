// 695. 岛屿的最大面积
// 给定一个包含了一些 0 和 1 的非空二维数组 grid 。

// 一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。

// 找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 0 。)

 

// 示例 1:

// [[0,0,1,0,0,0,0,1,0,0,0,0,0],
//  [0,0,0,0,0,0,0,1,1,1,0,0,0],
//  [0,1,1,0,1,0,0,0,0,0,0,0,0],
//  [0,1,0,0,1,1,0,0,1,0,1,0,0],
//  [0,1,0,0,1,1,0,0,1,1,1,0,0],
//  [0,0,0,0,0,0,0,0,0,0,1,0,0],
//  [0,0,0,0,0,0,0,1,1,1,0,0,0],
//  [0,0,0,0,0,0,0,1,1,0,0,0,0]]
// 对于上面这个给定矩阵应返回 6。注意答案不应该是 11 ，因为岛屿只能包含水平或垂直的四个方向的 1 。

// 示例 2:

// [[0,0,0,0,0,0,0,0]]
// 对于上面这个给定的矩阵, 返回 0。


/*
 * @lc app=leetcode.cn id=695 lang=javascript
 *
 * [695] 岛屿的最大面积
 */

// @lc code=start
/**
 * @param {number[][]} grid
 * @return {number}
 */
var maxAreaOfIsland = function(grid) {
    let max=0,dfsMax=0
    let dfs=(x,y)=>{
        if(x<0||x>=grid.length||y<0||y>=grid[0].length||!grid[x][y]){
            return ;
        }
        if(grid[x][y]){
            grid[x][y]=0
            dfsMax++;
        }
        dfs(x,y+1)
        dfs(x,y-1)
        dfs(x-1,y)
        dfs(x+1,y)
    }
    for(let i=0;i<grid.length;i++){
        for(let j=0;j<grid[0].length;j++){
            if(grid[i][j]){
                dfsMax=0
                dfs(i,j)
                max=Math.max(max,dfsMax)
            }
        }
    }
    return max
};

/**
 * 递归。这个做的蛮顺利的，比周一周二的题来说要简单一些
 * 只包含水平垂直四个方向了，递归方向减少，只要记得更新最大值就行
 * 时间复杂度:O(N^3)
 * 两层for循环，且递归遍历每一个元素
 * 空间复杂度：O(1)
 * 更新最大值
 */
// @lc code=end

var maxAreaOfIsland = function(grid) {
    let max=0,direction=[[0,-1],[0,1],[1,0],[-1,0]]
    for(let i=0;i<grid.length;i++){
        for(let j=0;j<grid[0].length;j++){
            let queue=[[i,j]],bfsMax=0
            while(queue.length>0){
                let [x,y]=queue.shift()
                if(x<0||x>=grid.length||y<0||y>=grid[0].length||!grid[x][y]){
                    continue;
                }
                    grid[x][y]=0
                    bfsMax++;
                    for(let k=0;k<direction.length;k++){
                        queue.push([x+direction[k][0],y+direction[k][1]])
                    }
            }
            max=Math.max(max,bfsMax)
        }
    }
    return max
};


/**
 * BFS.这种题BFS的模版构建还是不是很清晰的，还要多战几次呀
 * 队列遍历四个方向。边界控制好。
 * 时间复杂度:O(N^3)
 * 两层for循环，且递归遍历每一个元素
 * 空间复杂度：O(N)
 * 队列存储索引值代表位置。
 */