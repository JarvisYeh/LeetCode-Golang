package main

func canFinish207(numCourses int, prerequisites [][]int) bool {
	incomingNodes := make([]int, numCourses)
	outWards := make([][]int, numCourses)
	// construct adjacency list graph
	// and incoming edges array
	for _, pair := range prerequisites {
		prev := pair[1]
		post := pair[0]
		incomingNodes[post] += 1
		if outWards[prev] == nil {
			outWards[prev] = []int{post}
		} else {
			outWards[prev] = append(outWards[prev], post)
		}
	}

	// push no prerequisites course into queue
	q := []int{}
	for i, count := range incomingNodes {
		if count == 0 {
			q = append(q, i)
		}
	}

	complete := 0
	for len(q) != 0 {
		curr := q[0]
		q = q[1:]
		complete++
		if outWards[curr] == nil {
			continue
		}
		for _, nextCourse := range outWards[curr] {
			incomingNodes[nextCourse]--
			if incomingNodes[nextCourse] == 0 {
				q = append(q, nextCourse)
			}
		}
	}
	return complete == numCourses
}
