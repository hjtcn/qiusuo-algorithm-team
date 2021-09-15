/*
222. 完全二叉树的节点个数
给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。

完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。

 

示例 1：


输入：root = [1,2,3,4,5,6]
输出：6
示例 2：

输入：root = []
输出：0
示例 3：

输入：root = [1]
输出：1
 

提示：

树中节点的数目范围是[0, 5 * 104]
0 <= Node.val <= 5 * 104
题目数据保证输入的树是 完全二叉树
 

进阶：遍历树来统计节点是一种时间复杂度为 O(n) 的简单解决方案。你可以设计一个更快的算法吗？
*/

/*
    思路：遍历。时间复杂度O(N)

    题解：每次只查两端，即目标节点的左叶子结点和右叶子结点。如果都存在，则pow(2,depth)-1及为目标个数。
    如果不是都存在，就去递归左右子节点。
    目标个数为1+dfs(left)+dfs(right)

    比O(N)更快，我想到了O(logN),但是没想出来O(logN)*O(logN)，有时候没有思路的时候可以猜想时间复杂度去构建代码结构。

*/

var countNodes = function(root) {
    if(!root) return 0
    return countNodes(root.left)+1+countNodes(root.right)
}

var countNodes = function(root) {
    if(!root) return 0
    let queue=[root] 
    let count=0
    while(queue.length){
        let len=queue.length
        for(let i=0;i<len;i++){
            let curNode=queue.shift()
            if(curNode.left){
                queue.push(curNode.left)
            }
            if(curNode.right){
                queue.push(curNode.right)
            }
            count++;
        }
    }
    return count
}
/*
    时间复杂度：O(N)
    空间复杂度：O(n)
*/

var countNodes = function(root) {
    if(!root) return 0
    let dfs=(root)=>{
        let left=right=root
        let lDepth=0,rDepth=0
        while(left){
            left=left.left
            lDepth++;
        }
        while(right){
            right=right.right
            rDepth++;
        }
        if(lDepth==rDepth){
            return Math.pow(2,lDepth)-1
        }
        return 1+dfs(root.left)+dfs(root.right)
    }
    return dfs(root)
}
/*
    时间复杂度：O(logN*logN)
    空间复杂度：O(n)
*/

/*
    这个题写完没两周，又没思路了。
    大胆假设规律，摸索可能的最优解
*/