package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("x", 4)
	want := "xxxx"

	if got != want {
		t.Errorf("wanted %q but got %q", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}

func ExampleRepeat() {
	repeated := Repeat("g", 3)
	fmt.Println(repeated)
	// Output: ggg
}
