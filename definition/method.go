package definition

import "fmt"

// ==================== Method ====================

type MethodKindID int

const (
	MethodKindID_Get MethodKindID = iota
	MethodKindID_Set
)

func (m MethodKindID) IsGet() bool {
	return m == MethodKindID_Get
}

func (m MethodKindID) IsSet() bool {
	return m == MethodKindID_Set
}

var methodKindIDStringMap = map[MethodKindID]string{
	MethodKindID_Get: "get",
	MethodKindID_Set: "set",
}

func (m MethodKindID) String() string {
	if str, ok := methodKindIDStringMap[m]; ok {
		return str
	}
	return fmt.Sprintf("MethodKindID(%d)", m)
}

// ==================== Method ====================

type Method interface {
	Position
	GetMethodKind() MethodKindID
	GetMethodName() string
	GetMethodBelongs() *NormalField
	SetMethodBelongs(*NormalField)
}

var _ Method = (*GetMethod)(nil)

type GetMethod struct {
	BasePosition

	MethodName    string
	MethodRetType Type
	MethodExpr    Expr
	MethodBelongs *NormalField
}

func (m GetMethod) GetMethodKind() MethodKindID {
	return MethodKindID_Get
}

func (m GetMethod) GetMethodName() string {
	return m.MethodName
}

func (m GetMethod) GetMethodBelongs() *NormalField {
	return m.MethodBelongs
}

func (m *GetMethod) SetMethodBelongs(belongs *NormalField) {
	m.MethodBelongs = belongs
}

type SetMethod struct {
	BasePosition

	MethodName      string
	MethodParamType Type
	MethodExpr      Expr
	MethodBelongs   *NormalField
}

func (m SetMethod) GetMethodKind() MethodKindID {
	return MethodKindID_Set
}

func (m SetMethod) GetMethodName() string {
	return m.MethodName
}

func (m SetMethod) GetMethodBelongs() *NormalField {
	return m.MethodBelongs
}

func (m *SetMethod) SetMethodBelongs(belongs *NormalField) {
	m.MethodBelongs = belongs
}

// type Method struct {
// 	BasePosition

// 	MethodKind      MethodKindID
// 	MethodName      string
// 	MethodParamType Type
// 	MethodExpr      Expr
// 	MethodBelongs   *NormalField
// }

// func (m Method) String() string {
// 	typeStr := util.IndentSpace4NoFirst(m.MethodParamType)

// 	belongsStr := "<nil>"
// 	if m.MethodBelongs != nil {
// 		belongsStr = m.MethodBelongs.FieldName
// 	}

// 	return fmt.Sprintf("Method {\n    MethodKind: %s\n    MethodName: %s\n    MethodParamType: %s\n    MethodExpr: %s\n    MethodBelongs: %s\n}", m.MethodKind, m.MethodName, typeStr, m.MethodExpr, belongsStr)
// }
