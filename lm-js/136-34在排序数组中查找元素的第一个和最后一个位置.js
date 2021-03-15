
// 评论 (1.5k)
// 题解 (2.1k)
// 提交记录
// 34. 在排序数组中查找元素的第一个和最后一个位置
// 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

// 如果数组中不存在目标值 target，返回 [-1, -1]。

// 进阶：

// 你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？
 

// 示例 1：

// 输入：nums = [5,7,7,8,8,10], target = 8
// 输出：[3,4]
// 示例 2：

// 输入：nums = [5,7,7,8,8,10], target = 6
// 输出：[-1,-1]
// 示例 3：

// 输入：nums = [], target = 0
// 输出：[-1,-1]
 

// 提示：

// 0 <= nums.length <= 105
// -109 <= nums[i] <= 109
// nums 是一个非递减数组
// -109 <= target <= 109


/*
    思路：我首先还是先把架子搭好
    while(l<=r){
        let mid=parseInt((l+r)/2)
        if(nums[mid]<target){
           l++
        }
        else if(nums[mid]>target){
            r--
        }
        else if(nums[mid]==target){
            //这个时候不知道咋处理，因为mid不是我们要找的呀。
        }
    }

    后来我想了想，我们要找的其实是l和r呀，那就把l,r加一层提前判断。
    1.当l直接等于目标，说明起始位置找到了。打个标记，l不敢动了。
    2.当r直接等于目标，说明结束位置找到了，打个标记，r不敢动。
    如果l和r都找到，就直接返回。
    此时我运行了一下，发现陷入循环中了，为什么呢？因为打了标记.nums[mid]==target的情况我没处理。
    有可能出现虽然nums[mid]==target,但是l和r，其中一方并没有跑到target。
    因此在加一层nums[mid]==target的判断，如果某一方没跑到l++或者r--。


*/


var searchRange = function(nums, target) {
    let l=0,r=nums.length-1,res=[],flagL=0,flagR=0
    //对于长度为0也可以去掉，但发现加上耗时会少一些
    if(nums.length==1&&target==nums[0]){
        return [0,0]
    }
    while(l<=r){
        /*对于l,r边界的筛选 */
        if(nums[l]==target){
            res[0]=l
            flagL=1
        }
        if(nums[r]==target){
            res[1]=r
            flagR=1
        }
        if(flagR&&flagL){
            return res
        }
        //二分的框架
        let mid=parseInt((l+r)/2)
        if(nums[mid]<target&&!flagL){
            l=mid+1
        }
        else if(nums[mid]>target&&!flagR){
            r=mid-1
        }
        else if(nums[mid]==target){
            //其中一方没跑到
            if(!flagL){
               l++;
            }
            if(!flagR){
                r--;
            }

        }
    }
    return [-1,-1]
    
};

/*
    时间复杂度：O(logN)
    空间复杂度：O(1)
*/