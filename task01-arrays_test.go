package homework

import (
	"math"
	"testing"
)

const deltaTesting = 0.001

func equalFloat32DeltaFloat64(a, b float32, delta float64) bool {
	return math.Abs(float64(b-a)) < delta
}

func Test_average(t *testing.T) {
	data := []struct {
		In       [15]float32
		Expected float32
	}{
		{In: [15]float32{0}, Expected: 0},
		{In: [15]float32{1, 2, 3, 4, 5, 6}, Expected: 3.5},
		{In: [15]float32{150}, Expected: 10},
	}
	for _, q := range data {
		got := average(q.In)
		if !equalFloat32DeltaFloat64(got, q.Expected, deltaTesting) {
			t.Logf("average expected: %f, got: %f", q.Expected, got)
			t.Fail()
		}
	}
}
