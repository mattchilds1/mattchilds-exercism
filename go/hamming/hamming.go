package hamming

import "strings"
import "errors"

func Distance(a, b string) (int, error) {
	if len(a) == 0 || len(a) == 0 {
		return 0, nil
	}
	stringArrayA := splitString(a)
	stringArrayB := splitString(b)
	if len(stringArrayA) != len (stringArrayB) {
		return -1, errors.New("String length's different")
	}else {
		count := 0
		difference := 0
		stringlength := len(a)
		for count < stringlength {
			if stringArrayA[count] != stringArrayB[count] {
				difference = difference + 1

			}
			count = count + 1
		}
		return difference, nil
	}

}

func splitString(s string) []string {
	arrayofstrings := strings.Split(s,"")
	return arrayofstrings
}