package pkg

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// Write your code here
	if x1 == x2 && v1 == v2 {
		return "YES"
	}

	if x1 == x2 && v1 != v2 {
		return "NO"
	}

	if v1 == v2 {
		return "NO"
	}

	if (x1-x2)%(v2-v1) == 0 && (x1-x2)/(v2-v1) >= 0 {
		return "YES"
	}

	return "NO"
}
