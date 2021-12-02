package must

// ReduceIntSlicesToInt reduces int slices using an int accumulator.
func ReduceIntSlicesToInt(slices [][]int, acc int, f func(slice []int, acc int) int) int {
	for _, slice := range slices {
		acc = f(slice, acc)
	}
	return acc
}

// ReduceStringSlicesToIntSlice reduces string slices using an int slice accumulator.
func ReduceStringSlicesToIntSlice(lines []string, acc []int, f func(line string, acc []int) []int) []int {
	for _, line := range lines {
		acc = f(line, acc)
	}
	return acc
}
