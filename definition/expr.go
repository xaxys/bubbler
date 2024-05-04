package definition

import "fmt"

type ExprOp int

const (
	ExprOp_ADD    ExprOp = iota // +
	ExprOp_SUB                  // -
	ExprOp_MUL                  // *
	ExprOp_DIV                  // /
	ExprOp_MOD                  // %
	ExprOp_POW                  // **
	ExprOp_SHL                  // <<
	ExprOp_SHR                  // >>
	ExprOp_LT                   // <
	ExprOp_LE                   // <=
	ExprOp_GT                   // >
	ExprOp_GE                   // >=
	ExprOp_EQ                   // ==
	ExprOp_NE                   // !=
	ExprOp_BAND                 // &
	ExprOp_BXOR                 // ^
	ExprOp_BOR                  // |
	ExprOp_AND                  // &&
	ExprOp_OR                   // ||
	ExprOp_NOT                  // !
	ExprOp_BNOT                 // ~
	ExprOp_ASSIGN               // =
)

func (e ExprOp) IsAdd() bool {
	return e == ExprOp_ADD
}

func (e ExprOp) IsSub() bool {
	return e == ExprOp_SUB
}

func (e ExprOp) IsMul() bool {
	return e == ExprOp_MUL
}

func (e ExprOp) IsDiv() bool {
	return e == ExprOp_DIV
}

func (e ExprOp) IsMod() bool {
	return e == ExprOp_MOD
}

func (e ExprOp) IsPow() bool {
	return e == ExprOp_POW
}

func (e ExprOp) IsShl() bool {
	return e == ExprOp_SHL
}

func (e ExprOp) IsShr() bool {
	return e == ExprOp_SHR
}

func (e ExprOp) IsLt() bool {
	return e == ExprOp_LT
}

func (e ExprOp) IsLe() bool {
	return e == ExprOp_LE
}

func (e ExprOp) IsGt() bool {
	return e == ExprOp_GT
}

func (e ExprOp) IsGe() bool {
	return e == ExprOp_GE
}

func (e ExprOp) IsEq() bool {
	return e == ExprOp_EQ
}

func (e ExprOp) IsNe() bool {
	return e == ExprOp_NE
}

func (e ExprOp) IsBand() bool {
	return e == ExprOp_BAND
}

func (e ExprOp) IsBxor() bool {
	return e == ExprOp_BXOR
}

func (e ExprOp) IsBor() bool {
	return e == ExprOp_BOR
}

func (e ExprOp) IsAnd() bool {
	return e == ExprOp_AND
}

func (e ExprOp) IsOr() bool {
	return e == ExprOp_OR
}

func (e ExprOp) IsNot() bool {
	return e == ExprOp_NOT
}

func (e ExprOp) IsBnot() bool {
	return e == ExprOp_BNOT
}

func (e ExprOp) IsAssign() bool {
	return e == ExprOp_ASSIGN
}

var exprOpToString = map[ExprOp]string{
	ExprOp_ADD:    "+",
	ExprOp_SUB:    "-",
	ExprOp_MUL:    "*",
	ExprOp_DIV:    "/",
	ExprOp_MOD:    "%",
	ExprOp_POW:    "**",
	ExprOp_SHL:    "<<",
	ExprOp_SHR:    ">>",
	ExprOp_LT:     "<",
	ExprOp_LE:     "<=",
	ExprOp_GT:     ">",
	ExprOp_GE:     ">=",
	ExprOp_EQ:     "==",
	ExprOp_NE:     "!=",
	ExprOp_BAND:   "&",
	ExprOp_BXOR:   "^",
	ExprOp_BOR:    "|",
	ExprOp_AND:    "&&",
	ExprOp_OR:     "||",
	ExprOp_NOT:    "!",
	ExprOp_BNOT:   "~",
	ExprOp_ASSIGN: "=",
}

func (e ExprOp) String() string {
	if opString, ok := exprOpToString[e]; ok {
		return opString
	}
	return "unknown"
}

// ==================== Expr ====================

type ExprKindID int

const (
	ExprKindID_UnopExpr ExprKindID = iota
	ExprKindID_BinopExpr
	ExprKindID_CastExpr
	ExprKindID_TenaryExpr
	ExprKindID_ConstantExpr
	ExprKindID_ValueExpr
	ExprKindID_RawExpr
)

func (e ExprKindID) IsUnopExpr() bool {
	return e == ExprKindID_UnopExpr
}

func (e ExprKindID) IsBinopExpr() bool {
	return e == ExprKindID_BinopExpr
}

func (e ExprKindID) IsCastExpr() bool {
	return e == ExprKindID_CastExpr
}

func (e ExprKindID) IsTenaryExpr() bool {
	return e == ExprKindID_TenaryExpr
}

func (e ExprKindID) IsConstantExpr() bool {
	return e == ExprKindID_ConstantExpr
}

func (e ExprKindID) IsValueExpr() bool {
	return e == ExprKindID_ValueExpr
}

func (e ExprKindID) IsRawExpr() bool {
	return e == ExprKindID_RawExpr
}

type Expr interface {
	Position
	GetExprKind() ExprKindID
	GetExprType() *BasicType
}

type UnopExpr struct {
	BasePosition

	ExprType *BasicType
	Op       ExprOp
	Expr1    Expr
}

func (u UnopExpr) GetExprKind() ExprKindID {
	return ExprKindID_UnopExpr
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

func (b BinopExpr) GetExprKind() ExprKindID {
	return ExprKindID_BinopExpr
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

func (c CastExpr) GetExprKind() ExprKindID {
	return ExprKindID_CastExpr
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

func (t TenaryExpr) GetExprKind() ExprKindID {
	return ExprKindID_TenaryExpr
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

func (c ConstantExpr) GetExprKind() ExprKindID {
	return ExprKindID_ConstantExpr
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

func (v ValueExpr) GetExprKind() ExprKindID {
	return ExprKindID_ValueExpr
}

func (v ValueExpr) GetExprType() *BasicType {
	return v.ValueType
}

func (v ValueExpr) String() string {
	return "value"
}

// currently for internal use only

type RawExpr struct {
	BasePosition

	ExprType *BasicType
	Expr     string
}

func (r RawExpr) GetExprKind() ExprKindID {
	return ExprKindID_RawExpr
}

func (r RawExpr) GetExprType() *BasicType {
	return r.ExprType
}

func (r RawExpr) String() string {
	return r.Expr
}
