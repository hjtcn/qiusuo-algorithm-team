// 给定一个按非递减顺序排序的整数数组 A，返回每个数字的平方组成的新数组，要求也按非递减顺序排序。

//  

// 示例 1：

// 输入：[-4,-1,0,3,10]
// 输出：[0,1,9,16,100]
// 示例 2：

// 输入：[-7,-3,2,3,11]
// 输出：[4,9,9,49,121]
//  

// 提示：

// 1 <= A.length <= 10000
// -10000 <= A[i] <= 10000
// A 已按非递减顺序排序。



//自己第一次写的渣渣代码
/* @param {number[]} A
* @return {number[]}
*/
var sortedSquares = function (A) {
    let B = [], C = [], D = []
    //遍历区分出两个已经排好序的数组，存放原值大于0的值的数组和小于0的数组
    for (let i = 0; i < A.length; i++) {
        if (A[i] < 0) {
            B.unshift(A[i] * A[i])
        }
        else {
            C.push(A[i] * A[i])
        }
    }
    //两个数组都有值，找到小的值就放到新的数组中，一方数组没比较完毕就把剩余的数组全部插入到后面
    if (B.length && C.length) {
        let m = 0, n = 0
        for (let i = n; i < B.length; i++) {
            if (m == C.length) {
                break;
            }
            for (let j = m; j < C.length; j++) {
                if (B[i] <= C[j]) {
                    D.push(B[i])
                    n++
                    break;
                }
                else {
                    D.push(C[j])
                    m++;
                }
            }
        }
        if (m < C.length) {
            D = D.concat(C.slice(m))
        }
        if (n < B.length) {
            D = D.concat(B.slice(n))
        }
        return D
    }
    //如果一方有值，就返回另一方数组
    else if (!B.length) {
        return C
    }
    else if (!C.length) {
        return B
    }
};
/** 题解
有序数组的平方。遍历根据原值是否大于0，定义两个已经排好序的数组，
  1）如果两个数组都有值，将两个数组进行对比，出现最小的值就放入到D数组中，最后一方没有对比完毕，就把剩余的数组内容合并到D数组里，返回D数组
  2）一方数组没有值，则返回另一方数组

 复杂度分析：
    时间复杂度：O(n)
    虽然对比的时候是两层放循环，但是每一次一旦找到,就会跳出循环，一旦某一个数组执行完毕，会跳出整个循环
    空间复杂度：O(n)
    定义了很多数组
 */



//看完题解了解的方法。。。。。。。
/**
 * @param {number[]} A
 * @return {number[]}
 */
var sortedSquares = function (A) {
    //借助map，平方之后再排序
    return A.map(item => Math.pow(item, 2)).sort((a, b) => a - b)
};

/** 题解
有序数组的平方。借助平方之后再次排序，sort方法对于数组长度小于10，所用技术是插入排序，若数组长度大于10，所用技术为快排
复杂度分析：
   if(插入排序)
   最坏时间复杂度：O(n^2)：当数组是从大到小排列时，插入第一个元素需要移动一次，插入第二个需要移动两次，以此类推，所以一共为1+2+3+4+......+（n-1）,与冒泡排序相同
   最优时间复杂度：最好的情况是数组已经由小到大排好，则为O(n)
   稳定性：稳定
   if(快排)
   最坏时间复杂度：O（n^2） 当选择的基准值为最大值或最小值时
   平均时间复杂度：O(n*log2n)
   稳定性：不稳定


   空间复杂度：O(1)
   未定义多余变量
*/


/**
* @param {number[]} A
* @return {number[]}
*/
var sortedSquares = function (A) {
    let left = 0;
    let right = A.length - 1;
    let res = []
    let index = right
    while (left <= right) {
        //获取左右指针指向的平方值
        let leftValue = A[left] * A[left]
        let rightValue = A[right] * A[right]
        //左指针指向的值大，则left++，res赋值leftValue,赋值的索引减一
        if (leftValue >= rightValue) {
            res[index--] = leftValue
            left++;
        }
        else { //右指针指向的值大，则right--，res赋值rightValue,赋值的索引减一
            res[index--] = rightValue
            right--
        }
    }
    return res
};

/** 题解
有序数组的平方。利用双指针，从数组的最大长度开始赋值，左右指针比较大小，left指向的值大，就将left指向的值赋予res,left加一，right指向的值大，就将right的值赋予res，right减一，直到两个指针碰面
复杂度分析：
 时间复杂度：O(n)
 双指针只循环了一遍数组
 空间复杂度：O(n)
 定义的res数组
*/