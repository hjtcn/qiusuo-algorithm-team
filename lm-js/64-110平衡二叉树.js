// 110. 平衡二叉树
// 给定一个二叉树，判断它是否是高度平衡的二叉树。

// 本题中，一棵高度平衡二叉树定义为：

// 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。

 

// 示例 1：


// 输入：root = [3,9,20,null,null,15,7]
// 输出：true
// 示例 2：


// 输入：root = [1,2,2,3,3,null,null,4,4]
// 输出：false
// 示例 3：

// 输入：root = []
// 输出：true
 

// 提示：

// 树中的节点数在范围 [0, 5000] 内
// -104 <= Node.val <= 104

/**
 * @param {TreeNode} root
 * @return {boolean}
 */
var isBalanced = function (root) {
    if (!root) return true
    let getHeight = (node) => {
        if (!node) return -1;
        return 1 + Math.max(getHeight(node.left), getHeight(node.right))
    }
    return Math.abs(getHeight(root.left) - getHeight(root.right)) <= 1 && isBalanced(root.left) && isBalanced(root.right)
};


//看错题了，卡了好长时间，没读清题，一直以为是最深树长度和最浅树长度的差值大于1，好气呀。。。。。



/**
    暴力递归。声明一个获取深度的方法。返回左右节点的深度差值的绝对值是否小于1，递进调用左右节点。

    时间复杂度：O(nlogn)
    getHeight存在大量冗余操作
    
    空间复杂度：O(N)
    取决于队列节点数，不会超过n
 * 
 */


var isBalanced = function (root) {
    if (!root) return true
    let dps=(node)=>{
        if(!node) return 1
        let left=dps(node.left)
        let right=dps(node.right)
        //  左右子树不平衡或差值大于1，该二叉树不平衡
        if(!left||!right||Math.abs(left-right)>1){
            return 0
        }
        else{
            return 1+Math.max(left,right)
        }
    }
    return dps(root)
};

/**
 * 看题解get的
 * 后序遍历，从底至顶返回最大高度，判断每个子树是不是平衡树。
 * 如果平衡，则使用它们的高度判断父节点是否平衡，并计算父节点的高度，如果不平衡，返回0
 * 遍历比较二叉树每个节点的左右子树深度：
 
    1.判断子树是否是平衡树的方法：左右子树有一个不是平衡的，或左右子树差值大于 1 ，则二叉树不平衡
    
    2.若左右子树平衡，返回当前树的深度（左右子树的深度最大值 +1 ）


    时间复杂度：O(N)
    取决于树的深度
    
    空间复杂度：O(N)
    取决于队列节点数，不会超过n
 */