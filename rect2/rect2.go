package rect2

import (
	"fmt"

	rm "github.com/igadmg/raylib-go/raymath"
	"github.com/igadmg/raylib-go/raymath/vector2"
)

type Rectangle[T rm.SignedNumber] struct {
	Position vector2.Vector[T]
	Size     vector2.Vector[T]
}

type (
	Float64 = Rectangle[float64]
	Float32 = Rectangle[float32]
	Int     = Rectangle[int]
	Int64   = Rectangle[int64]
	Int32   = Rectangle[int32]
	Int16   = Rectangle[int16]
	Int8    = Rectangle[int8]
)

func New[T rm.SignedNumber](position vector2.Vector[T], size vector2.Vector[T]) Rectangle[T] {
	return Rectangle[T]{
		Position: position,
		Size:     size,
	}
}

func NewSize[T rm.SignedNumber](size vector2.Vector[T]) Rectangle[T] {
	return Rectangle[T]{
		Position: vector2.Zero[T](),
		Size:     size,
	}
}

func NewXYWH[T rm.SignedNumber](x, y, w, h T) Rectangle[T] {
	return Rectangle[T]{
		Position: vector2.New[T](x, y),
		Size:     vector2.New[T](w, h),
	}
}

func NewT[T rm.SignedNumber, PT, ST rm.SignedNumber](position vector2.Vector[PT], size vector2.Vector[ST]) Rectangle[T] {
	return Rectangle[T]{
		Position: vector2.To[T](position),
		Size:     vector2.To[T](size),
	}
}

func NewFloat64[PT, ST rm.SignedNumber](position vector2.Vector[PT], size vector2.Vector[ST]) Float64 {
	return NewT[float64](position, size)
}

func NewFloat32[PT, ST rm.SignedNumber](position vector2.Vector[PT], size vector2.Vector[ST]) Float32 {
	return NewT[float32](position, size)
}

func NewInt[PT, ST rm.SignedNumber](position vector2.Vector[PT], size vector2.Vector[ST]) Int {
	return NewT[int](position, size)
}

func NewInt64[PT, ST rm.SignedNumber](position vector2.Vector[PT], size vector2.Vector[ST]) Int64 {
	return NewT[int64](position, size)
}

func NewInt32[PT, ST rm.SignedNumber](position vector2.Vector[PT], size vector2.Vector[ST]) Int32 {
	return NewT[int32](position, size)
}

func NewInt16[PT, ST rm.SignedNumber](position vector2.Vector[PT], size vector2.Vector[ST]) Int16 {
	return NewT[int16](position, size)
}

func NewInt8[PT, ST rm.SignedNumber](position vector2.Vector[PT], size vector2.Vector[ST]) Int8 {
	return NewT[int8](position, size)
}

func Zero[T rm.SignedNumber]() Rectangle[T] {
	return Rectangle[T]{
		Position: vector2.Zero[T](),
		Size:     vector2.Zero[T](),
	}
}

func One[T rm.SignedNumber]() Rectangle[T] {
	return Rectangle[T]{
		Position: vector2.Zero[T](),
		Size:     vector2.One[T](),
	}
}

func (r Rectangle[T]) String() string {
	return fmt.Sprintf("Position: %v; Size: %v;", r.Position, r.Size)
}

func (r Rectangle[T]) A() vector2.Vector[T] {
	return r.Position
}

func (r Rectangle[T]) SetA(a vector2.Vector[T]) Rectangle[T] {
	dxy := a.Sub(r.Position)
	return Rectangle[T]{
		Position: a,
		Size:     r.Size.Sub(dxy),
	}
}

func (r Rectangle[T]) B() vector2.Vector[T] {
	return r.Position.Add(r.Size)
}

func (r Rectangle[T]) SetB(b vector2.Vector[T]) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position,
		Size:     b,
	}
}

func (r Rectangle[T]) HorizontalLine(y T) (vector2.Vector[T], vector2.Vector[T]) {
	return vector2.New(r.A().X, y), vector2.New(r.B().X, y)
}

func (r Rectangle[T]) VerticalLine(x T) (vector2.Vector[T], vector2.Vector[T]) {
	return vector2.New(x, r.A().Y), vector2.New(x, r.B().Y)
}

func (r Rectangle[T]) Center() vector2.Vector[T] {
	return r.Position.Add(r.Size.ScaleF(0.5))
}

func (v Rectangle[T]) ToFloat64() Rectangle[float64] {
	return Rectangle[float64]{
		Position: v.Position.ToFloat64(),
		Size:     v.Size.ToFloat64(),
	}
}

func (v Rectangle[T]) ToFloat32() Rectangle[float32] {
	return Rectangle[float32]{
		Position: v.Position.ToFloat32(),
		Size:     v.Size.ToFloat32(),
	}
}

func (v Rectangle[T]) ToInt() Rectangle[int] {
	return Rectangle[int]{
		Position: v.Position.ToInt(),
		Size:     v.Size.ToInt(),
	}
}

func (v Rectangle[T]) ToInt32() Rectangle[int32] {
	return Rectangle[int32]{
		Position: v.Position.ToInt32(),
		Size:     v.Size.ToInt32(),
	}
}

func (v Rectangle[T]) ToInt64() Rectangle[int64] {
	return Rectangle[int64]{
		Position: v.Position.ToInt64(),
		Size:     v.Size.ToInt64(),
	}
}

// X returns the x of the xy component
func (r Rectangle[T]) X() T {
	return r.Position.X
}

// SetX changes the x of the xy component of the rectangle
func (r Rectangle[T]) SetX(newX T) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position.SetX(newX),
		Size:     r.Size,
	}
}

func (r Rectangle[T]) AddX(dX T) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position.AddX(dX),
		Size:     r.Size,
	}
}

// Y returns the y of the xy component
func (r Rectangle[T]) Y() T {
	return r.Position.Y
}

// SetY changes the y of the xy component of the rectangle
func (r Rectangle[T]) SetY(newY T) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position.SetY(newY),
		Size:     r.Size,
	}
}

func (r Rectangle[T]) AddY(dY T) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position.AddY(dY),
		Size:     r.Size,
	}
}

// Width returns the x of the wh component
func (r Rectangle[T]) Width() T {
	return r.Size.X
}

// SetWidth changes the x of the wh component of the rectangle
func (r Rectangle[T]) SetWidth(newW T) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position,
		Size:     r.Size.SetX(newW),
	}
}

func (r Rectangle[T]) AddWidth(dW T) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position,
		Size:     r.Size.AddX(dW),
	}
}

// Y returns the y of the wh component
func (r Rectangle[T]) Height() T {
	return r.Size.Y
}

// SetHeight changes the y of the wh component of the rectangle
func (r Rectangle[T]) SetHeight(newH T) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position,
		Size:     r.Size.SetY(newH),
	}
}

func (r Rectangle[T]) AddHeight(dH T) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position,
		Size:     r.Size.AddY(dH),
	}
}

// ResetPosition zero the xy component of the rectangle
func (r Rectangle[T]) ResetPosition() Rectangle[T] {
	return Rectangle[T]{
		Position: vector2.Zero[T](),
		Size:     r.Size,
	}
}

// SetPosition changes the xy component of the rectangle
func (r Rectangle[T]) SetPosition(newXY vector2.Vector[T]) Rectangle[T] {
	return Rectangle[T]{
		Position: newXY,
		Size:     r.Size,
	}
}

// SetPosition changes the xy component of the rectangle
func (r Rectangle[T]) SetPositionXY(x, y T) Rectangle[T] {
	return Rectangle[T]{
		Position: vector2.New(x, y),
		Size:     r.Size,
	}
}

func (r Rectangle[T]) AddPosition(dXY vector2.Vector[T]) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position.Add(dXY),
		Size:     r.Size,
	}
}

func (r Rectangle[T]) AddPositionXY(x, y T) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position.AddXY(x, y),
		Size:     r.Size,
	}
}

// SetSize changes the wh component of the rectangle
func (r Rectangle[T]) SetSize(newWH vector2.Vector[T]) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position,
		Size:     newWH,
	}
}

// SetSizeXY changes the wh component of the rectangle
func (r Rectangle[T]) SetSizeXY(width, height T) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position,
		Size:     vector2.New(width, height),
	}
}

func (r Rectangle[T]) AddSize(dWH vector2.Vector[T]) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position,
		Size:     r.Size.Add(dWH),
	}
}

func (r Rectangle[T]) AddSizeXY(width, height T) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position,
		Size:     r.Size.AddXY(width, height),
	}
}

// Round takes each component of the rectangle and rounds it to the nearest whole
// number
func (v Rectangle[T]) Round() Rectangle[T] {
	return New(
		v.Position.Round(),
		v.Size.Round(),
	)
}

// RoundToInt takes each component of the rectangle and rounds it to the nearest
// whole number, and then casts it to a int
func (v Rectangle[T]) RoundToInt() Rectangle[int] {
	return New(
		v.Position.RoundToInt(),
		v.Size.RoundToInt(),
	)
}

// Ceil applies the ceil math operation to each component of the rectangle
func (v Rectangle[T]) Ceil() Rectangle[T] {
	return New(
		v.Position.Ceil(),
		v.Size.Ceil(),
	)
}

// CeilToInt applies the ceil math operation to each component of the rectangle,
// and then casts it to a int
func (v Rectangle[T]) CeilToInt() Rectangle[int] {
	return New(
		v.Position.CeilToInt(),
		v.Size.CeilToInt(),
	)
}

// Floor applies the floor math operation to each component of the rectangle
func (v Rectangle[T]) Floor() Rectangle[T] {
	return New(
		v.Position.Floor(),
		v.Size.Floor(),
	)
}

// FloorToInt applies the floor math operation to each component of the rectangle,
// and then casts it to a int
func (v Rectangle[T]) FloorToInt() Rectangle[int] {
	return New(
		v.Position.FloorToInt(),
		v.Size.FloorToInt(),
	)
}

func (r Rectangle[T]) Add(xy vector2.Vector[T], wh vector2.Vector[T]) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position.Add(xy),
		Size:     r.Size.Add(wh),
	}
}

func (r Rectangle[T]) AddXYWH(x, y, w, h T) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position.AddXY(x, y),
		Size:     r.Size.AddXY(w, h),
	}
}

func (r Rectangle[T]) GrowXYWH(left, top, right, bottom T) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position.AddXY(-left, -top),
		Size:     r.Size.AddXY(left+right, top+bottom),
	}
}

func (r Rectangle[T]) ShrinkXYWH(left, top, right, bottom T) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position.AddXY(left, top),
		Size:     r.Size.AddXY(-left-right, -top-bottom),
	}
}

func (r Rectangle[T]) Scale(f float64) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position,
		Size:     r.Size.Scale(f),
	}
}

func (r Rectangle[T]) ScaleF(f float32) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position,
		Size:     r.Size.ScaleF(f),
	}
}

func (r Rectangle[T]) ScaleByVector(f vector2.Float64) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position,
		Size:     r.Size.ScaleByVector(f),
	}
}

func (r Rectangle[T]) ScaleByVectorF(f vector2.Float32) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position,
		Size:     r.Size.ScaleByVectorF(f),
	}
}

func (r Rectangle[T]) ScaleByXY(x, y float64) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position,
		Size:     r.Size.ScaleByXY(x, y),
	}
}

func (r Rectangle[T]) ScaleByXYF(x, y float32) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position,
		Size:     r.Size.ScaleByXYF(x, y),
	}
}

func (r Rectangle[T]) Zoom(f float64) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position.Scale(f),
		Size:     r.Size.Scale(f),
	}
}

func (r Rectangle[T]) ZoomF(f float32) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position.ScaleF(f),
		Size:     r.Size.ScaleF(f),
	}
}

func (r Rectangle[T]) ZoomByVector(f vector2.Float64) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position.ScaleByVector(f),
		Size:     r.Size.ScaleByVector(f),
	}
}

func (r Rectangle[T]) ZoomByVectorF(f vector2.Float32) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position.ScaleByVectorF(f),
		Size:     r.Size.ScaleByVectorF(f),
	}
}

func (r Rectangle[T]) ZoomByXY(x, y float64) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position.ScaleByXY(x, y),
		Size:     r.Size.ScaleByXY(x, y),
	}
}

func (r Rectangle[T]) ZoomByXYF(x, y float32) Rectangle[T] {
	return Rectangle[T]{
		Position: r.Position.ScaleByXYF(x, y),
		Size:     r.Size.ScaleByXYF(x, y),
	}
}

func (r Rectangle[T]) Inverse(v vector2.Float64) vector2.Float64 {
	return r.InverseLerp(v).SubXY(0.5, 0.5)
}

func (r Rectangle[T]) InverseF(v vector2.Float32) vector2.Float32 {
	return r.InverseLerpF(v).SubXY(0.5, 0.5)
}

// InverseLerp calculates the inverse lerp of a point within the rectangle, returning a normalized vector2.Vector[T].
func (r Rectangle[T]) InverseLerp(v vector2.Float64) vector2.Float64 {
	return v.Sub(r.Position.ToFloat64()).ToFloat64().ScaleByVector(r.Size.Inv())
}

// InverseLerpF calculates the inverse lerp of a point within the rectangle, returning a normalized vector2.Float32.
func (r Rectangle[T]) InverseLerpF(v vector2.Float32) vector2.Float32 {
	return v.Sub(r.Position.ToFloat32()).ToFloat32().ScaleByVectorF(r.Size.InvF())
}

// InverseLerpXYF calculates the inverse lerp of a point within the rectangle using float32 x and y, returning a normalized vector2.Float32.
//func (r Rectangle[T]) InverseLerpXYF(x, y float32) vector2.Float32 {
//	return vector2.New(float32(x)-float32(r.Position.X), float32(y)-float32(r.Position.Y)).Div(r.Size.ToFloat32())
//}

func (r Rectangle[T]) Lerp(t vector2.Float64) vector2.Vector[T] {
	return r.Position.Add(r.Size.ScaleByVector(t))
}

func (r Rectangle[T]) LerpF(t vector2.Float32) vector2.Vector[T] {
	return r.Position.Add(r.Size.ScaleByVectorF(t))
}

func (r Rectangle[T]) LerpXYF(x, y float32) vector2.Vector[T] {
	return r.Position.Add(r.Size.ScaleByXYF(x, y))
}

func (r Rectangle[T]) Contains(v vector2.Vector[T]) bool {
	return vector2.GreaterEq(v, r.A()) && vector2.LessEq(v, r.B())
}

func (r Rectangle[T]) Pivot(anchor vector2.Vector[T], xy vector2.Vector[T]) Rectangle[T] {
	return Rectangle[T]{
		Position: xy.Sub(anchor.MultByVector(r.Size)),
		Size:     r.Size,
	}
}
