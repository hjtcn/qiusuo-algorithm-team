// 给你一个整数数组 arr，请你判断数组中是否存在连续三个元素都是奇数的情况：如果存在，请返回 true ；否则，返回 false 。

//  

// 示例 1：

// 输入：arr = [2,6,4,1]
// 输出：false
// 解释：不存在连续三个元素都是奇数的情况。
// 示例 2：

// 输入：arr = [1,2,34,3,4,5,7,23,12]
// 输出：true
// 解释：存在连续三个元素都是奇数的情况，即 [5,7,23] 。
//  

// 提示：

// 1 <= arr.length <= 1000
// 1 <= arr[i] <= 1000

/**
 * @param {number[]} arr
 * @return {boolean}
 */
var threeConsecutiveOdds = function (arr) {
    let k = 0//记录连续出现奇数的次数
    for (let i = 0; i < arr.length; i++) {
        if (arr[i] % 2) {//若是奇数，就增加k的值
            k++
        }
        else {//若非奇数，将k置为0
            k = 0
        }
        if (k == 3) {//出现连续三个返回true
            return true
        }
    }
    return false
};

/** 题解
存在连续三个奇数的数组。一层遍历，利用一个变量k记录是否存在连续的三个奇数，一旦存在不为奇数的值，k就置零，一旦k等于3，就返回true,没达到就返回false

 复杂度分析：
    平均时间复杂度是:O(N)
    一层for循环变量
    空间复杂度：O(1)
    定义k值记录连续出现奇数的次数
 */

//利用js的api写的代码
/**
 * @param {number[]} arr
 * @return {boolean}
 */
var threeConsecutiveOdds = function (arr) {
    return undefined !== arr.find((item, i) => (i <= arr.length - 2) && arr[i] % 2 && arr[i + 1] % 2 && arr[i + 2] % 2)
};

/** 题解
存在连续三个奇数的数组。利用arr遍历查询，是否存在arr[i],arr[i+1],arr[i+2]都为奇数，此时要记得判断i不能超过数组长度减去2防止数组溢出

 复杂度分析：
    平均时间复杂度是:O(N)
    一层for循环变量
    空间复杂度：O(1)
    未定义什么变量
 */