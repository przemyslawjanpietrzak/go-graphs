package euler

func Euler(v int, graph *[][]int) {
	var c []int
	var graphValue [][]int = *graph
	for index, w := range graphValue[v] {
		if graphValue[v][index] == 0 {
			graphValue[v][w] = 1
			graphValue[w][v] = 1
			c = append(c, index)
			Euler(index, &graphValue)
			break // i'm not sure
		}
	}
}
