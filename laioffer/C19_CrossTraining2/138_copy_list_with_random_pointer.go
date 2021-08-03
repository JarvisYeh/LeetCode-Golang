package main

type Node138 struct {
	Val    int
	Next   *Node138
	Random *Node138
}

func copyRandomList138(head *Node138) *Node138 {
	dummyHead := &Node138{-1, nil, nil}
	prev := dummyHead
	nodeMap := make(map[*Node138]*Node138)

	for head != nil {
		// copy node
		if node, ok := nodeMap[head]; ok {
			prev.Next = node
		} else {
			tmpNode := &Node138{head.Val, nil, nil}
			prev.Next = tmpNode
			nodeMap[head] = tmpNode
		}

		// copy random
		if head.Random != nil {
			if node, ok := nodeMap[head.Random]; ok {
				prev.Next.Random = node
			} else {
				tmp := &Node138{head.Random.Val, nil, nil}
				prev.Next.Random = tmp
				nodeMap[head.Random] = tmp
			}
		}

		// update reference
		head = head.Next
		prev = prev.Next
	}
	return dummyHead.Next
}
