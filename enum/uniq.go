package enum

// Uniq removes all duplicated elements.
func Uniq[T comparable](items []T) (ret []T) {
	seen := map[T]struct{}{}
	for _, item := range items {
		if _, ok := seen[item]; ok {
			continue
		}
		ret = append(ret, item)
		seen[item] = struct{}{}
	}

	return ret
}
