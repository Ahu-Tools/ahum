package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

func traverse(path string) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		fmt.Printf("Error occured during parsing file:\n%w\n", err)
		return
	}

	fmt.Println("Imports:")
	for _, i := range node.Imports {
		fmt.Println(i.Path.Value)
	}

	fmt.Println("Comments:")
	for _, i := range node.Comments {
		fmt.Println(i.Text())
	}

	fmt.Println("Functions:")
	for _, i := range node.Decls {
		fn, ok := i.(*ast.FuncDecl)
		if !ok {
			continue
		}
		fmt.Println(fn.Name.Name)
	}

	fmt.Println("Returns:")
	ast.Inspect(node, func(n ast.Node) bool {
		ret, ok := n.(*ast.ReturnStmt)
		if ok {
			fmt.Printf("Return statement found on line %d:\n\t", fset.Position(ret.Pos()).Line)
			printer.Fprint(os.Stdout, fset, ret)
		}
		return true
	})
}
