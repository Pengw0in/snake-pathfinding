package algorithms

import rl "github.com/gen2brain/raylib-go/raylib"

func Bfs(startRow, startCol, targetRow, targetCol, rows, columns int, segmentPos []rl.Vector2, cellSize int) []int{

	var directions = [][]int{
		{0, 1},  // Right : 0
		{1, 0},  // Down  : 1
		{0, -1}, // Left  : 2
		{-1, 0}, // Up    : 3
	}

	var visited = make([][]bool, rows)
	for i := range visited{
		visited[i] = make([]bool, columns)
	}

	for _, segmentCoord := range segmentPos{
		visited[int(segmentCoord.Y)/cellSize][int(segmentCoord.X)/ cellSize] = true
	}

	type node struct{
		r, c int
		path []int
	}

	var queue = []node{{startRow, startCol, []int{}}}
	visited[startRow][startCol] = true


	for len(queue) > 0 {
		var current = queue[0]
		queue = queue[1:]

		if current.r == targetRow && current.c == targetCol {
			return current.path
		}

		for dirIndex, d := range directions {
			var newRow = current.r + d[0]
			var newCol = current.c + d[1]

			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < columns && !visited[newRow][newCol]{
				visited[newRow][newCol] = true
				
				var newPath = append([]int{}, current.path...)
				newPath = append(newPath, dirIndex)

				queue = append(queue, node{newRow, newCol, newPath})
			}
		}
			
	}
	return nil
}