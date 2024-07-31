
package main

import (
	"os"
	"fmt"
	"github.com/karshPrime/trash/cmd"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Missing Arguments")
		printUsage()
		return
	}

	switch os.Args[1] {
		case "-l", "--list":
			cmd.RecordPrint()

		case "-r", "--restore":
			if len(os.Args) > 2 {
				for _, filePath := range os.Args[2:] {
					cmd.ActionRestore(&filePath)
				}
			} else {
				fmt.Println("Missing FileID(s)")
				printUsage()
			}

		case "-h", "--help":
			fmt.Println("delete files by using system trash")
			printUsage()

		default:
			for _, filePath := range os.Args[1:] {
				cmd.ActionDelete(&filePath)
			}
	}
}

func printUsage() {
	fmt.Println("\nUSAGE:")
	fmt.Println("    $ trash [OPTIONS] {files}")
	fmt.Println("\nOPTIONS:")
	fmt.Println("    -l	--list\t\tprint record of files deleted")
	fmt.Println("    -r	--restore\tmove files back to original dir from trash")
	fmt.Println("    -h	--help\t\tshow this help menu")
}

