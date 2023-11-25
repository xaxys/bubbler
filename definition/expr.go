package definition

import "fmt"

type ExprOp int

const (
	ExprOp_ADD ExprOp = iota
	ExprOp_SUB
	ExprOp_MUL
	ExprOp_DIV
	ExprOp_MOD
	ExprOp_POW
	ExprOp_SHL
	ExprOp_SHR
	ExprOp_LT
	ExprOp_LE
	ExprOp_GT
	ExprOp_GE
	ExprOp_EQ
	ExprOp_NE
	ExprOp_BAND
	ExprOp_BXOR
	ExprOp_BOR
	ExprOp_AND
	ExprOp_OR
	ExprOp_NOT
	ExprOp_BNOT
)

func (e ExprOp) String() string {
	switch e {
	case ExprOp_ADD:
		return "+"
	case ExprOp_SUB:
		return "-"
	case ExprOp_MUL:
		return "*"
	case ExprOp_DIV:
		return "/"
	case ExprOp_MOD:
		return "%"
	case ExprOp_POW:
		return "**"
	case ExprOp_SHL:
		return "<<"
	case ExprOp_SHR:
		return ">>"
	case ExprOp_LT:
		return "<"
	case ExprOp_LE:
		return "<="
	case ExprOp_GT:
		return ">"
	case ExprOp_GE:
		return ">="
	case ExprOp_EQ:
		return "=="
	case ExprOp_NE:
		return "!="
	case ExprOp_BAND:
		return "&"
	case ExprOp_BXOR:
		return "^"
	case ExprOp_BOR:
		return "|"
	case ExprOp_AND:
		return "&&"
	case ExprOp_OR:
		return "||"
	case ExprOp_NOT:
		return "!"
	case ExprOp_BNOT:
		return "~"
	default:
		return ""
	}
}

// ==================== Expr ====================

type Expr interface {
	Position
	GetExprType() *BasicType
}

type UnopExpr struct {
	BasePosition
	ExprType *BasicType
	Op       ExprOp
	Expr1    Expr
}

func (u UnopExpr) GetExprType() *BasicType {
	return u.ExprType
}

func (u UnopExpr) String() string {
	return fmt.Sprintf("(%s%s)", u.Op, u.Expr1)
}

type BinopExpr struct {
	BasePosition
	ExprType *BasicType
	Op       ExprOp
	Expr1    Expr
	Expr2    Expr
}

func (b BinopExpr) GetExprType() *BasicType {
	return b.ExprType
}

func (b BinopExpr) String() string {
	return fmt.Sprintf("(%s %s %s)", b.Expr1, b.Op, b.Expr2)
}

type CastExpr struct {
	BasePosition
	ToType *BasicType
	Expr1  Expr
}

func (c CastExpr) GetExprType() *BasicType {
	return c.ToType
}

func (c CastExpr) String() string {
	return fmt.Sprintf("(%s)%s", c.ToType, c.Expr1)
}

type TenaryExpr struct {
	BasePosition
	// ExprType is the type of Expr1 or Expr2
	// Expr1 and Expr2 must be the same type
	Cond  Expr
	Expr1 Expr
	Expr2 Expr
}

func (t TenaryExpr) GetExprType() *BasicType {
	return t.Expr1.GetExprType()
}

func (t TenaryExpr) String() string {
	return fmt.Sprintf("(%s ? %s : %s)", t.Cond, t.Expr1, t.Expr2)
}

type ConstantExpr struct {
	BasePosition
	ConstantType  *BasicType
	ConstantValue Literal
}

func (c ConstantExpr) GetExprType() *BasicType {
	return c.ConstantType
}

func (c ConstantExpr) String() string {
	return fmt.Sprintf("%s", c.ConstantValue)
}

type ValueExpr struct {
	BasePosition
	ValueType *BasicType
}

func (v ValueExpr) GetExprType() *BasicType {
	return v.ValueType
}

func (v ValueExpr) String() string {
	return "value"
}
