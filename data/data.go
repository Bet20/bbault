package data

import (
	"io/ioutil"
	"os"
	"strings"
)

func Higienize(s string) string {
	return strings.Trim(strings.Trim(s, "\n"), "\r")
}

func AppendPaths(s []string) string {
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

func AcessStorageFile(path string) []byte {
	jsonfl, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer jsonfl.Close()

	byteData, _ := ioutil.ReadAll(jsonfl)

	return byteData
}
