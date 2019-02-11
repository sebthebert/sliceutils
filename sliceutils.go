package sliceutils

//
// All(strings []string, filter func(string) bool) => bool
//
// Returns true if all items from the slice match the filter
//
func All(strings []string, filter func(string) bool) bool {
	for _, str := range strings {
		if !filter(str) {
			return false
		}
	}
	return true
}

//
// Any(strings []string, filter func(string) bool) => bool
//
// Returns true if at least one item from the slice matches the filter
//
func Any(strings []string, filter func(string) bool) bool {
	for _, str := range strings {
		if filter(str) {
			return true
		}
	}
	return false
}
