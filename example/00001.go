package main

import "fmt"

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

func main() {
    var nums = []int{2,7,11,15}
    var target int  = 9
    result := twoSum(nums, target)
    fmt.Println(result)
}