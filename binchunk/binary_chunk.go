package binchunk

const (
	LUA_SIGNATURE    = "\x1bLua"
	LUAC_VERSION     = 0x53
	LUAC_FORMAT      = 1
	LUAC_DATA        = "\x19\x93\r\n\x1a\n"
	CINT_SIZE        = 4
	INSTRUCTION_SIZE = 4
	LUA_INTEGER_SIZE = 8
	LUA_NUMBER_SIZE  = 8
	LUAC_INT         = 0x5678
	LUAC_NUM         = 370.5
)

// function prototype
type Prototype struct {
	Source          string // debug
	LineDefined     uint32
	LastLineDefined uint32
	NumParams       byte
	IsVararg        byte
	MaxStackSize    byte
	Code            []uint32
	Constants       []interface{}
	Upvalues        []Upvalue
	Protos          []*Prototype
	LineInfo        []uint32 // debug
	LocVars         []LocVar // debug
	UpvalueNames    []string // debug
}

type Upvalue struct {
	Instack byte
	Idx     byte
}

type LocVar struct {
	VarName string
	StartPC uint32
	EndPC   uint32
}

func IsBinaryChunk(data []byte) bool {
	return len(data) > 4 &&
		string(data[:4]) == LUA_SIGNATURE
}

func Undump(data []byte) *Prototype {
	reader := &reader{data}
	reader.checkHeader()
	reader.readByte() // size_upvalues
	return reader.readProto("")
}

func Dump(proto *Prototype) []byte {
	writer := &writer{}
	writer.writeHeader()
	writer.writeByte(byte(len(proto.Upvalues)))
	writer.writeProto(proto, "")
	return writer.data()
}

func List(proto *Prototype, full bool) string {
	printer := &printer{make([]string, 0, 64)}
	return printer.printFunc(proto, full)
}

func StripDebug(proto *Prototype) {
	proto.Source = ""
	proto.LineInfo = nil
	proto.LocVars = nil
	proto.UpvalueNames = nil
	for _, p := range proto.Protos {
		StripDebug(p)
	}
}
