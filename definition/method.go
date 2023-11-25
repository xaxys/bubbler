package definition

import (
	"fmt"

	"github.com/xaxys/bubbler/util"
)

// ==================== Method ====================

type MethodKindID int

const (
	MethodKindID_Get MethodKindID = iota
	MethodKindID_Set
)

func (m MethodKindID) String() string {
	switch m {
	case MethodKindID_Get:
		return "get"
	case MethodKindID_Set:
		return "set"
	}
	return ""
}

type Method struct {
	BasePosition
	MethodKind      MethodKindID
	MethodName      string
	MethodParamType Type
	MethodExpr      Expr
	MethodBelongs   *NormalField // TODO: fix relation
}

func (m Method) String() string {
	typeStr := util.IndentSpace4NoFirst(m.MethodParamType)

	belongsStr := "<nil>"
	if m.MethodBelongs != nil {
		belongsStr = m.MethodBelongs.FieldName
	}

	return fmt.Sprintf("Method {\n    MethodKind: %s\n    MethodName: %s\n    MethodParamType: %s\n    MethodExpr: %s\n    MethodBelongs: %s\n}", m.MethodKind, m.MethodName, typeStr, m.MethodExpr, belongsStr)
}
