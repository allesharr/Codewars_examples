package spinwords

import (
	"strings"
)

//Si longitude de la linea es 5 o mas de 5 symbolos -> tienes que volver la linea
func SpinWords(str string) string {
	ss := strings.Split(str, " ")
	var return_string string = ""
	for _, element := range ss {
		if len(element) > 4 {
			for i := 0; i < len(element); i++ {
				b := element[len(element)-i-1]
				return_string = return_string + string(b)
			}
			return_string = return_string + " "
		} else {
			return_string = return_string + element + " "
		}
	}
	return return_string
}
