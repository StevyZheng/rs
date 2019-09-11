package common

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"regexp"
	"strings"
)

func SplitString(src string, splitStr string) (strArr arraylist.List) {
	slice01 := strings.Split(src, splitStr)
	reg := regexp.MustCompile("\\s+")
	for i := range slice01 {
		reg.ReplaceAllString(slice01[i], "")
		if len(slice01[i]) != 0 {
			strArr.Add(slice01[i])
		}
	}
	return
}

func RegexStringLines(src string) (strArr arraylist.List) {
	return
}
