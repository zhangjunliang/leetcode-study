# -*- coding: UTF-8 -*-
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