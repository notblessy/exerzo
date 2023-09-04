package pkg

func BirthdayCakeCandles(candles []int32) int32 {
	var tallests []int32

	max := candles[0]
	for _, c := range candles {
		if c > max {
			max = c
		}
	}

	for _, c := range candles {
		if c == max {
			tallests = append(tallests, c)
		}
	}

	return int32(len(tallests))
}
