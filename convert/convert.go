package convert

import (
	"overUnderModel/fileio"
	"strconv"
	"strings"
)

func StringToInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func SliceIntToString(d []int) string {
	var slc []string

	for i := range d {
		slc = append(slc, fileio.IntToString(d[i]))
	}
	return strings.Join(slc, ",")
}
