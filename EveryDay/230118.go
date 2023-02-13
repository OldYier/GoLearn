package EveryDay

import "fmt"

func Func230118() {
	var mountainH = [8]int{0, 8, 7, 3, 6, 5, 2, 4}
	for {
		var j = 0
		var sum = 0
		for i := 0; i < 8; i++ {
			// mountainH: represents the height of one mountain.

			if mountainH[i] > j {
				sum = i
				j = mountainH[i]
			}
			fmt.Scan(&mountainH)
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		fmt.Println(sum) // The index of the mountain to fire on.
	}
}
