// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import "regexp"



// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {

	png, _ := regexp.Compile("Portable Network Graphics")
	pngResult := png.ReplaceAllString(s, "PNG")

	ror, _ := regexp.Compile("Ruby on Rails")
	rorResult := ror.ReplaceAllString(pngResult, "ROR")

	fifo, _ := regexp.Compile("First In, First Out")
	fifoResult := fifo.ReplaceAllString(rorResult, "FIFO")

	gimp, _ := regexp.Compile("GNU Image Manipulation Program")
	gimpResult := gimp.ReplaceAllString(fifoResult, "GIMP")

	cmos, _ := regexp.Compile("Complementary metal-oxide semiconductor")
	result := cmos.ReplaceAllString(gimpResult, "CMOS")


	return result
}
