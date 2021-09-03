/*
449. 序列化和反序列化二叉搜索树
序列化是将数据结构或对象转换为一系列位的过程，以便它可以存储在文件或内存缓冲区中，或通过网络连接链路传输，以便稍后在同一个或另一个计算机环境中重建。

设计一个算法来序列化和反序列化 二叉搜索树 。 对序列化/反序列化算法的工作方式没有限制。 您只需确保二叉搜索树可以序列化为字符串，并且可以将该字符串反序列化为最初的二叉搜索树。

编码的字符串应尽可能紧凑。

 

示例 1：

输入：root = [2,1,3]
输出：[2,1,3]
示例 2：

输入：root = []
输出：[]
 

提示：

树中节点数范围是 [0, 104]
0 <= Node.val <= 104
题目数据 保证 输入的树是一棵二叉搜索树。
 

注意：不要使用类成员/全局/静态变量来存储状态。 你的序列化和反序列化算法应该是无状态的。


*/

/*
    思路：写这道题真的经历满多的，拨丝抽茧才知道它到底想让干什么，先说说心理路程吧。
        1.二叉搜索树，我上去先搞了个中序遍历返回以逗号拼接的递增的字符串，然后将中序遍历转化为二叉树。
            结果：发现二叉树和输入的二叉树格式样式变了
        2.然后想着那不能动原二叉树的结构，我最后以迭代的形式一层一层将二叉树处理然后返回。
            结果：但是一层一层就无法确认根节点和左右子节点的关系。
        3.又回到了1接着拓展。中序遍历后的格式转化了，是因为我上去就自动二分了，你怎么确认你找到你直接二分的就是每个根节点呢？
        说明还需要前序遍历或后序遍历来确定根节点(这个过程看题解了，也查询了一些前序遍历和中序遍历生成二叉树的一些资料)

        let dfs=(preArr,innerArr)=>{
            if(!preArr.length||!innerArr.length) return null
            let node=preArr.shift()
            let index=innerArr.indexOf(node)
            let root=new NodeTree(node)
            root.left=dfs(preArr.slice(0,index),innerArr.slice(0,index))
            root.right=dfs(preArr.slice(index),innerArr.slice(index+1))
            return root
        }

        一个小疑问。root.right=dfs(preArr.slice(index),innerArr.slice(index+1))
        为什么前序数组要切割从preArr.slice(index)开始呢？
        怀疑和preArr.shift()有关系，索引变了。

*/

/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */

/**
 * Encodes a tree to a single string.
 *
 * @param {TreeNode} root
 * @return {string}
 */
 var serialize = function(root) {
    if(!root) return null
    let arr=[]
     let dfs=(root)=>{
        if(!root) return;
        arr.push(root.val)
        dfs(root.left)
        dfs(root.right)
    }
    dfs(root)
    return arr.toString()
};

/**
 * Decodes your encoded data to tree.
 *
 * @param {string} data
 * @return {TreeNode}
 */
var deserialize = function(data) {
    if(!data) return null
    let arr=data.split(',')
    let len=arr.length
    if(!len) return null;
    if(len==1) return new TreeNode(arr[0])
    let copyArr=[...arr].sort((a,b)=>a-b)
    let dfs=(copyArr,arr)=>{
            if(!copyArr.length||!arr.length) return null;
            let node=arr.shift()
            let mid=copyArr.indexOf(node)
            var root = new TreeNode(node);
            root.left=dfs(copyArr.slice(0,mid),arr.slice(0,mid))
            root.right=dfs(copyArr.slice(mid+1),arr.slice(mid))
            return root
    }
    return dfs(copyArr,arr)
};

/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

/*
    思考：越学越深。三种遍历构建二叉树的过程，我们也需要多深入练习

*/