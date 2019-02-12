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

//
// First(strings []string, filter func(string) bool) => string
//
// Returns the first string matching the filter
//
func First(strings []string, filter func(string) bool) (string, bool) {
	for _, str := range strings {
		if filter(str) {
			return str, true
		}
	}
	return "", false
}

//
// Last(strings []string, filter func(string) bool) => string
//
// Returns the last string matching the filter
//
func Last(strings []string, filter func(string) bool) (string, bool) {
	var last string
	var match bool
	for _, str := range strings {
		if filter(str) {
			last = str
			match = true
		}
	}
	if match {
		return last, true
	}
	return "", false
}
