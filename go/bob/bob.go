
package bob

import (
	"regexp"
	"bytes"
)


func removeWhiteSpace(remark string) string {
	data := []byte(remark)
	re := regexp.MustCompile("  +")
	replaced := re.ReplaceAll(bytes.TrimSpace(data), []byte(" "))
	returnedValue := string(replaced)
	return returnedValue
}

func Hey(remark string) string {

	result := "Whatever."
	processedString := removeWhiteSpace(remark)

	uppercaseFlag, _ := regexp.MatchString("^[A-Z \\d\\W]+$", processedString)
	questionFlag, _ := regexp.MatchString("[0-Z]*[?]$", processedString)
	nullFlag, _ := regexp.MatchString("^$",processedString)
	containsLettersFlag, _ :=regexp.MatchString("[a-zA-Z]", processedString)
	containsNumbersFlag, _ :=regexp.MatchString("[1-9].*", processedString)

	switch {
	case nullFlag == true :
		result = "Fine. Be that way!"
	case nullFlag == true :
		result = "Fine. Be that way!"
	case containsLettersFlag == false && containsNumbersFlag == true && questionFlag == true  :
		result = "Sure."
	case containsLettersFlag == false && questionFlag == true  :
		result = "Sure."
	case uppercaseFlag == true && containsLettersFlag == false :
		result = "Whatever."
	case questionFlag == true && uppercaseFlag == true :
		result = "Calm down, I know what I'm doing!"
	case questionFlag == true :
		result = "Sure."
	case uppercaseFlag == true :
		result = "Whoa, chill out!"
	default  :
		result = "Whatever."
	}



	return result

}

