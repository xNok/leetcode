package utils

func dfs(nodes []string, adjList map[string][]string, perf string, pos int, res *[]string) bool {
	//if v is already visited
	if pos == len(nodes) {
		*res = append(*res, perf)
		return true
	}

	for _, v := range adjList[nodes[pos]] {
		// for all neighbors x of v
		dfs(nodes, adjList, perf+v, pos+1, res)
	}

	return false
}

func edgeListToAdjList(edges [][]string) (adj map[string][]string) {
	adj = make(map[string][]string)

	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
	}

	return
}
