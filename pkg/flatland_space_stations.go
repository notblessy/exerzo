package pkg

// Complete the flatlandSpaceStations function below.
func flatlandSpaceStations(n int32, c []int32) int32 {
	if int32(len(c)) == n {
		return 0
	}

	for i := 0; i < len(c)-1; i++ {
		for j := i + 1; j < len(c); j++ {
			if c[i] > c[j] {
				c[i], c[j] = c[j], c[i]
			}
		}
	}

	maxDistance := c[0]
	lastStation := c[0]

	for i := 1; i < len(c); i++ {
		distance := (c[i] - lastStation) / 2
		if distance > maxDistance {
			maxDistance = distance
		}
		lastStation = c[i]
	}

	distanceToLastCity := n - 1 - lastStation
	if distanceToLastCity > maxDistance {
		maxDistance = distanceToLastCity
	}

	return maxDistance

}
