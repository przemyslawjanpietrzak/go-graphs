package data

import (
	"errors"
	"math/rand"
)

func connect3Points(graph [][]int, pointsCount int) ([3]int, error) {
	var firstPoint = rand.Intn(pointsCount)
	var nonConnectedPoints = [3]int{}
	var step1Completed = false
	for index, point := range graph[firstPoint] {
		if point == 0 && nonConnectedPoints[0] != 1 && !step1Completed { // find first path
			nonConnectedPoints[0] = index
			step1Completed = true
		} else if point == 0 && nonConnectedPoints[1] != 1 && graph[nonConnectedPoints[0]][nonConnectedPoints[1]] == 0 { //find second patch AND can create third
			nonConnectedPoints[1] = index
			nonConnectedPoints[2] = firstPoint

			return nonConnectedPoints, nil
		}
	}
	return nonConnectedPoints, errors.New("Points not found")
}

func GetData() [][]int {
	var pointsCount = 10

	var graph = make([][]int, pointsCount)
	for row := range graph {
		graph[row] = make([]int, pointsCount)
	} // mean [pointsCount][pointsCount]int

	var edgeCount = 0
	var maxEdgeCount = pointsCount * (pointsCount - 2) / 2

	//create matrix
	for x := 0; x < pointsCount; x++ {
		for y := 0; y < pointsCount; y++ {
			graph[x][y] = 0
		}
	}

	//create basic paths
	for x := 0; x < pointsCount-1; x++ {
		graph[x][x+1] = 1
		graph[x+1][x] = 1
	}
	edgeCount += pointsCount - 1

	//create new paths
	for edgeCount < maxEdgeCount/3 {

		var connect3PointsResult, error = connect3Points(graph, pointsCount)
		for error != nil { // call function until he find 3 points
			connect3PointsResult, error = connect3Points(graph, pointsCount)
		}

		graph[connect3PointsResult[0]][connect3PointsResult[1]] = 1
		graph[connect3PointsResult[1]][connect3PointsResult[0]] = 1

		graph[connect3PointsResult[1]][connect3PointsResult[2]] = 1
		graph[connect3PointsResult[2]][connect3PointsResult[1]] = 1

		graph[connect3PointsResult[2]][connect3PointsResult[0]] = 1
		graph[connect3PointsResult[0]][connect3PointsResult[2]] = 1

		edgeCount += 3
	}
	// fmt.Printf("%v", graph)
	return graph
}
