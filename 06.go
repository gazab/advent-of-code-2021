package main

import "fmt"

func Solve06A(shoal []int) {

	for i := 0; i < 80; i++ {

		var length = len(shoal)
		for f := 0; f < length; f++ {
			fish := shoal[f]

			if fish > 0 {
				shoal[f]--
			} else if fish == 0 {
				shoal[f] = 6
				shoal = append(shoal, 8)
			}
		}

		if debug {
			fmt.Println(shoal)
		}
	}
	fmt.Println(len(shoal))
}

func Solve06B(shoal []int) {

	fishMap := make(map[int]int)

	for _, fish := range shoal {
		fishMap[fish]++
	}
	if debug {
		fmt.Println("Initial state ", fishMap)
	}

	days := 256

	for day := 1; day <= days; day++ {
		newFish := 0
		for i := 0; i <= 8; i++ {
			if i == 0 {
				newFish = fishMap[0]
				fishMap[7] = fishMap[7] + newFish
			}
			fishMap[i] = fishMap[i+1]
		}
		fishMap[8] = newFish

		if debug {
			fmt.Println("After", day, "days", countFishies(fishMap))
			fmt.Println(fishMap)
		}
	}

	fmt.Println(fishMap)
	fmt.Println("After", days, "days", countFishies(fishMap))
}

func countFishies(fishies map[int]int) int {
	sum := 0
	for _, age := range fishies {
		sum += age
	}
	return sum
}
