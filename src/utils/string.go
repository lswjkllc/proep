package utils

import "strings"

func JoinStrings(strs ...string) string {
	bder := strings.Builder{}
	for _, str := range strs {
		bder.WriteString(str)
	}
	return bder.String()
}
