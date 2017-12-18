package greeting

import "testing"

// Define a function named HelloWorld that takes no arguments,
// and returns a string.
// In other words, define a function with the following signature:
// HelloWorld() string

func TestHelloWorld(t *testing.T) {
	for _, testCase := range addCases {
		actual := HelloWorld(testCase.in)
		if actual != testCase.expected {
			t.Errorf("HelloWorld(%s): expected %s, actual %s", testCase.in, testCase.expected, actual)
		}
	}
}


// func BenchmarkHelloWorld(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		HelloWorld(testCase.in)
// 	}
// }
