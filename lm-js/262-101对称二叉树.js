/*
101. 对称二叉树
给定一个二叉树，检查它是否是镜像对称的。

 

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
 

但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3
 

进阶：

你可以运用递归和迭代两种方法解决这个问题吗？
*/

/*
    思路：看到这个题，我第一个想法是只能迭代广搜，每一层构建数组，判断是否回文。
    题解：找规律。left.left=right.right
                left.right=right.left
                保证所有节点都符合

        迭代思路的时候还踩坑了。递归转换为迭代还是需要注意细节的
        if(!left&&!right) continue;
        我却直接return true了。


*/
var isSymmetric = function(root) {
    let queue=[root]
    let isPalindrome=(arr)=>{
        let newArr=[...arr]
        return arr.reverse().toString()==newArr.toString()
    }
    while(queue.length>0){
        let len=queue.length
        let arr=[]
        for(let i=0;i<len;i++){
            let curNode=queue.shift()
            if(curNode)
            {
                arr.push(curNode.val)
                queue.push(curNode.left)
                queue.push(curNode.right)
            }
            else{
                arr.push(null)
            }
        }
        if(!isPalindrome(arr)){
            return false
        }
    }
    return true
}
/*
    时间复杂度：O(N*N)
    空间复杂度：O(N)
*/

var isSymmetric = function(root) {
    let dfs=(left,right)=>{
        if(!left&&!right) return true
        else if(!left||!right) return false
        return left.val==right.val&&dfs(left.left,right.right)&&dfs(left.right,right.left)
    }
    return dfs(root,root)
};

var isSymmetric = function(root) {
    let queue=[[root,root]]
    while(queue.length>0){
        let len=queue.length
        for(let i=0;i<len;i++){
            let [left,right]=queue.shift()
            if(!left&&!right) continue
            else if(!left||!right) return false
            if(left.val!=right.val) return false
            queue.push([left.left,right.right])
            queue.push([left.right,right.left])
        }
    }
    return true
};

/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

/*
    分成小模块，敢于摸索规律
*/