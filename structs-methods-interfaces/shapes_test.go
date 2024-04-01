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
		got := Rectangle{Width: 10.0, Height: 10.0}
		want := 100.0
		checkArea(t, got, want)
	})
	t.Run("test the area of a circle", func(t *testing.T) {
		got := Circle{Radious: 10.0}
		want := 100.0
		checkArea(t, got, want)
	})
}

func TestAreaTableDriven(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{Radious: 10}, want: 100.0},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		assertEquals(t, got, tt.want)
	}
}

func assertEquals(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
