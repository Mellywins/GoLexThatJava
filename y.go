// Code generated by goyacc syn.y. DO NOT EDIT.

//line syn.y:3
package main

import __yyfmt__ "fmt"

//line syn.y:3
import (
	"github.com/timtadh/lexmachine"
)

//line syn.y:9
type yySymType struct {
	yys   int
	token *lexmachine.Token
	ast   *Node
}

const BANG = 57346
const INTEGER_LITERAL = 57347
const BOOLEAN_LITERAL = 57348
const IDENTIFIER = 57349
const CLASS = 57350
const PUBLIC = 57351
const STATIC = 57352
const VOID = 57353
const MAIN = 57354
const STRING = 57355
const SYSTEMOUTPRINTLN = 57356
const RETURN = 57357
const INT = 57358
const IF = 57359
const FOR = 57360
const ELSE = 57361
const WHILE = 57362
const THIS = 57363
const NEW = 57364
const BOOLEAN = 57365
const LENGTH = 57366
const EXTENDS = 57367
const LEFTBRACKET = 57368
const RIGHTBRACKET = 57369
const LEFTANGLEBRACKET = 57370
const RIGHTANGLEBRACKET = 57371
const COMMA = 57372
const SEMICOLON = 57373
const COLON = 57374
const LEFTPARENTHESIS = 57375
const RIGHTPARENTHESIS = 57376
const PERIOD = 57377
const PLUS = 57378
const ASTERIX = 57379
const DIVISION = 57380
const MODULO = 57381
const DOUBLEQUAL = 57382
const DIFFERENT = 57383
const LESS = 57384
const LESSOREQUALS = 57385
const GREATER = 57386
const GREATEROREQUALS = 57387
const LOGICALAND = 57388
const LOGICALOR = 57389
const EQUAL = 57390
const MINUS = 57391

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"BANG",
	"INTEGER_LITERAL",
	"BOOLEAN_LITERAL",
	"IDENTIFIER",
	"CLASS",
	"PUBLIC",
	"STATIC",
	"VOID",
	"MAIN",
	"STRING",
	"SYSTEMOUTPRINTLN",
	"RETURN",
	"INT",
	"IF",
	"FOR",
	"ELSE",
	"WHILE",
	"THIS",
	"NEW",
	"BOOLEAN",
	"LENGTH",
	"EXTENDS",
	"LEFTBRACKET",
	"RIGHTBRACKET",
	"LEFTANGLEBRACKET",
	"RIGHTANGLEBRACKET",
	"COMMA",
	"SEMICOLON",
	"COLON",
	"LEFTPARENTHESIS",
	"RIGHTPARENTHESIS",
	"PERIOD",
	"PLUS",
	"ASTERIX",
	"DIVISION",
	"MODULO",
	"DOUBLEQUAL",
	"DIFFERENT",
	"LESS",
	"LESSOREQUALS",
	"GREATER",
	"GREATEROREQUALS",
	"LOGICALAND",
	"LOGICALOR",
	"EQUAL",
	"MINUS",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line syn.y:80

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 14

var yyAct = [...]int8{
	6, 3, 4, 5, 2, 1, 0, 8, 0, 0,
	0, 0, 0, 7,
}

var yyPact = [...]int16{
	-3, -1000, -1000, -36, -1000, -3, -1000, -1000, -1000,
}

var yyPgo = [...]int8{
	0, 5, 4, 1, 3,
}

var yyR1 = [...]int8{
	0, 1, 2, 3, 4, 4,
}

var yyR2 = [...]int8{
	0, 1, 3, 1, 1, 1,
}

var yyChk = [...]int16{
	-1000, -1, -2, -3, 5, -4, 36, 49, -3,
}

var yyDef = [...]int8{
	0, -2, 1, 0, 3, 0, 4, 5, 2,
}

var yyTok1 = [...]int8{
	1,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
//line syn.y:63
		{
			yylex.(*golex).stmts = append(yylex.(*golex).stmts, yyDollar[1].ast)
		}
	case 2:
		yyDollar = yyS[yypt-3 : yypt+1]
//line syn.y:67
		{

			yyVAL.ast = NewNode("Statement", nil).AddKid(yyDollar[2].ast).AddKid(yyDollar[1].ast).AddKid(yyDollar[3].ast)

			__yyfmt__.Printf("%s + %s", yyDollar[1].token.Lexeme, yyDollar[3].token.Lexeme)

		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line syn.y:75
		{
			yyVAL.ast = NewNode("INTEGER_LITERAL", yyDollar[1].token)
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line syn.y:77
		{
			yyVAL.ast = NewNode("PLUS", yyDollar[1].token)
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line syn.y:78
		{
			yyVAL.ast = NewNode("MINUS", yyDollar[1].token)
		}
	}
	goto yystack /* stack new state and value */
}
