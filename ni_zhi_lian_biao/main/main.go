package main
import (
	"fmt"
	."github.com/isdamir/gotype"
)
//这是带头节点的 面试谁这么写啊。。。。
func Reverse(node *LNode)  {
	if node == nil || node.Next == nil{
		return
	}
	var pre *LNode
	var cur *LNode
	next := node.Next
	for next != nil {
		cur = next.Next
		next.Next =pre
		pre =next
		next =cur
	}
	node.Next =pre
}
//不带头节点的
func reverseList(head *LNode) *LNode {
	var prev *LNode
	cur := head
	for cur != nil {
		next := cur.Next  //简单的记忆  脑残嫖娼 右上=左下
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}


func main()  {
	head := &LNode{}
	fmt.Println("逆序")
	CreateNode(head,4)
	PrintNode("原本:" ,head)
	Reverse(head)
	PrintNode("现在 :" ,head)
}
