// Code generated by goyacc -v y.output -o parser.go -p mtail parser.y. DO NOT EDIT.

//line parser.y:5
package vm

import __yyfmt__ "fmt"

//line parser.y:5

import (
	"time"

	"github.com/golang/glog"
	"github.com/google/mtail/internal/metrics"
	"github.com/google/mtail/internal/vm/ast"
	"github.com/google/mtail/internal/vm/position"
)

//line parser.y:18
type mtailSymType struct {
	yys      int
	intVal   int64
	floatVal float64
	op       int
	text     string
	texts    []string
	flag     bool
	n        ast.Node
	kind     metrics.Kind
	duration time.Duration
}

const INVALID = 57346
const COUNTER = 57347
const GAUGE = 57348
const TIMER = 57349
const TEXT = 57350
const AFTER = 57351
const AS = 57352
const BY = 57353
const CONST = 57354
const HIDDEN = 57355
const DEF = 57356
const DEL = 57357
const NEXT = 57358
const OTHERWISE = 57359
const ELSE = 57360
const STOP = 57361
const BUILTIN = 57362
const REGEX = 57363
const STRING = 57364
const CAPREF = 57365
const CAPREF_NAMED = 57366
const ID = 57367
const DECO = 57368
const INTLITERAL = 57369
const FLOATLITERAL = 57370
const DURATIONLITERAL = 57371
const INC = 57372
const DEC = 57373
const DIV = 57374
const MOD = 57375
const MUL = 57376
const MINUS = 57377
const PLUS = 57378
const POW = 57379
const SHL = 57380
const SHR = 57381
const LT = 57382
const GT = 57383
const LE = 57384
const GE = 57385
const EQ = 57386
const NE = 57387
const BITAND = 57388
const XOR = 57389
const BITOR = 57390
const NOT = 57391
const AND = 57392
const OR = 57393
const ADD_ASSIGN = 57394
const ASSIGN = 57395
const CONCAT = 57396
const MATCH = 57397
const NOT_MATCH = 57398
const LCURLY = 57399
const RCURLY = 57400
const LPAREN = 57401
const RPAREN = 57402
const LSQUARE = 57403
const RSQUARE = 57404
const COMMA = 57405
const NL = 57406

var mtailToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"INVALID",
	"COUNTER",
	"GAUGE",
	"TIMER",
	"TEXT",
	"AFTER",
	"AS",
	"BY",
	"CONST",
	"HIDDEN",
	"DEF",
	"DEL",
	"NEXT",
	"OTHERWISE",
	"ELSE",
	"STOP",
	"BUILTIN",
	"REGEX",
	"STRING",
	"CAPREF",
	"CAPREF_NAMED",
	"ID",
	"DECO",
	"INTLITERAL",
	"FLOATLITERAL",
	"DURATIONLITERAL",
	"INC",
	"DEC",
	"DIV",
	"MOD",
	"MUL",
	"MINUS",
	"PLUS",
	"POW",
	"SHL",
	"SHR",
	"LT",
	"GT",
	"LE",
	"GE",
	"EQ",
	"NE",
	"BITAND",
	"XOR",
	"BITOR",
	"NOT",
	"AND",
	"OR",
	"ADD_ASSIGN",
	"ASSIGN",
	"CONCAT",
	"MATCH",
	"NOT_MATCH",
	"LCURLY",
	"RCURLY",
	"LPAREN",
	"RPAREN",
	"LSQUARE",
	"RSQUARE",
	"COMMA",
	"NL",
}
var mtailStatenames = [...]string{}

const mtailEofCode = 1
const mtailErrCode = 2
const mtailInitialStackSize = 16

//line parser.y:601

//  tokenpos returns the position of the current token.
func tokenpos(mtaillex mtailLexer) position.Position {
	return mtaillex.(*parser).t.Pos
}

// markedpos returns the position recorded from the most recent mark_pos
// production.
func markedpos(mtaillex mtailLexer) position.Position {
	return mtaillex.(*parser).pos
}

//line yacctab:1
var mtailExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	1, 1,
	14, 108,
	26, 108,
	32, 108,
	-2, 88,
	-1, 105,
	14, 108,
	26, 108,
	32, 108,
	-2, 88,
}

const mtailPrivate = 57344

const mtailLast = 221

var mtailAct = [...]int{

	21, 122, 63, 44, 28, 27, 43, 42, 26, 41,
	29, 47, 22, 104, 14, 103, 120, 88, 46, 25,
	19, 150, 148, 149, 149, 159, 32, 52, 35, 33,
	34, 45, 53, 37, 38, 84, 125, 13, 28, 27,
	85, 49, 92, 76, 77, 11, 24, 83, 20, 10,
	15, 87, 12, 32, 30, 35, 33, 34, 45, 157,
	37, 38, 79, 78, 32, 36, 35, 33, 34, 45,
	59, 37, 38, 111, 50, 51, 50, 51, 81, 82,
	113, 49, 40, 138, 114, 121, 121, 65, 67, 66,
	45, 115, 36, 40, 116, 117, 118, 16, 110, 119,
	156, 101, 124, 36, 129, 136, 27, 28, 27, 126,
	161, 2, 127, 160, 128, 102, 130, 109, 142, 27,
	27, 17, 137, 19, 141, 140, 147, 146, 145, 152,
	151, 143, 144, 139, 1, 32, 13, 35, 33, 34,
	45, 89, 37, 38, 11, 24, 75, 20, 10, 15,
	158, 12, 32, 96, 35, 33, 34, 45, 86, 37,
	38, 105, 95, 94, 40, 69, 70, 71, 72, 73,
	74, 98, 99, 97, 36, 123, 100, 60, 112, 90,
	91, 40, 155, 108, 39, 154, 107, 135, 134, 61,
	131, 36, 93, 48, 64, 59, 16, 80, 68, 90,
	91, 55, 56, 57, 58, 62, 18, 153, 132, 133,
	54, 9, 8, 7, 106, 6, 31, 23, 5, 4,
	3,
}
var mtailPact = [...]int{

	-1000, -1000, 33, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 65, -1000, -1000, 24, -16, -1000, -32, 196, 163,
	6, 41, -1000, -1000, -1000, 125, -1000, -12, 10, 40,
	11, -26, -19, -1000, -1000, -1000, 44, -1000, -1000, 149,
	44, 127, -1000, -1000, 139, -1000, -1000, 97, -51, -1000,
	-1000, -1000, -1000, -1000, 161, -1000, -1000, -1000, -1000, -1000,
	73, -16, 169, -1000, -51, -1000, -1000, -1000, -51, -1000,
	-1000, -1000, -1000, -1000, -1000, -51, -1000, -1000, -51, -51,
	-51, -1000, -1000, -51, 44, 115, -24, 26, 38, -1000,
	-1000, -1000, -1000, -51, -1000, -1000, -51, -1000, -1000, -1000,
	-1000, 11, -16, 44, -1000, 132, 177, -1000, -1000, 84,
	-16, -1000, 54, 44, 44, 6, 44, 44, 44, 65,
	-40, 41, -1000, -1000, -39, -1000, 44, 44, -1000, 41,
	-1000, -1000, -1000, -1000, 160, 78, 27, -1000, -1000, 125,
	40, -1000, -1000, 26, 26, 127, -1000, -1000, -1000, 44,
	-1000, 139, -1000, -38, -1000, -1000, -1000, -1000, 41, 88,
	-1000, -1000,
}
var mtailPgo = [...]int{

	0, 111, 220, 16, 11, 219, 218, 121, 2, 3,
	9, 184, 1, 217, 19, 10, 0, 14, 216, 6,
	54, 8, 215, 214, 213, 212, 7, 12, 211, 210,
	209, 208, 207, 206, 198, 197, 194, 193, 192, 153,
	146, 141, 134, 15, 17, 117,
}
var mtailR1 = [...]int{

	0, 42, 1, 1, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 5, 5, 5, 6, 6, 4,
	7, 13, 13, 13, 17, 17, 17, 17, 37, 37,
	16, 16, 36, 36, 36, 14, 14, 34, 34, 34,
	34, 34, 34, 15, 15, 35, 35, 10, 10, 27,
	27, 27, 40, 40, 21, 20, 20, 20, 38, 38,
	9, 9, 39, 39, 39, 39, 12, 12, 11, 11,
	41, 41, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 18, 18, 19, 3, 3, 26, 22, 33, 33,
	23, 23, 23, 23, 29, 29, 29, 29, 31, 32,
	32, 32, 32, 30, 24, 25, 28, 28, 44, 45,
	43, 43,
}
var mtailR2 = [...]int{

	0, 1, 0, 2, 1, 1, 1, 1, 1, 1,
	1, 3, 1, 1, 4, 2, 2, 1, 2, 3,
	1, 1, 4, 4, 1, 1, 4, 4, 1, 1,
	1, 4, 1, 1, 1, 1, 4, 1, 1, 1,
	1, 1, 1, 1, 4, 1, 1, 1, 4, 1,
	4, 4, 1, 1, 1, 1, 4, 4, 1, 1,
	1, 4, 1, 1, 1, 1, 1, 2, 1, 2,
	1, 1, 1, 3, 4, 1, 1, 1, 3, 1,
	1, 1, 4, 1, 1, 3, 5, 3, 0, 1,
	2, 2, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 3, 3, 2, 4, 3, 4, 2, 0, 0,
	0, 1,
}
var mtailChk = [...]int{

	-1000, -42, -1, -2, -5, -6, -22, -24, -25, -28,
	16, 12, 19, 4, -17, 17, 64, -7, -33, -44,
	15, -16, -27, -13, 13, -14, -21, -8, -12, -15,
	-20, -18, 20, 23, 24, 22, 59, 27, 28, -11,
	49, -10, -26, -19, -9, 25, -19, -4, -37, 57,
	50, 51, -4, 64, -29, 5, 6, 7, 8, 32,
	14, 26, -11, -8, -36, 46, 48, 47, -34, 40,
	41, 42, 43, 44, 45, -40, 55, 56, 53, 52,
	-35, 38, 39, 36, 61, 59, -7, -17, -44, -41,
	30, 31, -12, -38, 36, 35, -39, 34, 32, 33,
	37, -20, 18, -43, 64, -1, -23, 25, 22, -45,
	25, -4, 9, -43, -43, -43, -43, -43, -43, -43,
	-3, -16, -12, 60, -3, 60, -43, -43, -4, -16,
	-27, 58, -31, -30, 11, 10, 21, -4, 29, -14,
	-15, -21, -8, -17, -17, -10, -26, -19, 62, 63,
	60, -9, -12, -32, 25, 22, 22, 32, -16, 63,
	25, 22,
}
var mtailDef = [...]int{

	2, -2, -2, 3, 4, 5, 6, 7, 8, 9,
	10, 0, 12, 13, 21, 0, 17, 0, 0, 0,
	0, 24, 25, 20, 89, 30, 49, 68, 60, 35,
	54, 72, 0, 75, 76, 77, 108, 79, 80, 66,
	0, 43, 55, 81, 47, 83, 108, 15, 110, 2,
	28, 29, 16, 18, 0, 94, 95, 96, 97, 109,
	0, 0, 107, 68, 110, 32, 33, 34, 110, 37,
	38, 39, 40, 41, 42, 110, 52, 53, 110, 110,
	110, 45, 46, 110, 0, 0, 0, 21, 0, 69,
	70, 71, 67, 110, 58, 59, 110, 62, 63, 64,
	65, 11, 0, 108, 111, -2, 87, 92, 93, 0,
	0, 105, 0, 0, 0, 108, 108, 108, 0, 108,
	0, 84, 60, 73, 0, 78, 0, 0, 14, 26,
	27, 19, 90, 91, 0, 0, 0, 104, 106, 31,
	36, 50, 51, 22, 23, 44, 56, 57, 82, 0,
	74, 48, 61, 98, 99, 100, 103, 86, 85, 0,
	101, 102,
}
var mtailTok1 = [...]int{

	1,
}
var mtailTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64,
}
var mtailTok3 = [...]int{
	0,
}

var mtailErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{
	{109, 4, "unexpected end of file"},
}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	mtailDebug        = 0
	mtailErrorVerbose = false
)

type mtailLexer interface {
	Lex(lval *mtailSymType) int
	Error(s string)
}

type mtailParser interface {
	Parse(mtailLexer) int
	Lookahead() int
}

type mtailParserImpl struct {
	lval  mtailSymType
	stack [mtailInitialStackSize]mtailSymType
	char  int
}

func (p *mtailParserImpl) Lookahead() int {
	return p.char
}

func mtailNewParser() mtailParser {
	return &mtailParserImpl{}
}

const mtailFlag = -1000

func mtailTokname(c int) string {
	if c >= 1 && c-1 < len(mtailToknames) {
		if mtailToknames[c-1] != "" {
			return mtailToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func mtailStatname(s int) string {
	if s >= 0 && s < len(mtailStatenames) {
		if mtailStatenames[s] != "" {
			return mtailStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func mtailErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !mtailErrorVerbose {
		return "syntax error"
	}

	for _, e := range mtailErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + mtailTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := mtailPact[state]
	for tok := TOKSTART; tok-1 < len(mtailToknames); tok++ {
		if n := base + tok; n >= 0 && n < mtailLast && mtailChk[mtailAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if mtailDef[state] == -2 {
		i := 0
		for mtailExca[i] != -1 || mtailExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; mtailExca[i] >= 0; i += 2 {
			tok := mtailExca[i]
			if tok < TOKSTART || mtailExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if mtailExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += mtailTokname(tok)
	}
	return res
}

func mtaillex1(lex mtailLexer, lval *mtailSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = mtailTok1[0]
		goto out
	}
	if char < len(mtailTok1) {
		token = mtailTok1[char]
		goto out
	}
	if char >= mtailPrivate {
		if char < mtailPrivate+len(mtailTok2) {
			token = mtailTok2[char-mtailPrivate]
			goto out
		}
	}
	for i := 0; i < len(mtailTok3); i += 2 {
		token = mtailTok3[i+0]
		if token == char {
			token = mtailTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = mtailTok2[1] /* unknown char */
	}
	if mtailDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", mtailTokname(token), uint(char))
	}
	return char, token
}

func mtailParse(mtaillex mtailLexer) int {
	return mtailNewParser().Parse(mtaillex)
}

func (mtailrcvr *mtailParserImpl) Parse(mtaillex mtailLexer) int {
	var mtailn int
	var mtailVAL mtailSymType
	var mtailDollar []mtailSymType
	_ = mtailDollar // silence set and not used
	mtailS := mtailrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	mtailstate := 0
	mtailrcvr.char = -1
	mtailtoken := -1 // mtailrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		mtailstate = -1
		mtailrcvr.char = -1
		mtailtoken = -1
	}()
	mtailp := -1
	goto mtailstack

ret0:
	return 0

ret1:
	return 1

mtailstack:
	/* put a state and value onto the stack */
	if mtailDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", mtailTokname(mtailtoken), mtailStatname(mtailstate))
	}

	mtailp++
	if mtailp >= len(mtailS) {
		nyys := make([]mtailSymType, len(mtailS)*2)
		copy(nyys, mtailS)
		mtailS = nyys
	}
	mtailS[mtailp] = mtailVAL
	mtailS[mtailp].yys = mtailstate

mtailnewstate:
	mtailn = mtailPact[mtailstate]
	if mtailn <= mtailFlag {
		goto mtaildefault /* simple state */
	}
	if mtailrcvr.char < 0 {
		mtailrcvr.char, mtailtoken = mtaillex1(mtaillex, &mtailrcvr.lval)
	}
	mtailn += mtailtoken
	if mtailn < 0 || mtailn >= mtailLast {
		goto mtaildefault
	}
	mtailn = mtailAct[mtailn]
	if mtailChk[mtailn] == mtailtoken { /* valid shift */
		mtailrcvr.char = -1
		mtailtoken = -1
		mtailVAL = mtailrcvr.lval
		mtailstate = mtailn
		if Errflag > 0 {
			Errflag--
		}
		goto mtailstack
	}

mtaildefault:
	/* default state action */
	mtailn = mtailDef[mtailstate]
	if mtailn == -2 {
		if mtailrcvr.char < 0 {
			mtailrcvr.char, mtailtoken = mtaillex1(mtaillex, &mtailrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if mtailExca[xi+0] == -1 && mtailExca[xi+1] == mtailstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			mtailn = mtailExca[xi+0]
			if mtailn < 0 || mtailn == mtailtoken {
				break
			}
		}
		mtailn = mtailExca[xi+1]
		if mtailn < 0 {
			goto ret0
		}
	}
	if mtailn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			mtaillex.Error(mtailErrorMessage(mtailstate, mtailtoken))
			Nerrs++
			if mtailDebug >= 1 {
				__yyfmt__.Printf("%s", mtailStatname(mtailstate))
				__yyfmt__.Printf(" saw %s\n", mtailTokname(mtailtoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for mtailp >= 0 {
				mtailn = mtailPact[mtailS[mtailp].yys] + mtailErrCode
				if mtailn >= 0 && mtailn < mtailLast {
					mtailstate = mtailAct[mtailn] /* simulate a shift of "error" */
					if mtailChk[mtailstate] == mtailErrCode {
						goto mtailstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if mtailDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", mtailS[mtailp].yys)
				}
				mtailp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if mtailDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", mtailTokname(mtailtoken))
			}
			if mtailtoken == mtailEofCode {
				goto ret1
			}
			mtailrcvr.char = -1
			mtailtoken = -1
			goto mtailnewstate /* try again in the same state */
		}
	}

	/* reduction by production mtailn */
	if mtailDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", mtailn, mtailStatname(mtailstate))
	}

	mtailnt := mtailn
	mtailpt := mtailp
	_ = mtailpt // guard against "declared and not used"

	mtailp -= mtailR2[mtailn]
	// mtailp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if mtailp+1 >= len(mtailS) {
		nyys := make([]mtailSymType, len(mtailS)*2)
		copy(nyys, mtailS)
		mtailS = nyys
	}
	mtailVAL = mtailS[mtailp+1]

	/* consult goto table to find next state */
	mtailn = mtailR1[mtailn]
	mtailg := mtailPgo[mtailn]
	mtailj := mtailg + mtailS[mtailp].yys + 1

	if mtailj >= mtailLast {
		mtailstate = mtailAct[mtailg]
	} else {
		mtailstate = mtailAct[mtailj]
		if mtailChk[mtailstate] != -mtailn {
			mtailstate = mtailAct[mtailg]
		}
	}
	// dummy call; replaced with literal code
	switch mtailnt {

	case 1:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:83
		{
			mtaillex.(*parser).root = mtailDollar[1].n
		}
	case 2:
		mtailDollar = mtailS[mtailpt-0 : mtailpt+1]
//line parser.y:90
		{
			mtailVAL.n = &ast.StmtList{}
		}
	case 3:
		mtailDollar = mtailS[mtailpt-2 : mtailpt+1]
//line parser.y:94
		{
			mtailVAL.n = mtailDollar[1].n
			if mtailDollar[2].n != nil {
				mtailVAL.n.(*ast.StmtList).Children = append(mtailVAL.n.(*ast.StmtList).Children, mtailDollar[2].n)
			}
		}
	case 4:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:104
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 5:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:106
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 6:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:108
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 7:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:110
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 8:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:112
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 9:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:114
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 10:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:116
		{
			mtailVAL.n = &ast.NextNode{tokenpos(mtaillex)}
		}
	case 11:
		mtailDollar = mtailS[mtailpt-3 : mtailpt+1]
//line parser.y:120
		{
			mtailVAL.n = &ast.PatternFragmentDefNode{Id: mtailDollar[2].n, Expr: mtailDollar[3].n}
		}
	case 12:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:124
		{
			mtailVAL.n = &ast.StopNode{tokenpos(mtaillex)}
		}
	case 13:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:128
		{
			mtailVAL.n = &ast.ErrorNode{tokenpos(mtaillex), mtailDollar[1].text}
		}
	case 14:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:135
		{
			mtailVAL.n = &ast.Cond{mtailDollar[1].n, mtailDollar[2].n, mtailDollar[4].n, nil}
		}
	case 15:
		mtailDollar = mtailS[mtailpt-2 : mtailpt+1]
//line parser.y:139
		{
			if mtailDollar[1].n != nil {
				mtailVAL.n = &ast.Cond{mtailDollar[1].n, mtailDollar[2].n, nil, nil}
			} else {
				mtailVAL.n = mtailDollar[2].n
			}
		}
	case 16:
		mtailDollar = mtailS[mtailpt-2 : mtailpt+1]
//line parser.y:147
		{
			o := &ast.OtherwiseNode{tokenpos(mtaillex)}
			mtailVAL.n = &ast.Cond{o, mtailDollar[2].n, nil, nil}
		}
	case 17:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:155
		{
			mtailVAL.n = nil
		}
	case 18:
		mtailDollar = mtailS[mtailpt-2 : mtailpt+1]
//line parser.y:157
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 19:
		mtailDollar = mtailS[mtailpt-3 : mtailpt+1]
//line parser.y:162
		{
			mtailVAL.n = mtailDollar[2].n
		}
	case 20:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:169
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 21:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:174
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 22:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:178
		{
			mtailVAL.n = &ast.BinaryExpr{Lhs: mtailDollar[1].n, Rhs: mtailDollar[4].n, Op: mtailDollar[2].op}
		}
	case 23:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:182
		{
			mtailVAL.n = &ast.BinaryExpr{Lhs: mtailDollar[1].n, Rhs: mtailDollar[4].n, Op: mtailDollar[2].op}
		}
	case 24:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:189
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 25:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:191
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 26:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:193
		{
			mtailVAL.n = &ast.BinaryExpr{Lhs: mtailDollar[1].n, Rhs: mtailDollar[4].n, Op: mtailDollar[2].op}
		}
	case 27:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:197
		{
			mtailVAL.n = &ast.BinaryExpr{Lhs: mtailDollar[1].n, Rhs: mtailDollar[4].n, Op: mtailDollar[2].op}
		}
	case 28:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:204
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 29:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:206
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 30:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:211
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 31:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:213
		{
			mtailVAL.n = &ast.BinaryExpr{Lhs: mtailDollar[1].n, Rhs: mtailDollar[4].n, Op: mtailDollar[2].op}
		}
	case 32:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:220
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 33:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:222
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 34:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:224
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 35:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:229
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 36:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:231
		{
			mtailVAL.n = &ast.BinaryExpr{Lhs: mtailDollar[1].n, Rhs: mtailDollar[4].n, Op: mtailDollar[2].op}
		}
	case 37:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:238
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 38:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:240
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 39:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:242
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 40:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:244
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 41:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:246
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 42:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:248
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 43:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:253
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 44:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:255
		{
			mtailVAL.n = &ast.BinaryExpr{Lhs: mtailDollar[1].n, Rhs: mtailDollar[4].n, Op: mtailDollar[2].op}
		}
	case 45:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:262
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 46:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:264
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 47:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:269
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 48:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:271
		{
			mtailVAL.n = &ast.BinaryExpr{Lhs: mtailDollar[1].n, Rhs: mtailDollar[4].n, Op: mtailDollar[2].op}
		}
	case 49:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:278
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 50:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:280
		{
			mtailVAL.n = &ast.BinaryExpr{Lhs: mtailDollar[1].n, Rhs: mtailDollar[4].n, Op: mtailDollar[2].op}
		}
	case 51:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:284
		{
			mtailVAL.n = &ast.BinaryExpr{Lhs: mtailDollar[1].n, Rhs: mtailDollar[4].n, Op: mtailDollar[2].op}
		}
	case 52:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:291
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 53:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:293
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 54:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:298
		{
			mtailVAL.n = &ast.PatternExpr{Expr: mtailDollar[1].n}
		}
	case 55:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:305
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 56:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:307
		{
			mtailVAL.n = &ast.BinaryExpr{Lhs: mtailDollar[1].n, Rhs: mtailDollar[4].n, Op: CONCAT}
		}
	case 57:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:311
		{
			mtailVAL.n = &ast.BinaryExpr{Lhs: mtailDollar[1].n, Rhs: mtailDollar[4].n, Op: CONCAT}
		}
	case 58:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:318
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 59:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:320
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 60:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:325
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 61:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:327
		{
			mtailVAL.n = &ast.BinaryExpr{Lhs: mtailDollar[1].n, Rhs: mtailDollar[4].n, Op: mtailDollar[2].op}
		}
	case 62:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:334
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 63:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:336
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 64:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:338
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 65:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:340
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 66:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:345
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 67:
		mtailDollar = mtailS[mtailpt-2 : mtailpt+1]
//line parser.y:347
		{
			mtailVAL.n = &ast.UnaryExpr{P: tokenpos(mtaillex), Expr: mtailDollar[2].n, Op: mtailDollar[1].op}
		}
	case 68:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:354
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 69:
		mtailDollar = mtailS[mtailpt-2 : mtailpt+1]
//line parser.y:356
		{
			mtailVAL.n = &ast.UnaryExpr{P: tokenpos(mtaillex), Expr: mtailDollar[1].n, Op: mtailDollar[2].op}
		}
	case 70:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:363
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 71:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:365
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 72:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:370
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 73:
		mtailDollar = mtailS[mtailpt-3 : mtailpt+1]
//line parser.y:372
		{
			mtailVAL.n = &ast.BuiltinNode{P: tokenpos(mtaillex), Name: mtailDollar[1].text, Args: nil}
		}
	case 74:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:376
		{
			mtailVAL.n = &ast.BuiltinNode{P: tokenpos(mtaillex), Name: mtailDollar[1].text, Args: mtailDollar[3].n}
		}
	case 75:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:380
		{
			mtailVAL.n = &ast.CaprefNode{tokenpos(mtaillex), mtailDollar[1].text, false, nil}
		}
	case 76:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:384
		{
			mtailVAL.n = &ast.CaprefNode{tokenpos(mtaillex), mtailDollar[1].text, true, nil}
		}
	case 77:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:388
		{
			mtailVAL.n = &ast.StringConst{tokenpos(mtaillex), mtailDollar[1].text}
		}
	case 78:
		mtailDollar = mtailS[mtailpt-3 : mtailpt+1]
//line parser.y:392
		{
			mtailVAL.n = mtailDollar[2].n
		}
	case 79:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:396
		{
			mtailVAL.n = &ast.IntConst{tokenpos(mtaillex), mtailDollar[1].intVal}
		}
	case 80:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:400
		{
			mtailVAL.n = &ast.FloatConst{tokenpos(mtaillex), mtailDollar[1].floatVal}
		}
	case 81:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:407
		{
			mtailVAL.n = &ast.IndexedExpr{Lhs: mtailDollar[1].n, Index: &ast.ExprList{}}
		}
	case 82:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:411
		{
			mtailVAL.n = mtailDollar[1].n
			mtailVAL.n.(*ast.IndexedExpr).Index.(*ast.ExprList).Children = append(
				mtailVAL.n.(*ast.IndexedExpr).Index.(*ast.ExprList).Children,
				mtailDollar[3].n.(*ast.ExprList).Children...)
		}
	case 83:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:421
		{
			mtailVAL.n = &ast.Id{tokenpos(mtaillex), mtailDollar[1].text, nil, false}
		}
	case 84:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:428
		{
			mtailVAL.n = &ast.ExprList{}
			mtailVAL.n.(*ast.ExprList).Children = append(mtailVAL.n.(*ast.ExprList).Children, mtailDollar[1].n)
		}
	case 85:
		mtailDollar = mtailS[mtailpt-3 : mtailpt+1]
//line parser.y:433
		{
			mtailVAL.n = mtailDollar[1].n
			mtailVAL.n.(*ast.ExprList).Children = append(mtailVAL.n.(*ast.ExprList).Children, mtailDollar[3].n)
		}
	case 86:
		mtailDollar = mtailS[mtailpt-5 : mtailpt+1]
//line parser.y:441
		{
			mp := markedpos(mtaillex)
			tp := tokenpos(mtaillex)
			pos := ast.MergePosition(&mp, &tp)
			mtailVAL.n = &ast.PatternConst{P: *pos, Pattern: mtailDollar[4].text}
		}
	case 87:
		mtailDollar = mtailS[mtailpt-3 : mtailpt+1]
//line parser.y:451
		{
			mtailVAL.n = mtailDollar[3].n
			d := mtailVAL.n.(*ast.DeclNode)
			d.Kind = mtailDollar[2].kind
			d.Hidden = mtailDollar[1].flag
		}
	case 88:
		mtailDollar = mtailS[mtailpt-0 : mtailpt+1]
//line parser.y:461
		{
			mtailVAL.flag = false
		}
	case 89:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:465
		{
			mtailVAL.flag = true
		}
	case 90:
		mtailDollar = mtailS[mtailpt-2 : mtailpt+1]
//line parser.y:472
		{
			mtailVAL.n = mtailDollar[1].n
			mtailVAL.n.(*ast.DeclNode).Keys = mtailDollar[2].texts
		}
	case 91:
		mtailDollar = mtailS[mtailpt-2 : mtailpt+1]
//line parser.y:477
		{
			mtailVAL.n = mtailDollar[1].n
			mtailVAL.n.(*ast.DeclNode).ExportedName = mtailDollar[2].text
		}
	case 92:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:482
		{
			mtailVAL.n = &ast.DeclNode{P: tokenpos(mtaillex), Name: mtailDollar[1].text}
		}
	case 93:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:486
		{
			mtailVAL.n = &ast.DeclNode{P: tokenpos(mtaillex), Name: mtailDollar[1].text}
		}
	case 94:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:493
		{
			mtailVAL.kind = metrics.Counter
		}
	case 95:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:497
		{
			mtailVAL.kind = metrics.Gauge
		}
	case 96:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:501
		{
			mtailVAL.kind = metrics.Timer
		}
	case 97:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:505
		{
			mtailVAL.kind = metrics.Text
		}
	case 98:
		mtailDollar = mtailS[mtailpt-2 : mtailpt+1]
//line parser.y:512
		{
			mtailVAL.texts = mtailDollar[2].texts
		}
	case 99:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:519
		{
			mtailVAL.texts = make([]string, 0)
			mtailVAL.texts = append(mtailVAL.texts, mtailDollar[1].text)
		}
	case 100:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:524
		{
			mtailVAL.texts = make([]string, 0)
			mtailVAL.texts = append(mtailVAL.texts, mtailDollar[1].text)
		}
	case 101:
		mtailDollar = mtailS[mtailpt-3 : mtailpt+1]
//line parser.y:529
		{
			mtailVAL.texts = mtailDollar[1].texts
			mtailVAL.texts = append(mtailVAL.texts, mtailDollar[3].text)
		}
	case 102:
		mtailDollar = mtailS[mtailpt-3 : mtailpt+1]
//line parser.y:534
		{
			mtailVAL.texts = mtailDollar[1].texts
			mtailVAL.texts = append(mtailVAL.texts, mtailDollar[3].text)
		}
	case 103:
		mtailDollar = mtailS[mtailpt-2 : mtailpt+1]
//line parser.y:542
		{
			mtailVAL.text = mtailDollar[2].text
		}
	case 104:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:549
		{
			mtailVAL.n = &ast.DecoDefNode{P: markedpos(mtaillex), Name: mtailDollar[3].text, Block: mtailDollar[4].n}
		}
	case 105:
		mtailDollar = mtailS[mtailpt-3 : mtailpt+1]
//line parser.y:556
		{
			mtailVAL.n = &ast.DecoNode{markedpos(mtaillex), mtailDollar[2].text, mtailDollar[3].n, nil, nil}
		}
	case 106:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:563
		{
			mtailVAL.n = &ast.DelNode{P: tokenpos(mtaillex), N: mtailDollar[2].n, Expiry: mtailDollar[4].duration}
		}
	case 107:
		mtailDollar = mtailS[mtailpt-2 : mtailpt+1]
//line parser.y:567
		{
			mtailVAL.n = &ast.DelNode{P: tokenpos(mtaillex), N: mtailDollar[2].n}
		}
	case 108:
		mtailDollar = mtailS[mtailpt-0 : mtailpt+1]
//line parser.y:577
		{
			glog.V(2).Infof("position marked at %v", tokenpos(mtaillex))
			mtaillex.(*parser).pos = tokenpos(mtaillex)
		}
	case 109:
		mtailDollar = mtailS[mtailpt-0 : mtailpt+1]
//line parser.y:587
		{
			mtaillex.(*parser).inRegex()
		}
	}
	goto mtailstack /* stack new state and value */
}
