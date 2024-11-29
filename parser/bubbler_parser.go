// Code generated from bubbler.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // bubbler
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type bubblerParser struct {
	*antlr.BaseParser
}

var BubblerParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func bubblerParserInit() {
	staticData := &BubblerParserStaticData
	staticData.LiteralNames = []string{
		"", "'syntax'", "'import'", "'get'", "'set'", "'value'", "'package'",
		"'option'", "'void'", "'int8'", "'int16'", "'int32'", "'int64'", "'uint8'",
		"'uint16'", "'uint32'", "'uint64'", "'float32'", "'float64'", "'bool'",
		"'string'", "'bytes'", "'enum'", "'struct'", "'#'", "';'", "'='", "'?'",
		"'('", "')'", "'['", "']'", "'{'", "'}'", "'<'", "'<='", "'>'", "'>='",
		"'=='", "'!='", "'.'", "','", "':'", "'+'", "'-'", "'*'", "'/'", "'%'",
		"'**'", "'<<'", "'>>'", "'&'", "'|'", "'^'", "'~'", "'&&'", "'||'",
		"'!'",
	}
	staticData.SymbolicNames = []string{
		"", "SYNTAX", "IMPORT", "GET", "SET", "VALUE", "PACKAGE", "OPTION",
		"VOID", "INT8", "INT16", "INT32", "INT64", "UINT8", "UINT16", "UINT32",
		"UINT64", "FLOAT32", "FLOAT64", "BOOL", "STRING", "BYTES", "ENUM", "STRUCT",
		"HASH", "SEMI", "ASSIGN", "QUESTION", "LP", "RP", "LB", "RB", "LC",
		"RC", "LT", "LE", "GT", "GE", "EQ", "NE", "DOT", "COMMA", "COLON", "ADD",
		"SUB", "MUL", "DIV", "MOD", "POW", "SHL", "SHR", "BAND", "BOR", "BXOR",
		"BNOT", "AND", "OR", "NOT", "STR_LIT", "BOOL_LIT", "FLOAT_LIT", "INT_LIT",
		"IDENTIFIER", "WS", "LINE_COMMENT", "COMMENT", "KEYWORDS",
	}
	staticData.RuleNames = []string{
		"proto", "importStatement", "topLevelDef", "size_", "byteSize", "bitSize",
		"packageStatement", "optionStatement", "optionName", "field", "fieldVoid",
		"fieldConstant", "fieldEmbedded", "fieldNormal", "fieldOptions", "fieldOption",
		"fieldMethods", "fieldMethod", "type_", "basicType", "arrayElementType",
		"arrayType", "enumDef", "enumBody", "enumElement", "enumValue", "enumValueOptions",
		"enumValueOption", "structDef", "structBody", "structElement", "expr",
		"value", "constant", "emptyStatement_", "ident", "fullIdent", "fieldName",
		"methodName", "structName", "enumName", "enumValueName", "structType",
		"enumType", "intLit", "strLit", "boolLit", "floatLit",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 66, 429, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 102, 8, 0, 10, 0, 12, 0, 105,
		9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 3, 2, 115, 8, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 121, 8, 3, 1, 3, 1, 3, 3, 3, 125, 8, 3, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 149, 8,
		9, 1, 10, 1, 10, 1, 10, 3, 10, 154, 8, 10, 1, 10, 1, 10, 1, 11, 1, 11,
		3, 11, 160, 8, 11, 1, 11, 3, 11, 163, 8, 11, 1, 11, 1, 11, 1, 11, 3, 11,
		168, 8, 11, 1, 11, 3, 11, 171, 8, 11, 1, 11, 1, 11, 1, 12, 1, 12, 3, 12,
		177, 8, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 3, 13, 184, 8, 13, 1, 13,
		3, 13, 187, 8, 13, 1, 13, 3, 13, 190, 8, 13, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 14, 1, 14, 5, 14, 198, 8, 14, 10, 14, 12, 14, 201, 9, 14, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 5, 16, 211, 8, 16, 10, 16,
		12, 16, 214, 9, 16, 1, 16, 1, 16, 1, 17, 1, 17, 3, 17, 220, 8, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 3, 18, 236, 8, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 3, 20, 246, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 5, 23, 259, 8, 23, 10, 23,
		12, 23, 262, 9, 23, 1, 23, 1, 23, 1, 24, 1, 24, 3, 24, 268, 8, 24, 1, 25,
		1, 25, 1, 25, 1, 25, 3, 25, 274, 8, 25, 3, 25, 276, 8, 25, 1, 25, 3, 25,
		279, 8, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 5, 26, 287, 8, 26,
		10, 26, 12, 26, 290, 9, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		28, 1, 28, 3, 28, 300, 8, 28, 1, 28, 1, 28, 1, 29, 1, 29, 5, 29, 306, 8,
		29, 10, 29, 12, 29, 309, 9, 29, 1, 29, 1, 29, 1, 30, 1, 30, 3, 30, 315,
		8, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 331, 8, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 5, 31, 372, 8, 31, 10, 31, 12, 31, 375,
		9, 31, 1, 32, 1, 32, 1, 33, 3, 33, 380, 8, 33, 1, 33, 1, 33, 3, 33, 384,
		8, 33, 1, 33, 1, 33, 1, 33, 3, 33, 389, 8, 33, 1, 34, 1, 34, 1, 35, 1,
		35, 1, 36, 1, 36, 1, 36, 5, 36, 398, 8, 36, 10, 36, 12, 36, 401, 9, 36,
		1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1,
		41, 1, 41, 1, 42, 1, 42, 3, 42, 417, 8, 42, 1, 43, 1, 43, 1, 44, 1, 44,
		1, 45, 1, 45, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 0, 1, 62, 48, 0, 2, 4,
		6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78,
		80, 82, 84, 86, 88, 90, 92, 94, 0, 9, 1, 0, 3, 4, 1, 0, 9, 19, 2, 0, 25,
		25, 41, 41, 3, 0, 43, 44, 54, 54, 57, 57, 1, 0, 45, 47, 1, 0, 43, 44, 1,
		0, 49, 50, 1, 0, 34, 37, 1, 0, 38, 39, 446, 0, 103, 1, 0, 0, 0, 2, 108,
		1, 0, 0, 0, 4, 114, 1, 0, 0, 0, 6, 116, 1, 0, 0, 0, 8, 128, 1, 0, 0, 0,
		10, 130, 1, 0, 0, 0, 12, 132, 1, 0, 0, 0, 14, 136, 1, 0, 0, 0, 16, 142,
		1, 0, 0, 0, 18, 148, 1, 0, 0, 0, 20, 150, 1, 0, 0, 0, 22, 157, 1, 0, 0,
		0, 24, 174, 1, 0, 0, 0, 26, 180, 1, 0, 0, 0, 28, 193, 1, 0, 0, 0, 30, 204,
		1, 0, 0, 0, 32, 208, 1, 0, 0, 0, 34, 217, 1, 0, 0, 0, 36, 235, 1, 0, 0,
		0, 38, 237, 1, 0, 0, 0, 40, 245, 1, 0, 0, 0, 42, 247, 1, 0, 0, 0, 44, 252,
		1, 0, 0, 0, 46, 256, 1, 0, 0, 0, 48, 267, 1, 0, 0, 0, 50, 269, 1, 0, 0,
		0, 52, 282, 1, 0, 0, 0, 54, 293, 1, 0, 0, 0, 56, 297, 1, 0, 0, 0, 58, 303,
		1, 0, 0, 0, 60, 314, 1, 0, 0, 0, 62, 330, 1, 0, 0, 0, 64, 376, 1, 0, 0,
		0, 66, 388, 1, 0, 0, 0, 68, 390, 1, 0, 0, 0, 70, 392, 1, 0, 0, 0, 72, 394,
		1, 0, 0, 0, 74, 402, 1, 0, 0, 0, 76, 404, 1, 0, 0, 0, 78, 406, 1, 0, 0,
		0, 80, 409, 1, 0, 0, 0, 82, 412, 1, 0, 0, 0, 84, 416, 1, 0, 0, 0, 86, 418,
		1, 0, 0, 0, 88, 420, 1, 0, 0, 0, 90, 422, 1, 0, 0, 0, 92, 424, 1, 0, 0,
		0, 94, 426, 1, 0, 0, 0, 96, 102, 3, 2, 1, 0, 97, 102, 3, 12, 6, 0, 98,
		102, 3, 14, 7, 0, 99, 102, 3, 4, 2, 0, 100, 102, 3, 68, 34, 0, 101, 96,
		1, 0, 0, 0, 101, 97, 1, 0, 0, 0, 101, 98, 1, 0, 0, 0, 101, 99, 1, 0, 0,
		0, 101, 100, 1, 0, 0, 0, 102, 105, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103,
		104, 1, 0, 0, 0, 104, 106, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 106, 107,
		5, 0, 0, 1, 107, 1, 1, 0, 0, 0, 108, 109, 5, 2, 0, 0, 109, 110, 3, 90,
		45, 0, 110, 111, 5, 25, 0, 0, 111, 3, 1, 0, 0, 0, 112, 115, 3, 44, 22,
		0, 113, 115, 3, 56, 28, 0, 114, 112, 1, 0, 0, 0, 114, 113, 1, 0, 0, 0,
		115, 5, 1, 0, 0, 0, 116, 124, 5, 30, 0, 0, 117, 120, 3, 8, 4, 0, 118, 119,
		5, 24, 0, 0, 119, 121, 3, 10, 5, 0, 120, 118, 1, 0, 0, 0, 120, 121, 1,
		0, 0, 0, 121, 125, 1, 0, 0, 0, 122, 123, 5, 24, 0, 0, 123, 125, 3, 10,
		5, 0, 124, 117, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0,
		126, 127, 5, 31, 0, 0, 127, 7, 1, 0, 0, 0, 128, 129, 3, 88, 44, 0, 129,
		9, 1, 0, 0, 0, 130, 131, 3, 88, 44, 0, 131, 11, 1, 0, 0, 0, 132, 133, 5,
		6, 0, 0, 133, 134, 3, 72, 36, 0, 134, 135, 5, 25, 0, 0, 135, 13, 1, 0,
		0, 0, 136, 137, 5, 7, 0, 0, 137, 138, 3, 16, 8, 0, 138, 139, 5, 26, 0,
		0, 139, 140, 3, 66, 33, 0, 140, 141, 5, 25, 0, 0, 141, 15, 1, 0, 0, 0,
		142, 143, 3, 70, 35, 0, 143, 17, 1, 0, 0, 0, 144, 149, 3, 20, 10, 0, 145,
		149, 3, 22, 11, 0, 146, 149, 3, 24, 12, 0, 147, 149, 3, 26, 13, 0, 148,
		144, 1, 0, 0, 0, 148, 145, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148, 147,
		1, 0, 0, 0, 149, 19, 1, 0, 0, 0, 150, 151, 5, 8, 0, 0, 151, 153, 3, 6,
		3, 0, 152, 154, 3, 28, 14, 0, 153, 152, 1, 0, 0, 0, 153, 154, 1, 0, 0,
		0, 154, 155, 1, 0, 0, 0, 155, 156, 5, 25, 0, 0, 156, 21, 1, 0, 0, 0, 157,
		159, 3, 38, 19, 0, 158, 160, 3, 74, 37, 0, 159, 158, 1, 0, 0, 0, 159, 160,
		1, 0, 0, 0, 160, 162, 1, 0, 0, 0, 161, 163, 3, 6, 3, 0, 162, 161, 1, 0,
		0, 0, 162, 163, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 167, 5, 26, 0, 0,
		165, 168, 3, 66, 33, 0, 166, 168, 3, 70, 35, 0, 167, 165, 1, 0, 0, 0, 167,
		166, 1, 0, 0, 0, 168, 170, 1, 0, 0, 0, 169, 171, 3, 28, 14, 0, 170, 169,
		1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 173, 5, 25,
		0, 0, 173, 23, 1, 0, 0, 0, 174, 176, 3, 36, 18, 0, 175, 177, 3, 28, 14,
		0, 176, 175, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178,
		179, 5, 25, 0, 0, 179, 25, 1, 0, 0, 0, 180, 181, 3, 36, 18, 0, 181, 183,
		3, 74, 37, 0, 182, 184, 3, 6, 3, 0, 183, 182, 1, 0, 0, 0, 183, 184, 1,
		0, 0, 0, 184, 186, 1, 0, 0, 0, 185, 187, 3, 28, 14, 0, 186, 185, 1, 0,
		0, 0, 186, 187, 1, 0, 0, 0, 187, 189, 1, 0, 0, 0, 188, 190, 3, 32, 16,
		0, 189, 188, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191,
		192, 5, 25, 0, 0, 192, 27, 1, 0, 0, 0, 193, 194, 5, 30, 0, 0, 194, 199,
		3, 30, 15, 0, 195, 196, 5, 41, 0, 0, 196, 198, 3, 30, 15, 0, 197, 195,
		1, 0, 0, 0, 198, 201, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 199, 200, 1, 0,
		0, 0, 200, 202, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 202, 203, 5, 31, 0, 0,
		203, 29, 1, 0, 0, 0, 204, 205, 3, 16, 8, 0, 205, 206, 5, 26, 0, 0, 206,
		207, 3, 66, 33, 0, 207, 31, 1, 0, 0, 0, 208, 212, 5, 32, 0, 0, 209, 211,
		3, 34, 17, 0, 210, 209, 1, 0, 0, 0, 211, 214, 1, 0, 0, 0, 212, 210, 1,
		0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 215, 1, 0, 0, 0, 214, 212, 1, 0, 0,
		0, 215, 216, 5, 33, 0, 0, 216, 33, 1, 0, 0, 0, 217, 219, 7, 0, 0, 0, 218,
		220, 3, 76, 38, 0, 219, 218, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 221,
		1, 0, 0, 0, 221, 222, 5, 28, 0, 0, 222, 223, 3, 38, 19, 0, 223, 224, 5,
		29, 0, 0, 224, 225, 5, 42, 0, 0, 225, 226, 3, 62, 31, 0, 226, 227, 5, 25,
		0, 0, 227, 35, 1, 0, 0, 0, 228, 236, 3, 38, 19, 0, 229, 236, 5, 20, 0,
		0, 230, 236, 5, 21, 0, 0, 231, 236, 3, 42, 21, 0, 232, 236, 3, 84, 42,
		0, 233, 236, 3, 86, 43, 0, 234, 236, 3, 70, 35, 0, 235, 228, 1, 0, 0, 0,
		235, 229, 1, 0, 0, 0, 235, 230, 1, 0, 0, 0, 235, 231, 1, 0, 0, 0, 235,
		232, 1, 0, 0, 0, 235, 233, 1, 0, 0, 0, 235, 234, 1, 0, 0, 0, 236, 37, 1,
		0, 0, 0, 237, 238, 7, 1, 0, 0, 238, 39, 1, 0, 0, 0, 239, 246, 3, 38, 19,
		0, 240, 246, 5, 20, 0, 0, 241, 246, 5, 21, 0, 0, 242, 246, 3, 84, 42, 0,
		243, 246, 3, 86, 43, 0, 244, 246, 3, 70, 35, 0, 245, 239, 1, 0, 0, 0, 245,
		240, 1, 0, 0, 0, 245, 241, 1, 0, 0, 0, 245, 242, 1, 0, 0, 0, 245, 243,
		1, 0, 0, 0, 245, 244, 1, 0, 0, 0, 246, 41, 1, 0, 0, 0, 247, 248, 3, 40,
		20, 0, 248, 249, 5, 34, 0, 0, 249, 250, 3, 88, 44, 0, 250, 251, 5, 36,
		0, 0, 251, 43, 1, 0, 0, 0, 252, 253, 3, 80, 40, 0, 253, 254, 3, 6, 3, 0,
		254, 255, 3, 46, 23, 0, 255, 45, 1, 0, 0, 0, 256, 260, 5, 32, 0, 0, 257,
		259, 3, 48, 24, 0, 258, 257, 1, 0, 0, 0, 259, 262, 1, 0, 0, 0, 260, 258,
		1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 261, 263, 1, 0, 0, 0, 262, 260, 1, 0,
		0, 0, 263, 264, 5, 33, 0, 0, 264, 47, 1, 0, 0, 0, 265, 268, 3, 50, 25,
		0, 266, 268, 3, 68, 34, 0, 267, 265, 1, 0, 0, 0, 267, 266, 1, 0, 0, 0,
		268, 49, 1, 0, 0, 0, 269, 275, 3, 82, 41, 0, 270, 273, 5, 26, 0, 0, 271,
		274, 3, 66, 33, 0, 272, 274, 3, 70, 35, 0, 273, 271, 1, 0, 0, 0, 273, 272,
		1, 0, 0, 0, 274, 276, 1, 0, 0, 0, 275, 270, 1, 0, 0, 0, 275, 276, 1, 0,
		0, 0, 276, 278, 1, 0, 0, 0, 277, 279, 3, 52, 26, 0, 278, 277, 1, 0, 0,
		0, 278, 279, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 281, 7, 2, 0, 0, 281,
		51, 1, 0, 0, 0, 282, 283, 5, 30, 0, 0, 283, 288, 3, 54, 27, 0, 284, 285,
		5, 41, 0, 0, 285, 287, 3, 54, 27, 0, 286, 284, 1, 0, 0, 0, 287, 290, 1,
		0, 0, 0, 288, 286, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289, 291, 1, 0, 0,
		0, 290, 288, 1, 0, 0, 0, 291, 292, 5, 31, 0, 0, 292, 53, 1, 0, 0, 0, 293,
		294, 3, 16, 8, 0, 294, 295, 5, 26, 0, 0, 295, 296, 3, 66, 33, 0, 296, 55,
		1, 0, 0, 0, 297, 299, 3, 78, 39, 0, 298, 300, 3, 6, 3, 0, 299, 298, 1,
		0, 0, 0, 299, 300, 1, 0, 0, 0, 300, 301, 1, 0, 0, 0, 301, 302, 3, 58, 29,
		0, 302, 57, 1, 0, 0, 0, 303, 307, 5, 32, 0, 0, 304, 306, 3, 60, 30, 0,
		305, 304, 1, 0, 0, 0, 306, 309, 1, 0, 0, 0, 307, 305, 1, 0, 0, 0, 307,
		308, 1, 0, 0, 0, 308, 310, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 310, 311,
		5, 33, 0, 0, 311, 59, 1, 0, 0, 0, 312, 315, 3, 18, 9, 0, 313, 315, 3, 68,
		34, 0, 314, 312, 1, 0, 0, 0, 314, 313, 1, 0, 0, 0, 315, 61, 1, 0, 0, 0,
		316, 317, 6, 31, -1, 0, 317, 331, 3, 64, 32, 0, 318, 331, 3, 66, 33, 0,
		319, 320, 5, 28, 0, 0, 320, 321, 3, 62, 31, 0, 321, 322, 5, 29, 0, 0, 322,
		331, 1, 0, 0, 0, 323, 324, 7, 3, 0, 0, 324, 331, 3, 62, 31, 13, 325, 326,
		5, 28, 0, 0, 326, 327, 3, 38, 19, 0, 327, 328, 5, 29, 0, 0, 328, 329, 3,
		62, 31, 12, 329, 331, 1, 0, 0, 0, 330, 316, 1, 0, 0, 0, 330, 318, 1, 0,
		0, 0, 330, 319, 1, 0, 0, 0, 330, 323, 1, 0, 0, 0, 330, 325, 1, 0, 0, 0,
		331, 373, 1, 0, 0, 0, 332, 333, 10, 14, 0, 0, 333, 334, 5, 48, 0, 0, 334,
		372, 3, 62, 31, 15, 335, 336, 10, 11, 0, 0, 336, 337, 7, 4, 0, 0, 337,
		372, 3, 62, 31, 12, 338, 339, 10, 10, 0, 0, 339, 340, 7, 5, 0, 0, 340,
		372, 3, 62, 31, 11, 341, 342, 10, 9, 0, 0, 342, 343, 7, 6, 0, 0, 343, 372,
		3, 62, 31, 10, 344, 345, 10, 8, 0, 0, 345, 346, 7, 7, 0, 0, 346, 372, 3,
		62, 31, 9, 347, 348, 10, 7, 0, 0, 348, 349, 7, 8, 0, 0, 349, 372, 3, 62,
		31, 8, 350, 351, 10, 6, 0, 0, 351, 352, 5, 51, 0, 0, 352, 372, 3, 62, 31,
		7, 353, 354, 10, 5, 0, 0, 354, 355, 5, 53, 0, 0, 355, 372, 3, 62, 31, 6,
		356, 357, 10, 4, 0, 0, 357, 358, 5, 52, 0, 0, 358, 372, 3, 62, 31, 5, 359,
		360, 10, 3, 0, 0, 360, 361, 5, 55, 0, 0, 361, 372, 3, 62, 31, 4, 362, 363,
		10, 2, 0, 0, 363, 364, 5, 56, 0, 0, 364, 372, 3, 62, 31, 3, 365, 366, 10,
		1, 0, 0, 366, 367, 5, 27, 0, 0, 367, 368, 3, 62, 31, 0, 368, 369, 5, 42,
		0, 0, 369, 370, 3, 62, 31, 2, 370, 372, 1, 0, 0, 0, 371, 332, 1, 0, 0,
		0, 371, 335, 1, 0, 0, 0, 371, 338, 1, 0, 0, 0, 371, 341, 1, 0, 0, 0, 371,
		344, 1, 0, 0, 0, 371, 347, 1, 0, 0, 0, 371, 350, 1, 0, 0, 0, 371, 353,
		1, 0, 0, 0, 371, 356, 1, 0, 0, 0, 371, 359, 1, 0, 0, 0, 371, 362, 1, 0,
		0, 0, 371, 365, 1, 0, 0, 0, 372, 375, 1, 0, 0, 0, 373, 371, 1, 0, 0, 0,
		373, 374, 1, 0, 0, 0, 374, 63, 1, 0, 0, 0, 375, 373, 1, 0, 0, 0, 376, 377,
		5, 5, 0, 0, 377, 65, 1, 0, 0, 0, 378, 380, 7, 5, 0, 0, 379, 378, 1, 0,
		0, 0, 379, 380, 1, 0, 0, 0, 380, 381, 1, 0, 0, 0, 381, 389, 3, 88, 44,
		0, 382, 384, 7, 5, 0, 0, 383, 382, 1, 0, 0, 0, 383, 384, 1, 0, 0, 0, 384,
		385, 1, 0, 0, 0, 385, 389, 3, 94, 47, 0, 386, 389, 3, 90, 45, 0, 387, 389,
		3, 92, 46, 0, 388, 379, 1, 0, 0, 0, 388, 383, 1, 0, 0, 0, 388, 386, 1,
		0, 0, 0, 388, 387, 1, 0, 0, 0, 389, 67, 1, 0, 0, 0, 390, 391, 5, 25, 0,
		0, 391, 69, 1, 0, 0, 0, 392, 393, 5, 62, 0, 0, 393, 71, 1, 0, 0, 0, 394,
		399, 3, 70, 35, 0, 395, 396, 5, 40, 0, 0, 396, 398, 3, 70, 35, 0, 397,
		395, 1, 0, 0, 0, 398, 401, 1, 0, 0, 0, 399, 397, 1, 0, 0, 0, 399, 400,
		1, 0, 0, 0, 400, 73, 1, 0, 0, 0, 401, 399, 1, 0, 0, 0, 402, 403, 3, 70,
		35, 0, 403, 75, 1, 0, 0, 0, 404, 405, 3, 70, 35, 0, 405, 77, 1, 0, 0, 0,
		406, 407, 5, 23, 0, 0, 407, 408, 3, 70, 35, 0, 408, 79, 1, 0, 0, 0, 409,
		410, 5, 22, 0, 0, 410, 411, 3, 70, 35, 0, 411, 81, 1, 0, 0, 0, 412, 413,
		3, 70, 35, 0, 413, 83, 1, 0, 0, 0, 414, 417, 3, 78, 39, 0, 415, 417, 3,
		56, 28, 0, 416, 414, 1, 0, 0, 0, 416, 415, 1, 0, 0, 0, 417, 85, 1, 0, 0,
		0, 418, 419, 3, 80, 40, 0, 419, 87, 1, 0, 0, 0, 420, 421, 5, 61, 0, 0,
		421, 89, 1, 0, 0, 0, 422, 423, 5, 58, 0, 0, 423, 91, 1, 0, 0, 0, 424, 425,
		5, 59, 0, 0, 425, 93, 1, 0, 0, 0, 426, 427, 5, 60, 0, 0, 427, 95, 1, 0,
		0, 0, 37, 101, 103, 114, 120, 124, 148, 153, 159, 162, 167, 170, 176, 183,
		186, 189, 199, 212, 219, 235, 245, 260, 267, 273, 275, 278, 288, 299, 307,
		314, 330, 371, 373, 379, 383, 388, 399, 416,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// bubblerParserInit initializes any static state used to implement bubblerParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewbubblerParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BubblerParserInit() {
	staticData := &BubblerParserStaticData
	staticData.once.Do(bubblerParserInit)
}

// NewbubblerParser produces a new parser instance for the optional input antlr.TokenStream.
func NewbubblerParser(input antlr.TokenStream) *bubblerParser {
	BubblerParserInit()
	this := new(bubblerParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &BubblerParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "bubbler.g4"

	return this
}

// bubblerParser tokens.
const (
	bubblerParserEOF          = antlr.TokenEOF
	bubblerParserSYNTAX       = 1
	bubblerParserIMPORT       = 2
	bubblerParserGET          = 3
	bubblerParserSET          = 4
	bubblerParserVALUE        = 5
	bubblerParserPACKAGE      = 6
	bubblerParserOPTION       = 7
	bubblerParserVOID         = 8
	bubblerParserINT8         = 9
	bubblerParserINT16        = 10
	bubblerParserINT32        = 11
	bubblerParserINT64        = 12
	bubblerParserUINT8        = 13
	bubblerParserUINT16       = 14
	bubblerParserUINT32       = 15
	bubblerParserUINT64       = 16
	bubblerParserFLOAT32      = 17
	bubblerParserFLOAT64      = 18
	bubblerParserBOOL         = 19
	bubblerParserSTRING       = 20
	bubblerParserBYTES        = 21
	bubblerParserENUM         = 22
	bubblerParserSTRUCT       = 23
	bubblerParserHASH         = 24
	bubblerParserSEMI         = 25
	bubblerParserASSIGN       = 26
	bubblerParserQUESTION     = 27
	bubblerParserLP           = 28
	bubblerParserRP           = 29
	bubblerParserLB           = 30
	bubblerParserRB           = 31
	bubblerParserLC           = 32
	bubblerParserRC           = 33
	bubblerParserLT           = 34
	bubblerParserLE           = 35
	bubblerParserGT           = 36
	bubblerParserGE           = 37
	bubblerParserEQ           = 38
	bubblerParserNE           = 39
	bubblerParserDOT          = 40
	bubblerParserCOMMA        = 41
	bubblerParserCOLON        = 42
	bubblerParserADD          = 43
	bubblerParserSUB          = 44
	bubblerParserMUL          = 45
	bubblerParserDIV          = 46
	bubblerParserMOD          = 47
	bubblerParserPOW          = 48
	bubblerParserSHL          = 49
	bubblerParserSHR          = 50
	bubblerParserBAND         = 51
	bubblerParserBOR          = 52
	bubblerParserBXOR         = 53
	bubblerParserBNOT         = 54
	bubblerParserAND          = 55
	bubblerParserOR           = 56
	bubblerParserNOT          = 57
	bubblerParserSTR_LIT      = 58
	bubblerParserBOOL_LIT     = 59
	bubblerParserFLOAT_LIT    = 60
	bubblerParserINT_LIT      = 61
	bubblerParserIDENTIFIER   = 62
	bubblerParserWS           = 63
	bubblerParserLINE_COMMENT = 64
	bubblerParserCOMMENT      = 65
	bubblerParserKEYWORDS     = 66
)

// bubblerParser rules.
const (
	bubblerParserRULE_proto            = 0
	bubblerParserRULE_importStatement  = 1
	bubblerParserRULE_topLevelDef      = 2
	bubblerParserRULE_size_            = 3
	bubblerParserRULE_byteSize         = 4
	bubblerParserRULE_bitSize          = 5
	bubblerParserRULE_packageStatement = 6
	bubblerParserRULE_optionStatement  = 7
	bubblerParserRULE_optionName       = 8
	bubblerParserRULE_field            = 9
	bubblerParserRULE_fieldVoid        = 10
	bubblerParserRULE_fieldConstant    = 11
	bubblerParserRULE_fieldEmbedded    = 12
	bubblerParserRULE_fieldNormal      = 13
	bubblerParserRULE_fieldOptions     = 14
	bubblerParserRULE_fieldOption      = 15
	bubblerParserRULE_fieldMethods     = 16
	bubblerParserRULE_fieldMethod      = 17
	bubblerParserRULE_type_            = 18
	bubblerParserRULE_basicType        = 19
	bubblerParserRULE_arrayElementType = 20
	bubblerParserRULE_arrayType        = 21
	bubblerParserRULE_enumDef          = 22
	bubblerParserRULE_enumBody         = 23
	bubblerParserRULE_enumElement      = 24
	bubblerParserRULE_enumValue        = 25
	bubblerParserRULE_enumValueOptions = 26
	bubblerParserRULE_enumValueOption  = 27
	bubblerParserRULE_structDef        = 28
	bubblerParserRULE_structBody       = 29
	bubblerParserRULE_structElement    = 30
	bubblerParserRULE_expr             = 31
	bubblerParserRULE_value            = 32
	bubblerParserRULE_constant         = 33
	bubblerParserRULE_emptyStatement_  = 34
	bubblerParserRULE_ident            = 35
	bubblerParserRULE_fullIdent        = 36
	bubblerParserRULE_fieldName        = 37
	bubblerParserRULE_methodName       = 38
	bubblerParserRULE_structName       = 39
	bubblerParserRULE_enumName         = 40
	bubblerParserRULE_enumValueName    = 41
	bubblerParserRULE_structType       = 42
	bubblerParserRULE_enumType         = 43
	bubblerParserRULE_intLit           = 44
	bubblerParserRULE_strLit           = 45
	bubblerParserRULE_boolLit          = 46
	bubblerParserRULE_floatLit         = 47
)

// IProtoContext is an interface to support dynamic dispatch.
type IProtoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllImportStatement() []IImportStatementContext
	ImportStatement(i int) IImportStatementContext
	AllPackageStatement() []IPackageStatementContext
	PackageStatement(i int) IPackageStatementContext
	AllOptionStatement() []IOptionStatementContext
	OptionStatement(i int) IOptionStatementContext
	AllTopLevelDef() []ITopLevelDefContext
	TopLevelDef(i int) ITopLevelDefContext
	AllEmptyStatement_() []IEmptyStatement_Context
	EmptyStatement_(i int) IEmptyStatement_Context

	// IsProtoContext differentiates from other interfaces.
	IsProtoContext()
}

type ProtoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProtoContext() *ProtoContext {
	var p = new(ProtoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_proto
	return p
}

func InitEmptyProtoContext(p *ProtoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_proto
}

func (*ProtoContext) IsProtoContext() {}

func NewProtoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProtoContext {
	var p = new(ProtoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_proto

	return p
}

func (s *ProtoContext) GetParser() antlr.Parser { return s.parser }

func (s *ProtoContext) EOF() antlr.TerminalNode {
	return s.GetToken(bubblerParserEOF, 0)
}

func (s *ProtoContext) AllImportStatement() []IImportStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImportStatementContext); ok {
			len++
		}
	}

	tst := make([]IImportStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImportStatementContext); ok {
			tst[i] = t.(IImportStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProtoContext) ImportStatement(i int) IImportStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportStatementContext)
}

func (s *ProtoContext) AllPackageStatement() []IPackageStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPackageStatementContext); ok {
			len++
		}
	}

	tst := make([]IPackageStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPackageStatementContext); ok {
			tst[i] = t.(IPackageStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProtoContext) PackageStatement(i int) IPackageStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPackageStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPackageStatementContext)
}

func (s *ProtoContext) AllOptionStatement() []IOptionStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOptionStatementContext); ok {
			len++
		}
	}

	tst := make([]IOptionStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOptionStatementContext); ok {
			tst[i] = t.(IOptionStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProtoContext) OptionStatement(i int) IOptionStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionStatementContext)
}

func (s *ProtoContext) AllTopLevelDef() []ITopLevelDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITopLevelDefContext); ok {
			len++
		}
	}

	tst := make([]ITopLevelDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITopLevelDefContext); ok {
			tst[i] = t.(ITopLevelDefContext)
			i++
		}
	}

	return tst
}

func (s *ProtoContext) TopLevelDef(i int) ITopLevelDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITopLevelDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITopLevelDefContext)
}

func (s *ProtoContext) AllEmptyStatement_() []IEmptyStatement_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			len++
		}
	}

	tst := make([]IEmptyStatement_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEmptyStatement_Context); ok {
			tst[i] = t.(IEmptyStatement_Context)
			i++
		}
	}

	return tst
}

func (s *ProtoContext) EmptyStatement_(i int) IEmptyStatement_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *ProtoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProtoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProtoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitProto(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) Proto() (localctx IProtoContext) {
	localctx = NewProtoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, bubblerParserRULE_proto)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&46137540) != 0 {
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case bubblerParserIMPORT:
			{
				p.SetState(96)
				p.ImportStatement()
			}

		case bubblerParserPACKAGE:
			{
				p.SetState(97)
				p.PackageStatement()
			}

		case bubblerParserOPTION:
			{
				p.SetState(98)
				p.OptionStatement()
			}

		case bubblerParserENUM, bubblerParserSTRUCT:
			{
				p.SetState(99)
				p.TopLevelDef()
			}

		case bubblerParserSEMI:
			{
				p.SetState(100)
				p.EmptyStatement_()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(106)
		p.Match(bubblerParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImportStatementContext is an interface to support dynamic dispatch.
type IImportStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IMPORT() antlr.TerminalNode
	StrLit() IStrLitContext
	SEMI() antlr.TerminalNode

	// IsImportStatementContext differentiates from other interfaces.
	IsImportStatementContext()
}

type ImportStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportStatementContext() *ImportStatementContext {
	var p = new(ImportStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_importStatement
	return p
}

func InitEmptyImportStatementContext(p *ImportStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_importStatement
}

func (*ImportStatementContext) IsImportStatementContext() {}

func NewImportStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStatementContext {
	var p = new(ImportStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_importStatement

	return p
}

func (s *ImportStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStatementContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(bubblerParserIMPORT, 0)
}

func (s *ImportStatementContext) StrLit() IStrLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrLitContext)
}

func (s *ImportStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(bubblerParserSEMI, 0)
}

func (s *ImportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitImportStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) ImportStatement() (localctx IImportStatementContext) {
	localctx = NewImportStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, bubblerParserRULE_importStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(bubblerParserIMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.StrLit()
	}
	{
		p.SetState(110)
		p.Match(bubblerParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITopLevelDefContext is an interface to support dynamic dispatch.
type ITopLevelDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EnumDef() IEnumDefContext
	StructDef() IStructDefContext

	// IsTopLevelDefContext differentiates from other interfaces.
	IsTopLevelDefContext()
}

type TopLevelDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopLevelDefContext() *TopLevelDefContext {
	var p = new(TopLevelDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_topLevelDef
	return p
}

func InitEmptyTopLevelDefContext(p *TopLevelDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_topLevelDef
}

func (*TopLevelDefContext) IsTopLevelDefContext() {}

func NewTopLevelDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLevelDefContext {
	var p = new(TopLevelDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_topLevelDef

	return p
}

func (s *TopLevelDefContext) GetParser() antlr.Parser { return s.parser }

func (s *TopLevelDefContext) EnumDef() IEnumDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumDefContext)
}

func (s *TopLevelDefContext) StructDef() IStructDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDefContext)
}

func (s *TopLevelDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopLevelDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitTopLevelDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) TopLevelDef() (localctx ITopLevelDefContext) {
	localctx = NewTopLevelDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, bubblerParserRULE_topLevelDef)
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case bubblerParserENUM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(112)
			p.EnumDef()
		}

	case bubblerParserSTRUCT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(113)
			p.StructDef()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISize_Context is an interface to support dynamic dispatch.
type ISize_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LB() antlr.TerminalNode
	RB() antlr.TerminalNode
	ByteSize() IByteSizeContext
	HASH() antlr.TerminalNode
	BitSize() IBitSizeContext

	// IsSize_Context differentiates from other interfaces.
	IsSize_Context()
}

type Size_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySize_Context() *Size_Context {
	var p = new(Size_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_size_
	return p
}

func InitEmptySize_Context(p *Size_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_size_
}

func (*Size_Context) IsSize_Context() {}

func NewSize_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Size_Context {
	var p = new(Size_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_size_

	return p
}

func (s *Size_Context) GetParser() antlr.Parser { return s.parser }

func (s *Size_Context) LB() antlr.TerminalNode {
	return s.GetToken(bubblerParserLB, 0)
}

func (s *Size_Context) RB() antlr.TerminalNode {
	return s.GetToken(bubblerParserRB, 0)
}

func (s *Size_Context) ByteSize() IByteSizeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IByteSizeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IByteSizeContext)
}

func (s *Size_Context) HASH() antlr.TerminalNode {
	return s.GetToken(bubblerParserHASH, 0)
}

func (s *Size_Context) BitSize() IBitSizeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBitSizeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBitSizeContext)
}

func (s *Size_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Size_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Size_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitSize_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) Size_() (localctx ISize_Context) {
	localctx = NewSize_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, bubblerParserRULE_size_)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Match(bubblerParserLB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case bubblerParserINT_LIT:
		{
			p.SetState(117)
			p.ByteSize()
		}
		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == bubblerParserHASH {
			{
				p.SetState(118)
				p.Match(bubblerParserHASH)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(119)
				p.BitSize()
			}

		}

	case bubblerParserHASH:
		{
			p.SetState(122)
			p.Match(bubblerParserHASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.BitSize()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(126)
		p.Match(bubblerParserRB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IByteSizeContext is an interface to support dynamic dispatch.
type IByteSizeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IntLit() IIntLitContext

	// IsByteSizeContext differentiates from other interfaces.
	IsByteSizeContext()
}

type ByteSizeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyByteSizeContext() *ByteSizeContext {
	var p = new(ByteSizeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_byteSize
	return p
}

func InitEmptyByteSizeContext(p *ByteSizeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_byteSize
}

func (*ByteSizeContext) IsByteSizeContext() {}

func NewByteSizeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ByteSizeContext {
	var p = new(ByteSizeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_byteSize

	return p
}

func (s *ByteSizeContext) GetParser() antlr.Parser { return s.parser }

func (s *ByteSizeContext) IntLit() IIntLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntLitContext)
}

func (s *ByteSizeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ByteSizeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ByteSizeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitByteSize(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) ByteSize() (localctx IByteSizeContext) {
	localctx = NewByteSizeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, bubblerParserRULE_byteSize)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.IntLit()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBitSizeContext is an interface to support dynamic dispatch.
type IBitSizeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IntLit() IIntLitContext

	// IsBitSizeContext differentiates from other interfaces.
	IsBitSizeContext()
}

type BitSizeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBitSizeContext() *BitSizeContext {
	var p = new(BitSizeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_bitSize
	return p
}

func InitEmptyBitSizeContext(p *BitSizeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_bitSize
}

func (*BitSizeContext) IsBitSizeContext() {}

func NewBitSizeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BitSizeContext {
	var p = new(BitSizeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_bitSize

	return p
}

func (s *BitSizeContext) GetParser() antlr.Parser { return s.parser }

func (s *BitSizeContext) IntLit() IIntLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntLitContext)
}

func (s *BitSizeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitSizeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BitSizeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitBitSize(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) BitSize() (localctx IBitSizeContext) {
	localctx = NewBitSizeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, bubblerParserRULE_bitSize)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.IntLit()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPackageStatementContext is an interface to support dynamic dispatch.
type IPackageStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PACKAGE() antlr.TerminalNode
	FullIdent() IFullIdentContext
	SEMI() antlr.TerminalNode

	// IsPackageStatementContext differentiates from other interfaces.
	IsPackageStatementContext()
}

type PackageStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageStatementContext() *PackageStatementContext {
	var p = new(PackageStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_packageStatement
	return p
}

func InitEmptyPackageStatementContext(p *PackageStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_packageStatement
}

func (*PackageStatementContext) IsPackageStatementContext() {}

func NewPackageStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageStatementContext {
	var p = new(PackageStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_packageStatement

	return p
}

func (s *PackageStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageStatementContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(bubblerParserPACKAGE, 0)
}

func (s *PackageStatementContext) FullIdent() IFullIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFullIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *PackageStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(bubblerParserSEMI, 0)
}

func (s *PackageStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitPackageStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) PackageStatement() (localctx IPackageStatementContext) {
	localctx = NewPackageStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, bubblerParserRULE_packageStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Match(bubblerParserPACKAGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)
		p.FullIdent()
	}
	{
		p.SetState(134)
		p.Match(bubblerParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOptionStatementContext is an interface to support dynamic dispatch.
type IOptionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPTION() antlr.TerminalNode
	OptionName() IOptionNameContext
	ASSIGN() antlr.TerminalNode
	Constant() IConstantContext
	SEMI() antlr.TerminalNode

	// IsOptionStatementContext differentiates from other interfaces.
	IsOptionStatementContext()
}

type OptionStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionStatementContext() *OptionStatementContext {
	var p = new(OptionStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_optionStatement
	return p
}

func InitEmptyOptionStatementContext(p *OptionStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_optionStatement
}

func (*OptionStatementContext) IsOptionStatementContext() {}

func NewOptionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionStatementContext {
	var p = new(OptionStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_optionStatement

	return p
}

func (s *OptionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionStatementContext) OPTION() antlr.TerminalNode {
	return s.GetToken(bubblerParserOPTION, 0)
}

func (s *OptionStatementContext) OptionName() IOptionNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *OptionStatementContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(bubblerParserASSIGN, 0)
}

func (s *OptionStatementContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *OptionStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(bubblerParserSEMI, 0)
}

func (s *OptionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitOptionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) OptionStatement() (localctx IOptionStatementContext) {
	localctx = NewOptionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, bubblerParserRULE_optionStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(bubblerParserOPTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
		p.OptionName()
	}
	{
		p.SetState(138)
		p.Match(bubblerParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.Constant()
	}
	{
		p.SetState(140)
		p.Match(bubblerParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOptionNameContext is an interface to support dynamic dispatch.
type IOptionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext

	// IsOptionNameContext differentiates from other interfaces.
	IsOptionNameContext()
}

type OptionNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionNameContext() *OptionNameContext {
	var p = new(OptionNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_optionName
	return p
}

func InitEmptyOptionNameContext(p *OptionNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_optionName
}

func (*OptionNameContext) IsOptionNameContext() {}

func NewOptionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionNameContext {
	var p = new(OptionNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_optionName

	return p
}

func (s *OptionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *OptionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitOptionName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) OptionName() (localctx IOptionNameContext) {
	localctx = NewOptionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, bubblerParserRULE_optionName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Ident()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FieldVoid() IFieldVoidContext
	FieldConstant() IFieldConstantContext
	FieldEmbedded() IFieldEmbeddedContext
	FieldNormal() IFieldNormalContext

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_field
	return p
}

func InitEmptyFieldContext(p *FieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_field
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) FieldVoid() IFieldVoidContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldVoidContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldVoidContext)
}

func (s *FieldContext) FieldConstant() IFieldConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldConstantContext)
}

func (s *FieldContext) FieldEmbedded() IFieldEmbeddedContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldEmbeddedContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldEmbeddedContext)
}

func (s *FieldContext) FieldNormal() IFieldNormalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNormalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNormalContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, bubblerParserRULE_field)
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(144)
			p.FieldVoid()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(145)
			p.FieldConstant()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(146)
			p.FieldEmbedded()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(147)
			p.FieldNormal()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldVoidContext is an interface to support dynamic dispatch.
type IFieldVoidContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VOID() antlr.TerminalNode
	Size_() ISize_Context
	SEMI() antlr.TerminalNode
	FieldOptions() IFieldOptionsContext

	// IsFieldVoidContext differentiates from other interfaces.
	IsFieldVoidContext()
}

type FieldVoidContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldVoidContext() *FieldVoidContext {
	var p = new(FieldVoidContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_fieldVoid
	return p
}

func InitEmptyFieldVoidContext(p *FieldVoidContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_fieldVoid
}

func (*FieldVoidContext) IsFieldVoidContext() {}

func NewFieldVoidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldVoidContext {
	var p = new(FieldVoidContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_fieldVoid

	return p
}

func (s *FieldVoidContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldVoidContext) VOID() antlr.TerminalNode {
	return s.GetToken(bubblerParserVOID, 0)
}

func (s *FieldVoidContext) Size_() ISize_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISize_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISize_Context)
}

func (s *FieldVoidContext) SEMI() antlr.TerminalNode {
	return s.GetToken(bubblerParserSEMI, 0)
}

func (s *FieldVoidContext) FieldOptions() IFieldOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *FieldVoidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldVoidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldVoidContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitFieldVoid(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) FieldVoid() (localctx IFieldVoidContext) {
	localctx = NewFieldVoidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, bubblerParserRULE_fieldVoid)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Match(bubblerParserVOID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)
		p.Size_()
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == bubblerParserLB {
		{
			p.SetState(152)
			p.FieldOptions()
		}

	}
	{
		p.SetState(155)
		p.Match(bubblerParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldConstantContext is an interface to support dynamic dispatch.
type IFieldConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BasicType() IBasicTypeContext
	ASSIGN() antlr.TerminalNode
	SEMI() antlr.TerminalNode
	Constant() IConstantContext
	Ident() IIdentContext
	FieldName() IFieldNameContext
	Size_() ISize_Context
	FieldOptions() IFieldOptionsContext

	// IsFieldConstantContext differentiates from other interfaces.
	IsFieldConstantContext()
}

type FieldConstantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldConstantContext() *FieldConstantContext {
	var p = new(FieldConstantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_fieldConstant
	return p
}

func InitEmptyFieldConstantContext(p *FieldConstantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_fieldConstant
}

func (*FieldConstantContext) IsFieldConstantContext() {}

func NewFieldConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldConstantContext {
	var p = new(FieldConstantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_fieldConstant

	return p
}

func (s *FieldConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldConstantContext) BasicType() IBasicTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBasicTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBasicTypeContext)
}

func (s *FieldConstantContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(bubblerParserASSIGN, 0)
}

func (s *FieldConstantContext) SEMI() antlr.TerminalNode {
	return s.GetToken(bubblerParserSEMI, 0)
}

func (s *FieldConstantContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *FieldConstantContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *FieldConstantContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *FieldConstantContext) Size_() ISize_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISize_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISize_Context)
}

func (s *FieldConstantContext) FieldOptions() IFieldOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *FieldConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitFieldConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) FieldConstant() (localctx IFieldConstantContext) {
	localctx = NewFieldConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, bubblerParserRULE_fieldConstant)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.BasicType()
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == bubblerParserIDENTIFIER {
		{
			p.SetState(158)
			p.FieldName()
		}

	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == bubblerParserLB {
		{
			p.SetState(161)
			p.Size_()
		}

	}
	{
		p.SetState(164)
		p.Match(bubblerParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case bubblerParserADD, bubblerParserSUB, bubblerParserSTR_LIT, bubblerParserBOOL_LIT, bubblerParserFLOAT_LIT, bubblerParserINT_LIT:
		{
			p.SetState(165)
			p.Constant()
		}

	case bubblerParserIDENTIFIER:
		{
			p.SetState(166)
			p.Ident()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == bubblerParserLB {
		{
			p.SetState(169)
			p.FieldOptions()
		}

	}
	{
		p.SetState(172)
		p.Match(bubblerParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldEmbeddedContext is an interface to support dynamic dispatch.
type IFieldEmbeddedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() IType_Context
	SEMI() antlr.TerminalNode
	FieldOptions() IFieldOptionsContext

	// IsFieldEmbeddedContext differentiates from other interfaces.
	IsFieldEmbeddedContext()
}

type FieldEmbeddedContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldEmbeddedContext() *FieldEmbeddedContext {
	var p = new(FieldEmbeddedContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_fieldEmbedded
	return p
}

func InitEmptyFieldEmbeddedContext(p *FieldEmbeddedContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_fieldEmbedded
}

func (*FieldEmbeddedContext) IsFieldEmbeddedContext() {}

func NewFieldEmbeddedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldEmbeddedContext {
	var p = new(FieldEmbeddedContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_fieldEmbedded

	return p
}

func (s *FieldEmbeddedContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldEmbeddedContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *FieldEmbeddedContext) SEMI() antlr.TerminalNode {
	return s.GetToken(bubblerParserSEMI, 0)
}

func (s *FieldEmbeddedContext) FieldOptions() IFieldOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *FieldEmbeddedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldEmbeddedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldEmbeddedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitFieldEmbedded(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) FieldEmbedded() (localctx IFieldEmbeddedContext) {
	localctx = NewFieldEmbeddedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, bubblerParserRULE_fieldEmbedded)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Type_()
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == bubblerParserLB {
		{
			p.SetState(175)
			p.FieldOptions()
		}

	}
	{
		p.SetState(178)
		p.Match(bubblerParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldNormalContext is an interface to support dynamic dispatch.
type IFieldNormalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() IType_Context
	FieldName() IFieldNameContext
	SEMI() antlr.TerminalNode
	Size_() ISize_Context
	FieldOptions() IFieldOptionsContext
	FieldMethods() IFieldMethodsContext

	// IsFieldNormalContext differentiates from other interfaces.
	IsFieldNormalContext()
}

type FieldNormalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNormalContext() *FieldNormalContext {
	var p = new(FieldNormalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_fieldNormal
	return p
}

func InitEmptyFieldNormalContext(p *FieldNormalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_fieldNormal
}

func (*FieldNormalContext) IsFieldNormalContext() {}

func NewFieldNormalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNormalContext {
	var p = new(FieldNormalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_fieldNormal

	return p
}

func (s *FieldNormalContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNormalContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *FieldNormalContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *FieldNormalContext) SEMI() antlr.TerminalNode {
	return s.GetToken(bubblerParserSEMI, 0)
}

func (s *FieldNormalContext) Size_() ISize_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISize_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISize_Context)
}

func (s *FieldNormalContext) FieldOptions() IFieldOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *FieldNormalContext) FieldMethods() IFieldMethodsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldMethodsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldMethodsContext)
}

func (s *FieldNormalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNormalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNormalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitFieldNormal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) FieldNormal() (localctx IFieldNormalContext) {
	localctx = NewFieldNormalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, bubblerParserRULE_fieldNormal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Type_()
	}
	{
		p.SetState(181)
		p.FieldName()
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(182)
			p.Size_()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == bubblerParserLB {
		{
			p.SetState(185)
			p.FieldOptions()
		}

	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == bubblerParserLC {
		{
			p.SetState(188)
			p.FieldMethods()
		}

	}
	{
		p.SetState(191)
		p.Match(bubblerParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldOptionsContext is an interface to support dynamic dispatch.
type IFieldOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LB() antlr.TerminalNode
	AllFieldOption() []IFieldOptionContext
	FieldOption(i int) IFieldOptionContext
	RB() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFieldOptionsContext differentiates from other interfaces.
	IsFieldOptionsContext()
}

type FieldOptionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldOptionsContext() *FieldOptionsContext {
	var p = new(FieldOptionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_fieldOptions
	return p
}

func InitEmptyFieldOptionsContext(p *FieldOptionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_fieldOptions
}

func (*FieldOptionsContext) IsFieldOptionsContext() {}

func NewFieldOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldOptionsContext {
	var p = new(FieldOptionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_fieldOptions

	return p
}

func (s *FieldOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldOptionsContext) LB() antlr.TerminalNode {
	return s.GetToken(bubblerParserLB, 0)
}

func (s *FieldOptionsContext) AllFieldOption() []IFieldOptionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldOptionContext); ok {
			len++
		}
	}

	tst := make([]IFieldOptionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldOptionContext); ok {
			tst[i] = t.(IFieldOptionContext)
			i++
		}
	}

	return tst
}

func (s *FieldOptionsContext) FieldOption(i int) IFieldOptionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldOptionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldOptionContext)
}

func (s *FieldOptionsContext) RB() antlr.TerminalNode {
	return s.GetToken(bubblerParserRB, 0)
}

func (s *FieldOptionsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(bubblerParserCOMMA)
}

func (s *FieldOptionsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(bubblerParserCOMMA, i)
}

func (s *FieldOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldOptionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitFieldOptions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) FieldOptions() (localctx IFieldOptionsContext) {
	localctx = NewFieldOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, bubblerParserRULE_fieldOptions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		p.Match(bubblerParserLB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(194)
		p.FieldOption()
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == bubblerParserCOMMA {
		{
			p.SetState(195)
			p.Match(bubblerParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(196)
			p.FieldOption()
		}

		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(202)
		p.Match(bubblerParserRB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldOptionContext is an interface to support dynamic dispatch.
type IFieldOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OptionName() IOptionNameContext
	ASSIGN() antlr.TerminalNode
	Constant() IConstantContext

	// IsFieldOptionContext differentiates from other interfaces.
	IsFieldOptionContext()
}

type FieldOptionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldOptionContext() *FieldOptionContext {
	var p = new(FieldOptionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_fieldOption
	return p
}

func InitEmptyFieldOptionContext(p *FieldOptionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_fieldOption
}

func (*FieldOptionContext) IsFieldOptionContext() {}

func NewFieldOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldOptionContext {
	var p = new(FieldOptionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_fieldOption

	return p
}

func (s *FieldOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldOptionContext) OptionName() IOptionNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *FieldOptionContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(bubblerParserASSIGN, 0)
}

func (s *FieldOptionContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *FieldOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldOptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitFieldOption(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) FieldOption() (localctx IFieldOptionContext) {
	localctx = NewFieldOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, bubblerParserRULE_fieldOption)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.OptionName()
	}
	{
		p.SetState(205)
		p.Match(bubblerParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.Constant()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldMethodsContext is an interface to support dynamic dispatch.
type IFieldMethodsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LC() antlr.TerminalNode
	RC() antlr.TerminalNode
	AllFieldMethod() []IFieldMethodContext
	FieldMethod(i int) IFieldMethodContext

	// IsFieldMethodsContext differentiates from other interfaces.
	IsFieldMethodsContext()
}

type FieldMethodsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldMethodsContext() *FieldMethodsContext {
	var p = new(FieldMethodsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_fieldMethods
	return p
}

func InitEmptyFieldMethodsContext(p *FieldMethodsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_fieldMethods
}

func (*FieldMethodsContext) IsFieldMethodsContext() {}

func NewFieldMethodsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldMethodsContext {
	var p = new(FieldMethodsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_fieldMethods

	return p
}

func (s *FieldMethodsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldMethodsContext) LC() antlr.TerminalNode {
	return s.GetToken(bubblerParserLC, 0)
}

func (s *FieldMethodsContext) RC() antlr.TerminalNode {
	return s.GetToken(bubblerParserRC, 0)
}

func (s *FieldMethodsContext) AllFieldMethod() []IFieldMethodContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldMethodContext); ok {
			len++
		}
	}

	tst := make([]IFieldMethodContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldMethodContext); ok {
			tst[i] = t.(IFieldMethodContext)
			i++
		}
	}

	return tst
}

func (s *FieldMethodsContext) FieldMethod(i int) IFieldMethodContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldMethodContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldMethodContext)
}

func (s *FieldMethodsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldMethodsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldMethodsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitFieldMethods(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) FieldMethods() (localctx IFieldMethodsContext) {
	localctx = NewFieldMethodsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, bubblerParserRULE_fieldMethods)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.Match(bubblerParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == bubblerParserGET || _la == bubblerParserSET {
		{
			p.SetState(209)
			p.FieldMethod()
		}

		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(215)
		p.Match(bubblerParserRC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldMethodContext is an interface to support dynamic dispatch.
type IFieldMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	LP() antlr.TerminalNode
	BasicType() IBasicTypeContext
	RP() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Expr() IExprContext
	SEMI() antlr.TerminalNode
	GET() antlr.TerminalNode
	SET() antlr.TerminalNode
	MethodName() IMethodNameContext

	// IsFieldMethodContext differentiates from other interfaces.
	IsFieldMethodContext()
}

type FieldMethodContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyFieldMethodContext() *FieldMethodContext {
	var p = new(FieldMethodContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_fieldMethod
	return p
}

func InitEmptyFieldMethodContext(p *FieldMethodContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_fieldMethod
}

func (*FieldMethodContext) IsFieldMethodContext() {}

func NewFieldMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldMethodContext {
	var p = new(FieldMethodContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_fieldMethod

	return p
}

func (s *FieldMethodContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldMethodContext) GetOp() antlr.Token { return s.op }

func (s *FieldMethodContext) SetOp(v antlr.Token) { s.op = v }

func (s *FieldMethodContext) LP() antlr.TerminalNode {
	return s.GetToken(bubblerParserLP, 0)
}

func (s *FieldMethodContext) BasicType() IBasicTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBasicTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBasicTypeContext)
}

func (s *FieldMethodContext) RP() antlr.TerminalNode {
	return s.GetToken(bubblerParserRP, 0)
}

func (s *FieldMethodContext) COLON() antlr.TerminalNode {
	return s.GetToken(bubblerParserCOLON, 0)
}

func (s *FieldMethodContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FieldMethodContext) SEMI() antlr.TerminalNode {
	return s.GetToken(bubblerParserSEMI, 0)
}

func (s *FieldMethodContext) GET() antlr.TerminalNode {
	return s.GetToken(bubblerParserGET, 0)
}

func (s *FieldMethodContext) SET() antlr.TerminalNode {
	return s.GetToken(bubblerParserSET, 0)
}

func (s *FieldMethodContext) MethodName() IMethodNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodNameContext)
}

func (s *FieldMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldMethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldMethodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitFieldMethod(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) FieldMethod() (localctx IFieldMethodContext) {
	localctx = NewFieldMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, bubblerParserRULE_fieldMethod)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(217)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*FieldMethodContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == bubblerParserGET || _la == bubblerParserSET) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*FieldMethodContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == bubblerParserIDENTIFIER {
		{
			p.SetState(218)
			p.MethodName()
		}

	}
	{
		p.SetState(221)
		p.Match(bubblerParserLP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(222)
		p.BasicType()
	}
	{
		p.SetState(223)
		p.Match(bubblerParserRP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(224)
		p.Match(bubblerParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(225)
		p.expr(0)
	}
	{
		p.SetState(226)
		p.Match(bubblerParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BasicType() IBasicTypeContext
	STRING() antlr.TerminalNode
	BYTES() antlr.TerminalNode
	ArrayType() IArrayTypeContext
	StructType() IStructTypeContext
	EnumType() IEnumTypeContext
	Ident() IIdentContext

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_type_
	return p
}

func InitEmptyType_Context(p *Type_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_type_
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) BasicType() IBasicTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBasicTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBasicTypeContext)
}

func (s *Type_Context) STRING() antlr.TerminalNode {
	return s.GetToken(bubblerParserSTRING, 0)
}

func (s *Type_Context) BYTES() antlr.TerminalNode {
	return s.GetToken(bubblerParserBYTES, 0)
}

func (s *Type_Context) ArrayType() IArrayTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayTypeContext)
}

func (s *Type_Context) StructType() IStructTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructTypeContext)
}

func (s *Type_Context) EnumType() IEnumTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumTypeContext)
}

func (s *Type_Context) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitType_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) Type_() (localctx IType_Context) {
	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, bubblerParserRULE_type_)
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(228)
			p.BasicType()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(229)
			p.Match(bubblerParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(230)
			p.Match(bubblerParserBYTES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(231)
			p.ArrayType()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(232)
			p.StructType()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(233)
			p.EnumType()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(234)
			p.Ident()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBasicTypeContext is an interface to support dynamic dispatch.
type IBasicTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BOOL() antlr.TerminalNode
	INT8() antlr.TerminalNode
	INT16() antlr.TerminalNode
	INT32() antlr.TerminalNode
	INT64() antlr.TerminalNode
	UINT8() antlr.TerminalNode
	UINT16() antlr.TerminalNode
	UINT32() antlr.TerminalNode
	UINT64() antlr.TerminalNode
	FLOAT32() antlr.TerminalNode
	FLOAT64() antlr.TerminalNode

	// IsBasicTypeContext differentiates from other interfaces.
	IsBasicTypeContext()
}

type BasicTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBasicTypeContext() *BasicTypeContext {
	var p = new(BasicTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_basicType
	return p
}

func InitEmptyBasicTypeContext(p *BasicTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_basicType
}

func (*BasicTypeContext) IsBasicTypeContext() {}

func NewBasicTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BasicTypeContext {
	var p = new(BasicTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_basicType

	return p
}

func (s *BasicTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BasicTypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(bubblerParserBOOL, 0)
}

func (s *BasicTypeContext) INT8() antlr.TerminalNode {
	return s.GetToken(bubblerParserINT8, 0)
}

func (s *BasicTypeContext) INT16() antlr.TerminalNode {
	return s.GetToken(bubblerParserINT16, 0)
}

func (s *BasicTypeContext) INT32() antlr.TerminalNode {
	return s.GetToken(bubblerParserINT32, 0)
}

func (s *BasicTypeContext) INT64() antlr.TerminalNode {
	return s.GetToken(bubblerParserINT64, 0)
}

func (s *BasicTypeContext) UINT8() antlr.TerminalNode {
	return s.GetToken(bubblerParserUINT8, 0)
}

func (s *BasicTypeContext) UINT16() antlr.TerminalNode {
	return s.GetToken(bubblerParserUINT16, 0)
}

func (s *BasicTypeContext) UINT32() antlr.TerminalNode {
	return s.GetToken(bubblerParserUINT32, 0)
}

func (s *BasicTypeContext) UINT64() antlr.TerminalNode {
	return s.GetToken(bubblerParserUINT64, 0)
}

func (s *BasicTypeContext) FLOAT32() antlr.TerminalNode {
	return s.GetToken(bubblerParserFLOAT32, 0)
}

func (s *BasicTypeContext) FLOAT64() antlr.TerminalNode {
	return s.GetToken(bubblerParserFLOAT64, 0)
}

func (s *BasicTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasicTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BasicTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitBasicType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) BasicType() (localctx IBasicTypeContext) {
	localctx = NewBasicTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, bubblerParserRULE_basicType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1048064) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayElementTypeContext is an interface to support dynamic dispatch.
type IArrayElementTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BasicType() IBasicTypeContext
	STRING() antlr.TerminalNode
	BYTES() antlr.TerminalNode
	StructType() IStructTypeContext
	EnumType() IEnumTypeContext
	Ident() IIdentContext

	// IsArrayElementTypeContext differentiates from other interfaces.
	IsArrayElementTypeContext()
}

type ArrayElementTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayElementTypeContext() *ArrayElementTypeContext {
	var p = new(ArrayElementTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_arrayElementType
	return p
}

func InitEmptyArrayElementTypeContext(p *ArrayElementTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_arrayElementType
}

func (*ArrayElementTypeContext) IsArrayElementTypeContext() {}

func NewArrayElementTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayElementTypeContext {
	var p = new(ArrayElementTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_arrayElementType

	return p
}

func (s *ArrayElementTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayElementTypeContext) BasicType() IBasicTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBasicTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBasicTypeContext)
}

func (s *ArrayElementTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(bubblerParserSTRING, 0)
}

func (s *ArrayElementTypeContext) BYTES() antlr.TerminalNode {
	return s.GetToken(bubblerParserBYTES, 0)
}

func (s *ArrayElementTypeContext) StructType() IStructTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructTypeContext)
}

func (s *ArrayElementTypeContext) EnumType() IEnumTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumTypeContext)
}

func (s *ArrayElementTypeContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *ArrayElementTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayElementTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayElementTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitArrayElementType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) ArrayElementType() (localctx IArrayElementTypeContext) {
	localctx = NewArrayElementTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, bubblerParserRULE_arrayElementType)
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case bubblerParserINT8, bubblerParserINT16, bubblerParserINT32, bubblerParserINT64, bubblerParserUINT8, bubblerParserUINT16, bubblerParserUINT32, bubblerParserUINT64, bubblerParserFLOAT32, bubblerParserFLOAT64, bubblerParserBOOL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(239)
			p.BasicType()
		}

	case bubblerParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(240)
			p.Match(bubblerParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case bubblerParserBYTES:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(241)
			p.Match(bubblerParserBYTES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case bubblerParserSTRUCT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(242)
			p.StructType()
		}

	case bubblerParserENUM:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(243)
			p.EnumType()
		}

	case bubblerParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(244)
			p.Ident()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayTypeContext is an interface to support dynamic dispatch.
type IArrayTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ArrayElementType() IArrayElementTypeContext
	LT() antlr.TerminalNode
	IntLit() IIntLitContext
	GT() antlr.TerminalNode

	// IsArrayTypeContext differentiates from other interfaces.
	IsArrayTypeContext()
}

type ArrayTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayTypeContext() *ArrayTypeContext {
	var p = new(ArrayTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_arrayType
	return p
}

func InitEmptyArrayTypeContext(p *ArrayTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_arrayType
}

func (*ArrayTypeContext) IsArrayTypeContext() {}

func NewArrayTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayTypeContext {
	var p = new(ArrayTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_arrayType

	return p
}

func (s *ArrayTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayTypeContext) ArrayElementType() IArrayElementTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayElementTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayElementTypeContext)
}

func (s *ArrayTypeContext) LT() antlr.TerminalNode {
	return s.GetToken(bubblerParserLT, 0)
}

func (s *ArrayTypeContext) IntLit() IIntLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntLitContext)
}

func (s *ArrayTypeContext) GT() antlr.TerminalNode {
	return s.GetToken(bubblerParserGT, 0)
}

func (s *ArrayTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitArrayType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) ArrayType() (localctx IArrayTypeContext) {
	localctx = NewArrayTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, bubblerParserRULE_arrayType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.ArrayElementType()
	}
	{
		p.SetState(248)
		p.Match(bubblerParserLT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(249)
		p.IntLit()
	}
	{
		p.SetState(250)
		p.Match(bubblerParserGT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnumDefContext is an interface to support dynamic dispatch.
type IEnumDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EnumName() IEnumNameContext
	Size_() ISize_Context
	EnumBody() IEnumBodyContext

	// IsEnumDefContext differentiates from other interfaces.
	IsEnumDefContext()
}

type EnumDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumDefContext() *EnumDefContext {
	var p = new(EnumDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_enumDef
	return p
}

func InitEmptyEnumDefContext(p *EnumDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_enumDef
}

func (*EnumDefContext) IsEnumDefContext() {}

func NewEnumDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumDefContext {
	var p = new(EnumDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_enumDef

	return p
}

func (s *EnumDefContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumDefContext) EnumName() IEnumNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumNameContext)
}

func (s *EnumDefContext) Size_() ISize_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISize_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISize_Context)
}

func (s *EnumDefContext) EnumBody() IEnumBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

func (s *EnumDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitEnumDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) EnumDef() (localctx IEnumDefContext) {
	localctx = NewEnumDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, bubblerParserRULE_enumDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		p.EnumName()
	}
	{
		p.SetState(253)
		p.Size_()
	}
	{
		p.SetState(254)
		p.EnumBody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnumBodyContext is an interface to support dynamic dispatch.
type IEnumBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LC() antlr.TerminalNode
	RC() antlr.TerminalNode
	AllEnumElement() []IEnumElementContext
	EnumElement(i int) IEnumElementContext

	// IsEnumBodyContext differentiates from other interfaces.
	IsEnumBodyContext()
}

type EnumBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumBodyContext() *EnumBodyContext {
	var p = new(EnumBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_enumBody
	return p
}

func InitEmptyEnumBodyContext(p *EnumBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_enumBody
}

func (*EnumBodyContext) IsEnumBodyContext() {}

func NewEnumBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumBodyContext {
	var p = new(EnumBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_enumBody

	return p
}

func (s *EnumBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumBodyContext) LC() antlr.TerminalNode {
	return s.GetToken(bubblerParserLC, 0)
}

func (s *EnumBodyContext) RC() antlr.TerminalNode {
	return s.GetToken(bubblerParserRC, 0)
}

func (s *EnumBodyContext) AllEnumElement() []IEnumElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnumElementContext); ok {
			len++
		}
	}

	tst := make([]IEnumElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnumElementContext); ok {
			tst[i] = t.(IEnumElementContext)
			i++
		}
	}

	return tst
}

func (s *EnumBodyContext) EnumElement(i int) IEnumElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumElementContext)
}

func (s *EnumBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitEnumBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) EnumBody() (localctx IEnumBodyContext) {
	localctx = NewEnumBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, bubblerParserRULE_enumBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		p.Match(bubblerParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == bubblerParserSEMI || _la == bubblerParserIDENTIFIER {
		{
			p.SetState(257)
			p.EnumElement()
		}

		p.SetState(262)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(263)
		p.Match(bubblerParserRC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnumElementContext is an interface to support dynamic dispatch.
type IEnumElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EnumValue() IEnumValueContext
	EmptyStatement_() IEmptyStatement_Context

	// IsEnumElementContext differentiates from other interfaces.
	IsEnumElementContext()
}

type EnumElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumElementContext() *EnumElementContext {
	var p = new(EnumElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_enumElement
	return p
}

func InitEmptyEnumElementContext(p *EnumElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_enumElement
}

func (*EnumElementContext) IsEnumElementContext() {}

func NewEnumElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumElementContext {
	var p = new(EnumElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_enumElement

	return p
}

func (s *EnumElementContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumElementContext) EnumValue() IEnumValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumValueContext)
}

func (s *EnumElementContext) EmptyStatement_() IEmptyStatement_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *EnumElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitEnumElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) EnumElement() (localctx IEnumElementContext) {
	localctx = NewEnumElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, bubblerParserRULE_enumElement)
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case bubblerParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(265)
			p.EnumValue()
		}

	case bubblerParserSEMI:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(266)
			p.EmptyStatement_()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnumValueContext is an interface to support dynamic dispatch.
type IEnumValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EnumValueName() IEnumValueNameContext
	SEMI() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	EnumValueOptions() IEnumValueOptionsContext
	Constant() IConstantContext
	Ident() IIdentContext

	// IsEnumValueContext differentiates from other interfaces.
	IsEnumValueContext()
}

type EnumValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueContext() *EnumValueContext {
	var p = new(EnumValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_enumValue
	return p
}

func InitEmptyEnumValueContext(p *EnumValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_enumValue
}

func (*EnumValueContext) IsEnumValueContext() {}

func NewEnumValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueContext {
	var p = new(EnumValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_enumValue

	return p
}

func (s *EnumValueContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueContext) EnumValueName() IEnumValueNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumValueNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumValueNameContext)
}

func (s *EnumValueContext) SEMI() antlr.TerminalNode {
	return s.GetToken(bubblerParserSEMI, 0)
}

func (s *EnumValueContext) COMMA() antlr.TerminalNode {
	return s.GetToken(bubblerParserCOMMA, 0)
}

func (s *EnumValueContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(bubblerParserASSIGN, 0)
}

func (s *EnumValueContext) EnumValueOptions() IEnumValueOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumValueOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumValueOptionsContext)
}

func (s *EnumValueContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *EnumValueContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *EnumValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitEnumValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) EnumValue() (localctx IEnumValueContext) {
	localctx = NewEnumValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, bubblerParserRULE_enumValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(269)
		p.EnumValueName()
	}
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == bubblerParserASSIGN {
		{
			p.SetState(270)
			p.Match(bubblerParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case bubblerParserADD, bubblerParserSUB, bubblerParserSTR_LIT, bubblerParserBOOL_LIT, bubblerParserFLOAT_LIT, bubblerParserINT_LIT:
			{
				p.SetState(271)
				p.Constant()
			}

		case bubblerParserIDENTIFIER:
			{
				p.SetState(272)
				p.Ident()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

	}
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == bubblerParserLB {
		{
			p.SetState(277)
			p.EnumValueOptions()
		}

	}
	{
		p.SetState(280)
		_la = p.GetTokenStream().LA(1)

		if !(_la == bubblerParserSEMI || _la == bubblerParserCOMMA) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnumValueOptionsContext is an interface to support dynamic dispatch.
type IEnumValueOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LB() antlr.TerminalNode
	AllEnumValueOption() []IEnumValueOptionContext
	EnumValueOption(i int) IEnumValueOptionContext
	RB() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsEnumValueOptionsContext differentiates from other interfaces.
	IsEnumValueOptionsContext()
}

type EnumValueOptionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueOptionsContext() *EnumValueOptionsContext {
	var p = new(EnumValueOptionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_enumValueOptions
	return p
}

func InitEmptyEnumValueOptionsContext(p *EnumValueOptionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_enumValueOptions
}

func (*EnumValueOptionsContext) IsEnumValueOptionsContext() {}

func NewEnumValueOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueOptionsContext {
	var p = new(EnumValueOptionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_enumValueOptions

	return p
}

func (s *EnumValueOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueOptionsContext) LB() antlr.TerminalNode {
	return s.GetToken(bubblerParserLB, 0)
}

func (s *EnumValueOptionsContext) AllEnumValueOption() []IEnumValueOptionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnumValueOptionContext); ok {
			len++
		}
	}

	tst := make([]IEnumValueOptionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnumValueOptionContext); ok {
			tst[i] = t.(IEnumValueOptionContext)
			i++
		}
	}

	return tst
}

func (s *EnumValueOptionsContext) EnumValueOption(i int) IEnumValueOptionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumValueOptionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumValueOptionContext)
}

func (s *EnumValueOptionsContext) RB() antlr.TerminalNode {
	return s.GetToken(bubblerParserRB, 0)
}

func (s *EnumValueOptionsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(bubblerParserCOMMA)
}

func (s *EnumValueOptionsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(bubblerParserCOMMA, i)
}

func (s *EnumValueOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueOptionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitEnumValueOptions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) EnumValueOptions() (localctx IEnumValueOptionsContext) {
	localctx = NewEnumValueOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, bubblerParserRULE_enumValueOptions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		p.Match(bubblerParserLB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(283)
		p.EnumValueOption()
	}
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == bubblerParserCOMMA {
		{
			p.SetState(284)
			p.Match(bubblerParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(285)
			p.EnumValueOption()
		}

		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(291)
		p.Match(bubblerParserRB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnumValueOptionContext is an interface to support dynamic dispatch.
type IEnumValueOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OptionName() IOptionNameContext
	ASSIGN() antlr.TerminalNode
	Constant() IConstantContext

	// IsEnumValueOptionContext differentiates from other interfaces.
	IsEnumValueOptionContext()
}

type EnumValueOptionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueOptionContext() *EnumValueOptionContext {
	var p = new(EnumValueOptionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_enumValueOption
	return p
}

func InitEmptyEnumValueOptionContext(p *EnumValueOptionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_enumValueOption
}

func (*EnumValueOptionContext) IsEnumValueOptionContext() {}

func NewEnumValueOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueOptionContext {
	var p = new(EnumValueOptionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_enumValueOption

	return p
}

func (s *EnumValueOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueOptionContext) OptionName() IOptionNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *EnumValueOptionContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(bubblerParserASSIGN, 0)
}

func (s *EnumValueOptionContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *EnumValueOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueOptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitEnumValueOption(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) EnumValueOption() (localctx IEnumValueOptionContext) {
	localctx = NewEnumValueOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, bubblerParserRULE_enumValueOption)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		p.OptionName()
	}
	{
		p.SetState(294)
		p.Match(bubblerParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(295)
		p.Constant()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructDefContext is an interface to support dynamic dispatch.
type IStructDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StructName() IStructNameContext
	StructBody() IStructBodyContext
	Size_() ISize_Context

	// IsStructDefContext differentiates from other interfaces.
	IsStructDefContext()
}

type StructDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDefContext() *StructDefContext {
	var p = new(StructDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_structDef
	return p
}

func InitEmptyStructDefContext(p *StructDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_structDef
}

func (*StructDefContext) IsStructDefContext() {}

func NewStructDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDefContext {
	var p = new(StructDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_structDef

	return p
}

func (s *StructDefContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDefContext) StructName() IStructNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructNameContext)
}

func (s *StructDefContext) StructBody() IStructBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructBodyContext)
}

func (s *StructDefContext) Size_() ISize_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISize_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISize_Context)
}

func (s *StructDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitStructDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) StructDef() (localctx IStructDefContext) {
	localctx = NewStructDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, bubblerParserRULE_structDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(297)
		p.StructName()
	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == bubblerParserLB {
		{
			p.SetState(298)
			p.Size_()
		}

	}
	{
		p.SetState(301)
		p.StructBody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructBodyContext is an interface to support dynamic dispatch.
type IStructBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LC() antlr.TerminalNode
	RC() antlr.TerminalNode
	AllStructElement() []IStructElementContext
	StructElement(i int) IStructElementContext

	// IsStructBodyContext differentiates from other interfaces.
	IsStructBodyContext()
}

type StructBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructBodyContext() *StructBodyContext {
	var p = new(StructBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_structBody
	return p
}

func InitEmptyStructBodyContext(p *StructBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_structBody
}

func (*StructBodyContext) IsStructBodyContext() {}

func NewStructBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructBodyContext {
	var p = new(StructBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_structBody

	return p
}

func (s *StructBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *StructBodyContext) LC() antlr.TerminalNode {
	return s.GetToken(bubblerParserLC, 0)
}

func (s *StructBodyContext) RC() antlr.TerminalNode {
	return s.GetToken(bubblerParserRC, 0)
}

func (s *StructBodyContext) AllStructElement() []IStructElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructElementContext); ok {
			len++
		}
	}

	tst := make([]IStructElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructElementContext); ok {
			tst[i] = t.(IStructElementContext)
			i++
		}
	}

	return tst
}

func (s *StructBodyContext) StructElement(i int) IStructElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructElementContext)
}

func (s *StructBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitStructBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) StructBody() (localctx IStructBodyContext) {
	localctx = NewStructBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, bubblerParserRULE_structBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(303)
		p.Match(bubblerParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4611686018477719296) != 0 {
		{
			p.SetState(304)
			p.StructElement()
		}

		p.SetState(309)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(310)
		p.Match(bubblerParserRC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructElementContext is an interface to support dynamic dispatch.
type IStructElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Field() IFieldContext
	EmptyStatement_() IEmptyStatement_Context

	// IsStructElementContext differentiates from other interfaces.
	IsStructElementContext()
}

type StructElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructElementContext() *StructElementContext {
	var p = new(StructElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_structElement
	return p
}

func InitEmptyStructElementContext(p *StructElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_structElement
}

func (*StructElementContext) IsStructElementContext() {}

func NewStructElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructElementContext {
	var p = new(StructElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_structElement

	return p
}

func (s *StructElementContext) GetParser() antlr.Parser { return s.parser }

func (s *StructElementContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *StructElementContext) EmptyStatement_() IEmptyStatement_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *StructElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitStructElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) StructElement() (localctx IStructElementContext) {
	localctx = NewStructElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, bubblerParserRULE_structElement)
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case bubblerParserVOID, bubblerParserINT8, bubblerParserINT16, bubblerParserINT32, bubblerParserINT64, bubblerParserUINT8, bubblerParserUINT16, bubblerParserUINT32, bubblerParserUINT64, bubblerParserFLOAT32, bubblerParserFLOAT64, bubblerParserBOOL, bubblerParserSTRING, bubblerParserBYTES, bubblerParserENUM, bubblerParserSTRUCT, bubblerParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(312)
			p.Field()
		}

	case bubblerParserSEMI:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(313)
			p.EmptyStatement_()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExprMulDivModContext struct {
	ExprContext
	op antlr.Token
}

func NewExprMulDivModContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprMulDivModContext {
	var p = new(ExprMulDivModContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprMulDivModContext) GetOp() antlr.Token { return s.op }

func (s *ExprMulDivModContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprMulDivModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprMulDivModContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprMulDivModContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprMulDivModContext) MUL() antlr.TerminalNode {
	return s.GetToken(bubblerParserMUL, 0)
}

func (s *ExprMulDivModContext) DIV() antlr.TerminalNode {
	return s.GetToken(bubblerParserDIV, 0)
}

func (s *ExprMulDivModContext) MOD() antlr.TerminalNode {
	return s.GetToken(bubblerParserMOD, 0)
}

func (s *ExprMulDivModContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitExprMulDivMod(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprTernaryContext struct {
	ExprContext
}

func NewExprTernaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprTernaryContext {
	var p = new(ExprTernaryContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprTernaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprTernaryContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprTernaryContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprTernaryContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(bubblerParserQUESTION, 0)
}

func (s *ExprTernaryContext) COLON() antlr.TerminalNode {
	return s.GetToken(bubblerParserCOLON, 0)
}

func (s *ExprTernaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitExprTernary(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprBitXorContext struct {
	ExprContext
}

func NewExprBitXorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprBitXorContext {
	var p = new(ExprBitXorContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprBitXorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprBitXorContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprBitXorContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprBitXorContext) BXOR() antlr.TerminalNode {
	return s.GetToken(bubblerParserBXOR, 0)
}

func (s *ExprBitXorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitExprBitXor(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprParensContext struct {
	ExprContext
}

func NewExprParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprParensContext {
	var p = new(ExprParensContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprParensContext) LP() antlr.TerminalNode {
	return s.GetToken(bubblerParserLP, 0)
}

func (s *ExprParensContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprParensContext) RP() antlr.TerminalNode {
	return s.GetToken(bubblerParserRP, 0)
}

func (s *ExprParensContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitExprParens(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprConstantContext struct {
	ExprContext
}

func NewExprConstantContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprConstantContext {
	var p = new(ExprConstantContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprConstantContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ExprConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitExprConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprPowerContext struct {
	ExprContext
}

func NewExprPowerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprPowerContext {
	var p = new(ExprPowerContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprPowerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprPowerContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprPowerContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprPowerContext) POW() antlr.TerminalNode {
	return s.GetToken(bubblerParserPOW, 0)
}

func (s *ExprPowerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitExprPower(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprLogicalOrContext struct {
	ExprContext
}

func NewExprLogicalOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprLogicalOrContext {
	var p = new(ExprLogicalOrContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprLogicalOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprLogicalOrContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprLogicalOrContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprLogicalOrContext) OR() antlr.TerminalNode {
	return s.GetToken(bubblerParserOR, 0)
}

func (s *ExprLogicalOrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitExprLogicalOr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprRelationalContext struct {
	ExprContext
	op antlr.Token
}

func NewExprRelationalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprRelationalContext {
	var p = new(ExprRelationalContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprRelationalContext) GetOp() antlr.Token { return s.op }

func (s *ExprRelationalContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprRelationalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprRelationalContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprRelationalContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprRelationalContext) LT() antlr.TerminalNode {
	return s.GetToken(bubblerParserLT, 0)
}

func (s *ExprRelationalContext) LE() antlr.TerminalNode {
	return s.GetToken(bubblerParserLE, 0)
}

func (s *ExprRelationalContext) GT() antlr.TerminalNode {
	return s.GetToken(bubblerParserGT, 0)
}

func (s *ExprRelationalContext) GE() antlr.TerminalNode {
	return s.GetToken(bubblerParserGE, 0)
}

func (s *ExprRelationalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitExprRelational(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprShiftContext struct {
	ExprContext
	op antlr.Token
}

func NewExprShiftContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprShiftContext {
	var p = new(ExprShiftContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprShiftContext) GetOp() antlr.Token { return s.op }

func (s *ExprShiftContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprShiftContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprShiftContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprShiftContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprShiftContext) SHL() antlr.TerminalNode {
	return s.GetToken(bubblerParserSHL, 0)
}

func (s *ExprShiftContext) SHR() antlr.TerminalNode {
	return s.GetToken(bubblerParserSHR, 0)
}

func (s *ExprShiftContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitExprShift(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprBitOrContext struct {
	ExprContext
}

func NewExprBitOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprBitOrContext {
	var p = new(ExprBitOrContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprBitOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprBitOrContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprBitOrContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprBitOrContext) BOR() antlr.TerminalNode {
	return s.GetToken(bubblerParserBOR, 0)
}

func (s *ExprBitOrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitExprBitOr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprAddSubContext struct {
	ExprContext
	op antlr.Token
}

func NewExprAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprAddSubContext {
	var p = new(ExprAddSubContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprAddSubContext) GetOp() antlr.Token { return s.op }

func (s *ExprAddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprAddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprAddSubContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprAddSubContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprAddSubContext) ADD() antlr.TerminalNode {
	return s.GetToken(bubblerParserADD, 0)
}

func (s *ExprAddSubContext) SUB() antlr.TerminalNode {
	return s.GetToken(bubblerParserSUB, 0)
}

func (s *ExprAddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitExprAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprValueContext struct {
	ExprContext
}

func NewExprValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprValueContext {
	var p = new(ExprValueContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprValueContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ExprValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitExprValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprCastContext struct {
	ExprContext
}

func NewExprCastContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprCastContext {
	var p = new(ExprCastContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprCastContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprCastContext) LP() antlr.TerminalNode {
	return s.GetToken(bubblerParserLP, 0)
}

func (s *ExprCastContext) BasicType() IBasicTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBasicTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBasicTypeContext)
}

func (s *ExprCastContext) RP() antlr.TerminalNode {
	return s.GetToken(bubblerParserRP, 0)
}

func (s *ExprCastContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprCastContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitExprCast(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprUnaryContext struct {
	ExprContext
	op antlr.Token
}

func NewExprUnaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprUnaryContext {
	var p = new(ExprUnaryContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprUnaryContext) GetOp() antlr.Token { return s.op }

func (s *ExprUnaryContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprUnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprUnaryContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprUnaryContext) ADD() antlr.TerminalNode {
	return s.GetToken(bubblerParserADD, 0)
}

func (s *ExprUnaryContext) SUB() antlr.TerminalNode {
	return s.GetToken(bubblerParserSUB, 0)
}

func (s *ExprUnaryContext) BNOT() antlr.TerminalNode {
	return s.GetToken(bubblerParserBNOT, 0)
}

func (s *ExprUnaryContext) NOT() antlr.TerminalNode {
	return s.GetToken(bubblerParserNOT, 0)
}

func (s *ExprUnaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitExprUnary(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprBitAndContext struct {
	ExprContext
}

func NewExprBitAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprBitAndContext {
	var p = new(ExprBitAndContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprBitAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprBitAndContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprBitAndContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprBitAndContext) BAND() antlr.TerminalNode {
	return s.GetToken(bubblerParserBAND, 0)
}

func (s *ExprBitAndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitExprBitAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprEqualityContext struct {
	ExprContext
	op antlr.Token
}

func NewExprEqualityContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprEqualityContext {
	var p = new(ExprEqualityContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprEqualityContext) GetOp() antlr.Token { return s.op }

func (s *ExprEqualityContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprEqualityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprEqualityContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprEqualityContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprEqualityContext) EQ() antlr.TerminalNode {
	return s.GetToken(bubblerParserEQ, 0)
}

func (s *ExprEqualityContext) NE() antlr.TerminalNode {
	return s.GetToken(bubblerParserNE, 0)
}

func (s *ExprEqualityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitExprEquality(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprLogicalAndContext struct {
	ExprContext
}

func NewExprLogicalAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprLogicalAndContext {
	var p = new(ExprLogicalAndContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprLogicalAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprLogicalAndContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprLogicalAndContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprLogicalAndContext) AND() antlr.TerminalNode {
	return s.GetToken(bubblerParserAND, 0)
}

func (s *ExprLogicalAndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitExprLogicalAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *bubblerParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 62
	p.EnterRecursionRule(localctx, 62, bubblerParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExprValueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(317)
			p.Value()
		}

	case 2:
		localctx = NewExprConstantContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(318)
			p.Constant()
		}

	case 3:
		localctx = NewExprParensContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(319)
			p.Match(bubblerParserLP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(320)
			p.expr(0)
		}
		{
			p.SetState(321)
			p.Match(bubblerParserRP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewExprUnaryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(323)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExprUnaryContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&162155974864404480) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExprUnaryContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(324)
			p.expr(13)
		}

	case 5:
		localctx = NewExprCastContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(325)
			p.Match(bubblerParserLP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(326)
			p.BasicType()
		}
		{
			p.SetState(327)
			p.Match(bubblerParserRP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(328)
			p.expr(12)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(373)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(371)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprPowerContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, bubblerParserRULE_expr)
				p.SetState(332)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(333)
					p.Match(bubblerParserPOW)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(334)
					p.expr(15)
				}

			case 2:
				localctx = NewExprMulDivModContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, bubblerParserRULE_expr)
				p.SetState(335)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(336)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprMulDivModContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&246290604621824) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprMulDivModContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(337)
					p.expr(12)
				}

			case 3:
				localctx = NewExprAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, bubblerParserRULE_expr)
				p.SetState(338)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(339)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprAddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == bubblerParserADD || _la == bubblerParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprAddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(340)
					p.expr(11)
				}

			case 4:
				localctx = NewExprShiftContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, bubblerParserRULE_expr)
				p.SetState(341)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(342)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprShiftContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == bubblerParserSHL || _la == bubblerParserSHR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprShiftContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(343)
					p.expr(10)
				}

			case 5:
				localctx = NewExprRelationalContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, bubblerParserRULE_expr)
				p.SetState(344)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(345)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprRelationalContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&257698037760) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprRelationalContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(346)
					p.expr(9)
				}

			case 6:
				localctx = NewExprEqualityContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, bubblerParserRULE_expr)
				p.SetState(347)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(348)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprEqualityContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == bubblerParserEQ || _la == bubblerParserNE) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprEqualityContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(349)
					p.expr(8)
				}

			case 7:
				localctx = NewExprBitAndContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, bubblerParserRULE_expr)
				p.SetState(350)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(351)
					p.Match(bubblerParserBAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(352)
					p.expr(7)
				}

			case 8:
				localctx = NewExprBitXorContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, bubblerParserRULE_expr)
				p.SetState(353)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(354)
					p.Match(bubblerParserBXOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(355)
					p.expr(6)
				}

			case 9:
				localctx = NewExprBitOrContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, bubblerParserRULE_expr)
				p.SetState(356)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(357)
					p.Match(bubblerParserBOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(358)
					p.expr(5)
				}

			case 10:
				localctx = NewExprLogicalAndContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, bubblerParserRULE_expr)
				p.SetState(359)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(360)
					p.Match(bubblerParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(361)
					p.expr(4)
				}

			case 11:
				localctx = NewExprLogicalOrContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, bubblerParserRULE_expr)
				p.SetState(362)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(363)
					p.Match(bubblerParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(364)
					p.expr(3)
				}

			case 12:
				localctx = NewExprTernaryContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, bubblerParserRULE_expr)
				p.SetState(365)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(366)
					p.Match(bubblerParserQUESTION)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(367)
					p.expr(0)
				}
				{
					p.SetState(368)
					p.Match(bubblerParserCOLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(369)
					p.expr(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(375)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VALUE() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) VALUE() antlr.TerminalNode {
	return s.GetToken(bubblerParserVALUE, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, bubblerParserRULE_value)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(376)
		p.Match(bubblerParserVALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IntLit() IIntLitContext
	SUB() antlr.TerminalNode
	ADD() antlr.TerminalNode
	FloatLit() IFloatLitContext
	StrLit() IStrLitContext
	BoolLit() IBoolLitContext

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_constant
	return p
}

func InitEmptyConstantContext(p *ConstantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_constant
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) IntLit() IIntLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntLitContext)
}

func (s *ConstantContext) SUB() antlr.TerminalNode {
	return s.GetToken(bubblerParserSUB, 0)
}

func (s *ConstantContext) ADD() antlr.TerminalNode {
	return s.GetToken(bubblerParserADD, 0)
}

func (s *ConstantContext) FloatLit() IFloatLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloatLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloatLitContext)
}

func (s *ConstantContext) StrLit() IStrLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrLitContext)
}

func (s *ConstantContext) BoolLit() IBoolLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolLitContext)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, bubblerParserRULE_constant)
	var _la int

	p.SetState(388)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(379)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == bubblerParserADD || _la == bubblerParserSUB {
			{
				p.SetState(378)
				_la = p.GetTokenStream().LA(1)

				if !(_la == bubblerParserADD || _la == bubblerParserSUB) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(381)
			p.IntLit()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(383)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == bubblerParserADD || _la == bubblerParserSUB {
			{
				p.SetState(382)
				_la = p.GetTokenStream().LA(1)

				if !(_la == bubblerParserADD || _la == bubblerParserSUB) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(385)
			p.FloatLit()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(386)
			p.StrLit()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(387)
			p.BoolLit()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEmptyStatement_Context is an interface to support dynamic dispatch.
type IEmptyStatement_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SEMI() antlr.TerminalNode

	// IsEmptyStatement_Context differentiates from other interfaces.
	IsEmptyStatement_Context()
}

type EmptyStatement_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmptyStatement_Context() *EmptyStatement_Context {
	var p = new(EmptyStatement_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_emptyStatement_
	return p
}

func InitEmptyEmptyStatement_Context(p *EmptyStatement_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_emptyStatement_
}

func (*EmptyStatement_Context) IsEmptyStatement_Context() {}

func NewEmptyStatement_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptyStatement_Context {
	var p = new(EmptyStatement_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_emptyStatement_

	return p
}

func (s *EmptyStatement_Context) GetParser() antlr.Parser { return s.parser }

func (s *EmptyStatement_Context) SEMI() antlr.TerminalNode {
	return s.GetToken(bubblerParserSEMI, 0)
}

func (s *EmptyStatement_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyStatement_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmptyStatement_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitEmptyStatement_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) EmptyStatement_() (localctx IEmptyStatement_Context) {
	localctx = NewEmptyStatement_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, bubblerParserRULE_emptyStatement_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(390)
		p.Match(bubblerParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentContext is an interface to support dynamic dispatch.
type IIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsIdentContext differentiates from other interfaces.
	IsIdentContext()
}

type IdentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentContext() *IdentContext {
	var p = new(IdentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_ident
	return p
}

func InitEmptyIdentContext(p *IdentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_ident
}

func (*IdentContext) IsIdentContext() {}

func NewIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentContext {
	var p = new(IdentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_ident

	return p
}

func (s *IdentContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(bubblerParserIDENTIFIER, 0)
}

func (s *IdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) Ident() (localctx IIdentContext) {
	localctx = NewIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, bubblerParserRULE_ident)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(392)
		p.Match(bubblerParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFullIdentContext is an interface to support dynamic dispatch.
type IFullIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdent() []IIdentContext
	Ident(i int) IIdentContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsFullIdentContext differentiates from other interfaces.
	IsFullIdentContext()
}

type FullIdentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFullIdentContext() *FullIdentContext {
	var p = new(FullIdentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_fullIdent
	return p
}

func InitEmptyFullIdentContext(p *FullIdentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_fullIdent
}

func (*FullIdentContext) IsFullIdentContext() {}

func NewFullIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FullIdentContext {
	var p = new(FullIdentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_fullIdent

	return p
}

func (s *FullIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *FullIdentContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *FullIdentContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *FullIdentContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(bubblerParserDOT)
}

func (s *FullIdentContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(bubblerParserDOT, i)
}

func (s *FullIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FullIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitFullIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) FullIdent() (localctx IFullIdentContext) {
	localctx = NewFullIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, bubblerParserRULE_fullIdent)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(394)
		p.Ident()
	}
	p.SetState(399)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == bubblerParserDOT {
		{
			p.SetState(395)
			p.Match(bubblerParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(396)
			p.Ident()
		}

		p.SetState(401)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_fieldName
	return p
}

func InitEmptyFieldNameContext(p *FieldNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_fieldName
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitFieldName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) FieldName() (localctx IFieldNameContext) {
	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, bubblerParserRULE_fieldName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(402)
		p.Ident()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMethodNameContext is an interface to support dynamic dispatch.
type IMethodNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext

	// IsMethodNameContext differentiates from other interfaces.
	IsMethodNameContext()
}

type MethodNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodNameContext() *MethodNameContext {
	var p = new(MethodNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_methodName
	return p
}

func InitEmptyMethodNameContext(p *MethodNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_methodName
}

func (*MethodNameContext) IsMethodNameContext() {}

func NewMethodNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodNameContext {
	var p = new(MethodNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_methodName

	return p
}

func (s *MethodNameContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *MethodNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitMethodName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) MethodName() (localctx IMethodNameContext) {
	localctx = NewMethodNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, bubblerParserRULE_methodName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(404)
		p.Ident()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructNameContext is an interface to support dynamic dispatch.
type IStructNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRUCT() antlr.TerminalNode
	Ident() IIdentContext

	// IsStructNameContext differentiates from other interfaces.
	IsStructNameContext()
}

type StructNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructNameContext() *StructNameContext {
	var p = new(StructNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_structName
	return p
}

func InitEmptyStructNameContext(p *StructNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_structName
}

func (*StructNameContext) IsStructNameContext() {}

func NewStructNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructNameContext {
	var p = new(StructNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_structName

	return p
}

func (s *StructNameContext) GetParser() antlr.Parser { return s.parser }

func (s *StructNameContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(bubblerParserSTRUCT, 0)
}

func (s *StructNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *StructNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitStructName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) StructName() (localctx IStructNameContext) {
	localctx = NewStructNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, bubblerParserRULE_structName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(406)
		p.Match(bubblerParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(407)
		p.Ident()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnumNameContext is an interface to support dynamic dispatch.
type IEnumNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENUM() antlr.TerminalNode
	Ident() IIdentContext

	// IsEnumNameContext differentiates from other interfaces.
	IsEnumNameContext()
}

type EnumNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumNameContext() *EnumNameContext {
	var p = new(EnumNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_enumName
	return p
}

func InitEmptyEnumNameContext(p *EnumNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_enumName
}

func (*EnumNameContext) IsEnumNameContext() {}

func NewEnumNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumNameContext {
	var p = new(EnumNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_enumName

	return p
}

func (s *EnumNameContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumNameContext) ENUM() antlr.TerminalNode {
	return s.GetToken(bubblerParserENUM, 0)
}

func (s *EnumNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *EnumNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitEnumName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) EnumName() (localctx IEnumNameContext) {
	localctx = NewEnumNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, bubblerParserRULE_enumName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(409)
		p.Match(bubblerParserENUM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(410)
		p.Ident()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnumValueNameContext is an interface to support dynamic dispatch.
type IEnumValueNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext

	// IsEnumValueNameContext differentiates from other interfaces.
	IsEnumValueNameContext()
}

type EnumValueNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueNameContext() *EnumValueNameContext {
	var p = new(EnumValueNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_enumValueName
	return p
}

func InitEmptyEnumValueNameContext(p *EnumValueNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_enumValueName
}

func (*EnumValueNameContext) IsEnumValueNameContext() {}

func NewEnumValueNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueNameContext {
	var p = new(EnumValueNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_enumValueName

	return p
}

func (s *EnumValueNameContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *EnumValueNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitEnumValueName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) EnumValueName() (localctx IEnumValueNameContext) {
	localctx = NewEnumValueNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, bubblerParserRULE_enumValueName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(412)
		p.Ident()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructTypeContext is an interface to support dynamic dispatch.
type IStructTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StructName() IStructNameContext
	StructDef() IStructDefContext

	// IsStructTypeContext differentiates from other interfaces.
	IsStructTypeContext()
}

type StructTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructTypeContext() *StructTypeContext {
	var p = new(StructTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_structType
	return p
}

func InitEmptyStructTypeContext(p *StructTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_structType
}

func (*StructTypeContext) IsStructTypeContext() {}

func NewStructTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructTypeContext {
	var p = new(StructTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_structType

	return p
}

func (s *StructTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *StructTypeContext) StructName() IStructNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructNameContext)
}

func (s *StructTypeContext) StructDef() IStructDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDefContext)
}

func (s *StructTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitStructType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) StructType() (localctx IStructTypeContext) {
	localctx = NewStructTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, bubblerParserRULE_structType)
	p.SetState(416)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(414)
			p.StructName()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(415)
			p.StructDef()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnumTypeContext is an interface to support dynamic dispatch.
type IEnumTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EnumName() IEnumNameContext

	// IsEnumTypeContext differentiates from other interfaces.
	IsEnumTypeContext()
}

type EnumTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumTypeContext() *EnumTypeContext {
	var p = new(EnumTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_enumType
	return p
}

func InitEmptyEnumTypeContext(p *EnumTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_enumType
}

func (*EnumTypeContext) IsEnumTypeContext() {}

func NewEnumTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumTypeContext {
	var p = new(EnumTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_enumType

	return p
}

func (s *EnumTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumTypeContext) EnumName() IEnumNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumNameContext)
}

func (s *EnumTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitEnumType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) EnumType() (localctx IEnumTypeContext) {
	localctx = NewEnumTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, bubblerParserRULE_enumType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(418)
		p.EnumName()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntLitContext is an interface to support dynamic dispatch.
type IIntLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT_LIT() antlr.TerminalNode

	// IsIntLitContext differentiates from other interfaces.
	IsIntLitContext()
}

type IntLitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntLitContext() *IntLitContext {
	var p = new(IntLitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_intLit
	return p
}

func InitEmptyIntLitContext(p *IntLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_intLit
}

func (*IntLitContext) IsIntLitContext() {}

func NewIntLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntLitContext {
	var p = new(IntLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_intLit

	return p
}

func (s *IntLitContext) GetParser() antlr.Parser { return s.parser }

func (s *IntLitContext) INT_LIT() antlr.TerminalNode {
	return s.GetToken(bubblerParserINT_LIT, 0)
}

func (s *IntLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitIntLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) IntLit() (localctx IIntLitContext) {
	localctx = NewIntLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, bubblerParserRULE_intLit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(420)
		p.Match(bubblerParserINT_LIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStrLitContext is an interface to support dynamic dispatch.
type IStrLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STR_LIT() antlr.TerminalNode

	// IsStrLitContext differentiates from other interfaces.
	IsStrLitContext()
}

type StrLitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrLitContext() *StrLitContext {
	var p = new(StrLitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_strLit
	return p
}

func InitEmptyStrLitContext(p *StrLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_strLit
}

func (*StrLitContext) IsStrLitContext() {}

func NewStrLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrLitContext {
	var p = new(StrLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_strLit

	return p
}

func (s *StrLitContext) GetParser() antlr.Parser { return s.parser }

func (s *StrLitContext) STR_LIT() antlr.TerminalNode {
	return s.GetToken(bubblerParserSTR_LIT, 0)
}

func (s *StrLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StrLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitStrLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) StrLit() (localctx IStrLitContext) {
	localctx = NewStrLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, bubblerParserRULE_strLit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(422)
		p.Match(bubblerParserSTR_LIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBoolLitContext is an interface to support dynamic dispatch.
type IBoolLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BOOL_LIT() antlr.TerminalNode

	// IsBoolLitContext differentiates from other interfaces.
	IsBoolLitContext()
}

type BoolLitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolLitContext() *BoolLitContext {
	var p = new(BoolLitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_boolLit
	return p
}

func InitEmptyBoolLitContext(p *BoolLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_boolLit
}

func (*BoolLitContext) IsBoolLitContext() {}

func NewBoolLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolLitContext {
	var p = new(BoolLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_boolLit

	return p
}

func (s *BoolLitContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolLitContext) BOOL_LIT() antlr.TerminalNode {
	return s.GetToken(bubblerParserBOOL_LIT, 0)
}

func (s *BoolLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitBoolLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) BoolLit() (localctx IBoolLitContext) {
	localctx = NewBoolLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, bubblerParserRULE_boolLit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(424)
		p.Match(bubblerParserBOOL_LIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFloatLitContext is an interface to support dynamic dispatch.
type IFloatLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FLOAT_LIT() antlr.TerminalNode

	// IsFloatLitContext differentiates from other interfaces.
	IsFloatLitContext()
}

type FloatLitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatLitContext() *FloatLitContext {
	var p = new(FloatLitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_floatLit
	return p
}

func InitEmptyFloatLitContext(p *FloatLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_floatLit
}

func (*FloatLitContext) IsFloatLitContext() {}

func NewFloatLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatLitContext {
	var p = new(FloatLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_floatLit

	return p
}

func (s *FloatLitContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatLitContext) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(bubblerParserFLOAT_LIT, 0)
}

func (s *FloatLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitFloatLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) FloatLit() (localctx IFloatLitContext) {
	localctx = NewFloatLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, bubblerParserRULE_floatLit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(426)
		p.Match(bubblerParserFLOAT_LIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *bubblerParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 31:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *bubblerParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
