// reverse обращает порядок чисел "на месте"
package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5}

	n := 0
	fmt.Scanf("%d", &n)
	shiftToRight(a[:], n)
	fmt.Printf("Shift to right to %d elements: %v\n", n, a)
	shiftToLeft(a[:], n)
	fmt.Printf("Shift to left to %d elements: %v\n", n, a)
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func shiftToLeft(a []int, k int) {
	k = k % len(a)
	reverse(a[:k])
	reverse(a[k:])
	reverse(a)
}

func shiftToRight(a []int, k int) {
	k = k % len(a)
	reverse(a)
	reverse(a[:k])
	reverse(a[k:])
}
