package util

import "regexp"

func OutLine(text string) [][]int {
	reg, err := regexp.Compile(`\n`)
	if err != nil {
		return nil
	}
	indexs := reg.FindAllStringIndex(text, -1)
	return indexs
}
