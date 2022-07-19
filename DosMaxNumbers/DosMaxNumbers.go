package dosmaxnumbers

func TwoOldestAges(ages []int) [2]int {
	var second_max, max int
	max = 0
	second_max = 0
	for _, element := range ages {
		if element > max {
			second_max = max
			max = element
		} else {
			if element > second_max && element <= max {
				second_max = element
			}
		}
	}
	var arr [2]int
	arr[0] = second_max
	arr[1] = max
	return arr
}
