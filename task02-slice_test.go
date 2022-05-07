package homework

import (
	"reflect"
	"testing"
)

func Test_reverse(t *testing.T) {
	data := []struct {
		In       []int64
		Expected []int64
	}{
		{In: []int64{0, 1, 2}, Expected: []int64{2, 1, 0}},
		{In: []int64{1, 2, 5, 15}, Expected: []int64{15, 5, 2, 1}},
	}
	for _, q := range data {
		got := reverse(q.In)

		if !reflect.DeepEqual(q.Expected, got) {
			t.Logf("reverse expected: %#v\n, got: %#v\n", q.Expected, got)
			t.Fail()
		}
	}
}
