package main

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	t.Run("Suming all arrays", func(t *testing.T) {
		num1 := []int{1, 2, 3, 4, 5, 6}
		num2 := []int{0, 0, 0, 0, 0, 0}
		got := SumAll(num1, num2)
		want := []int{21, 0}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
func TestSum(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want int, numbers []int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d, but wanted %d, for numbers %v", got, want, numbers)
		}
	}

	t.Run("collection of 6 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		// n := len(numbers)
		got := Sum(numbers)
		want := 21

		assertCorrectMessage(t, got, want, numbers)
	})

	t.Run("collection of N numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 9}
		// n := len(numbers)
		got := Sum(numbers)
		want := 30

		assertCorrectMessage(t, got, want, numbers)
	})

}
