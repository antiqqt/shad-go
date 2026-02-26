//go:build !solution

package hogwarts

// DFS is also possible here
func GetCourseList(prereqs map[string][]string) []string {
	adj := make(map[string][]string)
	indegree := make(map[string]int)

	// Build graph
	for course, deps := range prereqs {
		indegree[course] = len(deps)

		for _, dep := range deps {
			if _, ok := indegree[dep]; !ok {
				indegree[dep] = 0
			}
			adj[dep] = append(adj[dep], course)
		}
	}

	queue := make([]string, 0, len(indegree))
	for course, deg := range indegree {
		if deg == 0 {
			queue = append(queue, course)
		}
	}

	var result []string
	for i := 0; i < len(queue); i++ {
		course := queue[i]
		result = append(result, course)

		for _, nei := range adj[course] {
			indegree[nei]--
			if indegree[nei] == 0 {
				queue = append(queue, nei)
			}
		}
	}

	if len(result) != len(indegree) {
		panic("cycle")
	}

	return result
}
