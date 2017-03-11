package main

import (
	"fmt"
	"testing"
)

var srxhTests = []struct {
	key           int
	values        []int
	expectedPos   int
	expectedFound bool
}{
	{3, []int{}, 0, false},
	{7, []int{0, 1, 2, 3, 4, 5, 6}, 7, false},
	{-1, []int{0, 1, 2, 3, 4, 5, 6}, 0, false},
	{3, []int{0, 1, 2, 4, 5, 6}, 3, false},

	{0, []int{0}, 0, true},
	{1, []int{0, 1}, 1, true},
	{3, []int{0, 1, 2, 3, 4, 5, 6}, 3, true},
	{0, []int{0, 1, 2, 3, 4, 5, 6}, 0, true},
	{6, []int{0, 1, 2, 3, 4, 5, 6}, 6, true},
}

func TestSrxh(t *testing.T) {
	for _, tv := range srxhTests {
		actualPos, actualFound := binsrxh(tv.key, tv.values)
		if actualPos != tv.expectedPos || actualFound != tv.expectedFound {
			t.Errorf("Expected values (%d,%t) searching for %d in %s, but got (%d,%t) instead.",
				tv.expectedPos, tv.expectedFound, tv.key, fmt.Sprint(tv.values), actualPos, actualFound)
		}
	}
}
