package reverse_test

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func ExampleString() {
	fmt.Println(reverse.String("hello"))
	// Output: olleh
}
