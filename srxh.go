package main

import "fmt"

// binsrxh finds the position of key in values
func binsrxh(key int, values []int) (position int, found bool) {

	if len(values) == 0 {
		return 0, false
	}

	if len(values) == 1 && values[0] == key {
		return 0, true
	}

	mid := len(values) / 2

	if key > values[mid] {
		// key not below midpoint
		position, found = binsrxh(key, values[mid+1:])
		position += mid + 1
		return
	}

	if key < values[mid] {
		// key not above midpoint
		return binsrxh(key, values[0:mid])
	}

	// the key might, or might not be at the midpoint
	return mid, values[mid] == key
}

func main() {
	stuff := []int{}
	fmt.Println(binsrxh(3, stuff))

	stuff = []int{0}
	fmt.Println(binsrxh(0, stuff))

	stuff = []int{0, 1}
	fmt.Println(binsrxh(1, stuff))

	stuff = []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(binsrxh(3, stuff))

	stuff = []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(binsrxh(0, stuff))

	stuff = []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(binsrxh(6, stuff))

	stuff = []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(binsrxh(7, stuff))

	stuff = []int{0, 1, 2, 4, 5, 6}
	fmt.Println(binsrxh(3, stuff))
}
