//line gram.y:7
package main

import __yyfmt__ "fmt"

//line gram.y:7
import (
	"github.com/lufia/qsh/ast"
	"github.com/lufia/qsh/build"
)

//line gram.y:14
type yySymType struct {
	yys  int
	tree *ast.Node
}

const IF = 57346
const FOR = 57347
const IN = 57348
const WORD = 57349
const REDIR = 57350
const LOAD = 57351
const ANDAND = 57352
const OROR = 57353

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IF",
	"FOR",
	"IN",
	"WORD",
	"REDIR",
	"LOAD",
	"ANDAND",
	"OROR",
	"'|'",
	"'\\n'",
	"';'",
	"'&'",
	"'{'",
	"'}'",
	"'='",
	"'$'",
	"'('",
	"')'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 0,
	1, 1,
	-2, 13,
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 79

var yyAct = [...]int{

	22, 25, 11, 32, 37, 45, 11, 3, 4, 30,
	46, 23, 24, 26, 27, 15, 20, 14, 31, 41,
	11, 11, 11, 36, 1, 11, 33, 34, 35, 12,
	13, 38, 40, 10, 2, 42, 43, 5, 6, 21,
	14, 11, 7, 14, 48, 49, 8, 38, 40, 14,
	50, 28, 12, 13, 9, 12, 13, 44, 23, 14,
	29, 12, 13, 18, 19, 20, 47, 16, 17, 39,
	0, 12, 13, 18, 19, 20, 0, 16, 17,
}
var yyPact = [...]int{

	33, -1000, 2, 63, 33, -5, 10, 10, 52, -1000,
	-9, -1000, 10, -1000, -1000, -1000, -1000, -1000, 33, 33,
	33, -1000, -5, 33, 13, -1000, -1000, -1000, -1000, 10,
	10, -1000, 36, 4, 4, -1000, -1000, -7, 53, 33,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 42,
	-1000,
}
var yyPgo = [...]int{

	0, 34, 0, 4, 8, 69, 54, 51, 7, 46,
	33, 5, 1, 3, 24,
}
var yyR1 = [...]int{

	0, 14, 14, 1, 1, 3, 3, 4, 4, 5,
	5, 2, 6, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 9, 9, 9, 7, 10, 11, 12, 12,
	12, 13, 13,
}
var yyR2 = [...]int{

	0, 0, 2, 1, 2, 1, 2, 2, 2, 1,
	2, 3, 3, 0, 3, 5, 2, 1, 3, 3,
	3, 1, 1, 2, 2, 2, 1, 1, 2, 3,
	1, 0, 2,
}
var yyChk = [...]int{

	-1000, -14, -1, -8, -4, 4, 5, 9, -9, -6,
	-10, -12, 19, 20, 7, 13, 14, 15, 10, 11,
	12, -1, -2, 16, -11, -12, -11, -11, -7, 8,
	18, -11, -13, -8, -8, -8, -2, -3, -8, -5,
	-4, 6, -11, -11, 21, -11, 17, 13, -3, -13,
	-2,
}
var yyDef = [...]int{

	-2, -2, 0, 3, 13, 0, 0, 0, 17, 21,
	22, 26, 0, 31, 30, 2, 7, 8, 13, 13,
	13, 4, 0, 13, 0, 27, 16, 23, 24, 0,
	0, 28, 0, 18, 19, 20, 14, 0, 5, 13,
	9, 31, 25, 12, 29, 32, 11, 10, 6, 0,
	15,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	13, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 19, 3, 15, 3,
	20, 21, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 14,
	3, 18, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 16, 12, 17,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
}
var yyTok3 = [...]int{
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
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
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
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
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
	yyn = yyPact[yystate]
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
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
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
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
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
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
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

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line gram.y:22
		{
			return 1
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line gram.y:26
		{
			ast.Dump(yyDollar[1].tree)
			build.Compile(yyDollar[1].tree)
			return 0
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line gram.y:35
		{
			yyVAL.tree = ast.New(ast.LIST, yyDollar[1].tree, yyDollar[2].tree)
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line gram.y:42
		{
			yyVAL.tree = ast.New(ast.LIST, yyDollar[1].tree, yyDollar[2].tree)
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line gram.y:49
		{
			yyVAL.tree = ast.Async(yyDollar[1].tree)
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line gram.y:59
		{
			yyVAL.tree = ast.Block(yyDollar[2].tree)
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line gram.y:65
		{
			yyVAL.tree = ast.Assign(yyDollar[1].tree, yyDollar[3].tree)
		}
	case 13:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line gram.y:70
		{
			yyVAL.tree = nil
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line gram.y:74
		{
			yyVAL.tree = ast.New(ast.IF, yyDollar[2].tree, yyDollar[3].tree)
		}
	case 15:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line gram.y:78
		{
			p := ast.New(ast.LIST, yyDollar[2].tree, yyDollar[4].tree)
			yyVAL.tree = ast.New(ast.FOR, p, yyDollar[5].tree)
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line gram.y:83
		{
			yyVAL.tree = ast.Load(yyDollar[2].tree)
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line gram.y:87
		{
			yyVAL.tree = ast.Simple(yyDollar[1].tree)
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line gram.y:91
		{
			yyVAL.tree = ast.New(ast.ANDAND, yyDollar[1].tree, yyDollar[3].tree)
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line gram.y:95
		{
			yyVAL.tree = ast.New(ast.OROR, yyDollar[1].tree, yyDollar[3].tree)
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line gram.y:99
		{
			yyVAL.tree = ast.New(ast.PIPE, yyDollar[1].tree, yyDollar[3].tree)
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line gram.y:107
		{
			yyVAL.tree = ast.New(ast.LIST, yyDollar[1].tree, yyDollar[2].tree)
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line gram.y:111
		{
			yyVAL.tree = ast.New(ast.LIST, yyDollar[1].tree, yyDollar[2].tree)
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line gram.y:117
		{
			yyVAL.tree = ast.Redirect(yyDollar[1].tree, yyDollar[2].tree)
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line gram.y:129
		{
			yyVAL.tree = ast.Var(yyDollar[2].tree)
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line gram.y:133
		{
			yyVAL.tree = ast.Tuple(yyDollar[2].tree)
		}
	case 31:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line gram.y:139
		{
			yyVAL.tree = nil
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line gram.y:143
		{
			if yyDollar[1].tree == nil {
				yyVAL.tree = yyDollar[2].tree
			} else {
				yyVAL.tree = ast.New(ast.LIST, yyDollar[1].tree, yyDollar[2].tree)
			}
		}
	}
	goto yystack /* stack new state and value */
}
