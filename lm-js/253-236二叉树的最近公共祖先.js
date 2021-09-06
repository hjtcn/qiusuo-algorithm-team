/*
    236. 二叉树的最近公共祖先
    给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

    百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

    

    示例 1：


    输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
    输出：3
    解释：节点 5 和节点 1 的最近公共祖先是节点 3 。
    示例 2：


    输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
    输出：5
    解释：节点 5 和节点 4 的最近公共祖先是节点 5 。因为根据定义最近公共祖先节点可以为节点本身。
    示例 3：

    输入：root = [1,2], p = 1, q = 2
    输出：1
 

    提示：

    树中节点数目在范围 [2, 105] 内。
    -109 <= Node.val <= 109
    所有 Node.val 互不相同 。
    p != q
    p 和 q 均存在于给定的二叉树中。
*/

/*
    思路：一直缺少了递推的思路，从底向上，竟然妄想从上往下记录住祖先节点。
        这就牵扯到了后序遍历:左右根
        let dfs=(root)=>{
            if(!root) return;
            dfs(root.left)
            dfs(root.right)
            arr.push(root.val)
        }
    题解：1.遍历顺序：递推，自底向上，后序遍历
         2.递归的含义： 返回当前节点curNode的左右子节点是否含有p或者q。
         三种情况
         (1)如果左右子节点都含有p或者q,也就是一边一个，那么显然当前curNode就是我们要找的最近公共祖先
         (2)如果左边含有p或者q，右边没有，那就返回curNode的左节点
         (3)如果右边含有p或者q,左边没有，那就返回curNode的右节点。
        
*/

var lowestCommonAncestor = function(root, p, q) {
    if(!root||root==p||root==q){
        return root
    }
    let left=lowestCommonAncestor(root.left,p,q)
    let right=lowestCommonAncestor(root.right,p,q)
    return (left&&right)?root:(left||right)
}





/*
    思考：看了代码随想录的题解，发现之前第一遍做的时候，根本没有理解，就是从递归的角度考虑的。
        再次写到这个题，看了之前自己写的题解，连递归函数的返回值有什么意义都想不通。
    1.了解后序遍历
    2.尝试想了解动态规划dp的含义一样，去思考递归的意义和遍历顺序。
*/