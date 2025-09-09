package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter file path or none to exit: ")
	for scanner.Scan() {
		path := scanner.Text()
		if path == "" {
			break
		}
		path = filepath.Clean(path)
		traverse(path)
	}
	fmt.Println("Goodbye!")
}
