package pkg

import "fmt"

func Staircase(n int32) {

	// Loop through each step
	for i := int32(1); i <= n; i++ {
		// Print spaces for each step
		for j := int32(1); j <= n-i; j++ {
			fmt.Print(" ")
		}

		// Print the '#' character for each step
		for k := int32(1); k <= i; k++ {
			fmt.Print("#")
		}

		// Move to the next line for the next step
		fmt.Println()
	}
}
