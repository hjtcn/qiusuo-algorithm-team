
// 617. 合并二叉树
// 给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。

// 你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点。

// 示例 1:

// 输入: 
// 	Tree 1                     Tree 2                  
//           1                         2                             
//          / \                       / \                            
//         3   2                     1   3                        
//        /                           \   \                      
//       5                             4   7                  
// 输出: 
// 合并后的树:
// 	     3
// 	    / \
// 	   4   5
// 	  / \   \ 
// 	 5   4   7
// 注意: 合并必须从两个树的根节点开始。


// ============================================================
// ===                                                      ===
// ===       状态：通过,执行用时: 120 ms,内存消耗: 46.5 MB     ===
// ============================================================

/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     val: number
 *     left: TreeNode | null
 *     right: TreeNode | null
 *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *     }
 * }
 */

function mergeTrees(t1: TreeNode | null, t2: TreeNode | null): TreeNode | null {

    if(t1 || t2){
        if(t1 && t2){
            t1.val = t1.val + t2.val;

            t1.left = mergeTrees(t1.left || null,  t2.left || null);
            t1.right = mergeTrees(t1.right || null, t2.right || null);

            return t1;
        }else{
            // 此处t1是null
           return t1 || t2;
        }
    }
    return null;
};