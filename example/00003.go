package main

import (
    "fmt"
    "strings"
)

//实现1 字符串
func lengthOfLongestSubstring(s string) int {
    var not_repeat string
    r := 0
    l := len(s)
    for i :=0;i < l;i ++ {
        v := s[i:i+1]
        i_index := strings.Index(not_repeat,v)
        if i_index > -1 {
            not_repeat = not_repeat[i_index+1:]
            not_repeat += v
        } else {
            not_repeat += v
            n_l := len(not_repeat)
            if n_l > r {
                r = n_l
            }
        }
    }
    return r
}
//实现2
func lengthOfLongestSubstring2(s string) int {
    not_repeat := []byte{}
    r := 0
    l := len(s)
    for i :=0;i < l;i ++ {
        v := s[i]
        //v_i := indexOf(not_repeat,v)
        v_i := -1
        for r_i, r_v := range not_repeat {
            if r_v == v {
                v_i = r_i
                break
            }
        }
        if v_i > -1 {
            not_repeat = not_repeat[v_i+1:]
            not_repeat = append(not_repeat,v)
        } else {
           not_repeat = append(not_repeat,v)
           n_l := len(not_repeat)
           if n_l > r {
               r = n_l
           }
        }

    }
    return r
}

func indexOf(l []byte, s byte) int {
    r := -1
    for i, v := range l {
        if v == s {
            r = i
            break
        }
    }
    return r
}


//实现3
func lengthOfLongestSubstring3(s string) int {
    // 哈希集合，记录每个字符是否出现过
    m := map[byte]int{}
    n := len(s)
    // 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
    rk, ans := -1, 0
    for i := 0; i < n; i++ {
        if i != 0 {
            // 左指针向右移动一格，移除一个字符
            delete(m, s[i-1])
        }
        for rk + 1 < n && m[s[rk+1]] == 0 {
            // 不断地移动右指针
            m[s[rk+1]]++
            rk++
        }
        // 第 i 到 rk 个字符是一个极长的无重复字符子串
        ans = max(ans, rk - i + 1)
    }
    return ans
}

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

func main() {
   //s1 := "abcabcbb"
   var s2 string = "aaaab"
   //fmt.Println(lengthOfLongestSubstring2(s1))
   fmt.Println(lengthOfLongestSubstring2(s2))
}
