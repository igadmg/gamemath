package test

import (
	rm "github.com/igadmg/raylib-go/raymath"
	"github.com/igadmg/raylib-go/raymath/rect2"
	"github.com/igadmg/raylib-go/raymath/vector2"
	"github.com/igadmg/raylib-go/raymath/vector3"
	"github.com/stretchr/testify/assert"
)

func AssertVector2InDelta[T rm.SignedNumber](t assert.TestingT, expected, actual vector2.Vector[T], delta float64) {
	assert.InDelta(t, expected.X, actual.X, delta)
	assert.InDelta(t, expected.Y, actual.Y, delta)
}

func AssertVector3InDelta[T rm.SignedNumber](t assert.TestingT, expected, actual vector3.Vector[T], delta float64) {
	assert.InDelta(t, expected.X, actual.X, delta)
	assert.InDelta(t, expected.Y, actual.Y, delta)
	assert.InDelta(t, expected.Z, actual.Z, delta)
}

func AssertRectangle2InDelta[T rm.SignedNumber](t assert.TestingT, expected, actual rect2.Rectangle[T], delta float64) {
	AssertVector2InDelta(t, expected.Position, actual.Position, delta)
	AssertVector2InDelta(t, expected.Size, actual.Size, delta)
}
