package main

import (
	"fmt"
	"testing"
)

var srxhTests = []struct {
	key           int
	expectedPos   int
	expectedFound bool
	values        []int
}{
	{3, 0, false, []int{}},
	{7, 7, false, []int{0, 1, 2, 3, 4, 5, 6}},
	{-1, 0, false, []int{0, 1, 2, 3, 4, 5, 6}},
	{3, 3, false, []int{0, 1, 2, 4, 5, 6}},

	{0, 0, true, []int{0}},
	{1, 1, true, []int{0, 1}},
	{3, 3, true, []int{0, 1, 2, 3, 4, 5, 6}},
	{0, 0, true, []int{0, 1, 2, 3, 4, 5, 6}},
	{6, 6, true, []int{0, 1, 2, 3, 4, 5, 6}},
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
