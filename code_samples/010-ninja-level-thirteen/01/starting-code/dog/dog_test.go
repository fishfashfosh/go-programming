package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	got := Years(10)
	if got != 70 {
		t.Errorf("Years(10), want 70, got %d", got)
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	// Output:  70
}

func BenchmarkYears(b *testing.B) {

	for n := 0; n < b.N; n++ {
		Years(10)
	}
}

