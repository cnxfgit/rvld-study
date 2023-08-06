package main

import (
	"os"
	"rvld/pkg/linker"
	"rvld/pkg/utils"
)

func main() {
	if len(os.Args) < 2 {
		utils.Fatal("wrong args")
	}
	
	file := linker.MustNewFile(os.Args[1])

	if !linker.CheckMagic(file.Contents) {
		utils.Fatal("not an ELF file")
	}

	objFile := linker.NewObjectFile(file)
	utils.Assert(len(objFile.ElfSections) == 11)
	objFile.Parse()
	utils.Assert(objFile.FirstGlobal == 10)
}