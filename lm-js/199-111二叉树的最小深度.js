/*
111. 二叉树的最小深度
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明：叶子节点是指没有子节点的节点。

 

示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：2
示例 2：

输入：root = [2,null,3,null,4,null,5,null,6]
输出：5
 

提示：

树中节点数的范围在 [0, 105] 内
-1000 <= Node.val <= 1000
*/

/*
    思路：就是记录高度的过程。
    深搜：dfs(root,depth)
    如果root的左右子节点都为空，则更新最小深度。
    自己考虑的优化：如果已经当前的深度已经大于minHeight(最小深度)，则提前结束递归
    广搜:一层一层的搜索。
    写着写着，广搜的模板也慢慢清晰了。
    while(queue.length){
        let curNode=queue.shift()
        if(curNode.left){
            queue.push(curNode.left)
        }
        if(curNode.right){
            queue.push(curNode.right)
        }
    }
    这道题则需要同时记录节点和深度。因此初始化queue应为[[root,1]]
    let [curNode,depth]=queue.shift()

    题解：很多不同的题解都是dfs函数的返回值直接就是最小深度。无需新增depth参数

*/


var minDepth = function(root) {
    if(!root) return 0
    let minHeight=Infinity
    let dfs=(root,depth)=>{
        if(!root.left&&!root.right){
            //更新最小深度
            minHeight=Math.min(minHeight,depth)
            return;
        }
        if(depth>=minHeight){
            //提前结束递归
            return;
        }
        if(root.left){
            dfs(root.left,depth+1)
        }
        if(root.right){
            dfs(root.right,depth+1)
        }
      
    }
    dfs(root,1)
    return minHeight
};

var minDepth = function(root) {
    if(!root) return 0
    let minHeight=Infinity,floorArr=[[root,1]]
    while(floorArr.length){
        let [curNode,depth]=floorArr.shift()
        if(!curNode.left&&!curNode.right){
            //更新过程，直接返回
            minHeight=Math.min(minHeight,depth)
            return minHeight
        }
        if(curNode.left){
            floorArr.push([curNode.left,depth+1])
        }
        if(curNode.right){
            floorArr.push([curNode.right,depth+1])
        }
    }
};
/*
    时间复杂度:O(N)
    空间复杂度:O(N)。
*/
 
//题解:直接返回最小深度：注意if判断也是有先后顺序的。当然，直接返回深度可以不再借用dfs函数了，哈哈，看个人喜好。。。。
var minDepth = function(root) {
    if(!root) return 0
    let dfs=(root)=>{
        if(!root.left&&!root.right){
            return 1;
        }
        else if(root.left&&root.right){
            return 1+Math.min(dfs(root.left),dfs(root.right))
        }
        else if(root.left){
            return 1+dfs(root.left)
        }
        else if(root.right){
            return 1+dfs(root.right)
        }
    }
    return dfs(root)
}
var minDepth = function(root) {
    if(!root) return 0
    let left=minDepth(root.left),
        right=minDepth(root.right)
    if(left&&right){
        return 1+Math.min(left,right)
    }
    else if(left){
        return 1+left
    }
    else if(right){
        return 1+right
    }
    else{
        return 1
    }
}
/*
    时间复杂度:O(N)
    空间复杂度:O(H) 树的高度H,递归栈的开销
*/

/*
    思考：二叉树做的时间还是比较长的，记忆比较深刻，嘻嘻嘻
*/