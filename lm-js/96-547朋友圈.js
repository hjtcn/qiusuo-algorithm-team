// 547. 朋友圈
// 班上有 N 名学生。其中有些人是朋友，有些则不是。他们的友谊具有是传递性。如果已知 A 是 B 的朋友，B 是 C 的朋友，那么我们可以认为 A 也是 C 的朋友。所谓的朋友圈，是指所有朋友的集合。

// 给定一个 N * N 的矩阵 M，表示班级中学生之间的朋友关系。如果M[i][j] = 1，表示已知第 i 个和 j 个学生互为朋友关系，否则为不知道。你必须输出所有学生中的已知的朋友圈总数。

 

// 示例 1：

// 输入：
// [[1,1,0],
//  [1,1,0],
//  [0,0,1]]
// 输出：2 
// 解释：已知学生 0 和学生 1 互为朋友，他们在一个朋友圈。
// 第2个学生自己在一个朋友圈。所以返回 2 。
// 示例 2：

// 输入：
// [[1,1,0],
//  [1,1,1],
//  [0,1,1]]
// 输出：1
// 解释：已知学生 0 和学生 1 互为朋友，学生 1 和学生 2 互为朋友，所以学生 0 和学生 2 也是朋友，所以他们三个在一个朋友圈，返回 1 。
 

// 提示：

// 1 <= N <= 200
// M[i][i] == 1
// M[i][j] == M[j][i]



// @lc code=start
/**
 * @param {number[][]} M
 * @return {number}
 */
var findCircleNum = function (M) {
    let sum = 0, visited = []
    let dfs = (x) => {
        for (let j = 0; j < M[0].length; j++) {
            if (M[x][j] && !visited[j]) {
                visited[j] = true
                dfs(j)
            }
        }
    }
    for (let i = 0; i < M.length; i++) {
        if (!visited[i]) {
            sum++;
            dfs(i)
        }
    }
    return sum
};
// @lc code=end

/**
  DFS求连通分量。
  这个代码看着是很简单的，但看题解的时候，很多图的名词确实把我整懵了。
  连通图：若V(G)中任意两个不同的顶点V(i)和V(j)都连通，即有路径或者有边相连
  非连通的无向图：有n个子连通图
  任何连通图的连通分量只有一个，即其自身
  非连通的无向图有多个连通分量
  回到题目本身：用visited数组记录这个同学x被连通过没有，递归的条件则是遍历所有同学，判断x和他们是否是朋友且没被访问过。
  时间复杂度:O(N*M)
  每个位置只遍历一次
  空间复杂度:O(N)
  visited数组记录连通情况了
 */

var findCircleNum = function (M) {
    let visited = [], sum = 0
    let bfs = (x) => {
        let queue = [x]
        while (queue.length) {
            let curX = queue.shift()
            for (let j = 0; j < M[0].length; j++) {
                if (M[curX][j] && !visited[j]) {
                    visited[j] = true
                    queue.push(j)
                }
            }
        }
    }
    for (let i = 0; i < M.length; i++) {
        if (!visited[i]) {
            sum++;
            bfs(i)
        }
    }
    return sum

}

/**
    BFS。和上面的思路基本一致，只是进行是否连通过的遍历赋值选择了广搜的方法
    使用队列模拟，模版基本和之前的BFS一致
 */


 
 /**
  并查集三步走：
  1.创建并查集。自身就是一个并查集，一般就是for循环遍历
  2.find方法：查找根元素。优化方法：路径压缩：处于同一路径上的节点都直接连接到根节点上
  3.union.集合合并。优化方法：按秩合并：总是将更小的树连接到更大的树上，即改变小根节点到大根节点上
  */

 var findCircleNum = function (M) {
     let n=M.length,count=n
     let parent=[]
      //给每一个元素建立秩                
      let rank=new Array(n)
     //创建并查集
     for(let i=0;i<n;i++){
        parent[i]=i
        rank[i]=1
     }
     //查找根
     let find=(x)=>{
         //向上查找，如果没有查到，则自身及为根
         while(x!=parent[x]){
             //通过直接将当前节点的父节点直接指向爷爷节点
             parent[x]=parent[parent[x]]
             x=parent[x]
         }
         return x
     }
     //集合合并
     let union=(p,q)=>{
        let rootP=find(p),rootQ=find(q)
        //相交不合并
        if(rootP==rootQ){
            return 
        }
        //按秩合并
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

     for(let i=0;i<n;i++){
         for(let j=0;j<n;j++){
             if(M[i][j]){
                 union(i,j)
             }
         }
     }
     return count
 }

 /**
  * 神奇的是，我各种优化方法都用上了，耗时反而久了，难道是那会网不好，哈哈哈哈哈。
  * 学习了好久啊，多联系，熟悉整个模版流程
    时间复杂度:O(N*M)
    两层for循环
    空间复杂度:O(N)
    记录父子关系的parent数组，建立秩的rank数组。
  */