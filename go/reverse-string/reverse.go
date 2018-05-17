package reverse


func String(s string) string {
	c := 0
	rune :=make([]rune, len(s))
	for _, r := range s {
			rune[c] = r
			c++
	}

	rune = rune [0:c]

	for i := 0; i < c/2; i++ {
		rune[i], rune[c-1-i] = rune[c-1-i], rune[i]
	}

	returnvalue := string(rune)

	return returnvalue
}
