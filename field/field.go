package field

import (
	"database/sql/driver"
)

// Field a standard field struct
type Field = genericsField[driver.Valuer]

// AsString returns a String field
func (field Field) AsString() String {
	return newChars[string](field.expr)
}

// AsBytes returns a Bytes field
func (field Field) AsBytes() Bytes {
	return newChars[[]byte](field.expr)
}

// AsInt returns an Int field
func (field Field) AsInt() Int {
	return newNumber[int](field.expr)
}

// AsInt8 returns an Int8 field
func (field Field) AsInt8() Int8 {
	return newNumber[int8](field.expr)
}

// AsInt16 returns an Int16 field
func (field Field) AsInt16() Int16 {
	return newNumber[int16](field.expr)
}

// AsInt32 returns an Int32 field
func (field Field) AsInt32() Int32 {
	return newNumber[int32](field.expr)
}

// AsInt64 returns an Int64 field
func (field Field) AsInt64() Int64 {
	return newNumber[int64](field.expr)
}

// AsUint returns a Uint field
func (field Field) AsUint() Uint {
	return newNumber[uint](field.expr)
}

// AsUint8 returns a Uint8 field
func (field Field) AsUint8() Uint8 {
	return newNumber[uint8](field.expr)
}

// AsUint16 returns a Uint16 field
func (field Field) AsUint16() Uint16 {
	return newNumber[uint16](field.expr)
}

// AsUint32 returns a Uint32 field
func (field Field) AsUint32() Uint32 {
	return newNumber[uint32](field.expr)
}

// AsUint64 returns a Uint64 field
func (field Field) AsUint64() Uint64 {
	return newNumber[uint64](field.expr)
}

// AsFloat32 returns a Float32 field
func (field Field) AsFloat32() Float32 {
	return newNumber[float32](field.expr)
}

// AsFloat64 returns a Float64 field
func (field Field) AsFloat64() Float64 {
	return newNumber[float64](field.expr)
}
