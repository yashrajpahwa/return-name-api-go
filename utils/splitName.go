package utils

import "strings"

func SplitName(name string) string {
	if name == "" {
		return ""
	}
	return "[" + strings.Join(strings.Split(name, ""), ",") + "]"
}
