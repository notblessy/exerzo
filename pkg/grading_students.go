package pkg

func gradingStudents(grades []int32) []int32 {
	var res []int32
	for _, g := range grades {
		nextMultiple := (g + 4) / 5 * 5
		value := nextMultiple - g

		if g >= 38 && value < 3 {
			g = nextMultiple
		}

		res = append(res, g)
	}

	return res
}
