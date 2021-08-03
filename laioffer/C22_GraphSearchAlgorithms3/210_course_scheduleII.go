package main

// return array of courses for possible ordering of study all courses
func findOrder210(numCourses int, prerequisites [][]int) []int {
	// construct the directed graph
	// and incoming edges array for the graph
	incoming := make([]int, numCourses)
	graph := make([][]int, numCourses)
	for i, _ := range graph {
		graph[i] = []int{}
	}

	for _, pair := range prerequisites {
		end, start := pair[0], pair[1]
		graph[start] = append(graph[start], end)
		incoming[end]++
	}

	// push all 0 incoming edges courses
	res := []int{}
	q := []int{}
	for course, inNums := range incoming {
		if inNums == 0 {
			q = append(q, course)
		}
	}
	for len(q) != 0 {
		currCourse := q[0]
		q = q[1:]
		res = append(res, currCourse)
		// for all neighbors (directed graph neighbors)
		for _, nextCourse := range graph[currCourse] {
			incoming[nextCourse]--
			// if that course's incoming edges decreases to 0
			// which means it's ok to study that course
			// append it the queue
			if incoming[nextCourse] == 0 {
				q = append(q, nextCourse)
			}
		}
	}

	if len(res) == numCourses {
		return res
	}
	return []int{}
}
