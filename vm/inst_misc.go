package vm

import . "luago/api"

func move(i Instruction, vm LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	b += 1

	vm.Copy(b, a)
}
func jmp(i Instruction, vm LuaVM) {
	a, sBx := i.AsBx()

	vm.AddPC(sBx)
	if a != 0 {
		vm.CloseUpvalues(a)
	}
}
func self(i Instruction, vm LuaVM) {
	a, b, c := i.ABC()
	a += 1
	b += 1

	vm.Copy(b, a+1)

	vm.GetRK(c)
	vm.GetTable(b)
	vm.Replace(a)
}
