// 1337. 矩阵中战斗力最弱的 K 行
// 给你一个大小为 m * n 的矩阵 mat，矩阵由若干军人和平民组成，分别用 1 和 0 表示。

// 请你返回矩阵中战斗力最弱的 k 行的索引，按从最弱到最强排序。

// 如果第 i 行的军人数量少于第 j 行，或者两行军人数量相同但 i 小于 j，那么我们认为第 i 行的战斗力比第 j 行弱。

// 军人 总是 排在一行中的靠前位置，也就是说 1 总是出现在 0 之前。



// 示例 1：

// 输入：mat = 
// [[1,1,0,0,0],
//  [1,1,1,1,0],
//  [1,0,0,0,0],
//  [1,1,0,0,0],
//  [1,1,1,1,1]], 
// k = 3
// 输出：[2,0,3]
// 解释：
// 每行中的军人数目：
// 行 0 -> 2 
// 行 1 -> 4 
// 行 2 -> 1 
// 行 3 -> 2 
// 行 4 -> 5 
// 从最弱到最强对这些行排序后得到 [2,0,3,1,4]
// 示例 2：

// 输入：mat = 
// [[1,0,0,0],
//  [1,1,1,1],
//  [1,0,0,0],
//  [1,0,0,0]], 
// k = 2
// 输出：[0,2]
// 解释： 
// 每行中的军人数目：
// 行 0 -> 1 
// 行 1 -> 4 
// 行 2 -> 1 
// 行 3 -> 1 
// 从最弱到最强对这些行排序后得到 [0,2,3,1]


// 提示：

// m == mat.length
// n == mat[i].length
// 2 <= n, m <= 100
// 1 <= k <= m
// matrix[i][j] 不是 0 就是 1


/*
    思路：先说说自己想的思路吧。
         由于军人和平民用1和0表示，且军人一定在靠前位置
         所以可以思考的即遍历方式。一般情况下我会习惯横向遍历。
         但此题，可以看到，纵向遍历能最快的知道某列该行为0的就是最弱的。
         因此我开始竖着找0,同时，标记visited，该行已经被我放入目标数组。
         特殊情况判断：目标数组长度>k，切断
                            <k,补上（此时还是要用上之前的标记）
        看完题解，对于二分的优化程度并不是很确定，因为二分找到每一行的目标位(获取战斗力)，存储起来，依然需要排序。
        二分：由于军人一直在平民前面，因此，每一行是排序过的
        理解理解思路，尝试自己写的时候，又陷入细节的陷阱，后来冷静下来，把二分方法给封装起来使用。
        对于此题而言：1.二分获取每一行的战斗力，并存起来
                    2对存起来的数组根据战斗力排序，然后再获取目标位置




*/


var kWeakestRows = function (mat, k) {
    let m = mat.length, n = mat[0].length, res = [], visited = []
    for (let i = 0; i < n; i++) {
        let count = 0
        //竖着遍历找0
        for (let j = 0; j < m; j++) {
            if (mat[j][i] == 0 && !visited[j]) {
                res.push(j)
                visited[j] = 1
            }
        }
        if (res.length >= k) {
            return res.slice(0, k)
        }

    }
    for (let i = 0; i < k; i++) {
        if (!visited[i]) {
            res.push(i)
        }
    }
    return res.slice(0, k)
};
/*
    时间复杂度：O(M*N)
    空间复杂度：O(max(M,N))
*/


var kWeakestRows = function (mat, k) {
    let m = mat.length, n = mat[0].length, result = [], res = []
    //封装起来更清晰
    let dfs = (arr, left, right) => {
        while (left <= right) {
            let mid = parseInt((left + right) / 2)
            if (arr[mid] == 0) {
                right = mid - 1
            }
            if (arr[mid] === 1) {
                left = mid + 1
            }
        }
        return left
    }
    for (let i = 0; i < m; i++) {
        //获取每一行的战斗力
        res.push([i, dfs(mat[i], 0, n - 1)])
    }

    res.sort((a, b) => a[1] - b[1])
    res.forEach(item => {
        result.push(item[0])
    })
    return result.slice(0, k)
}


/*
    时间复杂度：O(M*logN)
    空间复杂度：O(M*N)
*/
