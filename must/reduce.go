package must

// ReduceIntSlicesToInt reduces int slices using and int accumulator.
func ReduceIntSlicesToInt(slices [][]int, acc int, f func(slice []int, acc int) int) int {
	for _, slice := range slices {
		acc = f(slice, acc)
	}
	return acc
}
