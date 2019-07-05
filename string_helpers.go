package main

import (
	"strings"
)

func parseStringOf(line, key string) string {
	joiner := []string{"|", key, " = "}
	return strings.TrimPrefix(line, strings.Join(joiner, ""))
}

func parseBooleanOf(line, key string) bool {
	joiner := []string{"|", key, " = "}
	return strings.Compare(strings.TrimPrefix(line, strings.Join(joiner, "")), "Yes") == 0
}

//func parseIntOf(line, key string) int {
	//joiner := []string{"|", key, " = "}
	//return -1
//}
