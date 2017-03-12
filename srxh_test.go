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

	{0, 4, true, []int{-9823, -983, -83, -3, 0, 83, 828, 7474}},
	{-9824, 0, false, []int{-9823, -983, -83, -3, 0, 83, 828, 7474}},
	{7475, 8, false, []int{-9823, -983, -83, -3, 0, 83, 828, 7474}},
}

func testOneLookup(t *testing.T, key int, values []int, expectedPos int, expectedFound bool) {

	actualPos, actualFound := binsrxh(key, values)

	if actualPos != expectedPos || actualFound != expectedFound {
		t.Errorf("Expected values (%d,%t) searching for %d in %s, but got (%d,%t) instead.",
			expectedPos, expectedFound, key, fmt.Sprint(values), actualPos, actualFound)
	}

	if actualFound && values[actualPos] != key {
		t.Errorf("Expected found value %d at position %d in %s, but see %d instead.",
			key, actualPos, fmt.Sprint(values), values[actualPos])
	}

}

// TestSrxh checks the specified results in srxhTests, including
// out of range scenarios.
func TestSrxh(t *testing.T) {
	for _, tv := range srxhTests {
		testOneLookup(t, tv.key, tv.values, tv.expectedPos, tv.expectedFound)
	}
}

// TestAllValues tests that every value can be correctly found
// in each of the test values slices.
func TestAllValues(t *testing.T) {
	for _, tv := range srxhTests {
		for expectedPos, val := range tv.values {
			testOneLookup(t, val, tv.values, expectedPos, true)
		}
	}
}
