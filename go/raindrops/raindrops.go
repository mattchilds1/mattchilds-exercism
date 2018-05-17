package raindrops

import "strconv"

func Convert(rain int) string {
	result := ""

	pling := rain % 3
	plang := rain % 5
	plong := rain % 7

	if pling == 0 {
		// append result pling
		result = result + "Pling"
	}
	if plang == 0 {
		// append result plang
		result = result + "Plang"
	}
	if plong == 0 {
		// append result plong
		result = result + "Plong"
	}

	if result == "" {
		result = strconv.Itoa(rain)
	}

	return result
}

