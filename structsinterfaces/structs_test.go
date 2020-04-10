package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %g, but want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		// got := rectangle.Perimeter()
		want := 40.0
		checkPerimeter(t, rectangle, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		// got := circle.Perimeter()
		want := 62.83185307179586
		checkPerimeter(t, circle, want)
	})

}

// func TestArea(t *testing.T) {
// 	checkArea := func(t *testing.T, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %g, but want %g", got, want)
// 		}
// 	}
// 	t.Run("rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{10.0, 10.0}
// 		got := rectangle.Area()
// 		want := 100.0
// 		checkArea(t, rectangle, want)
// 	})

// 	t.Run("circles", func(t *testing.T) {
// 		circle := Circle{10.0}
// 		got := circle.Area()
// 		want := 314.1592653589793
// 		checkArea(t, circle, want)
// 	})
// }

//Table based tests
func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{width: 10, height: 10}, hasArea: 100},
		{shape: Circle{radius: 10}, hasArea: 314.1592653589793},
		{shape: Triangle{base: 10, height: 2}, hasArea: 10},
	}
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g, want %g", tt.shape, got, tt.hasArea)
			}
		})

	}
}
