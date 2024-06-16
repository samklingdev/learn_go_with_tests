package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	testTable := []struct {
		char     string
		n        int
		expected string
	}{
		{"a", 5, "aaaaa"},
		{"b", 3, "bbb"},
		{"c", 2, "cc"},
		{"d", 1, "d"},
	}

	for _, tt := range testTable {
		title := fmt.Sprintf("repeat %s %d", tt.char, tt.n)
		t.Run(title, func(t *testing.T) {
			got := Repeat(tt.char, tt.n)
			if got != tt.expected {
				t.Errorf("got %q want %q", got, tt.expected)
			}
		})
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func BenchmarkRepeatArr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatWithBytesArray("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func ExampleRepeatWithBytesArray() {
	repeated := RepeatWithBytesArray("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}
