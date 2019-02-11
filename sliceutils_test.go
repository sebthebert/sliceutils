package sliceutils_test

import (
	"testing"

	su "github.com/sebthebert/sliceutils"
)

var numbers []string

func TestMain(m *testing.M) {
	numbers = []string{"One", "Two", "Three", "Four"}
	m.Run()
}

func length_greater_2(s string) bool {
	if len(s) > 2 {
		return true
	}
	return false
}

func length_greater_10(s string) bool {
	if len(s) > 10 {
		return true
	}
	return false
}

func length_is_3(s string) bool {
	if len(s) == 3 {
		return true
	}
	return false
}

func TestAll(t *testing.T) {
	t.Run("'All' on 'length_greater_2'", func(t *testing.T) {
		result := su.All(numbers, length_greater_2)
		if result == false {
			t.Errorf("'All' with 'length_greater_2' should return true !")
		}
	})
	t.Run("'All' on 'length_is_3'", func(t *testing.T) {
		result := su.All(numbers, length_is_3)
		if result == true {
			t.Errorf("'All' with 'length_is_3 should return false !")
		}
	})
}

func TestAny(t *testing.T) {
	t.Run("'Any' with 'length_is_3'", func(t *testing.T) {
		result := su.Any(numbers, length_is_3)
		if result == false {
			t.Errorf("'Any' with 'length_is_3' should return true !")
		}
	})
	t.Run("'Any' with 'length_greater_10'", func(t *testing.T) {
		result := su.Any(numbers, length_greater_10)
		if result == true {
			t.Errorf("'Any' with 'length_greater_10' should return false !")
		}
	})
}
