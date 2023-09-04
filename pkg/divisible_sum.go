package pkg

func divisibleSumPairs(n int, ar []int, k int) int {
	count := 0

	// Iterasi melalui array untuk mencari pasangan (i, j) di mana i < j
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if (ar[i]+ar[j])%k == 0 {
				count++
			}
		}
	}

	return count
}
