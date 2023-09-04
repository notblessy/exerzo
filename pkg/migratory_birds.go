package pkg

/*
 * Complete the 'migratoryBirds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func migratoryBirds(arr []int32) int32 {
	// Write your code here

	birdCount := make(map[int32]int32)

	var maxCount int32
	var mostCommonBird int32

	for _, birdType := range arr {
		birdCount[birdType]++
		if birdCount[birdType] > maxCount {
			maxCount = birdCount[birdType]
			mostCommonBird = birdType
		}

		if birdCount[birdType] == maxCount && birdType < mostCommonBird {
			mostCommonBird = birdType
		}
	}

	return mostCommonBird

}
