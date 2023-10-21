package pkg

/*
 * Complete the 'stones' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER a
 *  3. INTEGER b
 */

func stones(n int32, a int32, b int32) []int32 {
	// Write your code here
	var d1, d2 int32
	var v []int32

	if a == b {
		v = append(v, (n-1)*a)
		return v
	}

	for i := int32(0); i < n; i++ {
		d1 = i
		d2 = n - i - int32(1)
		v = append(v, d1*a+d2*b)
	}

	for i := 0; i < len(v)-1; i++ {
		for j := 0; j < len(v)-i-1; j++ {
			if v[j] > v[j+1] {
				v[j], v[j+1] = v[j+1], v[j]
			}
		}
	}

	return v
}
