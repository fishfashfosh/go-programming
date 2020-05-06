package mymath

import (
	"testing"
	"fmt"
)

func TestCenteredAvg(t *testing.T) {
	xi := []int {
		1,
		2,
		10,
		12,
	}

	got := CenteredAvg(xi)
	if got != 6.0 {
		t.Errorf("CenteredAvg of 1, 2, 10, 12 wanted 6.0, got %f", got)
	}
}

func ExampleCenteredAvg() {
	xi := []int {
		1,
		2,
		10,
		12,
	}

	fmt.Println(CenteredAvg(xi))
	// Output: 6
}

func BenchmarkCenteredAvg(t *testing.B) {
	xi := []int {
		1,
		2,
		10,
		12,
	}

	for i := 0; i < t.N; i++ {
		CenteredAvg(xi)
	}
}
