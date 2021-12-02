package must

import (
	"log"
	"strconv"
)

// Atoi converts a string to an int.
func Atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

// StringToInt takes a slice of strings and converts them to a
// slice of ints. It crashes if there are any conversion errors.
func StringToInt(lines []string) []int {
	result := make([]int, 0, len(lines))
	for _, line := range lines {
		v, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, v)
	}

	return result
}

// StringToInt64 takes a slice of strings and converts them to a
// slice of int64s. It crashes if there are any conversion errors.
func StringToInt64(lines []string) []int64 {
	result := make([]int64, 0, len(lines))
	for _, line := range lines {
		v, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, v)
	}

	return result
}
