package main

import (
	uol "CodeWars/UnitedOfLists"
)

func main() {
	// tocamelcase.Test()
	// fmt.Println(spinwords.SpinWords("Welcome my friend"))
	ss := uol.Unite("file1.txt", "file2.txt")
	uol.ReturnToFile(ss)
}
