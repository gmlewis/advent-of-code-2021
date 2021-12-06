package enum

// Zip zips corresponding elements from slice of slices.
//
// The zipping finishes as soon as any slice ends.
func Zip[T any](lists [][]T) (ret [][]T) {
	n := len(lists)
	var i int
outer:
	for {
		t := make([]T, 0, n)
		for _, list := range lists {
			if i >= len(list) {
				break outer
			}
			t = append(t, list[i])
		}
		ret = append(ret, t)
		i++
	}

	return ret
}

// Zip2 zips corresponding elements from different types into a slice of structs.
//
// The zipping finishes as soon as either slice ends.
func Zip2[S any, T any, KV any](sList []S, tList []T, f func(S, T) KV) (ret []KV) {
	for i := 0; i < len(sList) && i < len(tList); i++ {
		ret = append(ret, f(sList[i], tList[i]))
	}

	return ret
}
