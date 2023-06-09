package assist

import "gorm.io/gorm/clause"

// Bool boolean type field
type Bool Field

// NewBool new bool field.
func NewBool(table, column string, opts ...Option) Bool {
	return Bool{
		expr: expr{
			col: intoClauseColumn(table, column, opts...),
		},
	}
}

// IfNull use IFNULL(expr,?)
func (field Bool) IfNull(value bool) Expr {
	return field.ifNull(value)
}

// Eq equal to, use expr = ?
func (field Bool) Eq(value bool) Expr {
	return expr{
		col:       field.col,
		e:         clause.Eq{Column: field.RawExpr(), Value: value},
		buildOpts: field.buildOpts,
	}
}

// Neq not equal to, use expr <> ?
func (field Bool) Neq(value bool) Expr {
	return expr{
		col:       field.col,
		e:         clause.Neq{Column: field.RawExpr(), Value: value},
		buildOpts: field.buildOpts,
	}
}

// FindInSet use FIND_IN_SET(expr, ?)
func (field Bool) FindInSet(targetList string) Expr {
	return field.findInSet(targetList)
}

// Is use expr = ?
func (field Bool) Is(value bool) Expr {
	return field.Eq(value)
}

// Not use NOT expr
func (field Bool) Not() Expr {
	return expr{
		col:       field.col,
		e:         clause.Expr{SQL: "NOT ?", Vars: []any{field.RawExpr()}},
		buildOpts: field.buildOpts,
	}
}

// Xor use expr XOR ?
func (field Bool) Xor(value bool) Expr {
	return Bool{field.xor(value)}
}

// And use expr AND ?
func (field Bool) And(value bool) Expr {
	return Bool{field.and(value)}
}

// Or use expr OR ?
func (field Bool) Or(value bool) Expr {
	return Bool{field.or(value)}
}

// BitXor use expr expr^?
func (field Bool) BitXor(value bool) Expr {
	return Bool{field.bitXor(value)}
}

// BitAnd use expr expr&?
func (field Bool) BitAnd(value bool) Expr {
	return Bool{field.bitAnd(value)}
}

// BitOr use expr expr|?
func (field Bool) BitOr(value bool) Expr {
	return Bool{field.bitOr(value)}
}

// IntoColumns columns array with sub method
func (field Bool) IntoColumns() Columns {
	return NewColumns(field)
}
