package golang_united_school_homework

import (
	"errors"
	"fmt"
)

var (
	errorFullBox    = errors.New("box is full")
	errorOutOfRange = errors.New("index out of range")
	errorNilShape   = errors.New("shape by index doesn't exist")
	errorNoCircles  = errors.New("no circles in the box")
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) == b.shapesCapacity {
		return fmt.Errorf("could not add shape: %w", errorFullBox)
	}

	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, fmt.Errorf("could not get shape: %w", errorOutOfRange)
	}

	if b.shapes[i] == nil {
		return nil, fmt.Errorf("could not get shape: %w", errorNilShape)
	}

	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, fmt.Errorf("could not extract shape: %w", errorOutOfRange)
	}

	if b.shapes[i] == nil {
		return nil, fmt.Errorf("could not extract shape: %w", errorNilShape)
	}

	shape := b.shapes[i]
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)

	return shape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, fmt.Errorf("could not replace shape: %w", errorOutOfRange)
	}

	if b.shapes[i] == nil {
		return nil, fmt.Errorf("could not replace shape: %w", errorNilShape)
	}

	temp := b.shapes[i]
	b.shapes[i] = shape

	return temp, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64

	for _, shape := range b.shapes {
		sum += shape.CalcPerimeter()
	}

	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64

	for _, shape := range b.shapes {
		sum += shape.CalcArea()
	}

	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	contains := len(b.shapes)

	for i := 0; i < len(b.shapes); {
		if _, ok := b.shapes[i].(*Circle); ok {
			_, err := b.ExtractByIndex(i)
			if err != nil {
				return fmt.Errorf("could not remove circles: %w", err)
			}
		} else {
			i++
		}
	}

	if contains == len(b.shapes) {
		return fmt.Errorf("could not remove circles: %w", errorNoCircles)
	}

	return nil
}
