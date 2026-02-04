package main

import (
	"flag"
	"fmt"

	"github.com/LRz00/cli-lint/filesystem"
)

//os.args[0] = program name
//os.args[1] = user args

func main() {
	path := flag.String("path", ".", "Path to the file or directory to lint")
	lang := flag.String("lang", "", "Programming language of the CLI application (python or javascript)")

	flag.Parse()

	files, err := filesystem.CollectFiles(*path, *lang)
	if err != nil {
		fmt.Printf("Error collecting files: %v\n", err)
		return
	}

	fmt.Printf("%s files found: %d\n", *lang, len(files))

	fmt.Printf("path: %s\n", *path)
	fmt.Printf("lang: %s\n", *lang)

}
