package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	want := 40.0
	rect := Rectangle{10.0, 10.0}

	got := rect.Perimeter()

	if got != want {
		t.Errorf("Got %.2f, expected %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	assertArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("Got %.2f, expected %.2f", got, want)
		}
	}

	t.Run("Rectangle", func(t *testing.T) {
		want := 30.0
		rect := Rectangle{3.0, 10.0}

		assertArea(t, rect, want)
	})

	t.Run("Circle", func(t *testing.T) {
		want := 314.1592653589793
		circle := Circle{10.0}

		assertArea(t, circle, want)
	})
}
