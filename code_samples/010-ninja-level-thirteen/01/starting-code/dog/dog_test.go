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

func TestYearsTwo(t *testing.T) {
	got := YearsTwo(0)
	if got != 0 {
		t.Errorf("TearsTwo(0), want 0, got %d", got)
	}

	got = YearsTwo(10)
	if got != 70 {
		t.Errorf("TearsTwo(10), want 70, got %d", got)
	}
}


func ExampleYears() {
	fmt.Println(Years(10))
	// Output:  70
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	// Output:  70
}

func BenchmarkYears(b *testing.B) {

	for n := 0; n < b.N; n++ {
		Years(100)
	}
}

func BenchmarkYearsTwo(b *testing.B) {

	for n := 0; n < b.N; n++ {
		YearsTwo(100)
	}
}
