package hamilton

func getIndex(arr []int, point int) int {
	for index, value := range arr {
		if value == point {
			return index
		}
	}
	return -1
}

var arr []int

func Hamilton(p int, graph *[][]int) []int {

	var graphValue [][]int = *graph
	// v = append(p, v)
	arr = append(arr, p)

	for index, point := range graphValue[p] {
		if point == 0 {
			Hamilton(index, &graphValue)
		}
		if len(arr) == len(graphValue) && graphValue[p][0] == 1 {
			return arr
		} else {
			var pIndex = getIndex(arr, p)
			// fmt.Printf("%v %v %v \n", pIndex, p, len(arr))
			if pIndex != -1 {
				arr = append(arr[:pIndex], arr[pIndex+1:]...) // remove point from arr
			}

		}
	}
	return arr
}
