# leetcode-study

            ______________
            ___________(_)__  /
            ___  /____  /__  / 
            __  /____  / _  /  
            _____/__  /  /_/   
               /___/  2021 day day up
    
    力扣学习签到，活跃思维，自我成长
    来源：力扣（LeetCode）
 
#### 00001.两数之和 

> 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。

> 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

> 你可以按任意顺序返回答案。

示例 1：

    输入：nums = [2,7,11,15], target = 9
    输出：[0,1]
    解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

    输入：nums = [3,2,4], target = 6
    输出：[1,2]
示例 3：

    输入：nums = [3,3], target = 6
    输出：[0,1]
 
提示：

    2 <= nums.length <= 103
    -109 <= nums[i] <= 109
    -109 <= target <= 109
    只会存在一个有效答案

- 解题方法
    - 暴力枚举：时间复杂度O(N²),空间复杂度O(1)
    - 哈希表：时间复杂度O(N),空间复杂度O(N)
    
- 实例代码

    - golang *runtime:4 ms,memory:3 MB*
    ```
    func twoSum(nums []int, target int) []int {
       hash_table := map[int]int{}
       for i, v := range nums {
           var diff int = target - v
           if p, ok := hash_table[diff]; ok {
               return []int{p, i}
           }
           hash_table[v] = i
       }
       return []int{}
    }
    ```
    - python3 *runtime:44 ms,memory:14.6 MB*
    ```
    class Solution:
        def twoSum(self, nums: List[int], target: int) -> List[int]:
            """
            :type nums: List[int]
            :type target: int
            :rtype: List[int]
            """
            hash_table = {}
            for index,val in enumerate(nums):
                diff = target - val
                if diff in hash_table:
                    return [hash_table[diff],index]
                hash_table[val] = index     
            return []
    ```
    - php *runtime:20 ms,memory:15.3 MB*
    ```
    class Solution {
        /**
         * @param Integer[] $nums
         * @param Integer $target
         * @return Integer[]
         */
        function twoSum($nums, $target) {
            $hash_table = [];
            foreach($nums as $k=>$v){
                if(in_array($target - $v)){
                    return [$hash_table[$target - $v],$k]
                }
                $hash_table[$v] = $k;
            }
            return []
        }
    }
    ```
    
#### 00001.两数相加 
        
> 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

> 请你将两个数相加，并以相同形式返回一个表示和的链表。

> 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例 1：

![addtwonumber1](./data/images/addtwonumber1.jpg)

    输入：l1 = [2,4,3], l2 = [5,6,4]
    输出：[7,0,8]
    解释：342 + 465 = 807.
    
示例 2：

    输入：l1 = [0], l2 = [0]
    输出：[0]
    
示例 3：

    输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
    输出：[8,9,9,9,0,0,0,1]

提示：

    每个链表中的节点数在范围 [1, 100] 内
    0 <= Node.val <= 9
    题目数据保证列表表示的数字不含前导零