package vector2

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"

	rm "github.com/igadmg/raylib-go/raymath"
)

type Vector[T rm.SignedNumber] struct {
	X T
	Y T
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

func New[T rm.SignedNumber](x T, y T) Vector[T] {
	return Vector[T]{
		X: x,
		Y: y,
	}
}

func NewT[T rm.SignedNumber, XT, YT rm.Number](x XT, y YT) Vector[T] {
	return Vector[T]{
		X: T(x),
		Y: T(y),
	}
}

func NewFloat64[XT, YT rm.Number](x XT, y YT) Float64 {
	return NewT[float64](x, y)
}

func NewFloat32[XT, YT rm.Number](x XT, y YT) Float32 {
	return NewT[float32](x, y)
}

func NewInt[XT, YT rm.Number](x XT, y YT) Int {
	return NewT[int](x, y)
}

func NewInt64[XT, YT rm.Number](x XT, y YT) Int64 {
	return NewT[int64](x, y)
}

func NewInt32[XT, YT rm.Number](x XT, y YT) Int32 {
	return NewT[int32](x, y)
}

func NewInt16[XT, YT rm.Number](x XT, y YT) Int16 {
	return NewT[int16](x, y)
}

func NewInt8[XT, YT rm.Number](x XT, y YT) Int8 {
	return NewT[int8](x, y)
}

// Fill creates a vector where each component is equal to v
func Fill[T rm.SignedNumber](v T) Vector[T] {
	return Vector[T]{
		X: v,
		Y: v,
	}
}

func Zero[T rm.SignedNumber]() Vector[T] {
	return Vector[T]{
		X: 0,
		Y: 0,
	}
}

func Up[T rm.SignedNumber]() Vector[T] {
	return Vector[T]{
		X: 0,
		Y: 1,
	}
}

func Down[T rm.SignedNumber]() Vector[T] {
	return Vector[T]{
		X: 0,
		Y: -1,
	}
}

func Left[T rm.SignedNumber]() Vector[T] {
	return Vector[T]{
		X: -1,
		Y: 0,
	}
}

func Right[T rm.SignedNumber]() Vector[T] {
	return Vector[T]{
		X: 1,
		Y: 0,
	}
}

func One[T rm.SignedNumber]() Vector[T] {
	return Vector[T]{
		X: 1,
		Y: 1,
	}
}

func (v Vector[T]) SignI() Int {
	sv := Zero[int]()
	if v.X < 0 {
		sv.X = -1
	} else if v.X > 0 {
		sv.X = 1
	}
	if v.Y < 0 {
		sv.Y = -1
	} else if v.Y > 0 {
		sv.Y = 1
	}
	return sv
}

func (v Vector[T]) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

func (v Vector[T]) NearZeroF(epsilon float32) bool {
	fv := v.ToFloat32()
	return fv.X > -epsilon && fv.X < epsilon &&
		fv.Y > -epsilon && fv.Y < epsilon
}

func (v Vector[T]) ZeroF(epsilon float32) Float32 {
	fv := v.ToFloat32()
	if fv.X < epsilon && fv.X > -epsilon {
		fv = fv.SetX(0)
	}
	if fv.Y < epsilon && fv.Y > -epsilon {
		fv = fv.SetY(0)
	}

	return fv
}

func (v Vector[T]) Inv() Float64 {
	return Float64{
		X: 1.0 / float64(v.X),
		Y: 1.0 / float64(v.Y),
	}
}

func (v Vector[T]) InvF() Float32 {
	return Float32{
		X: 1.0 / float32(v.X),
		Y: 1.0 / float32(v.Y),
	}
}

// Lerp linearly interpolates between a and b by t
func Lerp[T rm.SignedNumber](t float32, a, b Vector[T]) Vector[T] {
	return Vector[T]{
		X: rm.Lerp(t, a.X, b.X),
		Y: rm.Lerp(t, a.Y, b.Y),
	}
}

func Min[T rm.SignedNumber](a, b Vector[T]) Vector[T] {
	return New(
		min(a.X, b.X),
		min(a.Y, b.Y),
	)
}

func Max[T rm.SignedNumber](a, b Vector[T]) Vector[T] {
	return New(
		max(a.X, b.X),
		max(a.Y, b.Y),
	)
}

func MaxX[T rm.SignedNumber](a, b Vector[T]) T {
	return max(a.X, b.X)
}

func MaxY[T rm.SignedNumber](a, b Vector[T]) T {
	return max(a.Y, b.Y)
}

func MinX[T rm.SignedNumber](a, b Vector[T]) T {
	return min(a.X, b.X)
}

func MinY[T rm.SignedNumber](a, b Vector[T]) T {
	return min(a.Y, b.Y)
}

func Less[T rm.SignedNumber](a, b Vector[T]) bool {
	return a.X < b.X && a.Y < b.Y
}

func LessEq[T rm.SignedNumber](a, b Vector[T]) bool {
	return a.X <= b.X && a.Y <= b.Y
}

func Greater[T rm.SignedNumber](a, b Vector[T]) bool {
	return a.X > b.X && a.Y > b.Y
}

func GreaterEq[T rm.SignedNumber](a, b Vector[T]) bool {
	return a.X >= b.X && a.Y >= b.Y
}

func Midpoint[T rm.SignedNumber](a, b Vector[T]) Vector[T] {
	// center = (b - a)0.5 + a
	// center = b0.5 - a0.5 + a
	// center = b0.5 + a0.5
	// center = 0.5(b + a)
	return Vector[T]{
		X: T(float64(a.X+b.X) * 0.5),
		Y: T(float64(a.Y+b.Y) * 0.5),
	}
}

func Index(xy Int, i int) Int {
	return Int{
		i % xy.X,
		i / xy.X,
	}
}

func (v Vector[T]) String() string {
	return fmt.Sprintf("X: %v; Y: %v;", v.X, v.Y)
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

	return v
}

func Rand(r *rand.Rand) Vector[float64] {
	return Vector[float64]{
		X: r.Float64(),
		Y: r.Float64(),
	}
}

func (v Vector[T]) MinComponent() T {
	return min(v.X, v.Y)
}

func (v Vector[T]) MaxComponent() T {
	return max(v.X, v.Y)
}

func (v Vector[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	}{
		X: float64(v.X),
		Y: float64(v.Y),
	})
}

func (v *Vector[T]) UnmarshalJSON(data []byte) error {
	aux := &struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	}{
		X: 0,
		Y: 0,
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	v.X = T(aux.X)
	v.Y = T(aux.Y)
	return nil
}

func (v Vector[T]) Format(format string) string {
	return fmt.Sprintf(format, v.X, v.Y)
}

// Sqrt applies the Sqrt to each component of the vector
func (v Vector[T]) Sqrt() Vector[T] {
	return New(
		rm.Sqrt(v.X),
		rm.Sqrt(v.Y),
	)
}

func (v Vector[T]) Clamp(vmin, vmax T) Vector[T] {
	return Vector[T]{
		X: rm.Clamp(v.X, vmin, vmax),
		Y: rm.Clamp(v.Y, vmin, vmax),
	}
}

func (v Vector[T]) ClampV(vmin, vmax Vector[T]) Vector[T] {
	return Vector[T]{
		X: rm.Clamp(v.X, vmin.X, vmax.X),
		Y: rm.Clamp(v.Y, vmin.Y, vmax.Y),
	}
}

func (v Vector[T]) Clamp0V(vmax Vector[T]) Vector[T] {
	return Vector[T]{
		X: rm.Clamp(v.X, 0, vmax.X),
		Y: rm.Clamp(v.Y, 0, vmax.Y),
	}
}

func (v Vector[T]) ToNpot() Vector[T] {
	return Vector[T]{
		X: rm.Npot(v.X),
		Y: rm.Npot(v.Y),
	}
}

func To[T, OT rm.SignedNumber](v Vector[OT]) Vector[T] {
	return Vector[T]{
		X: T(v.X),
		Y: T(v.Y),
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
func (v Vector[T]) SetX(X T) Vector[T] {
	return Vector[T]{
		X: X,
		Y: v.Y,
	}
}

// AddX adds to the x component of the vector
func (v Vector[T]) AddX(dX T) Vector[T] {
	return Vector[T]{
		X: v.X + dX,
		Y: v.Y,
	}
}

// SetY changes the y component of the vector
func (v Vector[T]) SetY(Y T) Vector[T] {
	return Vector[T]{
		X: v.X,
		Y: Y,
	}
}

// AddY adds to the y component of the vector
func (v Vector[T]) AddY(dY T) Vector[T] {
	return Vector[T]{
		X: v.X,
		Y: v.Y + dY,
	}
}

func (v Vector[T]) Axis(i int) T {
	switch i {
	case 0:
		return v.X
	case 1:
		return v.Y
	}

	return 0
}

func (v Vector[T]) SetAxis(i int, a T) Vector[T] {
	switch i {
	case 0:
		return Vector[T]{
			X: a,
			Y: v.Y,
		}
	case 1:
		return Vector[T]{
			X: v.X,
			Y: a,
		}
	}

	return v
}

func (v Vector[T]) YX() Vector[T] {
	return Vector[T]{
		X: v.Y,
		Y: v.X,
	}
}

// Angle return angle in radians between vector and other vector [float64]
func (v Vector[T]) Angle(other Vector[T]) float64 {
	denominator := rm.Sqrt(float64(v.LengthSquared() * other.LengthSquared()))
	if denominator < 1e-15 {
		return 0.
	}
	return rm.Acos(rm.Clamp(float64(v.Dot(other))/denominator, -1., 1.))
}

// AngleF return angle in radians between vector and other vector [float32]
func (v Vector[T]) AngleF(other Vector[T]) float32 {
	denominator := rm.Sqrt(float32(v.LengthSquared() * other.LengthSquared()))
	if denominator < 1e-15 {
		return 0.
	}
	return rm.Acos(rm.Clamp(float32(v.Dot(other))/denominator, -1., 1.))
}

// Midpoint returns the midpoint between this vector and the vector passed in.
func (v Vector[T]) Midpoint(o Vector[T]) Vector[T] {
	return o.Add(v).Scale(0.5)
}

// Dot return dot product between vector and other vector
func (v Vector[T]) Dot(other Vector[T]) T {
	return v.X*other.X + v.Y*other.Y
}

// Perpendicular creates a vector perpendicular to the one passed in with the
// same magnitude
func (v Vector[T]) Perpendicular() Vector[T] {
	return Vector[T]{
		X: v.Y,
		Y: -v.X,
	}
}

// Add returns a vector that is the result of two vectors added together
func (v Vector[T]) Add(other Vector[T]) Vector[T] {
	return Vector[T]{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v Vector[T]) AddXY(x, y T) Vector[T] {
	return Vector[T]{
		X: v.X + x,
		Y: v.Y + y,
	}
}

func (v Vector[T]) Sub(other Vector[T]) Vector[T] {
	return Vector[T]{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (v Vector[T]) SubXY(x, y T) Vector[T] {
	return Vector[T]{
		X: v.X - x,
		Y: v.Y - y,
	}
}

func (v Vector[T]) ReciprocalF() Vector[float32] {
	return Vector[float32]{
		X: 1.0 / float32(v.X),
		Y: 1.0 / float32(v.Y),
	}
}

func (v Vector[T]) Reciprocal() Vector[float64] {
	return Vector[float64]{
		X: 1.0 / float64(v.X),
		Y: 1.0 / float64(v.Y),
	}
}

func (v Vector[T]) Product() T {
	return v.X * v.Y
}

func (v Vector[T]) LengthSquared() T {
	return v.X*v.X + v.Y*v.Y
}

func (v Vector[T]) Length() float64 {
	return rm.Sqrt((float64)(v.LengthSquared()))
}

func (v Vector[T]) LengthF() float32 {
	return rm.Sqrt((float32)(v.LengthSquared()))
}

func (v Vector[T]) Normalized() Vector[T] {
	return v.DivByConstant(v.Length())
}

func (v Vector[T]) NormalizeF(a Vector[T]) Float32 {
	return Float32{
		X: rm.NormalizeF(a.X, 0, v.X),
		Y: rm.NormalizeF(a.Y, 0, v.Y),
	}
}

func (v Vector[T]) Normalize(a Vector[T]) Float64 {
	return Float64{
		X: float64(rm.Normalize(a.X, 0, v.X)),
		Y: float64(rm.Normalize(a.Y, 0, v.Y)),
	}
}

func (v Vector[T]) Negated() Vector[T] {
	return Vector[T]{
		X: -v.X,
		Y: -v.Y,
	}
}

func (v Vector[T]) Scale(t float64) Vector[T] {
	return Vector[T]{
		X: T(float64(v.X) * t),
		Y: T(float64(v.Y) * t),
	}
}

func (v Vector[T]) ScaleF(t float32) Vector[T] {
	return Vector[T]{
		X: T(float32(v.X) * t),
		Y: T(float32(v.Y) * t),
	}
}

func (v Vector[T]) ScaleByVector(o Float64) Vector[T] {
	return Vector[T]{
		X: T(float64(v.X) * o.X),
		Y: T(float64(v.Y) * o.Y),
	}
}

func (v Vector[T]) ScaleByVectorF(o Float32) Vector[T] {
	return Vector[T]{
		X: T(float32(v.X) * o.X),
		Y: T(float32(v.Y) * o.Y),
	}
}

func (v Vector[T]) ScaleByVectorI(o Int) Vector[T] {
	return Vector[T]{
		X: v.X * T(o.X),
		Y: v.Y * T(o.Y),
	}
}

func (v Vector[T]) ScaleByXY(x, y float64) Vector[T] {
	return Vector[T]{
		X: T(float64(v.X) * x),
		Y: T(float64(v.Y) * y),
	}
}

func (v Vector[T]) ScaleByXYF(x, y float32) Vector[T] {
	return Vector[T]{
		X: T(float32(v.X) * x),
		Y: T(float32(v.Y) * y),
	}
}

func (v Vector[T]) ScaleByXYI(x, y int) Vector[T] {
	return Vector[T]{
		X: v.X * T(x),
		Y: v.Y * T(y),
	}
}

func (v Vector[T]) MultByVector(o Vector[T]) Vector[T] {
	return Vector[T]{
		X: v.X * o.X,
		Y: v.Y * o.Y,
	}
}

func (v Vector[T]) DivByVector(o Vector[T]) Vector[T] {
	return Vector[T]{
		X: v.X / o.X,
		Y: v.Y / o.Y,
	}
}

func (v Vector[T]) DivByConstant(t float64) Vector[T] {
	return v.Scale(1.0 / t)
}

func (v Vector[T]) Project(normal Vector[T]) Vector[T] {
	vdn := float64(v.Dot(normal))
	ndn := float64(normal.Dot(normal))
	mag := vdn / ndn
	return normal.Scale(mag)
}

func (v Vector[T]) Reject(normal Vector[T]) Vector[T] {
	return v.Sub(v.Project(normal))
}

func (v Vector[T]) Reflect(normal Vector[T]) Vector[T] {
	return v.Sub(normal.Scale(2. * float64(v.Dot(normal))))
}

func (v Vector[T]) DistanceSquared(other Vector[T]) T {
	xDist := other.X - v.X
	yDist := other.Y - v.Y
	return (xDist * xDist) + (yDist * yDist)
}

// Distance is the euclidean distance between two points
func (v Vector[T]) Distance(other Vector[T]) float64 {
	return rm.Sqrt((float64)(v.DistanceSquared(other)))
}

// Round takes each component of the vector and rounds it to the nearest whole
// number
func (v Vector[T]) Round() Vector[T] {
	return Vector[T]{
		X: rm.Round(v.X),
		Y: rm.Round(v.Y),
	}
}

// RoundToInt takes each component of the vector and rounds it to the nearest
// whole number, and then casts it to a int
func (v Vector[T]) RoundToInt() Vector[int] {
	return New(
		int(rm.Round(v.X)),
		int(rm.Round(v.Y)),
	)
}

// Ceil applies the ceil math operation to each component of the vector
func (v Vector[T]) Ceil() Vector[T] {
	return Vector[T]{
		X: rm.Ceil(v.X),
		Y: rm.Ceil(v.Y),
	}
}

// CeilToInt applies the ceil math operation to each component of the vector,
// and then casts it to a int
func (v Vector[T]) CeilToInt() Vector[int] {
	return New(
		int(rm.Ceil(v.X)),
		int(rm.Ceil(v.Y)),
	)
}

func (v Vector[T]) Floor() Vector[T] {
	return Vector[T]{
		X: rm.Floor(v.X),
		Y: rm.Floor(v.Y),
	}
}

// FloorToInt applies the floor math operation to each component of the vector,
// and then casts it to a int
func (v Vector[T]) FloorToInt() Vector[int] {
	return New(
		int(rm.Floor(v.X)),
		int(rm.Floor(v.Y)),
	)
}

// Abs applies the Abs math operation to each component of the vector
func (v Vector[T]) Abs() Vector[T] {
	return Vector[T]{
		X: rm.Abs(v.X),
		Y: rm.Abs(v.Y),
	}
}

func (v Vector[T]) NearZero() bool {
	return rm.NearZero(v.X) && rm.NearZero(v.Y)
}

func (v Vector[T]) ContainsNaN() bool {
	if math.IsNaN(float64(v.X)) {
		return true
	}

	if math.IsNaN(float64(v.Y)) {
		return true
	}

	return false
}

func (v Vector[T]) Flip() Vector[T] {
	return Vector[T]{
		X: v.X * -1,
		Y: v.Y * -1,
	}
}

func (v Vector[T]) FlipX() Vector[T] {
	return Vector[T]{
		X: v.X * -1,
		Y: v.Y,
	}
}

func (v Vector[T]) FlipY() Vector[T] {
	return Vector[T]{
		X: v.X,
		Y: v.Y * -1,
	}
}

func (v Vector[T]) Pivot(anchor Vector[T], wh Vector[T]) Vector[T] {
	return Vector[T]{
		X: v.X - wh.X*anchor.X,
		Y: v.Y - wh.Y*anchor.Y,
	}
}

// Log returns the natural logarithm for each component
func (v Vector[T]) Log() Vector[T] {
	return Vector[T]{
		X: rm.Log(v.X),
		Y: rm.Log(v.Y),
	}
}

// Log10 returns the decimal logarithm for each component.
func (v Vector[T]) Log10() Vector[T] {
	return Vector[T]{
		X: rm.Log10(v.X),
		Y: rm.Log10(v.Y),
	}
}

// Log2 returns the binary logarithm for each component
func (v Vector[T]) Log2() Vector[T] {
	return Vector[T]{
		X: rm.Log2(v.X),
		Y: rm.Log2(v.Y),
	}
}

// Exp2 returns 2**x, the base-2 exponential for each component
func (v Vector[T]) Exp2() Vector[T] {
	return Vector[T]{
		X: rm.Exp2(v.X),
		Y: rm.Exp2(v.Y),
	}
}

// Exp returns e**x, the base-e exponential for each component
func (v Vector[T]) Exp() Vector[T] {
	return Vector[T]{
		X: rm.Exp(v.X),
		Y: rm.Exp(v.Y),
	}
}

// Expm1 returns e**x - 1, the base-e exponential for each component minus 1. It is more accurate than Exp(x) - 1 when the component is near zero
func (v Vector[T]) Expm1() Vector[T] {
	return Vector[T]{
		X: rm.Expm1(v.X),
		Y: rm.Expm1(v.Y),
	}
}

func (v Vector[T]) Values() (T, T) {
	return v.X, v.Y
}
