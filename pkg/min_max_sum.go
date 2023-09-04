package pkg

import "fmt"

func MiniMaxSum(arr []int32) {
	if len(arr) == 0 {
		fmt.Println("0 0")
	}

	var newArr []int64
	for _, v := range arr {
		newArr = append(newArr, int64(v))
	}

	n := len(newArr)
	swapped := true

	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if newArr[i-1] > newArr[i] {
				// Swap elements if they are in the wrong order
				newArr[i-1], newArr[i] = newArr[i], newArr[i-1]
				swapped = true
			}
		}
		n--
	}

	min := newArr[0] + newArr[1] + newArr[2] + newArr[3]

	max := newArr[1] + newArr[2] + newArr[3] + newArr[4]

	fmt.Printf("%d %d\n", min, max)
}
