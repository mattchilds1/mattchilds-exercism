package scrabble

import (
	"regexp"
	"strings"
)

func removeWhiteSpace(result string) string {
	processedResult := strings.Replace(result, " ", "", -1)
	return processedResult
}

func Score(word string) int {
	score := 0

	onepointtotal := calculateScore(word, "[A|a|E|e|I|i|O|o|U|u|L|l|N|n|R|r|S|s|T|t|]", 1)
	twopointtotal := calculateScore(word, "[D|d|G|g]", 2)
	threepointtotal := calculateScore(word, "[B|b|C|c|M|m|P|p]", 3)
	fourpointtotal := calculateScore(word, "[F|f|H|h|V|v|W|w|Y|y]", 4)
	fivepointtotal := calculateScore(word, "[k|K]", 5)
	eightpointtotal := calculateScore(word, "[J|j|X|x]", 8)
	tenpointtotal := calculateScore(word, "[Q|q|Z|z]", 10)

	score = onepointtotal + twopointtotal + threepointtotal + fourpointtotal + fivepointtotal + eightpointtotal + tenpointtotal
	return score
}


func calculateScore(word string, pattern string, pointvalue int) int {
	score := 0
	pointvalueregex, _ := regexp.Compile(pattern)
	countpointvalue:= pointvalueregex.FindAllString(word, -1)
	count := 0
	wordmatchlength := len(strings.Join(countpointvalue,""))

	for count < wordmatchlength {
		score = score + pointvalue
		count = count + 1
	}
	return score
}

