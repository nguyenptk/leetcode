// https://leetcode.com/problems/unique-paths-ii/
package medium

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	r := len(obstacleGrid)
	c := len(obstacleGrid[0])

	// If obstacle is at starting position
	if obstacleGrid[0][0] != 0 {
		return 0
	}

	//  Initializing starting position
	obstacleGrid[0][0] = 1

	// first row all are '1' until obstacle
	for j := 1; j < c; j++ {
		if obstacleGrid[0][j] == 0 {
			obstacleGrid[0][j] = obstacleGrid[0][j-1]
		} else {
			// No ways to reach at this index
			obstacleGrid[0][j] = 0
		}
	}

	// first column all are '1' until obstacle
	for i := 1; i < r; i++ {
		if obstacleGrid[i][0] == 0 {
			obstacleGrid[i][0] = obstacleGrid[i-1][0]
		} else {
			// No ways to reach at this index
			obstacleGrid[i][0] = 0
		}
	}

	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			// If current cell has no obstacle
			if obstacleGrid[i][j] == 0 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			} else {
				// No ways to reach at this index
				obstacleGrid[i][j] = 0
			}
		}
	}
	// returning the bottom right
	// corner of Grid
	return obstacleGrid[r-1][c-1]
}
