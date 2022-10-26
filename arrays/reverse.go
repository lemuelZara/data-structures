package arrays

import (
	"fmt"
	"strings"
)

func Reverse(s string) string {
	var arrName = strings.Split(s, "")
	var reversed strings.Builder

	for i := len(arrName) - 1; i >= 0; i-- {
		fmt.Fprintf(&reversed, arrName[i])
	}

	return reversed.String()
}
