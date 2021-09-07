/*
501. 二叉搜索树中的众数
给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。

假定 BST 有如下定义：

结点左子树中所含结点的值小于等于当前结点的值
结点右子树中所含结点的值大于等于当前结点的值
左子树和右子树都是二叉搜索树
例如：
给定 BST [1,null,2,2],

   1
    \
     2
    /
   2
返回[2].

提示：如果众数超过1个，不需考虑输出顺序

进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内） 

*/


/*
    思路：注意读题。踩了两个坑
        1.找出BST中的所有众数
        2.刚开始不知道题目要的返回结果到底是什么。是节点还是节点值。后来发现是[数组存放节点值]，哈哈哈，在提示代码中有说。
        找众数。二叉搜索树，中序遍历。计算持续出现的某个值的次数count，更新最大出现次数max.
        如果count>max,max=count,maxArr=[root.val]
        如果count==max,maxArr.push(root.val)
        遍历结束，返回maxArr
    题解：如果是非二叉搜索树呢？遍历存储排序

*/

/**
 * @param {TreeNode} root
 * @return {number[]}
 */
 var findMode = function(root) {
    let count=0,pre=null,max=-Infinity
    let dfs=(root)=>{
        if(!root) return ;
        dfs(root.left)
        if(root.val==pre){
            count++;
        }
        else{
            count=0
        }
        if(count>max){
            max=count
            maxArr=[root.val]
        }
        else if(count==max){
            maxArr.push(root.val)
        }
        pre=root.val
        dfs(root.right)
    }
    dfs(root)
    return maxArr
};

var findMode = function(root) {
    let count=0,pre=null,max=-Infinity
    let queue=[],curNode=root,maxTarget=[]
    while(curNode||queue.length>0){
        while(curNode){
            queue.push(curNode)
            curNode=curNode.left
        }
        curNode=queue.pop()
        if(pre==curNode.val){
            count++;
        }
        else{
            count=0
        }
        if(count>max){
            max=count
            maxTarget=[curNode.val]
        }
        else if(count==max){
            maxTarget.push(curNode.val)
        }
        pre=curNode.val
        curNode=curNode.right
    }
    return maxTarget
};
/*
    时间复杂度：O(N)
    空间复杂度：O(1)
*/

//普通二叉树
var findMode = function(root) {
    let nodeMap=new Map(),target=[]
    let dfs=(root)=>{
        if(!root) return;
        if(nodeMap.has(root.val)){
            nodeMap.set(root.val,nodeMap.get(root.val)+1)
        }
        else{
            nodeMap.set(root.val,1)
        }
        dfs(root.left)
        dfs(root.right)
    }
    dfs(root)
    var nodeArr=Array.from(nodeMap)
    nodeArr.sort((a,b)=>b[1]-a[1])
    target[0]=nodeArr[0][0]
    let targetCount=nodeArr[0][1]
    for(let i=1;i<nodeArr.length;i++){
        if(targetCount==nodeArr[i][1]){
            target.push(nodeArr[i][0])
        }
        else{
            break;
        }
    }
    return target
}

/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/
