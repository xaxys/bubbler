grammar bubbler;

proto:
    (
        importStatement
        | topLevelDef
        | emptyStatement_
    )* EOF;

// Import Statement

importStatement: IMPORT strLit SEMI;

// Top Level definitions

topLevelDef: enumDef | structDef;

// Size

size_:
	LB (
		( byteSize ( HASH bitSize )? )
		| ( HASH bitSize )
	) RB;

byteSize: intLit;

bitSize: intLit;

// Option

// optionName:
//     fullIdent
//     | LP fullIdent RP ( DOT fullIdent )?;
optionName: ident;

// Field

field:
    fieldVoid
    | fieldConstant
    | fieldEmbedded
    | fieldNormal
    ;

fieldVoid:
    VOID size_ fieldOptions? SEMI;

fieldConstant:
    basicType fieldName? size_? ASSIGN constant fieldOptions? SEMI;

fieldEmbedded:
    type_ fieldOptions? SEMI;

fieldNormal:
    type_ fieldName size_? fieldOptions? fieldMethods? SEMI;

fieldOptions: LB fieldOption ( COMMA fieldOption )* RB;

fieldOption: optionName ASSIGN constant;

fieldMethods : LC fieldMethod* RC;

fieldMethod: op = ( GET | SET ) methodName? LP basicType RP COLON expr SEMI;

// field types

type_:
    basicType
    | STRING
    | BYTES
    | arrayType
    | structType // old fashion: 'struct' ident
    | enumType   // old fashion: 'enum' ident
    | ident      // unknown type ident
    ;

basicType:
    BOOL
    | INT8
    | INT16
    | INT32
    | INT64
    | UINT8
    | UINT16
    | UINT32
    | UINT64
    | FLOAT32
    | FLOAT64
    ;

// enum

enumDef: enumName size_ enumBody;

enumBody: LC enumElement* RC;

enumElement: enumField | emptyStatement_;

enumField: ident ( ASSIGN intLit )? enumValueOptions? ( SEMI | COMMA );

enumValueOptions:
    LB enumValueOption ( COMMA enumValueOption )* RB;

enumValueOption: optionName ASSIGN constant;

// struct

structDef: structName size_? structBody;

structBody: LC structElement* RC;

structElement: field | emptyStatement_;

// expr

expr:
    value                                       # ExprValue
    | constant                                  # ExprConstant
    | LP expr RP                                # ExprParens
    | expr POW expr                             # ExprPower
    | op = ( ADD | SUB | BNOT | NOT ) expr      # ExprUnary
    | LP basicType RP expr                      # ExprCast
    | expr op = ( MUL | DIV | MOD ) expr        # ExprMulDivMod
    | expr op = ( ADD | SUB ) expr              # ExprAddSub
    | expr op = ( SHL | SHR ) expr              # ExprShift
    | expr op = ( LT | LE | GT | GE ) expr      # ExprRelational
    | expr op = ( EQ | NE ) expr                # ExprEquality
    | expr BAND expr                            # ExprBitAnd
    | expr BXOR expr                            # ExprBitXor
    | expr BOR expr                             # ExprBitOr
    | expr AND expr                             # ExprLogicalAnd
    | expr OR expr                              # ExprLogicalOr
	| expr QUESTION expr COLON expr             # ExprTernary
    ;

// lexical

value: VALUE;

constant:
//    fullIdent
    | ( SUB | ADD )? intLit
    | ( SUB | ADD )? floatLit
    | strLit
    | boolLit
//    | blockLit
    ;

// not specified in specification but used in tests
// blockLit: LC ( ident COLON constant )* RC;

emptyStatement_: SEMI;

// Lexical elements

ident: IDENTIFIER;
// fullIdent: ident ( DOT ident )*;
structName: STRUCT ident;
enumName: ENUM ident;
fieldName: ident;
methodName: ident;
// structType: ( DOT )? ( ident DOT )* structName;
// enumType: ( DOT )? ( ident DOT )* enumName;
structType: structName | structDef;
enumType: enumName;
arrayType: basicType LT intLit GT;

intLit: INT_LIT;
strLit: STR_LIT;
boolLit: BOOL_LIT;
floatLit: FLOAT_LIT;

// keywords
// SYNTAX: 'syntax';
IMPORT: 'import';
GET: 'get';
SET: 'set';
VALUE: 'value';
// WEAK: 'weak';
// PUBLIC: 'public';
// PACKAGE: 'package';
// OPTION: 'option';
// OPTIONAL: 'optional';
// REPEATED: 'repeated';
// ONEOF: 'oneof';
// MAP: 'map';
VOID: 'void';
INT8: 'int8';
INT16: 'int16';
INT32: 'int32';
INT64: 'int64';
UINT8: 'uint8';
UINT16: 'uint16';
UINT32: 'uint32';
UINT64: 'uint64';
FLOAT32: 'float32';
FLOAT64: 'float64';
BOOL: 'bool';
STRING: 'string';
BYTES: 'bytes';
ENUM: 'enum';
STRUCT: 'struct';
// SERVICE: 'service';
// EXTEND: 'extend';
// RPC: 'rpc';
// STREAM: 'stream';
// RETURNS: 'returns';
// symbols

HASH: '#';
SEMI: ';';
ASSIGN: '=';
QUESTION: '?';
LP: '(';
RP: ')';
LB: '[';
RB: ']';
LC: '{';
RC: '}';
LT: '<';
LE: '<=';
GT: '>';
GE: '>=';
EQ: '==';
NE: '!=';
DOT: '.';
COMMA: ',';
COLON: ':';
ADD: '+';
SUB: '-';
MUL: '*';
DIV: '/';
MOD: '%';
POW: '**';
SHL: '<<';
SHR: '>>';
BAND: '&';
BOR: '|';
BXOR: '^';
BNOT: '~';
AND: '&&';
OR: '||';
NOT: '!';

STR_LIT: ( '\'' ( CHAR_VALUE )*? '\'' ) | ( '"' ( CHAR_VALUE )*? '"' );
fragment CHAR_VALUE:
    HEX_ESCAPE
    | OCT_ESCAPE
    | CHAR_ESCAPE
    | ~[\u0000\n\\];
fragment HEX_ESCAPE: '\\' ( 'x' | 'X' ) HEX_DIGIT HEX_DIGIT;
fragment OCT_ESCAPE: '\\' OCTAL_DIGIT OCTAL_DIGIT OCTAL_DIGIT;
fragment CHAR_ESCAPE:
    '\\' (
        'a'
        | 'b'
        | 'f'
        | 'n'
        | 'r'
        | 't'
        | 'v'
        | '\\'
        | '\''
        | '"'
    );

BOOL_LIT: 'true' | 'false';

FLOAT_LIT: (
        DECIMALS DOT DECIMALS? EXPONENT?
        | DECIMALS EXPONENT
        | DOT DECIMALS EXPONENT?
    )
    | 'inf'
    | 'nan';
fragment EXPONENT: ( 'e' | 'E' ) ( ADD | SUB )? DECIMALS;
fragment DECIMALS: DECIMAL_DIGIT+;

INT_LIT: DECIMAL_LIT | OCTAL_LIT | HEX_LIT;
fragment DECIMAL_LIT: ( [1-9] ) DECIMAL_DIGIT*;
fragment OCTAL_LIT: '0' OCTAL_DIGIT*;
fragment HEX_LIT: '0' ( 'x' | 'X' ) HEX_DIGIT+;

IDENTIFIER: LETTER ( LETTER | DECIMAL_DIGIT )*;

fragment LETTER: [A-Za-z_];
fragment DECIMAL_DIGIT: [0-9];
fragment OCTAL_DIGIT: [0-7];
fragment HEX_DIGIT: [0-9A-Fa-f];

// comments
WS: [ \t\r\n\u000C]+ -> skip;
LINE_COMMENT: '//' ~[\r\n]* -> channel(HIDDEN);
COMMENT: '/*' .*? '*/' -> channel(HIDDEN);

KEYWORDS:
    IMPORT
    | GET
    | SET
    | VALUE
    | VOID
    | INT8
    | INT16
    | INT32
    | INT64
    | UINT8
    | UINT16
    | UINT32
    | UINT64
    | FLOAT32
    | FLOAT64
    | BOOL
    | STRING
    | BYTES
    | ENUM
    | STRUCT
    | BOOL_LIT
    ;