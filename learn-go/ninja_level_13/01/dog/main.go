// Package dog allows us to more fully understand the 'dog' species.
package dog

// Years will convert human years to dog years immediately
func Years(n int) int {
	return n * 7
}

// YearsAlt will convert human years to dog years by first creating a new variable
func YearsAlt(n int) int {
	years := n * 7
	return years
}
