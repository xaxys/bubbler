// Code generated from D:/GoProject/bubbler/tools/../bubbler.g4 by ANTLR 4.13.1. DO NOT EDIT.

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
		"", "'import'", "'get'", "'set'", "'value'", "'void'", "'int8'", "'int16'",
		"'int32'", "'int64'", "'uint8'", "'uint16'", "'uint32'", "'uint64'",
		"'float32'", "'float64'", "'bool'", "'string'", "'bytes'", "'enum'",
		"'struct'", "'#'", "';'", "'='", "'?'", "'('", "')'", "'['", "']'",
		"'{'", "'}'", "'<'", "'<='", "'>'", "'>='", "'=='", "'!='", "'.'", "','",
		"':'", "'+'", "'-'", "'*'", "'/'", "'%'", "'**'", "'<<'", "'>>'", "'&'",
		"'|'", "'^'", "'~'", "'&&'", "'||'", "'!'",
	}
	staticData.SymbolicNames = []string{
		"", "IMPORT", "GET", "SET", "VALUE", "VOID", "INT8", "INT16", "INT32",
		"INT64", "UINT8", "UINT16", "UINT32", "UINT64", "FLOAT32", "FLOAT64",
		"BOOL", "STRING", "BYTES", "ENUM", "STRUCT", "HASH", "SEMI", "ASSIGN",
		"QUESTION", "LP", "RP", "LB", "RB", "LC", "RC", "LT", "LE", "GT", "GE",
		"EQ", "NE", "DOT", "COMMA", "COLON", "ADD", "SUB", "MUL", "DIV", "MOD",
		"POW", "SHL", "SHR", "BAND", "BOR", "BXOR", "BNOT", "AND", "OR", "NOT",
		"STR_LIT", "BOOL_LIT", "FLOAT_LIT", "INT_LIT", "IDENTIFIER", "WS", "LINE_COMMENT",
		"COMMENT", "KEYWORDS",
	}
	staticData.RuleNames = []string{
		"proto", "importStatement", "topLevelDef", "size_", "byteSize", "bitSize",
		"optionName", "field", "fieldVoid", "fieldConstant", "fieldEmbedded",
		"fieldNormal", "fieldOptions", "fieldOption", "fieldMethods", "fieldMethod",
		"type_", "basicType", "enumDef", "enumBody", "enumElement", "enumField",
		"enumValueOptions", "enumValueOption", "structDef", "structBody", "structElement",
		"expr", "value", "constant", "emptyStatement_", "ident", "structName",
		"enumName", "fieldName", "methodName", "structType", "enumType", "arrayType",
		"intLit", "strLit", "boolLit", "floatLit",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 63, 384, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 1, 0, 1, 0, 1, 0, 5, 0, 90, 8, 0, 10, 0, 12, 0, 93, 9, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 3, 2, 103, 8, 2, 1, 3, 1,
		3, 1, 3, 1, 3, 3, 3, 109, 8, 3, 1, 3, 1, 3, 3, 3, 113, 8, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 127,
		8, 7, 1, 8, 1, 8, 1, 8, 3, 8, 132, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 3, 9,
		138, 8, 9, 1, 9, 3, 9, 141, 8, 9, 1, 9, 1, 9, 1, 9, 3, 9, 146, 8, 9, 1,
		9, 1, 9, 1, 10, 1, 10, 3, 10, 152, 8, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1,
		11, 3, 11, 159, 8, 11, 1, 11, 3, 11, 162, 8, 11, 1, 11, 3, 11, 165, 8,
		11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 173, 8, 12, 10, 12,
		12, 12, 176, 9, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1,
		14, 5, 14, 186, 8, 14, 10, 14, 12, 14, 189, 9, 14, 1, 14, 1, 14, 1, 15,
		1, 15, 3, 15, 195, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 211, 8, 16,
		1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 5, 19, 221, 8,
		19, 10, 19, 12, 19, 224, 9, 19, 1, 19, 1, 19, 1, 20, 1, 20, 3, 20, 230,
		8, 20, 1, 21, 1, 21, 1, 21, 3, 21, 235, 8, 21, 1, 21, 3, 21, 238, 8, 21,
		1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 246, 8, 22, 10, 22, 12,
		22, 249, 9, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24,
		3, 24, 259, 8, 24, 1, 24, 1, 24, 1, 25, 1, 25, 5, 25, 265, 8, 25, 10, 25,
		12, 25, 268, 9, 25, 1, 25, 1, 25, 1, 26, 1, 26, 3, 26, 274, 8, 26, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 3, 27, 290, 8, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 5, 27, 331, 8, 27, 10, 27, 12, 27, 334, 9, 27, 1, 28,
		1, 28, 1, 29, 1, 29, 3, 29, 340, 8, 29, 1, 29, 1, 29, 3, 29, 344, 8, 29,
		1, 29, 1, 29, 1, 29, 3, 29, 349, 8, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1,
		32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36,
		1, 36, 3, 36, 367, 8, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1,
		38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 0, 1,
		54, 43, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68,
		70, 72, 74, 76, 78, 80, 82, 84, 0, 9, 1, 0, 2, 3, 1, 0, 6, 16, 2, 0, 22,
		22, 38, 38, 3, 0, 40, 41, 51, 51, 54, 54, 1, 0, 42, 44, 1, 0, 40, 41, 1,
		0, 46, 47, 1, 0, 31, 34, 1, 0, 35, 36, 397, 0, 91, 1, 0, 0, 0, 2, 96, 1,
		0, 0, 0, 4, 102, 1, 0, 0, 0, 6, 104, 1, 0, 0, 0, 8, 116, 1, 0, 0, 0, 10,
		118, 1, 0, 0, 0, 12, 120, 1, 0, 0, 0, 14, 126, 1, 0, 0, 0, 16, 128, 1,
		0, 0, 0, 18, 135, 1, 0, 0, 0, 20, 149, 1, 0, 0, 0, 22, 155, 1, 0, 0, 0,
		24, 168, 1, 0, 0, 0, 26, 179, 1, 0, 0, 0, 28, 183, 1, 0, 0, 0, 30, 192,
		1, 0, 0, 0, 32, 210, 1, 0, 0, 0, 34, 212, 1, 0, 0, 0, 36, 214, 1, 0, 0,
		0, 38, 218, 1, 0, 0, 0, 40, 229, 1, 0, 0, 0, 42, 231, 1, 0, 0, 0, 44, 241,
		1, 0, 0, 0, 46, 252, 1, 0, 0, 0, 48, 256, 1, 0, 0, 0, 50, 262, 1, 0, 0,
		0, 52, 273, 1, 0, 0, 0, 54, 289, 1, 0, 0, 0, 56, 335, 1, 0, 0, 0, 58, 348,
		1, 0, 0, 0, 60, 350, 1, 0, 0, 0, 62, 352, 1, 0, 0, 0, 64, 354, 1, 0, 0,
		0, 66, 357, 1, 0, 0, 0, 68, 360, 1, 0, 0, 0, 70, 362, 1, 0, 0, 0, 72, 366,
		1, 0, 0, 0, 74, 368, 1, 0, 0, 0, 76, 370, 1, 0, 0, 0, 78, 375, 1, 0, 0,
		0, 80, 377, 1, 0, 0, 0, 82, 379, 1, 0, 0, 0, 84, 381, 1, 0, 0, 0, 86, 90,
		3, 2, 1, 0, 87, 90, 3, 4, 2, 0, 88, 90, 3, 60, 30, 0, 89, 86, 1, 0, 0,
		0, 89, 87, 1, 0, 0, 0, 89, 88, 1, 0, 0, 0, 90, 93, 1, 0, 0, 0, 91, 89,
		1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 94, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0,
		94, 95, 5, 0, 0, 1, 95, 1, 1, 0, 0, 0, 96, 97, 5, 1, 0, 0, 97, 98, 3, 80,
		40, 0, 98, 99, 5, 22, 0, 0, 99, 3, 1, 0, 0, 0, 100, 103, 3, 36, 18, 0,
		101, 103, 3, 48, 24, 0, 102, 100, 1, 0, 0, 0, 102, 101, 1, 0, 0, 0, 103,
		5, 1, 0, 0, 0, 104, 112, 5, 27, 0, 0, 105, 108, 3, 8, 4, 0, 106, 107, 5,
		21, 0, 0, 107, 109, 3, 10, 5, 0, 108, 106, 1, 0, 0, 0, 108, 109, 1, 0,
		0, 0, 109, 113, 1, 0, 0, 0, 110, 111, 5, 21, 0, 0, 111, 113, 3, 10, 5,
		0, 112, 105, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114,
		115, 5, 28, 0, 0, 115, 7, 1, 0, 0, 0, 116, 117, 3, 78, 39, 0, 117, 9, 1,
		0, 0, 0, 118, 119, 3, 78, 39, 0, 119, 11, 1, 0, 0, 0, 120, 121, 3, 62,
		31, 0, 121, 13, 1, 0, 0, 0, 122, 127, 3, 16, 8, 0, 123, 127, 3, 18, 9,
		0, 124, 127, 3, 20, 10, 0, 125, 127, 3, 22, 11, 0, 126, 122, 1, 0, 0, 0,
		126, 123, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 126, 125, 1, 0, 0, 0, 127,
		15, 1, 0, 0, 0, 128, 129, 5, 5, 0, 0, 129, 131, 3, 6, 3, 0, 130, 132, 3,
		24, 12, 0, 131, 130, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 133, 1, 0,
		0, 0, 133, 134, 5, 22, 0, 0, 134, 17, 1, 0, 0, 0, 135, 137, 3, 34, 17,
		0, 136, 138, 3, 68, 34, 0, 137, 136, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0,
		138, 140, 1, 0, 0, 0, 139, 141, 3, 6, 3, 0, 140, 139, 1, 0, 0, 0, 140,
		141, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 143, 5, 23, 0, 0, 143, 145,
		3, 58, 29, 0, 144, 146, 3, 24, 12, 0, 145, 144, 1, 0, 0, 0, 145, 146, 1,
		0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 148, 5, 22, 0, 0, 148, 19, 1, 0, 0,
		0, 149, 151, 3, 32, 16, 0, 150, 152, 3, 24, 12, 0, 151, 150, 1, 0, 0, 0,
		151, 152, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 154, 5, 22, 0, 0, 154,
		21, 1, 0, 0, 0, 155, 156, 3, 32, 16, 0, 156, 158, 3, 68, 34, 0, 157, 159,
		3, 6, 3, 0, 158, 157, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 161, 1, 0,
		0, 0, 160, 162, 3, 24, 12, 0, 161, 160, 1, 0, 0, 0, 161, 162, 1, 0, 0,
		0, 162, 164, 1, 0, 0, 0, 163, 165, 3, 28, 14, 0, 164, 163, 1, 0, 0, 0,
		164, 165, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 167, 5, 22, 0, 0, 167,
		23, 1, 0, 0, 0, 168, 169, 5, 27, 0, 0, 169, 174, 3, 26, 13, 0, 170, 171,
		5, 38, 0, 0, 171, 173, 3, 26, 13, 0, 172, 170, 1, 0, 0, 0, 173, 176, 1,
		0, 0, 0, 174, 172, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 177, 1, 0, 0,
		0, 176, 174, 1, 0, 0, 0, 177, 178, 5, 28, 0, 0, 178, 25, 1, 0, 0, 0, 179,
		180, 3, 12, 6, 0, 180, 181, 5, 23, 0, 0, 181, 182, 3, 58, 29, 0, 182, 27,
		1, 0, 0, 0, 183, 187, 5, 29, 0, 0, 184, 186, 3, 30, 15, 0, 185, 184, 1,
		0, 0, 0, 186, 189, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 187, 188, 1, 0, 0,
		0, 188, 190, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 190, 191, 5, 30, 0, 0, 191,
		29, 1, 0, 0, 0, 192, 194, 7, 0, 0, 0, 193, 195, 3, 70, 35, 0, 194, 193,
		1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 197, 5, 25,
		0, 0, 197, 198, 3, 34, 17, 0, 198, 199, 5, 26, 0, 0, 199, 200, 5, 39, 0,
		0, 200, 201, 3, 54, 27, 0, 201, 202, 5, 22, 0, 0, 202, 31, 1, 0, 0, 0,
		203, 211, 3, 34, 17, 0, 204, 211, 5, 17, 0, 0, 205, 211, 5, 18, 0, 0, 206,
		211, 3, 76, 38, 0, 207, 211, 3, 72, 36, 0, 208, 211, 3, 74, 37, 0, 209,
		211, 3, 62, 31, 0, 210, 203, 1, 0, 0, 0, 210, 204, 1, 0, 0, 0, 210, 205,
		1, 0, 0, 0, 210, 206, 1, 0, 0, 0, 210, 207, 1, 0, 0, 0, 210, 208, 1, 0,
		0, 0, 210, 209, 1, 0, 0, 0, 211, 33, 1, 0, 0, 0, 212, 213, 7, 1, 0, 0,
		213, 35, 1, 0, 0, 0, 214, 215, 3, 66, 33, 0, 215, 216, 3, 6, 3, 0, 216,
		217, 3, 38, 19, 0, 217, 37, 1, 0, 0, 0, 218, 222, 5, 29, 0, 0, 219, 221,
		3, 40, 20, 0, 220, 219, 1, 0, 0, 0, 221, 224, 1, 0, 0, 0, 222, 220, 1,
		0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 225, 1, 0, 0, 0, 224, 222, 1, 0, 0,
		0, 225, 226, 5, 30, 0, 0, 226, 39, 1, 0, 0, 0, 227, 230, 3, 42, 21, 0,
		228, 230, 3, 60, 30, 0, 229, 227, 1, 0, 0, 0, 229, 228, 1, 0, 0, 0, 230,
		41, 1, 0, 0, 0, 231, 234, 3, 62, 31, 0, 232, 233, 5, 23, 0, 0, 233, 235,
		3, 78, 39, 0, 234, 232, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 237, 1,
		0, 0, 0, 236, 238, 3, 44, 22, 0, 237, 236, 1, 0, 0, 0, 237, 238, 1, 0,
		0, 0, 238, 239, 1, 0, 0, 0, 239, 240, 7, 2, 0, 0, 240, 43, 1, 0, 0, 0,
		241, 242, 5, 27, 0, 0, 242, 247, 3, 46, 23, 0, 243, 244, 5, 38, 0, 0, 244,
		246, 3, 46, 23, 0, 245, 243, 1, 0, 0, 0, 246, 249, 1, 0, 0, 0, 247, 245,
		1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 250, 1, 0, 0, 0, 249, 247, 1, 0,
		0, 0, 250, 251, 5, 28, 0, 0, 251, 45, 1, 0, 0, 0, 252, 253, 3, 12, 6, 0,
		253, 254, 5, 23, 0, 0, 254, 255, 3, 58, 29, 0, 255, 47, 1, 0, 0, 0, 256,
		258, 3, 64, 32, 0, 257, 259, 3, 6, 3, 0, 258, 257, 1, 0, 0, 0, 258, 259,
		1, 0, 0, 0, 259, 260, 1, 0, 0, 0, 260, 261, 3, 50, 25, 0, 261, 49, 1, 0,
		0, 0, 262, 266, 5, 29, 0, 0, 263, 265, 3, 52, 26, 0, 264, 263, 1, 0, 0,
		0, 265, 268, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267,
		269, 1, 0, 0, 0, 268, 266, 1, 0, 0, 0, 269, 270, 5, 30, 0, 0, 270, 51,
		1, 0, 0, 0, 271, 274, 3, 14, 7, 0, 272, 274, 3, 60, 30, 0, 273, 271, 1,
		0, 0, 0, 273, 272, 1, 0, 0, 0, 274, 53, 1, 0, 0, 0, 275, 276, 6, 27, -1,
		0, 276, 290, 3, 56, 28, 0, 277, 290, 3, 58, 29, 0, 278, 279, 5, 25, 0,
		0, 279, 280, 3, 54, 27, 0, 280, 281, 5, 26, 0, 0, 281, 290, 1, 0, 0, 0,
		282, 283, 7, 3, 0, 0, 283, 290, 3, 54, 27, 13, 284, 285, 5, 25, 0, 0, 285,
		286, 3, 34, 17, 0, 286, 287, 5, 26, 0, 0, 287, 288, 3, 54, 27, 12, 288,
		290, 1, 0, 0, 0, 289, 275, 1, 0, 0, 0, 289, 277, 1, 0, 0, 0, 289, 278,
		1, 0, 0, 0, 289, 282, 1, 0, 0, 0, 289, 284, 1, 0, 0, 0, 290, 332, 1, 0,
		0, 0, 291, 292, 10, 14, 0, 0, 292, 293, 5, 45, 0, 0, 293, 331, 3, 54, 27,
		15, 294, 295, 10, 11, 0, 0, 295, 296, 7, 4, 0, 0, 296, 331, 3, 54, 27,
		12, 297, 298, 10, 10, 0, 0, 298, 299, 7, 5, 0, 0, 299, 331, 3, 54, 27,
		11, 300, 301, 10, 9, 0, 0, 301, 302, 7, 6, 0, 0, 302, 331, 3, 54, 27, 10,
		303, 304, 10, 8, 0, 0, 304, 305, 7, 7, 0, 0, 305, 331, 3, 54, 27, 9, 306,
		307, 10, 7, 0, 0, 307, 308, 7, 8, 0, 0, 308, 331, 3, 54, 27, 8, 309, 310,
		10, 6, 0, 0, 310, 311, 5, 48, 0, 0, 311, 331, 3, 54, 27, 7, 312, 313, 10,
		5, 0, 0, 313, 314, 5, 50, 0, 0, 314, 331, 3, 54, 27, 6, 315, 316, 10, 4,
		0, 0, 316, 317, 5, 49, 0, 0, 317, 331, 3, 54, 27, 5, 318, 319, 10, 3, 0,
		0, 319, 320, 5, 52, 0, 0, 320, 331, 3, 54, 27, 4, 321, 322, 10, 2, 0, 0,
		322, 323, 5, 53, 0, 0, 323, 331, 3, 54, 27, 3, 324, 325, 10, 1, 0, 0, 325,
		326, 5, 24, 0, 0, 326, 327, 3, 54, 27, 0, 327, 328, 5, 39, 0, 0, 328, 329,
		3, 54, 27, 2, 329, 331, 1, 0, 0, 0, 330, 291, 1, 0, 0, 0, 330, 294, 1,
		0, 0, 0, 330, 297, 1, 0, 0, 0, 330, 300, 1, 0, 0, 0, 330, 303, 1, 0, 0,
		0, 330, 306, 1, 0, 0, 0, 330, 309, 1, 0, 0, 0, 330, 312, 1, 0, 0, 0, 330,
		315, 1, 0, 0, 0, 330, 318, 1, 0, 0, 0, 330, 321, 1, 0, 0, 0, 330, 324,
		1, 0, 0, 0, 331, 334, 1, 0, 0, 0, 332, 330, 1, 0, 0, 0, 332, 333, 1, 0,
		0, 0, 333, 55, 1, 0, 0, 0, 334, 332, 1, 0, 0, 0, 335, 336, 5, 4, 0, 0,
		336, 57, 1, 0, 0, 0, 337, 349, 1, 0, 0, 0, 338, 340, 7, 5, 0, 0, 339, 338,
		1, 0, 0, 0, 339, 340, 1, 0, 0, 0, 340, 341, 1, 0, 0, 0, 341, 349, 3, 78,
		39, 0, 342, 344, 7, 5, 0, 0, 343, 342, 1, 0, 0, 0, 343, 344, 1, 0, 0, 0,
		344, 345, 1, 0, 0, 0, 345, 349, 3, 84, 42, 0, 346, 349, 3, 80, 40, 0, 347,
		349, 3, 82, 41, 0, 348, 337, 1, 0, 0, 0, 348, 339, 1, 0, 0, 0, 348, 343,
		1, 0, 0, 0, 348, 346, 1, 0, 0, 0, 348, 347, 1, 0, 0, 0, 349, 59, 1, 0,
		0, 0, 350, 351, 5, 22, 0, 0, 351, 61, 1, 0, 0, 0, 352, 353, 5, 59, 0, 0,
		353, 63, 1, 0, 0, 0, 354, 355, 5, 20, 0, 0, 355, 356, 3, 62, 31, 0, 356,
		65, 1, 0, 0, 0, 357, 358, 5, 19, 0, 0, 358, 359, 3, 62, 31, 0, 359, 67,
		1, 0, 0, 0, 360, 361, 3, 62, 31, 0, 361, 69, 1, 0, 0, 0, 362, 363, 3, 62,
		31, 0, 363, 71, 1, 0, 0, 0, 364, 367, 3, 64, 32, 0, 365, 367, 3, 48, 24,
		0, 366, 364, 1, 0, 0, 0, 366, 365, 1, 0, 0, 0, 367, 73, 1, 0, 0, 0, 368,
		369, 3, 66, 33, 0, 369, 75, 1, 0, 0, 0, 370, 371, 3, 34, 17, 0, 371, 372,
		5, 31, 0, 0, 372, 373, 3, 78, 39, 0, 373, 374, 5, 33, 0, 0, 374, 77, 1,
		0, 0, 0, 375, 376, 5, 58, 0, 0, 376, 79, 1, 0, 0, 0, 377, 378, 5, 55, 0,
		0, 378, 81, 1, 0, 0, 0, 379, 380, 5, 56, 0, 0, 380, 83, 1, 0, 0, 0, 381,
		382, 5, 57, 0, 0, 382, 85, 1, 0, 0, 0, 33, 89, 91, 102, 108, 112, 126,
		131, 137, 140, 145, 151, 158, 161, 164, 174, 187, 194, 210, 222, 229, 234,
		237, 247, 258, 266, 273, 289, 330, 332, 339, 343, 348, 366,
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
	bubblerParserIMPORT       = 1
	bubblerParserGET          = 2
	bubblerParserSET          = 3
	bubblerParserVALUE        = 4
	bubblerParserVOID         = 5
	bubblerParserINT8         = 6
	bubblerParserINT16        = 7
	bubblerParserINT32        = 8
	bubblerParserINT64        = 9
	bubblerParserUINT8        = 10
	bubblerParserUINT16       = 11
	bubblerParserUINT32       = 12
	bubblerParserUINT64       = 13
	bubblerParserFLOAT32      = 14
	bubblerParserFLOAT64      = 15
	bubblerParserBOOL         = 16
	bubblerParserSTRING       = 17
	bubblerParserBYTES        = 18
	bubblerParserENUM         = 19
	bubblerParserSTRUCT       = 20
	bubblerParserHASH         = 21
	bubblerParserSEMI         = 22
	bubblerParserASSIGN       = 23
	bubblerParserQUESTION     = 24
	bubblerParserLP           = 25
	bubblerParserRP           = 26
	bubblerParserLB           = 27
	bubblerParserRB           = 28
	bubblerParserLC           = 29
	bubblerParserRC           = 30
	bubblerParserLT           = 31
	bubblerParserLE           = 32
	bubblerParserGT           = 33
	bubblerParserGE           = 34
	bubblerParserEQ           = 35
	bubblerParserNE           = 36
	bubblerParserDOT          = 37
	bubblerParserCOMMA        = 38
	bubblerParserCOLON        = 39
	bubblerParserADD          = 40
	bubblerParserSUB          = 41
	bubblerParserMUL          = 42
	bubblerParserDIV          = 43
	bubblerParserMOD          = 44
	bubblerParserPOW          = 45
	bubblerParserSHL          = 46
	bubblerParserSHR          = 47
	bubblerParserBAND         = 48
	bubblerParserBOR          = 49
	bubblerParserBXOR         = 50
	bubblerParserBNOT         = 51
	bubblerParserAND          = 52
	bubblerParserOR           = 53
	bubblerParserNOT          = 54
	bubblerParserSTR_LIT      = 55
	bubblerParserBOOL_LIT     = 56
	bubblerParserFLOAT_LIT    = 57
	bubblerParserINT_LIT      = 58
	bubblerParserIDENTIFIER   = 59
	bubblerParserWS           = 60
	bubblerParserLINE_COMMENT = 61
	bubblerParserCOMMENT      = 62
	bubblerParserKEYWORDS     = 63
)

// bubblerParser rules.
const (
	bubblerParserRULE_proto            = 0
	bubblerParserRULE_importStatement  = 1
	bubblerParserRULE_topLevelDef      = 2
	bubblerParserRULE_size_            = 3
	bubblerParserRULE_byteSize         = 4
	bubblerParserRULE_bitSize          = 5
	bubblerParserRULE_optionName       = 6
	bubblerParserRULE_field            = 7
	bubblerParserRULE_fieldVoid        = 8
	bubblerParserRULE_fieldConstant    = 9
	bubblerParserRULE_fieldEmbedded    = 10
	bubblerParserRULE_fieldNormal      = 11
	bubblerParserRULE_fieldOptions     = 12
	bubblerParserRULE_fieldOption      = 13
	bubblerParserRULE_fieldMethods     = 14
	bubblerParserRULE_fieldMethod      = 15
	bubblerParserRULE_type_            = 16
	bubblerParserRULE_basicType        = 17
	bubblerParserRULE_enumDef          = 18
	bubblerParserRULE_enumBody         = 19
	bubblerParserRULE_enumElement      = 20
	bubblerParserRULE_enumField        = 21
	bubblerParserRULE_enumValueOptions = 22
	bubblerParserRULE_enumValueOption  = 23
	bubblerParserRULE_structDef        = 24
	bubblerParserRULE_structBody       = 25
	bubblerParserRULE_structElement    = 26
	bubblerParserRULE_expr             = 27
	bubblerParserRULE_value            = 28
	bubblerParserRULE_constant         = 29
	bubblerParserRULE_emptyStatement_  = 30
	bubblerParserRULE_ident            = 31
	bubblerParserRULE_structName       = 32
	bubblerParserRULE_enumName         = 33
	bubblerParserRULE_fieldName        = 34
	bubblerParserRULE_methodName       = 35
	bubblerParserRULE_structType       = 36
	bubblerParserRULE_enumType         = 37
	bubblerParserRULE_arrayType        = 38
	bubblerParserRULE_intLit           = 39
	bubblerParserRULE_strLit           = 40
	bubblerParserRULE_boolLit          = 41
	bubblerParserRULE_floatLit         = 42
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
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&5767170) != 0 {
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case bubblerParserIMPORT:
			{
				p.SetState(86)
				p.ImportStatement()
			}

		case bubblerParserENUM, bubblerParserSTRUCT:
			{
				p.SetState(87)
				p.TopLevelDef()
			}

		case bubblerParserSEMI:
			{
				p.SetState(88)
				p.EmptyStatement_()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(94)
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
		p.SetState(96)
		p.Match(bubblerParserIMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(97)
		p.StrLit()
	}
	{
		p.SetState(98)
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
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case bubblerParserENUM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(100)
			p.EnumDef()
		}

	case bubblerParserSTRUCT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(101)
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
		p.SetState(104)
		p.Match(bubblerParserLB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case bubblerParserINT_LIT:
		{
			p.SetState(105)
			p.ByteSize()
		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == bubblerParserHASH {
			{
				p.SetState(106)
				p.Match(bubblerParserHASH)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(107)
				p.BitSize()
			}

		}

	case bubblerParserHASH:
		{
			p.SetState(110)
			p.Match(bubblerParserHASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.BitSize()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(114)
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
		p.SetState(116)
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
		p.SetState(118)
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
	p.EnterRule(localctx, 12, bubblerParserRULE_optionName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
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
	p.EnterRule(localctx, 14, bubblerParserRULE_field)
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.FieldVoid()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(123)
			p.FieldConstant()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(124)
			p.FieldEmbedded()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(125)
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
	p.EnterRule(localctx, 16, bubblerParserRULE_fieldVoid)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(bubblerParserVOID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.Size_()
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == bubblerParserLB {
		{
			p.SetState(130)
			p.FieldOptions()
		}

	}
	{
		p.SetState(133)
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
	Constant() IConstantContext
	SEMI() antlr.TerminalNode
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

func (s *FieldConstantContext) SEMI() antlr.TerminalNode {
	return s.GetToken(bubblerParserSEMI, 0)
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
	p.EnterRule(localctx, 18, bubblerParserRULE_fieldConstant)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.BasicType()
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == bubblerParserIDENTIFIER {
		{
			p.SetState(136)
			p.FieldName()
		}

	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == bubblerParserLB {
		{
			p.SetState(139)
			p.Size_()
		}

	}
	{
		p.SetState(142)
		p.Match(bubblerParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
		p.Constant()
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == bubblerParserLB {
		{
			p.SetState(144)
			p.FieldOptions()
		}

	}
	{
		p.SetState(147)
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
	p.EnterRule(localctx, 20, bubblerParserRULE_fieldEmbedded)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Type_()
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == bubblerParserLB {
		{
			p.SetState(150)
			p.FieldOptions()
		}

	}
	{
		p.SetState(153)
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
	p.EnterRule(localctx, 22, bubblerParserRULE_fieldNormal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Type_()
	}
	{
		p.SetState(156)
		p.FieldName()
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(157)
			p.Size_()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == bubblerParserLB {
		{
			p.SetState(160)
			p.FieldOptions()
		}

	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == bubblerParserLC {
		{
			p.SetState(163)
			p.FieldMethods()
		}

	}
	{
		p.SetState(166)
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
	p.EnterRule(localctx, 24, bubblerParserRULE_fieldOptions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(bubblerParserLB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.FieldOption()
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == bubblerParserCOMMA {
		{
			p.SetState(170)
			p.Match(bubblerParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(171)
			p.FieldOption()
		}

		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(177)
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
	p.EnterRule(localctx, 26, bubblerParserRULE_fieldOption)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.OptionName()
	}
	{
		p.SetState(180)
		p.Match(bubblerParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(181)
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
	p.EnterRule(localctx, 28, bubblerParserRULE_fieldMethods)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Match(bubblerParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == bubblerParserGET || _la == bubblerParserSET {
		{
			p.SetState(184)
			p.FieldMethod()
		}

		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(190)
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
	p.EnterRule(localctx, 30, bubblerParserRULE_fieldMethod)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)

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
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == bubblerParserIDENTIFIER {
		{
			p.SetState(193)
			p.MethodName()
		}

	}
	{
		p.SetState(196)
		p.Match(bubblerParserLP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)
		p.BasicType()
	}
	{
		p.SetState(198)
		p.Match(bubblerParserRP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)
		p.Match(bubblerParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(200)
		p.expr(0)
	}
	{
		p.SetState(201)
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
	p.EnterRule(localctx, 32, bubblerParserRULE_type_)
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(203)
			p.BasicType()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(204)
			p.Match(bubblerParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(205)
			p.Match(bubblerParserBYTES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(206)
			p.ArrayType()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(207)
			p.StructType()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(208)
			p.EnumType()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(209)
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
	p.EnterRule(localctx, 34, bubblerParserRULE_basicType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&131008) != 0) {
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
	p.EnterRule(localctx, 36, bubblerParserRULE_enumDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.EnumName()
	}
	{
		p.SetState(215)
		p.Size_()
	}
	{
		p.SetState(216)
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
	p.EnterRule(localctx, 38, bubblerParserRULE_enumBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.Match(bubblerParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == bubblerParserSEMI || _la == bubblerParserIDENTIFIER {
		{
			p.SetState(219)
			p.EnumElement()
		}

		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(225)
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
	EnumField() IEnumFieldContext
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

func (s *EnumElementContext) EnumField() IEnumFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumFieldContext)
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
	p.EnterRule(localctx, 40, bubblerParserRULE_enumElement)
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case bubblerParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(227)
			p.EnumField()
		}

	case bubblerParserSEMI:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(228)
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

// IEnumFieldContext is an interface to support dynamic dispatch.
type IEnumFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext
	SEMI() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	IntLit() IIntLitContext
	EnumValueOptions() IEnumValueOptionsContext

	// IsEnumFieldContext differentiates from other interfaces.
	IsEnumFieldContext()
}

type EnumFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumFieldContext() *EnumFieldContext {
	var p = new(EnumFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_enumField
	return p
}

func InitEmptyEnumFieldContext(p *EnumFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bubblerParserRULE_enumField
}

func (*EnumFieldContext) IsEnumFieldContext() {}

func NewEnumFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumFieldContext {
	var p = new(EnumFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bubblerParserRULE_enumField

	return p
}

func (s *EnumFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumFieldContext) Ident() IIdentContext {
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

func (s *EnumFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(bubblerParserSEMI, 0)
}

func (s *EnumFieldContext) COMMA() antlr.TerminalNode {
	return s.GetToken(bubblerParserCOMMA, 0)
}

func (s *EnumFieldContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(bubblerParserASSIGN, 0)
}

func (s *EnumFieldContext) IntLit() IIntLitContext {
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

func (s *EnumFieldContext) EnumValueOptions() IEnumValueOptionsContext {
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

func (s *EnumFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case bubblerVisitor:
		return t.VisitEnumField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *bubblerParser) EnumField() (localctx IEnumFieldContext) {
	localctx = NewEnumFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, bubblerParserRULE_enumField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
		p.Ident()
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == bubblerParserASSIGN {
		{
			p.SetState(232)
			p.Match(bubblerParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(233)
			p.IntLit()
		}

	}
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == bubblerParserLB {
		{
			p.SetState(236)
			p.EnumValueOptions()
		}

	}
	{
		p.SetState(239)
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
	p.EnterRule(localctx, 44, bubblerParserRULE_enumValueOptions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.Match(bubblerParserLB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(242)
		p.EnumValueOption()
	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == bubblerParserCOMMA {
		{
			p.SetState(243)
			p.Match(bubblerParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(244)
			p.EnumValueOption()
		}

		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(250)
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
	p.EnterRule(localctx, 46, bubblerParserRULE_enumValueOption)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		p.OptionName()
	}
	{
		p.SetState(253)
		p.Match(bubblerParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(254)
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
	p.EnterRule(localctx, 48, bubblerParserRULE_structDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		p.StructName()
	}
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == bubblerParserLB {
		{
			p.SetState(257)
			p.Size_()
		}

	}
	{
		p.SetState(260)
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
	p.EnterRule(localctx, 50, bubblerParserRULE_structBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(262)
		p.Match(bubblerParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&576460752309714912) != 0 {
		{
			p.SetState(263)
			p.StructElement()
		}

		p.SetState(268)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(269)
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
	p.EnterRule(localctx, 52, bubblerParserRULE_structElement)
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case bubblerParserVOID, bubblerParserINT8, bubblerParserINT16, bubblerParserINT32, bubblerParserINT64, bubblerParserUINT8, bubblerParserUINT16, bubblerParserUINT32, bubblerParserUINT64, bubblerParserFLOAT32, bubblerParserFLOAT64, bubblerParserBOOL, bubblerParserSTRING, bubblerParserBYTES, bubblerParserENUM, bubblerParserSTRUCT, bubblerParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(271)
			p.Field()
		}

	case bubblerParserSEMI:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(272)
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
	_startState := 54
	p.EnterRecursionRule(localctx, 54, bubblerParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExprValueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(276)
			p.Value()
		}

	case 2:
		localctx = NewExprConstantContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(277)
			p.Constant()
		}

	case 3:
		localctx = NewExprParensContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(278)
			p.Match(bubblerParserLP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(279)
			p.expr(0)
		}
		{
			p.SetState(280)
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
			p.SetState(282)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExprUnaryContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&20269496858050560) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExprUnaryContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(283)
			p.expr(13)
		}

	case 5:
		localctx = NewExprCastContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(284)
			p.Match(bubblerParserLP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(285)
			p.BasicType()
		}
		{
			p.SetState(286)
			p.Match(bubblerParserRP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(287)
			p.expr(12)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(330)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprPowerContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, bubblerParserRULE_expr)
				p.SetState(291)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(292)
					p.Match(bubblerParserPOW)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(293)
					p.expr(15)
				}

			case 2:
				localctx = NewExprMulDivModContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, bubblerParserRULE_expr)
				p.SetState(294)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(295)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprMulDivModContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30786325577728) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprMulDivModContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(296)
					p.expr(12)
				}

			case 3:
				localctx = NewExprAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, bubblerParserRULE_expr)
				p.SetState(297)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(298)

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
					p.SetState(299)
					p.expr(11)
				}

			case 4:
				localctx = NewExprShiftContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, bubblerParserRULE_expr)
				p.SetState(300)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(301)

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
					p.SetState(302)
					p.expr(10)
				}

			case 5:
				localctx = NewExprRelationalContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, bubblerParserRULE_expr)
				p.SetState(303)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(304)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprRelationalContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32212254720) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprRelationalContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(305)
					p.expr(9)
				}

			case 6:
				localctx = NewExprEqualityContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, bubblerParserRULE_expr)
				p.SetState(306)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(307)

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
					p.SetState(308)
					p.expr(8)
				}

			case 7:
				localctx = NewExprBitAndContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, bubblerParserRULE_expr)
				p.SetState(309)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(310)
					p.Match(bubblerParserBAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(311)
					p.expr(7)
				}

			case 8:
				localctx = NewExprBitXorContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, bubblerParserRULE_expr)
				p.SetState(312)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(313)
					p.Match(bubblerParserBXOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(314)
					p.expr(6)
				}

			case 9:
				localctx = NewExprBitOrContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, bubblerParserRULE_expr)
				p.SetState(315)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(316)
					p.Match(bubblerParserBOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(317)
					p.expr(5)
				}

			case 10:
				localctx = NewExprLogicalAndContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, bubblerParserRULE_expr)
				p.SetState(318)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(319)
					p.Match(bubblerParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(320)
					p.expr(4)
				}

			case 11:
				localctx = NewExprLogicalOrContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, bubblerParserRULE_expr)
				p.SetState(321)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(322)
					p.Match(bubblerParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(323)
					p.expr(3)
				}

			case 12:
				localctx = NewExprTernaryContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, bubblerParserRULE_expr)
				p.SetState(324)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(325)
					p.Match(bubblerParserQUESTION)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(326)
					p.expr(0)
				}
				{
					p.SetState(327)
					p.Match(bubblerParserCOLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(328)
					p.expr(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(334)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 56, bubblerParserRULE_value)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(335)
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
	p.EnterRule(localctx, 58, bubblerParserRULE_constant)
	var _la int

	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(339)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == bubblerParserADD || _la == bubblerParserSUB {
			{
				p.SetState(338)
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
			p.SetState(341)
			p.IntLit()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(343)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == bubblerParserADD || _la == bubblerParserSUB {
			{
				p.SetState(342)
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
			p.SetState(345)
			p.FloatLit()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(346)
			p.StrLit()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(347)
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
	p.EnterRule(localctx, 60, bubblerParserRULE_emptyStatement_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(350)
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
	p.EnterRule(localctx, 62, bubblerParserRULE_ident)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(352)
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
	p.EnterRule(localctx, 64, bubblerParserRULE_structName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(354)
		p.Match(bubblerParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(355)
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
	p.EnterRule(localctx, 66, bubblerParserRULE_enumName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(357)
		p.Match(bubblerParserENUM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(358)
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
	p.EnterRule(localctx, 68, bubblerParserRULE_fieldName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(360)
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
	p.EnterRule(localctx, 70, bubblerParserRULE_methodName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(362)
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
	p.EnterRule(localctx, 72, bubblerParserRULE_structType)
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(364)
			p.StructName()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(365)
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
	p.EnterRule(localctx, 74, bubblerParserRULE_enumType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(368)
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

// IArrayTypeContext is an interface to support dynamic dispatch.
type IArrayTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BasicType() IBasicTypeContext
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

func (s *ArrayTypeContext) BasicType() IBasicTypeContext {
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
	p.EnterRule(localctx, 76, bubblerParserRULE_arrayType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(370)
		p.BasicType()
	}
	{
		p.SetState(371)
		p.Match(bubblerParserLT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(372)
		p.IntLit()
	}
	{
		p.SetState(373)
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
	p.EnterRule(localctx, 78, bubblerParserRULE_intLit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(375)
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
	p.EnterRule(localctx, 80, bubblerParserRULE_strLit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(377)
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
	p.EnterRule(localctx, 82, bubblerParserRULE_boolLit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(379)
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
	p.EnterRule(localctx, 84, bubblerParserRULE_floatLit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(381)
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
	case 27:
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
