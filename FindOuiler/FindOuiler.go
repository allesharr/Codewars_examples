package findouiler

//You are given an array (which will have a length of at least 3, but could be very large)\
// containing integers. The array is either entirely comprised of odd integers or entirely \
//comprised of even integers except for a single integer N.\
// Write a method that takes the array as an argument and returns this "outlier" N.

// func FindOutlier(integers []int) int {
// 	for _, element := range integers {
// 		if integers[0]%2 == 0 && integers[1]%2 == 0 {
// 			if element%2 == 1 {
// 				return element
// 			}
// 		} else {
// 			if element%2 == 0 {
// 				return element
// 			}
// 		}
// 		if (integers[0]%2 == 1 && integers[1]%2 == 0) || (integers[1]%2 == 1 && integers[0]%2 == 0) {
// 			switch integers[2] % 2 {
// 			case 0:
// 				if integers[1]%2 == 0 {
// 					return integers[0]
// 				} else {
// 					return integers[1]
// 				}
// 			case 1:
// 				if integers[1]%2 == 1 {
// 					return integers[0]
// 				} else {
// 					return integers[1]
// 				}
// 			}
// 		}
// 	}
// 	return 0
// }

/*
[2, 4, 0, 100, 4, 11, 2602, 36]
Should return: 11 (the only odd number)

[160, 3, 1719, 19, 11, 13, -21]
Should return: 160 (the only even number)
*/

func FindOutlier(integers []int) int {
	coping := make([]int, len(integers))
	copy(coping, integers)
	var DosZeroCount int = 0
	var DosZeroNotCount int = 0
	for index, _ := range coping {
		coping[index] = coping[index] % 2
		if coping[index] == 1 || coping[index] == -1 {
			DosZeroNotCount++
		} else {
			DosZeroCount++
		}
	}

	if DosZeroCount != 1 {
		for i, _ := range coping {
			if coping[i] == 1 || coping[i] == -1 {
				return integers[i]
			}
		}
	} else {
		for i, _ := range coping {
			if coping[i] == 0 {
				return integers[i]
			}
		}

	}
	// fmt.Println(strconv.Itoa(DosZeroCount) + "  " + strconv.Itoa(DosZeroNotCount))
	// var flag bool = false
	return 0
}
