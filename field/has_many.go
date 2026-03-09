package field

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type HasManyField[T any] struct {
	DB *gorm.DB

	RelationField
}

func NewHasManyField[T any](db *gorm.DB, fieldName string, fieldType string, relations ...Relation) *HasManyField[T] {
	return &HasManyField[T]{
		DB:            db.Session(&gorm.Session{}),
		RelationField: NewRelation(fieldName, fieldType, relations...),
	}
}

func (a HasManyField[T]) Clone(db *gorm.DB) *HasManyField[T] {
	a.DB = db.Session(&gorm.Session{Initialized: true})
	a.DB.Statement.ConnPool = db.Statement.ConnPool
	return &a
}

func (a HasManyField[T]) Where(conds ...Expr) *HasManyField[T] {
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

func (a HasManyField[T]) WithContext(ctx context.Context) *HasManyField[T] {
	a.DB = a.DB.WithContext(ctx)
	return &a
}

func (a HasManyField[T]) Session(session *gorm.Session) *HasManyField[T] {
	a.DB = a.DB.Session(session)
	return &a
}

func (a HasManyField[T]) Model(m *T) *HasManyFieldTx[T] {
	return &HasManyFieldTx[T]{tx: a.DB.Model(m).Association(a.Name())}
}

type HasManyFieldTx[T any] struct {
	tx *gorm.Association
}

func (a HasManyFieldTx[T]) Find() (result []*T, err error) {
	return result, a.tx.Find(&result)
}

func (a HasManyFieldTx[T]) Append(values ...*T) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a HasManyFieldTx[T]) Replace(values ...*T) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a HasManyFieldTx[T]) Delete(values ...*T) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a HasManyFieldTx[T]) Clear() error {
	return a.tx.Clear()
}

func (a HasManyFieldTx[T]) Count() int64 {
	return a.tx.Count()
}

func (a HasManyFieldTx[T]) Unscoped() *HasManyFieldTx[T] {
	a.tx = a.tx.Unscoped()
	return &a
}
