package main

// find lca with parent pointer

type TreeNodeP struct {
	Val    int
	Left   *TreeNodeP
	Right  *TreeNodeP
	Parent *TreeNodeP
}

// use hashset
// TC: O(log n)
// SC: O(log n)
func lowestCommonAncestor1650I(a *TreeNodeP, b *TreeNodeP) *TreeNodeP {
	set := make(map[*TreeNodeP]bool)
	for a != nil {
		set[a] = true
		a = a.Parent
	}

	for b != nil {
		if _, ok := set[b]; ok {
			return b
		}
		b = b.Parent
	}
	return nil
}

// move to same level, and together move up until lca is found
// TC: O(log n)
// SC: O(1)
func lowestCommonAncestor1650II(a *TreeNodeP, b *TreeNodeP) *TreeNodeP {
	depthA := getDepth(a)
	depthB := getDepth(b)

	deep := a
	shallow := b
	diff := depthA - depthB
	if depthB > depthA {
		deep = b
		shallow = a
		diff = depthB - depthA
	}

	// move to same level
	for i := 0; i < diff; i++ {
		deep = deep.Parent
	}

	for deep != shallow {
		deep = deep.Parent
		shallow = shallow.Parent
	}
	return deep
}

func getDepth(node *TreeNodeP) int {
	depth := 0
	for node != nil {
		depth++
		node = node.Parent
	}
	return depth
}
