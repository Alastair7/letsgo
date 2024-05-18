package iteration

import (
	"fmt"
	"testing"
)

const repeatTimes = 5;

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", repeatTimes)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", repeatTimes)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 3)
	fmt.Println(repeated)
	// Output: aaa
}