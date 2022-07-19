package squarecomparation

import "fmt"

//Have problems with tests
func Test() {
	// ss := uol.Unite("file1.txt", "file2.txt")
	// uol.ReturnToFile(ss)
	ss := Comp([]int{12, 144, 19, 161, 19, 144, 11}, []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19})
	fmt.Println(ss)
}
func Comp(array1 []int, array2 []int) bool {
	// var countPassed int = 0
	// var len int = len(array1)
	arr1Copy := array1

	for index, element := range array1 {
		square := element * element
		var isContainsElem bool = false
		for _, elem := range array2 {
			if Contains(arr1Copy, elem) {
				if elem == square {
					isContainsElem = true
					array1 = removeByIndex(array1, index)
					break
				}
				removeByValue(arr1Copy, elem)
				if !isContainsElem {
					return false
				}
			}

		}
	}
	return true
}

func removeByIndex(array []int, index int) []int {
	return append(array[:index], array[index+1:]...)
}

func removeByValue(array []int, value int) []int {

	for ind, elem := range array {
		if elem == value {
			return append(array[:ind], array[ind+1:]...)
		}
	}

	return array
}
func Contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
