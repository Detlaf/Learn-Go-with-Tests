package structsmethodsinterfaces

import (
	"testing"
)

type Shape interface {
	Area() float64
}

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t testing.TB, r Rectangle, want float64) {
		t.Helper()
		got := r.Perimeter()

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10, 3}
		checkPerimeter(t, rectangle, 26)
	})

}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Height: 12, Base: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
		}
	}

}
