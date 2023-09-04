package pkg

/*
 * Complete the 'breakingRecords' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY scores as parameter.
 */

func breakingRecords(scores []int32) []int32 {
	// Write your code here
	var high int32
	var low int32

	var highest int32
	var lowest int32
	for i, j := 0, 0; i < len(scores); i, j = i+1, j+1 {
		if i == 0 && j == 0 {
			highest = scores[i]
			lowest = scores[j]
			continue
		}

		if scores[i] > highest {
			high++
			highest = scores[i]
		}

		if scores[j] < lowest {
			low++
			lowest = scores[j]
		}
	}

	return []int32{high, low}
}
