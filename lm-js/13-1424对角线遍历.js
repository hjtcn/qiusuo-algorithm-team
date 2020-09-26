// 给你一个列表 nums ，里面每一个元素都是一个整数列表。请你依照下面各图的规则，按顺序返回 nums 中对角线上的整数。


//下面这个是自己想要暴力查询位置，但是会超时的代码，不舍得删除，记下来吧。。。。。。
/**
//  * @param {number[][]} nums
//  * @return {number[]}
//  */
// var findDiagonalOrder = function (nums) {
//     let row = nums.length, res = [],maxLen=0;
//     if(row==1){
//         return nums[0]
//     }
//     if(maxLen==1){
//         for(let i=0;i<row;i++){
//           if (nums[i][0]) {
//                 res.push(nums[i][0])
//             }
//         }
//         return res
//     }
//     for (let i = 0; i < row; i++) {
//         let k = i, w = 0,tag=1;
//         maxLen=Math.max(nums[i].length,maxLen)
//         while (tag) {
//             if (nums[k][w]) {
//                 res.push(nums[k][w])
//             }
//                 w++;
//                 k--;
//             if(k<0){
//                 tag=0
//             }
//         }
//     }
//     for (let i = 1; i <maxLen; i++) {
//         let k = i, w =row-1,tag=1;
//         while (tag) {
//             if (nums[w][k]) {
//                 res.push(nums[w][k])
//             }
//                 k++;
//                 w--;
//             if(k>maxLen||w<0){
//                 tag=0
//             }
//         }
//     }
//     return res
// };



//下面这个是看题解get到的解法，觉得js可真是厉害坏了。。。。。。。
/**
 * @param {number[][]} nums
 * @return {number[]}
 */
var findDiagonalOrder = function(nums) {
    let res=[];
    for (let i=0;i<nums.length;i++){
        for (let j=0;j<nums[i].length;j++){
            // 把和相同的放在一个数组中，从前面unshift到res[i+j]的数组，可以避免数组的一次反转
            res[i+j] ? res[i+j].unshift(nums[i][j]) : res[i+j]=[nums[i][j]];
        }
    }
    console.log(res)
    // 数组扁平化，跳过空元素，拉伸至一维数组。。。这个方法原来都不知道，这个真的是厉害了。。。。
    return res.flat(Infinity)
};

/**题解
   可以发现对角线元素nums[i][j]的索引之和（i+j）都相等，双层遍历后将和相同的元素放入一个数组。
   这样经过遍历赋值后的res数组就变为期望数组元素顺序排列的多维数组。
   然后利用flat方法进行无穷极降维到一维数组，并且返回
 
  复杂度分析：
    时间复杂度是:O(N^2)
    两层for循环遍历记录和相同的元素。
    空间复杂度：O(N*M）
    借用res数组存放多位数组的所有元素。
 */