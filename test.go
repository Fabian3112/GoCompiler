package main

import "fmt"

// Examples

func test() {

	fmt.Printf("\n")

	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
	ex6()
	ex7()
	ex8()
	ex9()
	ex10()
	ex11()
	simpleMethod()
	localVariablesDoNotOverrideGlobalVariables()
}

func run(e Exp) {
	s := make(map[string]*Val)
	t := make(map[string]*Type)
	fmt.Printf("\n ******* ")
	fmt.Printf("\n %s", e.pretty())
	fmt.Printf("\n %s", showVal(e.eval(s)))
	fmt.Printf("\n %s", showType(e.infer(t)))
}

func runStmt(stmt Stmt) {
	s := make(map[string]*Val)
	t := make(map[string]*Type)
	fmt.Printf("\n ******* ")
	fmt.Printf("\n%s \n \n ==> \n", stmt.pretty())
	stmt.eval(s)
	fmt.Printf("\n %t", stmt.check(t))
}

func ex1() {
	ast := plus(mult(number(1), number(2)), number(0))

	run(ast)
}

func ex2() {
	ast := and(boolean(false), number(0))
	run(ast)
}

func ex3() {
	ast := or(boolean(false), number(0))
	run(ast)
}

func ex4() {

}

func ex5() {
	ast := lesser(number(4), number(7))
	run(ast)
}

func ex6() {
	ast := lesser(number(7), number(7))
	run(ast)
}

func ex7() {
	ast := equal(number(4), number(4))
	run(ast)
}

func ex8() {
	ast := equal(number(4), number(8))
	run(ast)
}

func ex9() {
	ast := seq(decl("x", equal(number(4), number(8))), printStmt(number(3)))
	runStmt(ast)
}

func ex10() {
	ast := seq(decl("x", equal(number(4), number(8))), printStmt(variable("x")))
	runStmt(ast)
}

func ex11() {
	ast := seq(decl("x", equal(number(4), boolean(true))), printStmt(variable("x")))
	runStmt(ast)
}

/*
i := 0
while i < 10

	i = i + 1
	print(i)
*/
func simpleMethod() {
	ast := seq(
		decl("i", number(0)), //i := 0
		whileStmt(lesser(variable("i"), number(10)), //while i < 10
			block(seq(
				assign("i", plus(variable("i"), number(1))), // i = i +1
				printStmt(variable("i"))))))                 // print(i)

	runStmt(ast)
}

/*
i := false
x := 3
y := 1
if i

	i = true

else

	i := 3
	x := 0
	y = 2
	print(i)
	print(x)

print(i)
print(x)
print(y)
------------------
exspected:
3
0
false
3
2

true
*/
func localVariablesDoNotOverrideGlobalVariables() {
	ast := seq(
		decl("i", boolean(false)), //i := false
		seq(decl("x", number(3)), //x := 3
			seq(decl("y", number(1)), //y := 1
				seq(ifThenElse(Var("i"), // if i
					block(assign("i", boolean(true))), // { i = true }
					block(seq(decl("i", number(3)), // else { i:= 3
						seq(decl("x", number(0)), // x := 0
							seq(assign("y", number(2)), // y = 2
								seq(printStmt(Var("i")), //print(i)
									printStmt(Var("x")))))))), //print(x)}
					seq(printStmt(Var("i")), //print(i)
						seq(printStmt(Var("x")), //print(x)
							printStmt(Var("y")))))))) // print(y)

	runStmt(ast)
}
