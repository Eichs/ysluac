package vm

import "luago/api"

/* OpMode */
/* basic instruction format */
const (
	IABC  = iota // [  B:9  ][  C:9  ][ A:8  ][OP:6]
	IABx         // [      Bx:18     ][ A:8  ][OP:6]
	IAsBx        // [     sBx:18     ][ A:8  ][OP:6]
	IAx          // [           Ax:26        ][OP:6]
)

/* OpArgMask */
const (
	OpArgN = iota /* argument is not used */
	OpArgU        /* argument is used */
	OpArgR        /* argument is a register or a jump offset */
	OpArgK        /* argument is a constant or register/constant */
)

/* OpCode mhy*/
const (
	OP_ADD = iota
	OP_SUB
	OP_MUL
	OP_MOD
	OP_POW
	OP_DIV
	OP_IDIV
	OP_BAND
	OP_BOR
	OP_BXOR
	OP_SHL
	OP_SHR
	OP_UNM
	OP_BNOT
	OP_NOT
	OP_LEN
	OP_CONCAT
	OP_JMP
	OP_EQ
	OP_LT
	OP_LE
	OP_MOVE
	OP_LOADK
	OP_LOADKX
	OP_LOADBOOL
	OP_RETURN
	OP_GETUPVAL
	OP_SETLIST
	OP_TESTSET
	OP_SELF
	OP_FORLOOP
	OP_SETTABLE
	OP_TAILCALL
	OP_TFORLOOP
	OP_SETTABUP
	OP_EXTRAARG
	OP_GETTABUP
	OP_VARARG
	OP_LOADNIL
	OP_TFORCALL
	OP_SETUPVAL
	OP_GETTABLE
	OP_FORPREP
	OP_NEWTABLE
	OP_CALL
	OP_CLOSURE
	OP_TEST
)

type opcode struct {
	testFlag byte // operator is a test (next instruction must be a jump)
	setAFlag byte // instruction set register A
	argBMode byte // B arg mode
	argCMode byte // C arg mode
	opMode   byte // op mode
	name     string
	action   func(i Instruction, vm api.LuaVM)
}

var opcodes = []opcode{

	{0, 1, OpArgK, OpArgK, IABC, "ADD", add},
	{0, 1, OpArgK, OpArgK, IABC, "SUB", sub},
	{0, 1, OpArgK, OpArgK, IABC, "MUL", mul},
	{0, 1, OpArgK, OpArgK, IABC, "MOD", mod},
	{0, 1, OpArgK, OpArgK, IABC, "POW", pow},
	{0, 1, OpArgK, OpArgK, IABC, "DIV", div},
	{0, 1, OpArgK, OpArgK, IABC, "IDIV", idiv},
	{0, 1, OpArgK, OpArgK, IABC, "BAND", band},
	{0, 1, OpArgK, OpArgK, IABC, "BOR", bor},
	{0, 1, OpArgK, OpArgK, IABC, "BXOR", bxor},
	{0, 1, OpArgK, OpArgK, IABC, "SHL", shl},
	{0, 1, OpArgK, OpArgK, IABC, "SHR", shr},
	{0, 1, OpArgR, OpArgN, IABC, "UNM", unm},
	{0, 1, OpArgR, OpArgN, IABC, "BNOT", bnot},
	{0, 1, OpArgR, OpArgN, IABC, "NOT", not},
	{0, 1, OpArgR, OpArgN, IABC, "LEN", length},
	{0, 1, OpArgR, OpArgR, IABC, "CONCAT", concat},
	{0, 0, OpArgR, OpArgN, IAsBx, "JMP", jmp},
	{1, 0, OpArgK, OpArgK, IABC, "EQ", eq},
	{1, 0, OpArgK, OpArgK, IABC, "LT", lt},
	{1, 0, OpArgK, OpArgK, IABC, "LE", le},
	{0, 1, OpArgR, OpArgN, IABC, "MOVE", move},
	{0, 1, OpArgK, OpArgN, IABx, "LOADK", loadK},
	{0, 1, OpArgN, OpArgN, IABx, "LOADKX", loadKx},
	{0, 1, OpArgU, OpArgU, IABC, "LOADBOOL", loadBool},
	{0, 0, OpArgU, OpArgN, IABC, "RETURN", _return},
	{0, 1, OpArgU, OpArgN, IABC, "GETUPVAL", getUpval},
	{0, 0, OpArgU, OpArgU, IABC, "SETLIST", setList},
	{1, 1, OpArgR, OpArgU, IABC, "TESTSET", testSet},
	{0, 1, OpArgR, OpArgK, IABC, "SELF", self},
	{0, 1, OpArgR, OpArgN, IAsBx, "FORLOOP", forLoop},
	{0, 0, OpArgK, OpArgK, IABC, "SETTABLE", setTable},
	{0, 1, OpArgU, OpArgU, IABC, "TAILCALL", tailCall},
	{0, 1, OpArgR, OpArgN, IAsBx, "TFORLOOP", tForLoop},
	{0, 0, OpArgK, OpArgK, IABC, "SETTABUP", setTabUp},
	{0, 0, OpArgU, OpArgU, IAx, "EXTRAARG", nil},
	{0, 1, OpArgU, OpArgK, IABC, "GETTABUP", getTabUp},
	{0, 1, OpArgU, OpArgN, IABC, "VARARG", vararg},
	{0, 1, OpArgU, OpArgN, IABC, "LOADNIL", loadNil},
	{0, 0, OpArgN, OpArgU, IABC, "TFORCALL", tForCall},
	{0, 0, OpArgU, OpArgN, IABC, "SETUPVAL", setUpval},
	{0, 1, OpArgR, OpArgK, IABC, "GETTABLE", getTable},
	{0, 1, OpArgR, OpArgN, IAsBx, "FORPREP", forPrep},
	{0, 1, OpArgU, OpArgU, IABC, "NEWTABLE", newTable},
	{0, 1, OpArgU, OpArgU, IABC, "CALL", call},
	{0, 1, OpArgU, OpArgN, IABx, "CLOSURE", closure},
	{1, 0, OpArgN, OpArgU, IABC, "TEST", test},
}
