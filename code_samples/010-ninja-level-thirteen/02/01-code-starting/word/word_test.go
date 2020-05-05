package word

import (
	"testing"
)

func TestUseCount(t *testing.T) {
	mystring1 := "once upon a time"
	mymap1 := map[string]int {
		"once": 1,
		"upon": 1,
		"a": 1,
		"time": 1,
	}
	mystring2 := "one two two three three three four four four four"
	mymap2 := map[string]int {
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
	}

	got := UseCount(mystring1)
	if len(got) != len(mymap1) {
		t.Errorf("UseCount('Once upon a time')), want %v, got %v", mymap1, got)
	} else {
		for key, value := range mymap1 {
			i, ok := got[key]
			if !ok || i != value {
				t.Errorf("UseCount('Once upon a time')), want %v, got %v", mymap1, got)
			}
		}
	}
	got = UseCount(mystring2)
	if len(got) != len(mymap2) {
		t.Errorf("UseCount('Once upon a time')), want %v, got %v", mymap2, got)
	} else {
		for key, value := range mymap2 {
			i, ok := got[key]
			if !ok || i != value {
				t.Errorf("UseCount('Once upon a time')), want %v, got %v", mymap2, got)
			}
		}
	}
}

func TestCount(t *testing.T) {
	got := Count("one two three four")
	if got != 4 {
		t.Errorf("Count('one two three four'), want, %v, got %v", 4, got)
	}

}

