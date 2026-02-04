package main

import (
	"flag"
	"fmt"
)

//os.args[0] = program name
//os.args[1] = user args

func main() {
	path := flag.String("path", ".", "Path to the file or directory to lint")
	lang := flag.String("lang", "", "Programming language of the CLI application (python or javascript)")

	flag.Parse()

	fmt.Printf("path: %s\n", *path)
	fmt.Printf("lang: %s\n", *lang)

}
