/*
450. 删除二叉搜索树中的节点
给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

一般来说，删除节点可分为两个步骤：

首先找到需要删除的节点；
如果找到了，删除它。
说明： 要求算法时间复杂度为 O(h)，h 为树的高度。

示例:

root = [5,3,6,2,4,null,7]
key = 3

    5
   / \
  3   6
 / \   \
2   4   7

给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。

一个正确的答案是 [5,4,6,2,null,null,7], 如下图所示。

    5
   / \
  4   6
 /     \
2       7

另一个正确答案是 [5,2,6,null,4,null,7]。

    5
   / \
  2   6
   \   \
    4   7


*/

/*
    思路：我也尝试删除节点了，但是看过题解太感觉到自己格局小了。我只想到了要删除一个节点，而会上位的节点，上在那个位置呢，我判断不出来。

    题解：递归
        列举可能性。
        1.删除叶子结点，没有左右子树。
        2.删除的节点有左子树，没有右子树，左子树上位
        3.删除的节点有右子树，没有左子树，右子树上位
        4.删除的节点左右子树都有。将左子树整个移到右子树的左叶子(值最小)结点上。
        怎么找右子树right的左叶子结点
        while(right.left){
            right=right.left
        }
        right.left=left(左子树)
        root=root.right //右子树上位
        迭代
        代码是自己理解处理过程后自己写的。原因是不想调用题解中(代码随想录)删除一个节点的函数。然后就自己摸索着写了过程。
        1.记录根节点pre，并标记是左子树还是右子树
        2.当找到目标节点的时候，除了处理curNode节点，且拼接pre节点。
        如:if(flag==0){
            根节点就是目标节点。
            pre=curNode
            return pre
            }
            if(flag==1){
                目标节点存在与pre的左子树
                pre.left=curNode
            }
            if(flag==2){
                 目标节点存在与pre的右子树
                pre.right=curNode
            }
        我的迭代代码的时间度，不止遍历了每个节点，而且当被删除节点包含左右子节点的时候，还深层递归的寻找了右子树的左叶子结点。这个可能不太好。
*/

/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @param {number} key
 * @return {TreeNode}
 */
 /**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @param {number} key
 * @return {TreeNode}
 */

 var deleteNode = function(root, key) {
    let dfs=(root)=>{
        if(!root) return null;
        if(root.val==key){
            if(!root.left&&!root.right){
                root=null
            }
            else if(root.left&&!root.right){
                root=root.left
            }
            else if(!root.left&&root.right){
                root=root.right
            }
            else{
                let left=root.left
                let right=root.right
                while(right.left){
                    right=right.left
                }
                right.left=left
                root=root.right
            }
            return root
        }
        root.left=dfs(root.left)
        root.right=dfs(root.right)
        return root
    }
    return dfs(root)
}



var deleteNode = function(root, key) {
    if(!root) return root
    let pre=new TreeNode(null)
    let queue=[[root,pre,0]]
    while(queue.length){
        let len=queue.length
        for(let i=0;i<len;i++){
           [curNode,pre,flag]=queue.shift()
            if(curNode.val==key){
                if(!curNode.left&&!curNode.right){    
                    curNode=null
                }
                else if(curNode.left&&!curNode.right){
                    curNode=curNode.left
                }
                else if(curNode.right&&!curNode.left){
                    curNode=curNode.right
                }
                else{
                    let right=curNode.right
                    while(right.left){
                        right=right.left
                    }
                    right.left=curNode.left
                    curNode=curNode.right
                }
                if(flag==0){
                    //root节点就是目标节点
                    pre=curNode
                    return pre
                }
                if(flag==1){
                    //删除的是左子树中的节点
                    pre.left=curNode
                }
                if(flag==2){
                    //删除的是右子树的节点
                    pre.right=curNode
                }
            }
            if(curNode&&curNode.left){
                queue.push([curNode.left,curNode,1])
            }
            if(curNode&&curNode.right){
                queue.push([curNode.right,curNode,2])
            }
        }
    }
    return root
}

/*
    时间复杂度：O(N)
    空间复杂度：O(N)

*/

/*
    删除不容易呀。
    二叉搜索树就在于，删除节点的右子树上位的话，将左子树放在右子树的左叶子结点上。
    删除节点的左子树上位的话，需要将右子树放在左子树的右叶子结点上。去保证不改变二叉搜索树的特性。

*/