package main

import (
	"flag"
	"fmt"
	"luago/binchunk"
	"luago/compiler"
	"os"
)

const usage = `
usage: %s [options] [filename]
Available options are:
  -l       list (use -ll for full listing)
  -o name  output to file 'name' (default is "luac.out")
  -p       parse only
  -s       strip debug information
`

func main() {
	_l := flag.Bool("l", false, "")
	_ll := flag.Bool("ll", false, "")
	_o := flag.String("o", "luac.out", "")
	_p := flag.Bool("p", false, "")
	_s := flag.Bool("s", false, "")
	flag.Usage = printUsage
	flag.Parse()

	if len(flag.Args()) != 1 {
		printUsage()
		return
	}

	filename := flag.Args()[0]
	proto := loadOrCompile(filename)

	if *_p {
		return
	}
	if *_s {
		binchunk.StripDebug(proto)
	}
	if *_l || *_ll {
		output := binchunk.List(proto, *_ll)
		fmt.Println(output)
	} else {
		// write to disk
		data := binchunk.Dump(proto)
		os.WriteFile(*_o, data, 0644)
	}
}

func printUsage() {
	fmt.Printf(usage, os.Args[0])
}

func loadOrCompile(filename string) *binchunk.Prototype {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	if binchunk.IsBinaryChunk(data) {
		return binchunk.Undump(data)
	} else {
		return compiler.Compile(string(data), "@compiled_code")
	}
}
func CompileLuaCode(luaCode string) ([]byte, error) {

	proto := compiler.Compile(luaCode, "@compiled_code")
	binchunk.StripDebug(proto) //减少大小

	compiledBytes := binchunk.Dump(proto)

	// 返回编译后的字节码，和可能发生的错误。
	return compiledBytes, nil
}
