package shapes

import (
	"fmt"
	"math"
)

type Shape interface {
	Touches(Shape) (bool, error)
}

// NewBox creates a new Box Shape
func NewBox(x, y, height, width int) Shape {
	return &Box{
		x: x,
		y: y,
		h: height,
		w: width,
	}
}

type Box struct {
	w int
	h int
	x int
	y int
}

func (box *Box) Touches(shape Shape) (bool, error) {
	switch shape.(type) {
	case *Box:
		return BoxTouchesBox(box, shape.(*Box)), nil
	case *Circle:
		return CircleTouchesBox(shape.(*Circle), box), nil
	default:
		return false, fmt.Errorf("Unknown Shape")
	}

	return false, nil
}

// NewCircle creates a new Circle Shape
func NewCircle(x, y, radius int) Shape {
	return &Circle{
		x: x,
		y: y,
		r: radius,
	}
}

type Circle struct {
	x int
	y int
	r int
}

func (circle *Circle) Touches(shape Shape) (bool, error) {
	switch shape.(type) {
	case *Box:
		return CircleTouchesBox(circle, shape.(*Box)), nil
	case *Circle:
		return CircleTouchesCircle(shape.(*Circle), circle), nil
	default:
		return false, fmt.Errorf("Unknown Shape")
	}

	return false, nil
}

func BoxTouchesBox(box1 *Box, box2 *Box) bool {
	if box1 == nil || box2 == nil || (box2.h == 0 && box2.w == 0) {
		return false
	}

	c1 := box1.y > box2.y+box2.h // e is above e2
	c2 := box2.y > box1.y+box1.h // e2 is above e
	c3 := box2.x > box1.x+box1.w // e2 is to the right of e
	c4 := box1.x > box2.x+box2.w // e is to the right of e2

	return !c1 && !c2 && !c3 && !c4
}

func CircleTouchesCircle(circle1 *Circle, circle2 *Circle) bool {
	if circle1 == nil || circle2 == nil || circle2.r == 0 {
		return false
	}

	dx := circle1.x - circle2.x
	dy := circle1.y - circle2.y
	distance := math.Sqrt(float64(dx*dx + dy*dy))

	return distance < float64(circle1.r+circle2.r)
}

func CircleTouchesBox(circle *Circle, box *Box) bool {
	if circle == nil || box == nil {
		return false
	}

	// temporary variables to set edges for testing
	testX := circle.x
	testY := circle.y

	// which edge is closest?
	if circle.x < box.x {
		testX = box.x // test left edge
	} else if circle.x > box.x+box.w {
		testX = box.x + box.w // right edge
	}

	if circle.y < box.y {
		testY = box.y // top edge
	} else if circle.y > box.y+box.h {
		testY = box.y + box.h // bottom edge
	}

	// get distance from closest edges
	distX := circle.x - testX
	distY := circle.y - testY
	distance := math.Sqrt(float64((distX * distX) + (distY * distY)))

	// if the distance is less than the radius, collision!
	return distance <= float64(circle.r)
}
