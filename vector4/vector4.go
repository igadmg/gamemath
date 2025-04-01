package vector4

import (
	"encoding/json"
	"fmt"
	"image/color"
	"math"

	rm "github.com/igadmg/raylib-go/raymath"
	"github.com/igadmg/raylib-go/raymath/vector2"
	"github.com/igadmg/raylib-go/raymath/vector3"
)

// Vector contains 4 components
type Vector[T rm.SignedNumber] struct {
	X T
	Y T
	Z T
	W T
}

type (
	Float64 = Vector[float64]
	Float32 = Vector[float32]
	Int     = Vector[int]
	Int64   = Vector[int64]
	Int32   = Vector[int32]
	Int16   = Vector[int16]
	Int8    = Vector[int8]
)

// New creates a new vector with corresponding 3 components
func New[T rm.SignedNumber](x, y, z, w T) Vector[T] {
	return Vector[T]{
		X: x,
		Y: y,
		Z: z,
		W: w,
	}
}

func NewT[T rm.SignedNumber, XT, YT, ZT, WT rm.Number](x XT, y YT, z ZT, w WT) Vector[T] {
	return Vector[T]{
		X: T(x),
		Y: T(y),
		Z: T(z),
		W: T(w),
	}
}

func NewFloat64[XT, YT, ZT, WT rm.Number](x XT, y YT, z ZT, w WT) Float64 {
	return NewT[float64](x, y, z, w)
}

func NewFloat32[XT, YT, ZT, WT rm.Number](x XT, y YT, z ZT, w WT) Float32 {
	return NewT[float32](x, y, z, w)
}

func NewInt[XT, YT, ZT, WT rm.Number](x XT, y YT, z ZT, w WT) Int {
	return NewT[int](x, y, z, w)
}

func NewInt64[XT, YT, ZT, WT rm.Number](x XT, y YT, z ZT, w WT) Int64 {
	return NewT[int64](x, y, z, w)
}

func NewInt32[XT, YT, ZT, WT rm.Number](x XT, y YT, z ZT, w WT) Int32 {
	return NewT[int32](x, y, z, w)
}

func NewInt16[XT, YT, ZT, WT rm.Number](x XT, y YT, z ZT, w WT) Int16 {
	return NewT[int16](x, y, z, w)
}

func NewInt8[XT, YT, ZT, WT rm.Number](x XT, y YT, z ZT, w WT) Int8 {
	return NewT[int8](x, y, z, w)
}

// Fill creates a vector where each component is equal to v
func Fill[T rm.SignedNumber](v T) Vector[T] {
	return Vector[T]{
		X: v,
		Y: v,
		Z: v,
		W: v,
	}
}

func FromColor(c color.Color) Float64 {
	r, g, b, a := c.RGBA()
	return New(float64(r)/0xffff, float64(g)/0xffff, float64(b)/0xffff, float64(a)/0xffff)
}

// Zero is (0, 0, 0)
func Zero[T rm.SignedNumber]() Vector[T] {
	return New[T](0, 0, 0, 0)
}

// One is (1, 1, 1)
func One[T rm.SignedNumber]() Vector[T] {
	return New[T](1, 1, 1, 1)
}

// Average sums all vector4's components together and divides each
// component by the number of vectors added
func Average[T rm.SignedNumber](vectors []Vector[T]) Vector[T] {
	var center Vector[T]
	for _, v := range vectors {
		center = center.Add(v)
	}
	return center.DivByConstant(float64(len(vectors)))
}

// Lerp linearly interpolates between a and b by t
// func Lerp[T vector.Number](a, b Vector[T], t float64) Vector[T] {

// 	// (b - a) * t + a
// 	// bt - at + a
// 	// bt - a(1 - t)
// 	tm1 := 1. - t
// 	return Vector[T]{
// 		x: T((float64(b.x) * t) - (float64(a.x) * tm1)),
// 		y: T((float64(b.y) * t) - (float64(a.y) * tm1)),
// 		z: T((float64(b.z) * t) - (float64(a.z) * tm1)),
// 		w: T((float64(b.w) * t) - (float64(a.w) * tm1)),
// 	}
// }

// Lerp linearly interpolates between a and b by t
func Lerp[T rm.SignedNumber](a, b Vector[T], t float64) Vector[T] {

	// return b.Sub(a).Scale(t).Add(a)
	return Vector[T]{
		X: T((float64(b.X-a.X) * t) + float64(a.X)),
		Y: T((float64(b.Y-a.Y) * t) + float64(a.Y)),
		Z: T((float64(b.Z-a.Z) * t) + float64(a.Z)),
		W: T((float64(b.W-a.W) * t) + float64(a.W)),
	}
}

func (v Vector[T]) Negated() Vector[T] {
	return Vector[T]{
		X: -v.X,
		Y: -v.Y,
		Z: -v.Z,
		W: -v.W,
	}
}

func (v Vector[T]) Scale(t float64) Vector[T] {
	return Vector[T]{
		X: T(float64(v.X) * t),
		Y: T(float64(v.Y) * t),
		Z: T(float64(v.Z) * t),
		W: T(float64(v.W) * t),
	}
}

func (v Vector[T]) DivByConstant(t float64) Vector[T] {
	return Vector[T]{
		X: T(float64(v.X) / t),
		Y: T(float64(v.Y) / t),
		Z: T(float64(v.Z) / t),
		W: T(float64(v.W) / t),
	}
}

func Min[T rm.SignedNumber](a, b Vector[T]) Vector[T] {
	return New(
		min(a.X, b.X),
		min(a.Y, b.Y),
		min(a.Z, b.Z),
		min(a.W, b.W),
	)
}

func Max[T rm.SignedNumber](a, b Vector[T]) Vector[T] {
	return New(
		max(a.X, b.X),
		max(a.Y, b.Y),
		max(a.Z, b.Z),
		max(a.W, b.W),
	)
}

func MaxX[T rm.SignedNumber](a, b Vector[T]) T {
	return max(a.X, b.X)
}

func MaxY[T rm.SignedNumber](a, b Vector[T]) T {
	return max(a.Y, b.Y)
}

func MaxZ[T rm.SignedNumber](a, b Vector[T]) T {
	return max(a.Z, b.Z)
}

func MaxW[T rm.SignedNumber](a, b Vector[T]) T {
	return max(a.W, b.W)
}

func MinX[T rm.SignedNumber](a, b Vector[T]) T {
	return min(a.X, b.X)
}

func MinY[T rm.SignedNumber](a, b Vector[T]) T {
	return min(a.Y, b.Y)
}

func MinZ[T rm.SignedNumber](a, b Vector[T]) T {
	return min(a.Z, b.Z)
}

func MinW[T rm.SignedNumber](a, b Vector[T]) T {
	return min(a.W, b.W)
}

func Midpoint[T rm.SignedNumber](a, b Vector[T]) Vector[T] {
	// center = (b - a)0.5 + a
	// center = b0.5 - a0.5 + a
	// center = b0.5 + a0.5
	// center = 0.5(b + a)
	return Vector[T]{
		X: T(float64(a.X+b.X) * 0.5),
		Y: T(float64(a.Y+b.Y) * 0.5),
		Z: T(float64(a.Z+b.Z) * 0.5),
		W: T(float64(a.W+b.W) * 0.5),
	}
}

func (v Vector[T]) String() string {
	return fmt.Sprintf("X: %v; Y: %v; Z: %v; W: %v;", v.X, v.Y, v.Z, v.W)
}

// Builds a vector from the data found from the passed in array to the best of
// it's ability. If the length of the array is smaller than the vector itself,
// only those values will be used to build the vector, and the remaining vector
// components will remain the default value of the vector's data type (some
// version of 0).
func FromArray[T rm.SignedNumber](data []T) Vector[T] {
	v := Vector[T]{}

	if len(data) > 0 {
		v.X = data[0]
	}

	if len(data) > 1 {
		v.Y = data[1]
	}

	if len(data) > 2 {
		v.Z = data[2]
	}

	if len(data) > 3 {
		v.W = data[3]
	}

	return v
}

func (v Vector[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
		W float64 `json:"w"`
	}{
		X: float64(v.X),
		Y: float64(v.Y),
		Z: float64(v.Z),
		W: float64(v.W),
	})
}

func (v *Vector[T]) UnmarshalJSON(data []byte) error {
	aux := &struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
		W float64 `json:"w"`
	}{
		X: 0,
		Y: 0,
		Z: 0,
		W: 0,
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	v.X = T(aux.X)
	v.Y = T(aux.Y)
	v.Z = T(aux.Z)
	v.W = T(aux.W)
	return nil
}

func (v Vector[T]) Format(format string) string {
	return fmt.Sprintf(format, v.X, v.Y, v.Z, v.W)
}

func (v Vector[T]) MinComponent() T {
	return min(v.X, v.Y, v.Z, v.W)
}

func (v Vector[T]) MaxComponent() T {
	return max(v.X, v.Y, v.Z, v.W)
}

func To[T, OT rm.SignedNumber](v Vector[OT]) Vector[T] {
	return Vector[T]{
		X: T(v.X),
		Y: T(v.Y),
		Z: T(v.Z),
		W: T(v.W),
	}
}

func (v Vector[T]) ToFloat64() Vector[float64] {
	return To[float64](v)
}

func (v Vector[T]) ToFloat32() Vector[float32] {
	return To[float32](v)
}

func (v Vector[T]) ToInt() Vector[int] {
	return To[int](v)
}

func (v Vector[T]) ToInt64() Vector[int64] {
	return To[int64](v)
}

func (v Vector[T]) ToInt32() Vector[int32] {
	return To[int32](v)
}

func (v Vector[T]) ToInt16() Vector[int16] {
	return To[int16](v)
}

func (v Vector[T]) ToInt8() Vector[int8] {
	return To[int8](v)
}

// SetX changes the x component of the vector
func (v Vector[T]) SetX(newX T) Vector[T] {
	return Vector[T]{
		X: newX,
		Y: v.Y,
		Z: v.Z,
		W: v.W,
	}
}

func (v Vector[T]) AddX(dX T) Vector[T] {
	return Vector[T]{
		X: v.X + dX,
		Y: v.Y,
		Z: v.Z,
		W: v.W,
	}
}

// SetY changes the y component of the vector
func (v Vector[T]) SetY(newY T) Vector[T] {
	return Vector[T]{
		X: v.X,
		Y: newY,
		Z: v.Z,
		W: v.W,
	}
}

func (v Vector[T]) AddY(dY T) Vector[T] {
	return Vector[T]{
		X: v.X,
		Y: v.Y + dY,
		Z: v.Z,
		W: v.W,
	}
}

// SetZ changes the z component of the vector
func (v Vector[T]) SetZ(newZ T) Vector[T] {
	return Vector[T]{
		X: v.X,
		Y: v.Y,
		Z: newZ,
		W: v.W,
	}
}

func (v Vector[T]) AddZ(dZ T) Vector[T] {
	return Vector[T]{
		X: v.X,
		Y: v.Y,
		Z: v.Z + dZ,
		W: v.W,
	}
}

// SetW changes the w component of the vector
func (v Vector[T]) SetW(newW T) Vector[T] {
	return Vector[T]{
		X: v.X,
		Y: v.Y,
		Z: v.Z,
		W: newW,
	}
}

func (v Vector[T]) AddW(dW T) Vector[T] {
	return Vector[T]{
		X: v.X,
		Y: v.Y,
		Z: v.Z,
		W: v.W + dW,
	}
}

// Add takes each component of our vector and adds them to the vector passed
// in, returning a resulting vector
func (v Vector[T]) Add(other Vector[T]) Vector[T] {
	return Vector[T]{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
		W: v.W + other.W,
	}
}

func (v Vector[T]) Sub(other Vector[T]) Vector[T] {
	return Vector[T]{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
		W: v.W - other.W,
	}
}

func (v Vector[T]) ReciprocalF() Vector[float32] {
	return Vector[float32]{
		X: 1.0 / float32(v.X),
		Y: 1.0 / float32(v.Y),
		Z: 1.0 / float32(v.Z),
		W: 1.0 / float32(v.W),
	}
}

func (v Vector[T]) Reciprocal() Vector[float64] {
	return Vector[float64]{
		X: 1.0 / float64(v.X),
		Y: 1.0 / float64(v.Y),
		Z: 1.0 / float64(v.Z),
		W: 1.0 / float64(v.W),
	}
}

func (v Vector[T]) Product() T {
	return v.X * v.Y * v.Z * v.W
}

func (v Vector[T]) Dot(other Vector[T]) float64 {
	return float64((v.X * other.X) + (v.Y * other.Y) + (v.Z * other.Z) + (v.W * other.W))
}

func (v Vector[T]) Normalized() Vector[T] {
	return v.DivByConstant(v.Length())
}

func (v Vector[T]) Length() float64 {
	return rm.Sqrt(float64(v.LengthSquared()))
}

func (v Vector[T]) LengthF() float32 {
	return rm.Sqrt(float32(v.LengthSquared()))
}

func (v Vector[T]) LengthSquared() T {
	return (v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z) + (v.W * v.W)
}

// Sqrt applies the Sqrt to each component of the vector
func (v Vector[T]) Sqrt() Vector[T] {
	return New(
		rm.Sqrt(v.X),
		rm.Sqrt(v.Y),
		rm.Sqrt(v.Z),
		rm.Sqrt(v.W),
	)
}

// Abs applies the Abs math operation to each component of the vector
func (v Vector[T]) Abs() Vector[T] {
	return New(
		rm.Abs(v.X),
		rm.Abs(v.Y),
		rm.Abs(v.Z),
		rm.Abs(v.W),
	)
}

func (v Vector[T]) Clamp(vmin, vmax T) Vector[T] {
	return Vector[T]{
		X: rm.Clamp(v.X, vmin, vmax),
		Y: rm.Clamp(v.Y, vmin, vmax),
		Z: rm.Clamp(v.Z, vmin, vmax),
		W: rm.Clamp(v.W, vmin, vmax),
	}
}

// Round takes each component of the vector and rounds it to the nearest whole
// number
func (v Vector[T]) Round() Vector[T] {
	return New(
		rm.Round(v.X),
		rm.Round(v.Y),
		rm.Round(v.Z),
		rm.Round(v.W),
	)
}

// RoundToInt takes each component of the vector and rounds it to the nearest
// whole number, and then casts it to a int
func (v Vector[T]) RoundToInt() Vector[int] {
	return New(
		int(rm.Round(v.X)),
		int(rm.Round(v.Y)),
		int(rm.Round(v.Z)),
		int(rm.Round(v.W)),
	)
}

// Floor applies the floor math operation to each component of the vector
func (v Vector[T]) Floor() Vector[T] {
	return New(
		rm.Floor(v.X),
		rm.Floor(v.Y),
		rm.Floor(v.Z),
		rm.Floor(v.W),
	)
}

// FloorToInt applies the floor math operation to each component of the vector,
// and then casts it to a int
func (v Vector[T]) FloorToInt() Vector[int] {
	return New(
		int(rm.Floor(v.X)),
		int(rm.Floor(v.Y)),
		int(rm.Floor(v.Z)),
		int(rm.Floor(v.W)),
	)
}

// Ceil applies the ceil math operation to each component of the vector
func (v Vector[T]) Ceil() Vector[T] {
	return New(
		rm.Ceil(v.X),
		rm.Ceil(v.Y),
		rm.Ceil(v.Z),
		rm.Ceil(v.W),
	)
}

// CeilToInt applies the ceil math operation to each component of the vector,
// and then casts it to a int
func (v Vector[T]) CeilToInt() Vector[int] {
	return New(
		int(rm.Ceil(v.X)),
		int(rm.Ceil(v.Y)),
		int(rm.Ceil(v.Z)),
		int(rm.Ceil(v.W)),
	)
}

// MultByVector is component wise multiplication, also known as Hadamard product.
func (v Vector[T]) MultByVector(o Vector[T]) Vector[T] {
	return Vector[T]{
		X: v.X * o.X,
		Y: v.Y * o.Y,
		Z: v.Z * o.Z,
		W: v.W * o.W,
	}
}

func (v Vector[T]) ContainsNaN() bool {
	if math.IsNaN(float64(v.X)) {
		return true
	}

	if math.IsNaN(float64(v.Y)) {
		return true
	}

	if math.IsNaN(float64(v.Z)) {
		return true
	}

	if math.IsNaN(float64(v.W)) {
		return true
	}

	return false
}

func (v Vector[T]) NearZero() bool {
	return rm.NearZero(v.X) && rm.NearZero(v.Y) && rm.NearZero(v.Z) && rm.NearZero(v.W)
}

func (v Vector[T]) Flip() Vector[T] {
	return Vector[T]{
		X: v.X * -1,
		Y: v.Y * -1,
		Z: v.Z * -1,
		W: v.W * -1,
	}
}

func (v Vector[T]) FlipX() Vector[T] {
	return Vector[T]{
		X: v.X * -1,
		Y: v.Y,
		Z: v.Z,
		W: v.W,
	}
}

func (v Vector[T]) FlipY() Vector[T] {
	return Vector[T]{
		X: v.X,
		Y: v.Y * -1,
		Z: v.Z,
		W: v.W,
	}
}

func (v Vector[T]) FlipZ() Vector[T] {
	return Vector[T]{
		X: v.X,
		Y: v.Y,
		Z: v.Z * -1,
		W: v.W,
	}
}

func (v Vector[T]) FlipW() Vector[T] {
	return Vector[T]{
		X: v.X,
		Y: v.Y,
		Z: v.Z,
		W: v.W * -1,
	}
}

func (v Vector[T]) XYZ() vector3.Vector[T] {
	return vector3.New(v.X, v.Y, v.Z)
}

// XY returns vector2 with the x and y components
func (v Vector[T]) XY() vector2.Vector[T] {
	return vector2.New(v.X, v.Y)
}

// XZ returns vector2 with the x and z components
func (v Vector[T]) XZ() vector2.Vector[T] {
	return vector2.New(v.X, v.Z)
}

// YZ returns vector2 with the y and z components
func (v Vector[T]) YZ() vector2.Vector[T] {
	return vector2.New(v.Y, v.Z)
}

// YX returns vector2 with the y and x components
func (v Vector[T]) YX() vector2.Vector[T] {
	return vector2.New(v.Y, v.X)
}

// ZX returns vector2 with the z and x components
func (v Vector[T]) ZX() vector2.Vector[T] {
	return vector2.New(v.Z, v.X)
}

// ZY returns vector2 with the z and y components
func (v Vector[T]) ZY() vector2.Vector[T] {
	return vector2.New(v.Z, v.Y)
}

// Log returns the natural logarithm for each component
func (v Vector[T]) Log() Vector[T] {
	return Vector[T]{
		X: rm.Log(v.X),
		Y: rm.Log(v.Y),
		Z: rm.Log(v.Z),
		W: rm.Log(v.W),
	}
}

// Log10 returns the decimal logarithm for each component.
func (v Vector[T]) Log10() Vector[T] {
	return Vector[T]{
		X: rm.Log10(v.X),
		Y: rm.Log10(v.Y),
		Z: rm.Log10(v.Z),
		W: rm.Log10(v.W),
	}
}

// Log2 returns the binary logarithm for each component
func (v Vector[T]) Log2() Vector[T] {
	return Vector[T]{
		X: rm.Log2(v.X),
		Y: rm.Log2(v.Y),
		Z: rm.Log2(v.Z),
		W: rm.Log2(v.W),
	}
}

// Exp2 returns 2**x, the base-2 exponential for each component
func (v Vector[T]) Exp2() Vector[T] {
	return Vector[T]{
		X: rm.Exp2(v.X),
		Y: rm.Exp2(v.Y),
		Z: rm.Exp2(v.Z),
		W: rm.Exp2(v.W),
	}
}

// Exp returns e**x, the base-e exponential for each component
func (v Vector[T]) Exp() Vector[T] {
	return Vector[T]{
		X: rm.Exp(v.X),
		Y: rm.Exp(v.Y),
		Z: rm.Exp(v.Z),
		W: rm.Exp(v.W),
	}
}

// Expm1 returns e**x - 1, the base-e exponential for each component minus 1. It is more accurate than Exp(x) - 1 when the component is near zero
func (v Vector[T]) Expm1() Vector[T] {
	return Vector[T]{
		X: rm.Expm1(v.X),
		Y: rm.Expm1(v.Y),
		Z: rm.Expm1(v.Z),
		W: rm.Expm1(v.W),
	}
}

func (v Vector[T]) Values() (T, T, T, T) {
	return v.X, v.Y, v.Z, v.W
}
