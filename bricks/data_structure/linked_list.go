// варианты создания одно-связного списка

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	Head *ListNode
}

// через константы
var (
	HeadOne = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val:  7,
					Next: nil,
				},
			},
		},
	}
	HeadTwo = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val:  8,
					Next: nil,
				},
			},
		},
	}
)

func main() {
	current := getStaticList()
	for current != nil {
		fmt.Println(current)
		current = current.Next
	}
	fmt.Println("********************")
	current = HeadOne
	for current != nil {
		fmt.Println(current)
		current = current.Next
	}

	fmt.Println("********************")
	current = generateList()
	for current != nil {
		fmt.Println(current)
		current = current.Next
	}
}

// статично
func getStaticList() *ListNode {
	nodeSix := &ListNode{Val: 6, Next: nil}
	nodeFive := &ListNode{Val: 5, Next: nodeSix}
	nodeFour := &ListNode{Val: 4, Next: nodeFive}
	nodeThree := &ListNode{Val: 3, Next: nodeFour}
	nodeTwo := &ListNode{Val: 2, Next: nodeThree}
	nodeOne := &ListNode{Val: 1, Next: nodeTwo}

	return nodeOne
}

// через функцию
func (list *List) PushForward(val int) *ListNode {
	newNode := &ListNode{
		Val: val,
	}
	if list.Head == nil {
		list.Head = newNode
		return list.Head
	}
	newNode.Next = list.Head
	list.Head = newNode

	return list.Head
}

func generateList() *ListNode {
	list := List{}

	list.PushForward(10)
	list.PushForward(20)
	list.PushForward(30)
	list.PushForward(40)

	return list.Head
}
