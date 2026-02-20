package vector4

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"

	"github.com/Mishka-Squat/goex/mathex"
)

const componentCount = 4

func (v Vector[T]) Write(out io.Writer, endian binary.ByteOrder) (err error) {
	switch v := any(v).(type) {
	case Float64:
		bytes := make([]byte, 8*componentCount)
		endian.PutUint64(bytes, math.Float64bits(v.X))
		endian.PutUint64(bytes[8:], math.Float64bits(v.Y))
		endian.PutUint64(bytes[16:], math.Float64bits(v.Z))
		endian.PutUint64(bytes[24:], math.Float64bits(v.W))
		_, err = out.Write(bytes)
		return

	case Float32:
		bytes := make([]byte, 4*componentCount)
		endian.PutUint32(bytes, math.Float32bits(v.X))
		endian.PutUint32(bytes[4:], math.Float32bits(v.Y))
		endian.PutUint32(bytes[8:], math.Float32bits(v.Z))
		endian.PutUint32(bytes[12:], math.Float32bits(v.W))
		_, err = out.Write(bytes)
		return

	case Int8:
		_, err = out.Write([]byte{
			byte(v.X),
			byte(v.Y),
			byte(v.Z),
			byte(v.W),
		})
		return

	case Int16:
		bytes := make([]byte, 2*componentCount)
		endian.PutUint16(bytes, uint16(v.X))
		endian.PutUint16(bytes[2:], uint16(v.Y))
		endian.PutUint16(bytes[4:], uint16(v.Z))
		endian.PutUint16(bytes[6:], uint16(v.W))
		_, err = out.Write(bytes)
		return

	case Int32:
		bytes := make([]byte, 4*componentCount)
		endian.PutUint32(bytes, uint32(v.X))
		endian.PutUint32(bytes[4:], uint32(v.Y))
		endian.PutUint32(bytes[8:], uint32(v.Z))
		endian.PutUint32(bytes[12:], uint32(v.W))
		_, err = out.Write(bytes)
		return

	case Int64:
		bytes := make([]byte, 8*componentCount)
		endian.PutUint64(bytes, uint64(v.X))
		endian.PutUint64(bytes[8:], uint64(v.Y))
		endian.PutUint64(bytes[16:], uint64(v.Z))
		endian.PutUint64(bytes[24:], uint64(v.W))
		_, err = out.Write(bytes)
		return
	}

	panic(fmt.Errorf("write unimplemented type: %#v", v))
}

func Read[T mathex.SignedNumber](in io.Reader, endian binary.ByteOrder) (v Vector[T], err error) {
	switch any(v).(type) {
	case Float64:
		v, err := ReadFloat64(in, endian)
		return any(v).(Vector[T]), err

	case Float32:
		v, err := ReadFloat32(in, endian)
		return any(v).(Vector[T]), err

	case Int8:
		v, err := ReadInt8(in)
		return any(v).(Vector[T]), err

	case Int16:
		v, err := ReadInt16(in, endian)
		return any(v).(Vector[T]), err

	case Int32:
		v, err := ReadInt32(in, endian)
		return any(v).(Vector[T]), err

	case Int64:
		v, err := ReadInt64(in, endian)
		return any(v).(Vector[T]), err
	}

	panic(fmt.Errorf("read unimplemented type: %#v", v))
}

func ReadFloat64(in io.Reader, endian binary.ByteOrder) (Vector[float64], error) {
	buf := make([]byte, componentCount*8)
	_, err := io.ReadFull(in, buf)
	return Vector[float64]{
		X: math.Float64frombits(endian.Uint64(buf)),
		Y: math.Float64frombits(endian.Uint64(buf[8:])),
		Z: math.Float64frombits(endian.Uint64(buf[16:])),
		W: math.Float64frombits(endian.Uint64(buf[24:])),
	}, err
}

func ReadFloat32(in io.Reader, endian binary.ByteOrder) (Vector[float32], error) {
	buf := make([]byte, componentCount*4)
	_, err := io.ReadFull(in, buf)
	return Vector[float32]{
		X: math.Float32frombits(endian.Uint32(buf)),
		Y: math.Float32frombits(endian.Uint32(buf[4:])),
		Z: math.Float32frombits(endian.Uint32(buf[8:])),
		W: math.Float32frombits(endian.Uint32(buf[12:])),
	}, err
}

func ReadInt8(in io.Reader) (Vector[int8], error) {
	buf := make([]byte, componentCount)
	_, err := io.ReadFull(in, buf)
	return Vector[int8]{
		X: int8(buf[0]),
		Y: int8(buf[1]),
		Z: int8(buf[2]),
		W: int8(buf[3]),
	}, err
}

func ReadInt16(in io.Reader, endian binary.ByteOrder) (Vector[int16], error) {
	buf := make([]byte, componentCount*2)
	_, err := io.ReadFull(in, buf)
	return Vector[int16]{
		X: int16(endian.Uint16(buf)),
		Y: int16(endian.Uint16(buf[2:])),
		Z: int16(endian.Uint16(buf[4:])),
		W: int16(endian.Uint16(buf[6:])),
	}, err
}

func ReadInt32(in io.Reader, endian binary.ByteOrder) (Vector[int32], error) {
	buf := make([]byte, componentCount*4)
	_, err := io.ReadFull(in, buf)
	return Vector[int32]{
		X: int32(endian.Uint32(buf)),
		Y: int32(endian.Uint32(buf[4:])),
		Z: int32(endian.Uint32(buf[8:])),
		W: int32(endian.Uint32(buf[12:])),
	}, err
}

func ReadInt64(in io.Reader, endian binary.ByteOrder) (Vector[int64], error) {
	buf := make([]byte, componentCount*8)
	_, err := io.ReadFull(in, buf)
	return Vector[int64]{
		X: int64(endian.Uint64(buf)),
		Y: int64(endian.Uint64(buf[8:])),
		Z: int64(endian.Uint64(buf[16:])),
		W: int64(endian.Uint64(buf[24:])),
	}, err
}
