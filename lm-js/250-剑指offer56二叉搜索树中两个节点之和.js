/*
剑指 Offer II 056. 二叉搜索树中两个节点之和
给定一个二叉搜索树的 根节点 root 和一个整数 k , 请判断该二叉搜索树中是否存在两个节点它们的值之和等于 k 。假设二叉搜索树中节点的值均唯一。

 

示例 1：

输入: root = [8,6,10,5,7,9,11], k = 12
输出: true
解释: 节点 5 和节点 7 之和等于 12
示例 2：

输入: root = [8,6,10,5,7,9,11], k = 22
输出: false
解释: 不存在两个节点值之和为 22 的节点
 

提示：

二叉树的节点个数的范围是  [1, 104].
-104 <= Node.val <= 104
root 为二叉搜索树
-105 <= k <= 105
*/

/*
    思路：递归map标记节点值。arr存储节点值。map判断差值是否存在
    题解：看过题解，才知道自己陷入了两个个思维误区:
    1.如果一开始就判断差值是否出现过，那么有的还没放入map呀？
    2.二叉搜索树，中序遍历可以单调递增的,可以试试二分呀,又怕left++的过程中，错过掉某一个数据。
   
    后来想通了，就算第一遍curVal的差值没存到map中，那遍历到这个差值的时候curVal已经存进去了呀，不会出现自己想的问题。
    递归迭代都可以在遍历过程中，就去判断差值是否存在

    

*/

var findTarget = function(root, k) {

    let arr=[],map=new Map()
    let dfs=(root)=>{
        if(!root) return ;
        dfs(root.left)
        map.set(root.val,true)
        arr.push(root.val)
        dfs(root.right)
        return root
    }
    dfs(root)
    
    for(let i=0;i<arr.length;i++){
        if(arr[i]*2==k){
            return false
        }
        if(map.has(k-arr[i])){
            return true
        }
    }
    return false
};


//深搜
var findTarget = function(root, k) {
    let map=new Map()
    let dfs=(root)=>{
        if(!root) return false;
        if(map.has(k-root.val)){
            console.log(3)
            return true
        }
        map.set(root.val,true)    
        return dfs(root.left)||dfs(root.right)
    }
    return dfs(root)
};

//迭代
var findTarget = function(root, k) {
    let map=new Map()
    let queue=[root]
    while(queue.length){
        let len=queue.length
        for(let i=0;i<len;i++){
            let curNode=queue.shift()
            if(map.has(k-curNode.val)){
                return true
            }
            map.set(curNode.val,true)
            if(curNode.left){
                queue.push(curNode.left)
            }
            if(curNode.right){
                queue.push(curNode.right)
            }
        }
    }
    return false
};

var findTarget = function(root, k) {

    let arr=[],map=new Map()
    let dfs=(root)=>{
        if(!root) return ;
        dfs(root.left)
        map.set(root.val,true)
        arr.push(root.val)
        dfs(root.right)
        return root
    }
    dfs(root)
    
    let l=0;r=arr.length-1
    while(l<r){
        if(arr[l]+arr[r]<k){
           l++
        }
        else if(arr[l]+arr[r]>k){
            r--;
        }
        else{
            return true
        }
    }
    return false
};

/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/
/*
    思考：是我想太多了。
*/