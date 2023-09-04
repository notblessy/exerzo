package pkg

/*
 * Complete the 'birthday' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY s
 *  2. INTEGER d
 *  3. INTEGER m
 */

func birthday(s []int32, d int32, m int32) int32 {
	// Write your code here
	var ways int32
	for i := int32(0); i <= int32(len(s))-m; i++ {
		var segmentSum int32
		for j := i; j < i+m; j++ {
			segmentSum += s[j]
		}
		if segmentSum == d {
			ways++
		}
	}

	return ways
}
