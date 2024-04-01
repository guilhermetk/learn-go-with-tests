package arraysslices

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		expected := 15

		if got != expected {
			t.Errorf("got %d expected %d given %v", got, expected, numbers)
		}

	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	assertEquals(t, got, want)
}

func TestSumAllTail(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTail([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		assertEquals(t, got, want)
	})

	t.Run("safely sum emtpy slices", func(t *testing.T) {
		got := SumAllTail([]int{}, []int{})
		want := []int{0, 0}

		assertEquals(t, got, want)
	})
}

func assertEquals(t testing.TB, got, want []int) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
