package main

import (
	"math/rand"
)

func main() {
	var arr [][]int
	var pointsCount = 10

	//create matrix
	for x := 0; x < pointsCount; x++ {
		for y := 0; y < pointsCount; y++ {
			arr[x][y] = 0
		}
	}

	//create basic paths
	var base = 0
	for x := 0; x < pointsCount; x++ {
		var nextBase = rand.Intn(pointsCount)
		arr[base][nextBase] = 1
		arr[nextBase][base] = 1
		base = nextBase
	}

	//create new paths
	// var badPoints =  { rand.Intn(pointsCount), rand.Intn(pointsCount) }
	var firstPoint = rand.Intn(pointsCount)
	var nonConnectedPoints []int
	for index, point := range arr[firstPoint] {
		if point == 0 {
			nonConnectedPoints[index] = index
		}
	}
}
