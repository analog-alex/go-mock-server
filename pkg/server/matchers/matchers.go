package matchers

import (
	"reflect"
)

// AreMapsEqual same, but with URL interface for maps
func AreMapsEqual(m1 map[string][]string, m2 map[string][]string) bool {
	return reflect.DeepEqual(m1, m2)
}

// IsMapASubset check for sub sets matching i.e. if m1 is a superset of m2
func IsMapASubset(m1 map[string][]string, m2 map[string][]string) bool {
	for key, value := range m2 {
		if !reflect.DeepEqual(value, m1[key]) {
			return false
		}
	}
	return true
}
