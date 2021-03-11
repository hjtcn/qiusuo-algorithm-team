/*
你是产品经理，目前正在带领一个团队开发新的产品。不幸的是，你的产品的最新版本没有通过质量检测。由于每个版本都是基于之前的版本开发的，所以错误的版本之后的所有版本都是错的。
假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。
你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version 是否在单元测试中出错。实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。
示例:
给定 n = 5，并且 version = 4 是第一个错误的版本。
调用 isBadVersion(3) -> false
调用 isBadVersion(5) -> true
调用 isBadVersion(4) -> true
所以，4 是第一个错误的版本。 
*/
1. Clearfication:
记录一下我的反面教材哈：
1,2,3,4,5 找到第1个错的版本：
true,true,false,false,false
这个是不是很熟悉，类似我们周二找到数字边界的那道题，找到false的第一次出现的位置
将问题抽象程 1，2，3，3，3 找到第一个3出现的位置
伪代码：
left,right := 0,n
for left < right {
    mid := left + (right - left) / 2
    
    if nums[mid] == target {
        right = mid
    }else if nums[mid] < target {
        left = mid + 1
    }else if nums[mid] > target {
        right = mid
    }
}
if nums[left] == target {
    return left
}
return -1
写完伪代码以后，要验证边界逻辑，找特殊的case去进行验证
[1,2] target 3,  找比里面数字都大的数，看程序是否合理
[1,2] target 0   找比数组里面数字都小的数，看程序是否合理
[1,2,3,3,3] target 3  找数组中存在的数，看程序是否合理

2. Coding:
然后就写出了下面的代码
/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return            true if current version is bad 
 *                    false if current version is good
 * func isBadVersion(version int) bool;
 */
func firstBadVersion(n int) int {
    left,right := 0,n + 1
    for left < right {
        mid := left + (right - left) / 2
        if isBadVersion(mid) == false {
            right = mid
        }else {
            left = mid + 1
        }
    }
    if isBadVersion(left) == false {
        return left
    } 
    return -1
}

3. 看题解：
以为是自己分析错了，初始化条件改为了 left,right := 1,n + 1 还是没有过，就去看了题解：发现自己漏了暴力的解法
/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return            true if current version is bad 
 *                    false if current version is good
 * func isBadVersion(version int) bool;
 */
func firstBadVersion(n int) int {
    for i := 1;i <= n;i++ {
        if isBadVersion(i) == false {
            return i
        }
    }
    return -1
}
自己写了暴力的case，发现 run 不过去，就去仔细看了题意，发现题目是找第一个为 true 的地方，从一开始就分析错了，凉凉，然后就将上面的代码逻辑改了下判断就通过了
/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return            true if current version is bad 
 *                    false if current version is good
 * func isBadVersion(version int) bool;
 */
func firstBadVersion(n int) int {
    left,right := 1,n + 1
    for left < right {
        mid := left + (right - left) / 2
        if isBadVersion(mid) {
            right = mid
        }else {
            left = mid + 1
        }
    }
    if isBadVersion(left) {
        return left
    } 
    return -1
}
看了官方题解：
最后直接return left 的了，这道题目里面可以这样写，因为它的边界条件限制了，我试了一下，n =5 ,找 6 就报错了
case 是  [1,2,3,4,5] 我找等于5的，其实 left < right 的时候最后可能没有走到 5, 我们数组里面定义的是 1, len(nums)， 数组的最后一位下标是 len(nums) -1，所以这里如果要用 left < right 的判断的话，最好是写 left = 1,right = n + 1, 最后退出的时候对 left 做一下判断。
所以我们看别人的代码也是多要思考的，不能拿来都用，也不要有这个思维
public int firstBadVersion(int n) {
    int left = 1;
    int right = n;
    while (left < right) {
        int mid = left + (right - left) / 2;
        if (isBadVersion(mid)) {
            right = mid;
        } else {
            left = mid + 1;
        }
    }
    return left;
}

4. 复杂度分析：
时间复杂度：O(logN)
空间复杂度：O(1)

5. 总结：
5.1. 看题目，写完伪代码以后，要用题目中的case走一遍，不要以为自己真的理解了
5.2. 知道为什么比是什么更重要
5.3 做题的时候可以通过模拟面试，15分钟之内做不出来就去看题解
