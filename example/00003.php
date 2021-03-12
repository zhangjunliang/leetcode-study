<?php
class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    function lengthOfLongestSubstring($s) {
        $r = 0;
        $not_repeat = [];
        for($i = 0; $i< strlen($s); $i++){
            $v = $s[$i];
            if(in_array($v,$not_repeat)){
                #如果存在重复元素，删除到重复的位置，重新开始计算
                array_splice($not_repeat, 0, array_search($v,$not_repeat)+1);
                $not_repeat[] = $v;
            }else{
                #不存在重复元素，计算最大长度并且保存长度
                $not_repeat[] = $v;
                $l = count($not_repeat);
                $l > $r && $r = $l;
            }
        }
        return $r;
    }
}

$s1 = "abcabcbb";
$s2 = 'pwwkew';

$s = new Solution();
var_dump($s->lengthOfLongestSubstring($s1));
var_dump($s->lengthOfLongestSubstring($s2));