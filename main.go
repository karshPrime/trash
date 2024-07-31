
package main

import (
	"os"
	"fmt"
	"github.com/karshPrime/trash/cmd"
)

func main() {
	if len(os.Args) == 1 {
		noArgsError()
		return
	}

	switch os.Args[1] {
		case "-ls", "--list":
			cmd.RecordPrint()

		case "-r", "--restore":
			if len(os.Args) > 2 {
				for _, filePath := range os.Args[2:] {
					cmd.ActionRestore(&filePath)
				}
			} else {
				noArgsError()
			}

		case "-h", "--help":
			noArgsError()

		default:
			for _, filePath := range os.Args[1:] {
				cmd.ActionDelete(&filePath)
			}
	}
}
func noArgsError() {
	fmt.Println("Usage:")
}

