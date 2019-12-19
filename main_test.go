package shapes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBoxes(t *testing.T) {
	box1 := NewBox(0, 0, 5, 5)
	box2 := NewBox(3, 3, 5, 5)
	box3 := NewBox(6, 6, 5, 5)

	res, err := box1.Touches(box2)
	require.NoError(t, err)
	require.True(t, res)

	res, err = box2.Touches(box3)
	require.NoError(t, err)
	require.True(t, res)

	res, err = box1.Touches(box3)
	require.NoError(t, err)
	require.False(t, res)
}

func TestCircles(t *testing.T) {
	circle1 := NewCircle(0, 0, 5)
	circle2 := NewCircle(0, 3, 5)
	circle3 := NewCircle(0, 10, 5)

	res, err := circle1.Touches(circle2)
	require.NoError(t, err)
	require.True(t, res)

	res, err = circle2.Touches(circle3)
	require.NoError(t, err)
	require.True(t, res)

	res, err = circle1.Touches(circle3)
	require.NoError(t, err)
	require.False(t, res)
}

func TestCircleAndBox(t *testing.T) {
	circle1 := NewCircle(0, 0, 5)
	box1 := NewBox(0, 0, 5, 5)
	box2 := NewBox(0, 10, 5, 5)

	res, err := circle1.Touches(box1)
	require.NoError(t, err)
	require.True(t, res)

	res, err = circle1.Touches(box2)
	require.NoError(t, err)
	require.False(t, res)
}