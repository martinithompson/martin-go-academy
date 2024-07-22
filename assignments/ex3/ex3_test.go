package main

import "testing"

func TestIsBetweenLimits(t *testing.T) {
	checkValue := func(t testing.TB, value int, limit Limit, inLimits bool) {
		got := IsBetweenLimits(limit, value)

		if got != inLimits {
			t.Errorf("got %v want %v, for value %v in limits %v", got, inLimits, value, limit)
		}
	}
	t.Run("returns true for a number in ranges", func(t *testing.T) {
		value := 3
		limit := Limit{Min: 1, Max: 10}
		checkValue(t, value, limit, true)
	})
	t.Run("returns false for a number outside of the limits", func(t *testing.T) {
		value := -3
		limit := Limit{Min: 1, Max: 10}
		checkValue(t, value, limit, false)
	})
}
