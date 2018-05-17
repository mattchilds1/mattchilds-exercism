package isogram

import "strings"
import "sort"


func IsIsogram(s string) bool {
	isogram:= true
	cleanString := strings.ToLower(s)
	splitString := strings.Split(cleanString,"")
	sort.Strings(splitString)

	lastletter := ""
	count := 0
	splitstringlength := len(strings.Join(splitString,""))

	for count < splitstringlength {
	if isAlpha(splitString[count])  == true {
		if lastletter == splitString[count] {
			isogram = false
		}
	}
		lastletter = splitString[count]
		count = count +1
	}

	return isogram
}


func isAlpha(str string) bool {
	for i := range str {
		if str[i] < 'A' || str[i] > 'z' {
			return false
		} else if str[i] > 'Z' && str[i] < 'a' {
			return false
		}
	}
	return true
}