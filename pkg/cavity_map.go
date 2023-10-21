package pkg

/*
 * Complete the 'cavityMap' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts STRING_ARRAY grid as parameter.
 */

func cavityMap(grid []string) []string {
	// Write your code here
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if grid[i][j] > grid[i][j-1] && grid[i][j] > grid[i][j+1] &&
				grid[i][j] > grid[i-1][j] && grid[i][j] > grid[i+1][j] {
				grid[i] = grid[i][:j] + "X" + grid[i][j+1:]
			}
		}
	}
	return grid
}
