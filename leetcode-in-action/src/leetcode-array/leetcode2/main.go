package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	moveOne := l1
	moveTwo := l2
	var result *ListNode
	var move *ListNode
	suplus := 0
	for moveOne != nil && moveTwo != nil {
		calVal := (moveOne.Val + moveTwo.Val + suplus) % 10
		suplus = (moveOne.Val + moveTwo.Val + suplus) / 10
		if result == nil {
			result = &ListNode{
				Val:  calVal,
				Next: nil,
			}
			move = result
		} else {
			node := &ListNode{
				Val:  calVal,
				Next: nil,
			}
			move.Next = node
			move = node
		}
		moveOne = moveOne.Next
		moveTwo = moveTwo.Next
	}
	for ; moveOne != nil; moveOne = moveOne.Next {
		calVal := (moveOne.Val + suplus) % 10
		suplus = (moveOne.Val + suplus) / 10
		node := &ListNode{
			Val:  calVal,
			Next: nil,
		}
		move.Next = node
		move = node
	}
	for ; moveTwo != nil; moveTwo = moveTwo.Next {
		calVal := (moveTwo.Val + suplus) % 10
		suplus = (moveTwo.Val + suplus) / 10
		node := &ListNode{
			Val:  calVal,
			Next: nil,
		}
		move.Next = node
		move = node
	}
	for ; suplus != 0; suplus = suplus / 10 {
		calVal := suplus % 10
		node := &ListNode{
			Val:  calVal,
			Next: nil,
		}
		move.Next = node
		move = node
	}
	return result
}

func main() {

}
