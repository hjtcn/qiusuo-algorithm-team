/*
78. 子集
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

 

示例 1：

输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
示例 2：

输入：nums = [0]
输出：[[],[0]]
 

提示：

1 <= nums.length <= 10
-10 <= nums[i] <= 10
nums 中的所有元素 互不相同
*/

/*
    思路：勇气，哈哈哈
        首先一次循环固定子集的长度count。
        然后递归，统计目前子集长度是否达到目标长度了。
        let dfs=(arr,i,count)=>{
            //count目标子集长度，保持不变
            // 边界为
            if(arr.length==count){
                //要深copy不然会被修改
                target.push([...arr])
                return;
            }
            for(let j=i+1;j<len;j++){
                arr.push(nums[j])
                dfs(arr,j,count)
                arr.pop()
            }
        }
        //踩的坑：1.首先要用arr.push()返回值是新数组的长度，不能作为新数组放到递归参数中。
                2.深浅拷贝问题
                3.dfs函数的第二个值：当前查询到哪里了
    题解：看完题解，感觉自己写出来的喜悦没有了。
         人家的思维更简单。学习学习，进步进步
         方法1：递归枚举，是否选择当前值
         方法2:位运算迭代，牵扯到位元算就好难思考呀
         比如for循环控制for(mask=0;mask < (1 << n);++mask)
         变为二进制的循环：就会是从(0~2^n-1)
         这个就没遇到过。
         如：[1,2,3]对应的二进制为
            (0,0,0) ～  (1,1,1)     
            每个位置代表0，1代表是否取值。


*/

var subsets = function (nums) {
    let len = nums.length
    let target = []
    target.push([])
    let dfs = (arr, i,count) => {
        if (i == len||arr.length==count) {
            target.push([...arr])
            return;
        }
        for (let j = i + 1; j < len; j++) {
            arr.push(nums[j])
            dfs(arr,j,count)
            arr.pop()
        }
        return false
    }
    for (let i = 1; i <= len; i++) {
         let arr=[]
        for (let j = 0; j < len; j++) {
            arr.push(nums[j])
            dfs(arr, j,i)
            arr.pop()
        }
    }
    return target
};


var subsets = function (nums) {
    let len = nums.length,target=[],arr=[]
    let dfs=(curIndex)=>{
        if(curIndex==len){
            target.push([...arr])
            return ;
        }
        arr.push(nums[curIndex])
        dfs(curIndex+1)
        arr.pop()
        dfs(curIndex+1)
    }
    dfs(0)
    return target
}

var subsets = function(nums) {
    const ans = [];
    const n = nums.length;
    for (let mask = 0; mask < (1 << n); ++mask) {
        const t = [];
        for (let i = 0; i < n; ++i) {
            //先把1左移i位，然后把得到的结果再和x进行按位与运算。这里不懂，盲猜二进制下当前位置为1
            if (mask & (1 << i)) {
                t.push(nums[i]);
            }
        }
        ans.push(t);
    }
    return ans;
};

/*
    时间复杂度：O(N*2^N)
    空间复杂度：O(N)
*/

/*
    思考：不要怕，慢慢写。
        用递归只需要两步：
        1.边界
        2.子问题
*/