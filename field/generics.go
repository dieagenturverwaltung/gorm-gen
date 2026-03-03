package field

import (
	"gorm.io/gorm/clause"
)

// newGenerics create new generic field type
func newGenerics[T any](e expr) genericsField[T] {
	return genericsField[T]{e}
}

// genericsField a generics field struct
// serving as a base field type, offers a suite of fundamental methods/functions for database operations."
type genericsField[T any] struct{ expr }

// Eq judge equal
func (field genericsField[T]) Eq(value T) Expr {
	return expr{e: clause.Eq{Column: field.RawExpr(), Value: value}}
}

// Neq judge not equal
func (field genericsField[T]) Neq(value T) Expr {
	return expr{e: clause.Neq{Column: field.RawExpr(), Value: value}}
}

// In ...
func (field genericsField[T]) In(values ...T) Expr {
	return expr{e: clause.IN{Column: field.RawExpr(), Values: field.toSlice(values...)}}
}

// NotIn ...
func (field genericsField[T]) NotIn(values ...T) Expr {
	return expr{e: clause.Not(field.In(values...).expression())}
}

// Gt ...
func (field genericsField[T]) Gt(value T) Expr {
	return expr{e: clause.Gt{Column: field.RawExpr(), Value: value}}
}

// Gte ...
func (field genericsField[T]) Gte(value T) Expr {
	return expr{e: clause.Gte{Column: field.RawExpr(), Value: value}}
}

// Lt ...
func (field genericsField[T]) Lt(value T) Expr {
	return expr{e: clause.Lt{Column: field.RawExpr(), Value: value}}
}

// Lte ...
func (field genericsField[T]) Lte(value T) Expr {
	return expr{e: clause.Lte{Column: field.RawExpr(), Value: value}}
}

// Like ...
func (field genericsField[T]) Like(value string) Expr {
	return expr{e: clause.Like{Column: field.RawExpr(), Value: value}}
}

// NotLike ...
func (field genericsField[T]) NotLike(value string) Expr {
	return expr{e: clause.Not(field.Like(value).expression())}
}

// Value ...
func (field genericsField[T]) Value(value T) AssignExpr {
	return field.value(value)
}

// Sum ...
func (field genericsField[T]) Sum() genericsField[T] {
	return genericsField[T]{field.sum()}
}

// IfNull ...
func (field genericsField[T]) IfNull(value T) Expr {
	return field.ifNull(value)
}

// Field ...
func (field genericsField[T]) Field(value []interface{}) Expr {
	return field.field(value)
}

func (field genericsField[T]) toSlice(values ...T) []interface{} {
	slice := make([]interface{}, len(values))
	for i, v := range values {
		slice[i] = v
	}
	return slice
}

// AsString returns a String field
func (field genericsField[T]) AsString() String {
	return newChars[string](field.expr)
}

// AsBytes returns a Bytes field
func (field genericsField[T]) AsBytes() Bytes {
	return newChars[[]byte](field.expr)
}

// AsInt returns an Int field
func (field genericsField[T]) AsInt() Int {
	return newNumber[int](field.expr)
}

// AsInt8 returns an Int8 field
func (field genericsField[T]) AsInt8() Int8 {
	return newNumber[int8](field.expr)
}

// AsInt16 returns an Int16 field
func (field genericsField[T]) AsInt16() Int16 {
	return newNumber[int16](field.expr)
}

// AsInt32 returns an Int32 field
func (field genericsField[T]) AsInt32() Int32 {
	return newNumber[int32](field.expr)
}

// AsInt64 returns an Int64 field
func (field genericsField[T]) AsInt64() Int64 {
	return newNumber[int64](field.expr)
}

// AsUint returns a Uint field
func (field genericsField[T]) AsUint() Uint {
	return newNumber[uint](field.expr)
}

// AsUint8 returns a Uint8 field
func (field genericsField[T]) AsUint8() Uint8 {
	return newNumber[uint8](field.expr)
}

// AsUint16 returns a Uint16 field
func (field genericsField[T]) AsUint16() Uint16 {
	return newNumber[uint16](field.expr)
}

// AsUint32 returns a Uint32 field
func (field genericsField[T]) AsUint32() Uint32 {
	return newNumber[uint32](field.expr)
}

// AsUint64 returns a Uint64 field
func (field genericsField[T]) AsUint64() Uint64 {
	return newNumber[uint64](field.expr)
}

// AsFloat32 returns a Float32 field
func (field genericsField[T]) AsFloat32() Float32 {
	return newNumber[float32](field.expr)
}

// AsFloat64 returns a Float64 field
func (field genericsField[T]) AsFloat64() Float64 {
	return newNumber[float64](field.expr)
}
