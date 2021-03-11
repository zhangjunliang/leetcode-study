package main

import (
    "bytes"
    "encoding/json"
    "fmt"
)

//定义链表结构体
type ListNode struct {
    Val int
    Next *ListNode
}
//构造链表
func makeListNode(nums []int) *ListNode {
    if len(nums) == 0{
        return nil
    }
    res := &ListNode{
        Val:nums[0],
    }
    temp := res
    for i := 1; i < len(nums); i++ {
        temp.Next = &ListNode{Val:nums[i],}
        temp = temp.Next
    }
    return  res
}

func (head *ListNode) String() string {
    b, err := json.Marshal(*head)
    if err != nil {
        return fmt.Sprintf("%+v", *head)
    }
    var out bytes.Buffer
    err = json.Indent(&out, b, "", "    ")
    if err != nil {
        return fmt.Sprintf("%+v", *head)
    }
    return out.String()
}

// 遍历输出
func (head *ListNode) Traverse() {
    point := head
    for nil != point {
        fmt.Println(point.Val)
        point = point.Next
    }
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    //创建一个结点值为 None 的头结点, dummy 和 p 指向头结点, dummy 用来最后返回, p 用来遍历
    dummy :=  new(ListNode)
    curr := dummy
    carry := 0
    for (l1 != nil || l2 != nil || carry > 0) {
        curr.Next = new(ListNode)
        curr = curr.Next
        if l1 != nil {
            carry += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            carry += l2.Val
            l2 = l2.Next
        }
        curr.Val = carry % 10
        carry /= 10
    }
    return dummy.Next
}

func main() {

    l1 := makeListNode([]int{9,9,9,9,9,9,9})
    l2 := makeListNode([]int{9,9,9,9})
    l3 := addTwoNumbers(l1,l2)
    fmt.Println(l1.String(),l2.String(),l3.String())
}
