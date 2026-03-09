package field

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type HasOneField[T any] struct {
	DB *gorm.DB

	RelationField
}

func NewHasOneField[T any](db *gorm.DB, fieldName string, fieldType string, relations ...Relation) *HasOneField[T] {
	return &HasOneField[T]{
		DB:            db.Session(&gorm.Session{}),
		RelationField: NewRelation(fieldName, fieldType, relations...),
	}
}

func (a HasOneField[T]) Clone(db *gorm.DB) *HasOneField[T] {
	a.DB = db.Session(&gorm.Session{Initialized: true})
	a.DB.Statement.ConnPool = db.Statement.ConnPool
	return &a
}

func (a HasOneField[T]) Where(conds ...Expr) *HasOneField[T] {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.DB = a.DB.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a HasOneField[T]) WithContext(ctx context.Context) *HasOneField[T] {
	a.DB = a.DB.WithContext(ctx)
	return &a
}

func (a HasOneField[T]) Session(session *gorm.Session) *HasOneField[T] {
	a.DB = a.DB.Session(session)
	return &a
}

func (a HasOneField[T]) Model(m *T) *HasOneFieldTx[T] {
	return &HasOneFieldTx[T]{tx: a.DB.Model(m).Association(a.Name())}
}

type HasOneFieldTx[T any] struct {
	tx *gorm.Association
}

func (a HasOneFieldTx[T]) Find() (result *T, err error) {
	return result, a.tx.Find(&result)
}

func (a HasOneFieldTx[T]) Append(values ...*T) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a HasOneFieldTx[T]) Replace(values ...*T) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a HasOneFieldTx[T]) Delete(values ...*T) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a HasOneFieldTx[T]) Clear() error {
	return a.tx.Clear()
}

func (a HasOneFieldTx[T]) Count() int64 {
	return a.tx.Count()
}

func (a HasOneFieldTx[T]) Unscoped() *HasOneFieldTx[T] {
	a.tx = a.tx.Unscoped()
	return &a
}
