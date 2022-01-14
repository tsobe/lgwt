package structs

import (
	"reflect"
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
	areaTests := []struct {
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{Width: 3.0, Height: 10.0}, hasArea: 30.0},
		{shape: Circle{Radius: 10.0}, hasArea: 314.1592653589793},
		{shape: Triangle{Height: 12, Base: 3}, hasArea: 36.0},
	}

	for _, tc := range areaTests {
		t.Run(reflect.TypeOf(tc.shape).Name(), func(t *testing.T) {
			got := tc.shape.Area()
			if got != tc.hasArea {
				t.Errorf("%#v got %.2f, expected %.2f", tc.shape, got, tc.hasArea)
			}
		})
	}
}
