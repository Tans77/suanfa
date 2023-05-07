package main
import (
	"fmt"
	."github.com/isdamir/gotype"
)

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
func main()  {
	head := &LNode{}
	fmt.Println("逆序")
	CreateNode(head,10)
	PrintNode("原本:" ,head)
	Reverse(head)
	PrintNode("现在:" ,head)


}
