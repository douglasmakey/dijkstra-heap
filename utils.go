package main

func wasVisited(node string, nodes []string) bool {
	for _, n := range nodes {
		if n == node {
			return true
		}
	}

	return false
}
