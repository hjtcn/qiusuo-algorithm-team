/*
200. 岛屿数量
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

 

示例 1：

输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1
示例 2：

输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
输出：3
 

提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] 的值为 '0' 或 '1'
*/

/*
    思路：递归
         重点：只能由水平方向和/或竖直方向上相邻的陆地连接形成
         也就是扫描的过程会是四个方向
         let dirtion=[[1,0],[-1,0],[0,-1],[0,1]]
         然后是递归的模版
         let dfs=(x,y)=>{
             //边界
             if(x<0||y<0||x>=xLen||y>=yLen||choose[x][y]){
                 return;
             }
             choose[x][y]=1
             for(let i=0;i<dirtion.length;i++){
                 dfs(x+dirtion[i][0],y+dirtion[i][1])
             }
         }
         同时遍历查找每一个位置:1.走过的不走，为0的不走
         for（let i=0;i<columnLen;i++){
            for(let j=0;j<rowLen;j++){
                if(grid[i][j]&&!choose[i][j]){
                    dfs(i,j)
                }
            }
         }

         遇到的坑点：行列长度把我弄迷糊了：
         数据的初始化，遍历，边界都利用这两个值
         行i~~~~~~rowLen
         列j~~~~~columnLen
         优化点：本题中是否走过可以用grid[i][j]==0判断
         因此，可以不用choose数组，直接使用输入数组，节省空间
    题解：并查集
         1.创建并查集。自身是并查集，遍历赋值即可，但因为本题中主要是位置来确定的。两层i,j可转换为一维i*n+j
         for(let i=0;i<m;i++){
             for(let j=0;j<n;j++){
                 parent[i*n+j]=i*n+j
                 count++;
             }
         }
         2.find查找根元素
         let find=(x)=>{
            while(x!=parent[x]){
                x=parent[x]
            }
            return x
         }
         3.合并集合:可按秩合并
         let union=(p,q)=>{
             let rootP=find(p),rootQ=find(q)
             if(rootP==rootQ){
                 return;
             }
             parent[rootP]=rootQ
             count--
         }
         
*/

var numIslands = function(grid) {
    let dirtion=[[1,0],[-1,0],[0,1],[0,-1]]
    let rowLen=grid.length,columnLen=grid[0].length,count=0
    // let choose=Array.from(Array(rowLen),()=>Array(columnLen).fill(0))
    let dfs=(x,y)=>{
        if(x<0||y<0||x>=rowLen||y>=columnLen){
            return;
        }
        if(!Number(grid[x][y])){
            return;
        }
        // choose[x][y]=1
        grid[x][y]="0"
        for(let i=0;i<dirtion.length;i++){
            if(grid[x][y])
            {
                dfs(x+dirtion[i][0],y+dirtion[i][1])
            }
        }
    }
    for(let i=0;i<rowLen;i++){
        for(let j=0;j<columnLen;j++){
            // if(!choose[i][j]&&grid[i][j]=='1')
            if(grid[i][j]=='1')
            {
                count++;
                dfs(i,j)
            }
        }
    }
    return count
};
/*
    时间复杂度：O(m*n)
    空间复杂度：O(m*n) 看输入数组算不算啦，哈哈哈
*/



var numIslands = function(grid) {
    let rowLen=grid.length,columnLen=grid[0].length,count=0,parent=[],rank=[]
    for(let i=0;i<rowLen;i++){
        for(let j=0;j<columnLen;j++){
            if(grid[i][j] == 1){
            parent[i*columnLen+j]=i*columnLen+j
            count++;
            rank[i*columnLen+j]=1
            }
        }
    }
    let find=(x)=>{
        while(x!=parent[x]){
            parent[x] = parent[parent[x]];
            x = parent[x];
        }
        return x
     }
     let union=(p,q)=>{
        let rootP=find(p),rootQ=find(q)
        if(rootP==rootQ){
            return;
        }
        //这里还可以按秩合并：小树合到大树上
        if(rank[rootP]>rank[rootQ]){
            parent[rootQ]=rootP
        }
        else if(rank[rootP]<rank[rootQ]){
            parent[rootP]=rootQ
        }
        else{
            parent[rootQ]=rootP
            rank[rootQ]++
        }
        count--
    }
    for(let i=0;i<rowLen;i++){
        for(let j=0;j<columnLen;j++){
            if(grid[i][j] == 1){
                i-1>=0 && grid[i-1][j] == 1 && union(i*columnLen + j,(i-1)*columnLen + j);
                j-1>=0 && grid[i][j-1] == 1 && union(i*columnLen + j,i*columnLen + j-1);
                i+1<rowLen && grid[i+1][j] == 1 && union(i*columnLen + j,(i+1)*columnLen + j);
                j+1<columnLen && grid[i][j+1] == 1 && union(i*columnLen + j,i*columnLen + j+1);
            }
        }
    }
    return count
}

/*
    时间复杂度：O(mn*mn)
    空间复杂度：O(m*n)
*/

/*
    思考：第二次接触并查集少了畏惧感，这个结构特别清晰，我们如果多做两遍，模版构建应该会比较快
*/