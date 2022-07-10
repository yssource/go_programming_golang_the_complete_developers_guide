/* -- Summary:
 * Write a generic function to complete the existing code.
 *
 * -- Requirements:
 * 1. Write a single function named clamp that can restrict a value to a
 * specific value
 * - the function should work with integers, floating point numbers and
 *   arbitrary type aliases.
 * - use existing clamp function signature and comments as starting point
 */
package generics

// TODO: finish this and write tests

type Distance int32
type Velocity float64

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

func Clamp[T Number](value, min, max T) T {
	if value < min {
		return min
	} else if value > max {
		return max
	} else {
		return value
	}
}
