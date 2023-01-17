package main

import "fmt"

func run(e Exp) {
	s := make(map[string]*Val)
	t := make(map[string]*Type)
	fmt.Printf("\n ******* ")
	fmt.Printf("\n %s", e.pretty())
	fmt.Printf("\n %s", showVal(e.eval(s)))
	fmt.Printf("\n %s", showType(e.infer(t)))
}

func runStmt(stmt Stmt) string {
	s := make(map[string]*Val)
	t := make(map[string]*Type)
	fmt.Printf("\n ******* ")
	fmt.Printf("\n%s \n \n ==> \n", stmt.pretty())
	result := stmt.eval(s)
	fmt.Printf("\n %t", stmt.check(t))
	return result
}

func parseAndRun(prog string) string {
	ast := parseProgramm(prog)
	return runStmt(ast)
}
