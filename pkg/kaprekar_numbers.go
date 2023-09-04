package pkg

import "fmt"

/*
 * Complete the 'kaprekarNumbers' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER p
 *  2. INTEGER q
 */

func kaprekarNumbers(p int32, q int32) {
	// Write your code here
	foundOne := false

	for i := p; i <= q; i++ {
		if isKaprekarNumber(int64(i)) {
			foundOne = true
			fmt.Printf("%d ", i)
		}
	}

	if !foundOne {
		fmt.Println("INVALID RANGE")
	} else {
		fmt.Println()
	}

}

func isKaprekarNumber(i int64) bool {
	square := i * i
	sq := fmt.Sprint(square)
	leftString := ""
	rightString := ""
	leftDigit := int64(0)
	rightDigit := int64(0)

	if len(sq) > 1 {
		mid := len(sq) / 2
		leftString = sq[:mid]
		rightString = sq[mid:]
	} else {
		rightString = sq
	}

	if leftString != "" {
		fmt.Sscanf(leftString, "%d", &leftDigit)
	}

	if rightString != "" {
		fmt.Sscanf(rightString, "%d", &rightDigit)
	}

	total := leftDigit + rightDigit

	return total == i
}
