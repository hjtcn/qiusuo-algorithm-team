// 110. 平衡二叉树
// 给定一个二叉树，判断它是否是高度平衡的二叉树。

// 本题中，一棵高度平衡二叉树定义为：

// 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。

// 输入：root = [3,9,20,null,null,15,7]
// 输出：true

// 输入：root = [1,2,2,3,3,null,null,4,4]
// 输出：false

/* 
思路:
递归，获取高度并返回。
1.如果当前节点为空，返回1
2.递归调用左右子节点，获取左右子树的高度，如果存在子树高度差大于1的情况，返回0
3.若都没结束，就从下至上的返回左右子树深度的最大值+1。
*/
var isBalanced = function (root) {
    let getHeight=root=>{
        if(!root) return 1
        let left=getHeight(root.left),right=getHeight(root.right)
        if(!left||!right||Math.abs(left-right)>1) return 0
        return Math.max(left,right)+1
    }
    return getHeight(root)
}




/* 
    思路
    队列广搜。
    1.广搜一层一层存节点。借用一个数组，利用unshift尝试从底到高一层一层存储节点
    2.从底到高更新val为深度。
     1）深度从0开始，如果当前节点没有左右子节点，则深度为0
     2）如果左右子树的深度差大于1，则返回fals
    3.遍历完毕还没有返回值，要返回true
*/
var isBalanced=function(root){
    //这个地方一定要记得，二叉树为空，返回true
    if(!root) return true 
    let queue=[root],nodes=[]
    while(queue.length){
        let curNode=queue.shift()
        nodes.unshift(curNode)
        if(curNode.left){
            queue.push(curNode.left)
        }
        if(curNode.right){
            queue.push(curNode.right)
        }
    }
    for(node of nodes){
        let left=node.left?node.left.val:0
        let right=node.right?node.right.val:0
        if(Math.abs(left-right)>1){
            return false
        }
        node.val=Math.max(left,right)+1
    }
    return true
}

/*
    两种方式的时间复杂度和空间复杂度一致
    时间复杂度：O(N)
    取决于节点个数

    空间复杂度：O(N)
    取决于节点个数
*/

/** 
 1.提前去构建真的好难扭转思路啊，上去就想试探着写代码，尤其还是写原题，上去就会想着我上次是什么思路来着。
 2.昨天做了一遍，发现思路还是不够清晰，这个题遇到至少三次了，还是这样，所以昨天就没去交，只是把题做了，今天又把题做了两遍。
 如果有时间，五道题我还是反复多做几遍吧，死记硬背也要把每一个加深一下印象
*/