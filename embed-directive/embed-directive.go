package main

/*
 * go:embed is a compiler directive that allows programs to include
 * arbitrary files and folders in the Go binary at build time.
 *
 * embed directives accept paths relative to the directory containing
 * the Go source file.
 * We can also embed multiple files or even folders with wildcards.
 * This uses a variable of the embed.FS type, which implements a
 * simple virtual file system.
 */
import (
	"embed"
)

//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {
	print(fileString)
	print(string(fileByte))

	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
