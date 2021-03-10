<?php
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