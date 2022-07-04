package unitedoflists

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Unite(filepath1 string, filepath2 string) []string {
	var sa []string
	sa = append(sa, " ")

	file1, err := os.Open(filepath1)
	defer file1.Close()
	if err != nil {
		log.Fatal("Caanot open file with file path " + filepath1)
	}
	file2, err := os.Open(filepath2)
	defer file2.Close()
	if err != nil {
		log.Fatal("Caanot open file with file path " + filepath2)
	}

	scanner1 := bufio.NewScanner(file1)
	scanner2 := bufio.NewScanner(file2)
	for scanner1.Scan() {
		var ip1 string = scanner1.Text()
		ip1 = ip1 + "\n"

		for index, element := range sa {
			if index == len(sa)-1 {
				sa = append(sa, ip1)
			}
			if strings.Contains(element, ip1) {
				break
			}

		}

	}
	for scanner2.Scan() {

		var ip2 string = scanner1.Text()
		ip2 = ip2 + "\n"
		for index, element := range sa {
			if index == len(sa)-1 {
				sa = append(sa, ip2)
			}
			if strings.Contains(element, ip2) {
				break
			}

		}
	}
	sa = sa[1:]
	return sa
}

func ReturnToFile(ss []string) {
	file, err := os.Create("result_File.txt")
	defer file.Close()

	if err != nil {
		log.Fatal("Cannot create file")
	}

	var count int = 0
	for _, element := range ss {
		count++
		file.WriteString(element + "\n")

	}
	fmt.Println(count)
}
