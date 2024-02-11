package api

type LuaType = int
type ArithOp = int
type CompareOp = int
type ThreadStatus = int

type GoFunction func(LuaState) int
type UserData interface{}

type LuaState interface {
	BasicAPI
	String() string // debug
}
type BasicAPI interface {
	/* state manipulation */
	Close()                               //
	AtPanic(panicf GoFunction) GoFunction // set panic function
	Version() float64                     // get version number
	/* basic stack manipulation */
	GetTop() int              // stack.top
	AbsIndex(idx int) int     // abs(idx)
	UpvalueIndex(idx int) int //
	CheckStack(n int) bool    //
	Pop(n int)                // pop(n)
	Copy(fromIdx, toIdx int)  // r[toIdx] = r[fromidx]
	PushValue(idx int)        // push(r[idx])
	Replace(idx int)          // r[idx] = pop()
	Insert(idx int)           // r[idx, -1] >> 1
	Remove(idx int)           // remove(r[idx])
	Rotate(idx, n int)        // r[idx, -1] >> n
	SetTop(idx int)           // stack.top = idx
	XMove(to LuaState, n int) // to.push(pop(n))
	/* access functions (stack -> C) */
	TypeName(tp LuaType) string        // r[idx].type.name
	Type(idx int) LuaType              // r[idx].type
	IsNone(idx int) bool               // r[idx].type ==
	IsNil(idx int) bool                // r[idx].type ==
	IsNoneOrNil(idx int) bool          // r[idx].type ==
	IsBoolean(idx int) bool            // r[idx].type ==
	IsInteger(idx int) bool            // r[idx].type ==
	IsNumber(idx int) bool             // r[idx] ~= LuaNumber
	IsString(idx int) bool             // r[idx] ~= LuaString
	IsTable(idx int) bool              // r[idx].type ==
	IsThread(idx int) bool             // r[idx].type ==
	IsFunction(idx int) bool           // r[idx].type ==
	IsGoFunction(idx int) bool         // r[idx].type ==
	IsUserData(idx int) bool           // r[idx].type ==
	ToBoolean(idx int) bool            // r[idx] as bool
	ToInteger(idx int) int64           // r[idx] as LuaInteger
	ToIntegerX(idx int) (int64, bool)  // r[idx] as LuaInteger
	ToNumber(idx int) float64          // r[idx] as LuaNumber
	ToNumberX(idx int) (float64, bool) // r[idx] as LuaNumber
	ToString(idx int) string           // r[idx] as string
	ToStringX(idx int) (string, bool)  // r[idx] as string
	ToGoFunction(idx int) GoFunction   // r[idx] as
	ToThread(idx int) LuaState         // r[idx] as LuaThread
	ToUserData(idx int) UserData       // r[idx] as
	ToPointer(idx int) interface{}     // r[idx] as interface{}
	RawLen(idx int) uint               // len(r[idx])
	/* push functions (C -> stack) */
	PushNil()                                        // push(nil)
	PushBoolean(b bool)                              // push(b)
	PushInteger(n int64)                             // push(n)
	PushNumber(n float64)                            // push(n)
	PushString(s string)                             // push(s)
	PushFString(fmt string, a ...interface{}) string // push(fmt*a)
	PushGoFunction(f GoFunction)                     // push(f)
	PushGoClosure(f GoFunction, n int)               // push(f)
	PushLightUserData(d UserData)                    // push(d)
	PushGlobalTable()                                // push(global)
	PushThread() bool                                // push(thread)
	/* Comparison and arithmetic functions */
	Arith(op ArithOp)                          // b=pop(); a=pop(); push(a op b)
	Compare(idx1, idx2 int, op CompareOp) bool // r[idx1] op r[idx2]
	RawEqual(idx1, idx2 int) bool              // r[idx1] == r[idx2]
	/* get functions (Lua -> stack) */
	NewTable()                           // push({})
	CreateTable(nArr, nRec int)          // push({})
	GetTable(idx int) LuaType            // push(r[idx][pop()])
	GetField(idx int, k string) LuaType  // push(r[idx][k])
	GetI(idx int, i int64) LuaType       // push(r[idx][i])
	RawGet(idx int) LuaType              // push(r[idx][pop()])
	RawGetI(idx int, i int64) LuaType    // push(r[idx][i])
	RawGetP(idx int, p UserData) LuaType // push(r[idx][p])
	GetGlobal(name string) LuaType       // push(global[name])
	GetMetatable(idx int) bool           // push(r[idx].metatable)?
	GetUserValue(idx int) LuaType        // push(r[idx].userValue)
	/* set functions (stack -> Lua) */
	SetTable(idx int)                   // v=pop(); k=pop(); r[idx][k] = v
	SetField(idx int, k string)         // r[idx][k] = pop()
	SetI(idx int, i int64)              // r[idx][i] = pop()
	RawSet(idx int)                     // v=pop(); k=pop(); r[idx][k] = v
	RawSetI(idx int, i int64)           // r[idx][i] = pop()
	RawSetP(idx int, p UserData)        // r[idx][p] = pop()
	Register(name string, f GoFunction) // global[name] = f
	SetGlobal(name string)              // global[name] = pop()
	SetMetatable(idx int)               // r[idx].metatable = pop()
	SetUserValue(idx int)               // r[idx].userValue = pop()
	/* 'load' and 'call' functions (load and run Lua code) */
	Dump(strip bool) []byte                                 // todo
	Load(chunk []byte, chunkName, mode string) ThreadStatus // push(compile(chunk))
	Call(nArgs, nResults int)                               // args=pop(nArgs); f=pop(); f(args)
	CallK()                                                 //
	PCall(nArgs, nResults, msgh int) ThreadStatus           // call(nArgs, nResults) || push(err)
	PCallK()                                                //
	/* miscellaneous functions */
	Concat(n int)                 // push(concat(pop(n)))
	Len(idx int)                  // push(len(r[idx]))
	Next(idx int) bool            // key=pop(); k,v=next(r[idx]); push(k,v);
	StringToNumber(s string) bool // push(number(s))
	Error() int                   // panic(r[-1])
	/* coroutine functions */
	NewThread() LuaState                          // todo
	Resume(from LuaState, nArgs int) ThreadStatus // todo
	Yield(nResults int) int                       // todo
	YieldK()                                      // todo
	Status() ThreadStatus                         // todo
	IsYieldable() bool                            // todo
	/* garbage-collection function and options */
	GC(what, data int) int //
}

// no pseudo-index
// lua_insert
// lua_remove
// lua_rotate

// type LuaBoolean bool
// type LuaInteger int64
// type LuaNumber float64

// type LuaLightUserData UserData
// type LuaKContext int
// type LuaKFunction int
// type LuaReader int
// type LuaWriter int

// GetAllocf()
// GetExtraSpace()
// IsLightUserData(idx int)
// NewUserData(size uint)
// PushLiteral
// PushLString
// PushVFString()
// SetAllocf()
// ToLString(idx int) (string, uint)
