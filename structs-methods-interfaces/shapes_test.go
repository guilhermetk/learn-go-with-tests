package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	assertEquals(t, got, want)
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		assertEquals(t, got, want)
	}

	t.Run("test the area of a rectangle", func(t *testing.T) {
		got := Rectangle{10.0, 10.0}
		want := 100.0
		checkArea(t, got, want)
	})
	t.Run("test the area of a circle", func(t *testing.T) {
		got := Circle{10.0}
		want := 100.0
		checkArea(t, got, want)
	})
}

func assertEquals(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
