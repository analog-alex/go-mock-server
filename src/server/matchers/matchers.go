package matchers

import (
	"reflect"
)

// EqualMaps same, but with URL interface for maps
func EqualMaps(m1 map[string][]string, m2 map[string][]string) bool {
	return reflect.DeepEqual(m1, m2)
}

// MatchMaps check for sub sets matching i.e. if m1 is a superset of m2
func MatchMaps(m1 map[string][]string, m2 map[string][]string) bool {
	for key, value := range m1 {
		arrValue := m2[key]
		if arrValue != nil && !reflect.DeepEqual(value, arrValue) {
			return false
		}
	}
	return true
}
