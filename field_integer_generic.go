package assist

import (
	"strings"

	"golang.org/x/exp/constraints"
	"gorm.io/gorm/clause"
)

// Integer type field
type Integer[T constraints.Integer] Field

// NewInt new Integer
func NewInteger[T constraints.Integer](table, column string, opts ...Option) Integer[T] {
	return Integer[T]{expr: expr{col: intoColumn(table, column, opts...)}}
}

// IfNull use IFNULL(expr,?)
func (field Integer[T]) IfNull(value T) Expr {
	return field.ifNull(value)
}

// Eq equal to, use expr = ?
func (field Integer[T]) Eq(value T) Expr {
	return expr{e: clause.Eq{Column: field.RawExpr(), Value: value}}
}

// Neq not equal to, use expr <> ?
func (field Integer[T]) Neq(value T) Expr {
	return expr{e: clause.Neq{Column: field.RawExpr(), Value: value}}
}

// Gt greater than, use expr > ?
func (field Integer[T]) Gt(value T) Expr {
	return expr{e: clause.Gt{Column: field.RawExpr(), Value: value}}
}

// Gte greater or equal to, use expr >= ?
func (field Integer[T]) Gte(value T) Expr {
	return expr{e: clause.Gte{Column: field.RawExpr(), Value: value}}
}

// Lt less than, use expr < ?
func (field Integer[T]) Lt(value T) Expr {
	return expr{e: clause.Lt{Column: field.RawExpr(), Value: value}}
}

// Lte less or equal to, use expr <= ?
func (field Integer[T]) Lte(value T) Expr {
	return expr{e: clause.Lte{Column: field.RawExpr(), Value: value}}
}

// Between use expr BETWEEN ? AND ?
func (field Integer[T]) Between(left T, right T) Expr {
	return field.between([]any{left, right})
}

// NotBetween use NOT (expr BETWEEN ? AND ?)
func (field Integer[T]) NotBetween(left T, right T) Expr {
	return field.notBetween([]any{left, right})
}

// In use expr IN (?)
func (field Integer[T]) In(values ...T) Expr {
	return expr{e: clause.IN{Column: field.RawExpr(), Values: intoSlice(values...)}}
}

// NotIn use expr NOT IN (?)
func (field Integer[T]) NotIn(values ...T) Expr {
	return expr{e: clause.Not(clause.IN{Column: field.RawExpr(), Values: intoSlice(values...)})}
}

// Like use expr LIKE ?
func (field Integer[T]) Like(value T) Expr {
	return expr{e: clause.Like{Column: field.RawExpr(), Value: value}}
}

// NotLike use expr NOT LIKE ?
func (field Integer[T]) NotLike(value T) Expr {
	return expr{e: clause.Not(clause.Like{Column: field.RawExpr(), Value: value})}
}

// Sum use SUM(expr)
func (field Integer[T]) Sum() Integer[T] {
	return Integer[T]{field.sum()}
}

// Add use expr+?
func (field Integer[T]) Add(value T) Integer[T] {
	return Integer[T]{field.add(value)}
}

// Add use expr-?
func (field Integer[T]) Sub(value T) Integer[T] {
	return Integer[T]{field.sub(value)}
}

// Mul use expr*?
func (field Integer[T]) Mul(value T) Integer[T] {
	return Integer[T]{field.mul(value)}
}

// Div use expr/?
func (field Integer[T]) Div(value T) Integer[T] {
	return Integer[T]{field.div(value)}
}

// Mod use expr%?
func (field Integer[T]) Mod(value T) Integer[T] {
	return Integer[T]{field.mod(value)}
}

// FloorDiv use expr DIV ?
func (field Integer[T]) FloorDiv(value T) Integer[T] {
	return Integer[T]{field.floorDiv(value)}
}

// Round use ROUND(expr, ?)
func (field Integer[T]) Round(value int) Integer[T] {
	return Integer[T]{field.round(value)}
}

// RightShift use expr>>?
func (field Integer[T]) RightShift(value T) Integer[T] {
	return Integer[T]{field.rightShift(value)}
}

// LeftShift use expr<<?
func (field Integer[T]) LeftShift(value T) Integer[T] {
	return Integer[T]{field.leftShift(value)}
}

// BitXor use expr expr^?
func (field Integer[T]) BitXor(value T) Integer[T] {
	return Integer[T]{field.bitXor(value)}
}

// BitAnd use expr expr&?
func (field Integer[T]) BitAnd(value T) Integer[T] {
	return Integer[T]{field.bitAnd(value)}
}

// BitOr use expr expr|?
func (field Integer[T]) BitOr(value T) Integer[T] {
	return Integer[T]{field.bitOr(value)}
}

// BitFlip use expr ~expr
func (field Integer[T]) BitFlip() Integer[T] {
	return Integer[T]{field.bitFlip()}
}

// FromUnixTime use FromUnixTime(unix_timestamp[, format])
func (field Integer[T]) FromUnixTime(format ...string) String {
	if len(format) > 0 && strings.TrimSpace(format[0]) != "" {
		return String{expr{e: clause.Expr{SQL: "FROM_UNIXTIME(?, ?)", Vars: []any{field.RawExpr(), format[0]}}}}
	}
	return String{expr{e: clause.Expr{SQL: "FROM_UNIXTIME(?)", Vars: []any{field.RawExpr()}}}}
}

// FromDays use FROM_DAYS(value)
func (field Integer[T]) FromDays() Time {
	return Time{expr{e: clause.Expr{SQL: "FROM_DAYS(?)", Vars: []any{field.RawExpr()}}}}
}

// IntoColumns columns array with sub method
func (field Integer[T]) IntoColumns() Columns {
	return NewColumns(field)
}
