// Package common provides structs that will be used across the
// different packages.
package common

// Block represents a census block group.
type Block struct {
	// Value represents the median house value in the block.
	Value float64

	// Income represents the median income of households in the block, in units of $10,000.
	Income float64

	// Age indicates the median age of houses in the block.
	Age float64

	// Rooms is the total number of rooms in the block.
	Rooms float64

	// Bedrooms is the total number of bedrooms in the block.
	Bedrooms float64

	// Population is the total number of people residing in the block.
	Population float64

	// Households represents the total number of households in the block (?).
	Households float64
}
