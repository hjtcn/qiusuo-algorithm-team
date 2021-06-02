/*
剑指 Offer 40. 最小的k个数
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

 

示例 1：

输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]
示例 2：

输入：arr = [0,1,2,1], k = 1
输出：[0]
 

限制：

0 <= k <= arr.length <= 10000
0 <= arr[i] <= 10000
*/

/*
    思路：先排序，然后返回包含前k个最小数的数组

    题解：快排:
         quickSort方法：输入为数组arr
                       输出为数组
         如果arr的长度<=1，则返回arr
         如果>1,假设base=arr[0]
         for(let i=1;i<len;i++){
             left放置比base小的值
             right放置比base大的值
         }
        返回递归过程：return [...quickSort(left),...[base],...quickSort(right)]

         
        堆排序:大顶堆，根结点>其子节点的值(子节点无大小比较)
              完全二叉树。
              1.交换方法：借用数组利用参数(arr,i,j)交换在数组中的位置。
              2.构建大顶堆：如果i是根节点 parentNode=arr[i]
                    i            这是一棵小树
                2*i+1 2*i+2
                目的：使parentNode最大？
                1)比较子节点的大小。循环叶子节点设置为j(2*i+1)
                如果arr[j]<arr[j+1]
                j=j+1//只是为了得到大一点的索引
                2)父节点对比刚获取到的较大子节点
                如果小于arr[j],进行交换。同时i=j(这个交换的过程有点懵)
             3.排序：1)遍历每一个父节点，进行构建大顶堆
                    2)交换第一个元素(最大值)，和目标开始子节点。
                      并继续构建大顶堆
            
            堆排序只能比对着理解，但是目前自己还是写不出来，父子节点的i,j变化太灵活了。


*/
var getLeastNumbers = function(arr, k) {
    arr.sort((a,b)=>a-b)
    return arr.splice(0,k)
};

var getLeastNumbers = function(arr, k) {
    let quickSort=(arr)=>{
        let left=[],right=[]
        let len=arr.length
        if(arr.length<=1) return arr
        let base=arr[0]
        for(let i=1;i<len;i++){
            if(arr[i]<base){
                left.push(arr[i])
            }
            else{
                right.push(arr[i])
            }
        }
        return [...quickSort(left),...[base],...quickSort(right)]
    }
    return quickSort(arr).slice(0,k)
}
/*
    时间复杂度：O(nlogn) 最差(n^2)
    空间复杂度：O(n)
*/


var getLeastNumbers = function(arr, k) {
    let swap=(arr,i,j)=>{
        let tmp=arr[i]
        arr[i]=arr[j]
        arr[j]=tmp
    }
    
    let shiftDown=(arr,i,len)=>{
        let pNode=arr[i]
        for(let j=2*i+1;j<len;j=2*j+1){
            pNode=arr[i]
            //比较子节点
            if(j+1<len&&arr[j]<arr[j+1]){
                //说明右节点大一点
                j++;
            }
            
            if(pNode<arr[j]){//此时arr[j]是子节点中比较大的那个
                swap(arr,i,j)
                i=j//交换后，父节点的索引变为j!!!这个地方很核心呀。
            }
            else{
                break;
            }
        }
    }

    let heapSort=(arr)=>{
        let len=arr.length
        for(let i=Math.floor((len-1)/2);i>=0;i--){
            shiftDown(arr,i,len)
        }
        for(let i=Math.floor(len-1);i>0;i--){
            swap(arr,0,i)
            shiftDown(arr,0,i)
        }
        return arr
    }
    return heapSort(arr).slice(0,k)
}

/*
    时间复杂度：O(nlogn)
    空间复杂度：O(n)
*/

/*
    思考：这应该就是我们想要找的排序题，可挑战，练习
*/