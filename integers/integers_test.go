package integers

import (
	"fmt"
	"testing"
)

func TestIntegers(t *testing.T) {
	sum := Add(22, 8)
	expected := 30

	if sum != expected {
		t.Errorf("Expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleIntegers(t *testing.T) {
	sum := Add(22, 8)
	fmt.Println(sum)
	// Output:30
}
