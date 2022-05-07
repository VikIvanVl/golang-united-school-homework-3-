package homework

import (
	"reflect"
	"testing"
)

func Test_sortMapValues(t *testing.T) {
	data := []struct {
		In       map[int]string
		Expected []string
	}{
		{In: map[int]string{2: "a", 0: "b", 1: "c"}, Expected: []string{"b", "c", "a"}},
		{In: map[int]string{10: "aa", 0: "bb", 500: "cc"}, Expected: []string{"bb", "aa", "cc"}},
	}
	for _, q := range data {
		got := sortMapValues(q.In)

		if !reflect.DeepEqual(q.Expected, got) {
			t.Logf("sortMapValues expected: %#v\n, got: %#v\n", q.Expected, got)
			t.Fail()
		}
	}
}
