package matchers

import (
	"io.analog.alex.mockserver/pkg/server/matchers"
	"testing"
)

func TestEqualMaps(t *testing.T) {

	testSubjectOne := make(map[string][]string)
	testSubjectOne["one"] = []string{"first"}
	testSubjectOne["two"] = []string{"second"}

	testSubjectTwo := make(map[string][]string)
	testSubjectTwo["one"] = []string{"first"}
	testSubjectTwo["two"] = []string{"second"}

	result := matchers.AreMapsEqual(testSubjectOne, testSubjectTwo)

	if !result {
		t.Fatalf("Maps should equal one another.")
	}
}

func TestUnequalMaps(t *testing.T) {

	testSubjectOne := make(map[string][]string)
	testSubjectOne["one"] = []string{"first"}
	testSubjectOne["two"] = []string{"second"}

	testSubjectTwo := make(map[string][]string)
	testSubjectTwo["one"] = []string{"first"}
	testSubjectTwo["three"] = []string{"third"}

	result := matchers.AreMapsEqual(testSubjectOne, testSubjectTwo)

	if result {
		t.Fatalf("Maps should not equal one another.")
	}
}

func TestMatchingMaps(t *testing.T) {

	testSubjectOne := make(map[string][]string)
	testSubjectOne["one"] = []string{"first"}
	testSubjectOne["two"] = []string{"second"}

	testSubjectTwo := make(map[string][]string)
	testSubjectTwo["one"] = []string{"first"}

	result := matchers.IsMapASubset(testSubjectOne, testSubjectTwo)

	if !result {
		t.Fatalf("Maps should match one another.")
	}
}

func TestNonMatchingMaps(t *testing.T) {

	testSubjectOne := make(map[string][]string)
	testSubjectOne["one"] = []string{"first"}
	testSubjectOne["two"] = []string{"second"}

	testSubjectTwo := make(map[string][]string)
	testSubjectTwo["three"] = []string{"third"}

	result := matchers.IsMapASubset(testSubjectOne, testSubjectTwo)

	if result {
		t.Fatalf("Maps should not match one another.")
	}
}
