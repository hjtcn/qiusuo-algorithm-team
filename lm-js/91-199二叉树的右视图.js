// 199. 二叉树的右视图
// 给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

// 示例:

// 输入: [1,2,3,null,5,null,4]
// 输出: [1, 3, 4]
// 解释:

//    1            <---
//  /   \
// 2     3         <---
//  \     \
//   5     4       <---

//这道题一开始用递归写了一会，没写出来,虽然我也想到了记录深度，还是有点可惜啊。。
//然后我就尝试用BFS去解决了，这道题广搜视觉上还是比较直观的

var rightSideView=function(root){
    if(!root){
        return []
    }
    let stack=[root],res=[root.val]
    while(stack.length){
        let len=stack.length
        for(let i=0;i<len;i++){
            let curNode=stack.shift()
            if(curNode.left){
                stack.push(curNode.left)
            }
            if(curNode.right){
                stack.push(curNode.right)
            }  
            //横向for循环中最后一个值。且栈不为空
            if(i==len-1&&stack[stack.length-1]){
                // console.log(stack[stack.length-1])
                res.push(stack[stack.length-1].val)
            }       
        }      
    }
    return res
}

//感觉隔一段时间没写二叉树，广搜写的也不是那么顺手了，还有root节点判断为空的条件也会忘了
/**
 * BFS，勤写养成习惯构建模版
 * 时间复杂度：O(N)
    取决于树的节点个数，每个节点只会被入栈和出栈一次。
    空间复杂度：O(N)
    取决于树的深度
 */


var rightSideView = function(root) {
    if(!root){
        return []
    }
    let res=[],deep=0
    let helper=(root,deep)=>{
        if(!root){
            return false
        }
        if(deep==res.length){
            res.push(root.val)
        }
            helper(root.right,deep+1)
            helper(root.left,deep+1)
        
    }
    helper(root,deep)
    return res
};

//递归，没写出来，记录深度的参数设置的不对，我把deep++放到helper函数处理过程中了
//主要还是deep来控制，有右节点，就push右节点的值，push完res的长度增加了，再走左节点也不会进入push处理的判断条件中。
// 时间复杂度：O(N)
// 取决于树的节点个数，每个节点只会被入栈和出栈一次。
// 空间复杂度：O(N)
// 除去结果数组，取决于树的节点个数