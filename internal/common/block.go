// Package common provides structs that will be used across the
// different packages.
package common

// Block represents a census block group in California.
type Block struct {
	// Value represents the median house value in the block.
	Value float32

	// Income represents the median income of households in the block, in units of $10,000.
	Income float32

	// Age indicates the median age of houses in the block.
	Age float32

	// Rooms is the total number of rooms in the block.
	Rooms float32

	// Bedrooms is the total number of bedrooms in the block.
	Bedrooms float32

	// Population is the total number of people residing in the block.
	Population float32

	// Households represents the total number of households in the block (?).
	Households float32
}
