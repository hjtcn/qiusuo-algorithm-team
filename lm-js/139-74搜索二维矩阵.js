// 74. 搜索二维矩阵
// 编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

// 每行中的整数从左到右按升序排列。
// 每行的第一个整数大于前一行的最后一个整数。
 

// 示例 1：


// 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// 输出：true
// 示例 2：


// 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
// 输出：false
 

// 提示：

// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 100
// -104 <= matrix[i][j], target <= 104

/* 
    思路：刚开始胡思乱想，想着控制x--，y--,保证不超出边界。
         写着写着，想起来，我没有使用二分呀，这还是双层遍历的概念


         后来又读了读题，灵机一动，我想到之前用过一次flat(),降维的api，嘿，降维结束，我的数组不就变为一维的递增数组了。
         然后走二分框架，咔咔滴。


         后来又看了看题解，题解中核心为x=(mid/rowLen|0),y=(mid%rowLen)
         这样找出对比target的数组。还是很巧的，比我想到的有技术含量，哈哈哈


         题解中也有这样的想法，两层数组，直接作两次二分。
         1.纵向0---->martrix.length，找到目标所在行
         2.然后0---->martrix.length,找到目标所在元素
         自己用这种方式也写了一遍，发现寻找目标行的细节控制还是非常重要的
         1.while(rowTop<rowBottom)  能不能等于
         2.拿每一行末尾位置去对比的话，当matrix[mid][colLen]>target， rowBottom=mid而不是 rowBottom=mid-1.
         3.返回rowTop作为目标行，而不是mid.
         又是努力扒细节的一天
*/


// @lc code=start
/**
 * @param {number[][]} matrix
 * @param {number} target
 * @return {boolean}
 */
 var searchMatrix = function(matrix, target) {
     //数组降维
    matrix=matrix.flat()
    let l=0,r=matrix.length-1
    while(l<=r){
        let mid=parseInt((l+r)/2)
        if(matrix[mid]<target){
           l=mid+1
        }
        else if(matrix[mid]>target){
            r=mid-1
        }
        else{
            return true
        }
    }
    return false
};

/*将中间值换成对应的x,y轴 */
var searchMatrix = function(matrix, target) {
    let colLen=matrix.length,rowLen=matrix[0].length
    let l=0,r=(colLen*rowLen)-1
    while(l<=r){
        let mid=parseInt((l+r)/2)
        //mid对应的x,y轴
        let x=(mid/rowLen|0),y=(mid%rowLen)
        if(matrix[x][y]<target){
           l=mid+1
        }
        else if(matrix[x][y]>target){
            r=mid-1
        }
        else{
            return true
        }
    }
    return false
};



/*先找行数，再找某一行的目标值 */
var searchMatrix = function(matrix, target) {
    let rowLen=matrix.length-1,colLen=matrix[0].length-1
    //查找所在行
    let rowTop=0,rowBottom=rowLen
    while(rowTop<rowBottom){
        let mid=parseInt((rowTop+rowBottom)/2)
        if(matrix[mid][colLen]>target){
            rowBottom=mid
        }
        else if(matrix[mid][colLen]<target){
            rowTop=mid+1
        }
        else{
            return true
        }
    }
    //此时rowTop会是目标行
    let colLeft=0,colRight=colLen
    while(colLeft<=colRight){
        let mid=parseInt((colLeft+colRight)/2)
        if(matrix[rowTop][mid]<target){
             colLeft=mid+1
        }
        else if(matrix[rowTop][mid]>target){
             colRight=mid-1
        }
        else{
            return true
        }
    }
    return false
};



/*
    前两个方法时间复杂度为：O(log(rowLen*colLen)),
    我在想先找行再找目标值的方法：
            时间复杂度为：O(log(Max(rowLen,colLen)))

    空间复杂度：O(1)
    当然降维的赋值不确定需不需要额外的空间


*/