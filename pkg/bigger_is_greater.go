package pkg

/*
 * Complete the 'biggerIsGreater' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING w as parameter.
 */

func biggerIsGreater(w string) string {
	// Write your code here
	chars := []byte(w)

	i := len(chars) - 1
	for i > 0 && chars[i-1] >= chars[i] {
		i--
	}

	if i <= 0 {
		return "no answer"
	}

	j := len(chars) - 1
	for chars[j] <= chars[i-1] {
		j--
	}

	chars[i-1], chars[j] = chars[j], chars[i-1]

	// Reverse the suffix manually
	left, right := i, len(chars)-1
	for left < right {
		chars[left], chars[right] = chars[right], chars[left]
		left++
		right--
	}

	return string(chars)
}
