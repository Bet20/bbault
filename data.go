package main

import (
	"strings"
)

func higianize(s string) string {
	return strings.Trim(strings.Trim(s, "\n"), "\r")
}

func appendPaths(s []string) string {
	var total string
	for i, o := range s {
		if i != 0 {
			total += "/" + o
		} else {
			total += o
		}
	}
	return total
}
