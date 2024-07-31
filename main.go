
package main

import (
	"fmt"
	"github.com/karshPrime/trash/cmd"
)

func main() {
	// parse user arguments
	// comprehend action
	// comprehend file(s) path
	var lAllFiles []string // list of all files

	for _, filePath := range lAllFiles {
		// $ trash -ls
		cmd.RecordPrint()

		// $ trash -r
		cmd.ActionRestore(&filePath)

		// $ trash file.eg file2.eg .....
		cmd.ActionDelete(&filePath)
	}
}

func noArgsError() {
	fmt.Println("Usage:")
}

