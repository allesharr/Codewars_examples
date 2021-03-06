package tocamelcase

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

func Test() {
	ss := ToCamelCase("the-stealth-warrior")
	fmt.Println(ss)
	ss = ToCamelCase("The_Stealth_Warrior")
	fmt.Println(ss)
}

//Esribe un funtion que hablas los portas en el tipo (palabridadPalabridad 0 PalabridadPalabridad). Entre es los padabras divididos por un de symbolos -,_
func ToCamelCase(s string) string {
	reg, _ := regexp.MatchString(`\w\-\w`, s)
	returnString := ""
	var ss []string
	//Выявляем разделители
	if reg {
		ss = strings.Split(s, "-")
	} else {
		ss = strings.Split(s, "_")
	}

	//разделяем
	for index, element := range ss {
		if index == 0 {
			returnString = returnString + element
			continue
		}
		if len(element) > 0 {
			b, _ := []byte(element), 1
			b[0] = bytes.ToUpper(b)[0]
			element = string(b[0]) + element[1:]
			returnString = returnString + element
		}
	}

	return returnString
}
