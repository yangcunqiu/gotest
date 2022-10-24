package word

import (
	"fmt"
)

//
//func TestIsEven(t *testing.T) {
//	tests := []struct {
//		input  int
//		expect bool
//	}{
//		{0, true},
//		{1, false},
//		{2, true},
//	}
//
//	for _, test := range tests {
//		if actual := isEven(test.input); actual != test.expect {
//			t.Errorf("isEven(%v) = %v, expect = %v", test.input, actual, test.expect)
//		}
//	}
//}
//
//func BenchmarkIsEven(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		isEven(1)
//	}
//}

func ExampleIsEven() {
	fmt.Println(isEven(0))
	fmt.Println(isEven(1))
	// Output:
	// false
	// true
}
