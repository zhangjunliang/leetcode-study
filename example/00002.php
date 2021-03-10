<?php
class ListNode {
    public $val = 0;
    public $next = null;
    function __construct($val = 0) {
        if(is_int($val)){
            $this->val = $val;
            $this->next = null;
        }else if(is_array($val)){
            $this->val = array_shift($val);
            $this->next = $val ? new ListNode($val) : null;
        }
    }
}

class Solution {
    /**
     * @param ListNode $l1
     * @param ListNode $l2
     * @return ListNode
     */
    function addTwoNumbers($l1, $l2) {
        $result = null;
        $s = 0;
        while($l1 or $l2 or $s) {
            $l1_val = $l1 != null ? $l1->val : 0;
            $l2_val = $l2 != null ? $l2->val : 0;
            $tmp_val = $l1_val + $l2_val + $s;
            $s = intval($tmp_val/10);
            $tmp = new ListNode($tmp_val - $s * 10);
            if (is_null($result)) {
                $result = $tmp;
            } else {
                $next->next = $tmp;
            }
            $next = $tmp;
            $l1 && $l1 = $l1->next;
            $l2 && $l2 = $l2->next;
        }
        return $result;
    }
}

$l1 = new ListNode([9,9,9,9,9,9,9]);
$l2 = new ListNode([9,9,9,9]);

$a = new Solution();
$r = $a->addTwoNumbers($l1,$l2);
var_dump($r);