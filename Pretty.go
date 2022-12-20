package main

import "strconv"

/////////////////////////
// Stmt instances

// pretty print

func (stmt Seq) pretty() string {
	return stmt[0].pretty() + "; \n" + stmt[1].pretty()
}

func (decl Decl) pretty() string {
	return decl.lhs + " := " + decl.rhs.pretty()
}

/////////////////////////
// Exp instances

func (x Var) pretty() string {
	return (string)(x)
}

func (x Bool) pretty() string {
	if x {
		return "true"
	} else {
		return "false"
	}

}

func (x Num) pretty() string {
	return strconv.Itoa(int(x))
}

func (e Mult) pretty() string {

	var x string
	x = "("
	x += e[0].pretty()
	x += "*"
	x += e[1].pretty()
	x += ")"

	return x
}

func (e Plus) pretty() string {

	var x string
	x = "("
	x += e[0].pretty()
	x += "+"
	x += e[1].pretty()
	x += ")"

	return x
}

func (e And) pretty() string {

	var x string
	x = "("
	x += e[0].pretty()
	x += "&&"
	x += e[1].pretty()
	x += ")"

	return x
}

func (e Or) pretty() string {

	var x string
	x = "("
	x += e[0].pretty()
	x += "||"
	x += e[1].pretty()
	x += ")"

	return x
}

/*-------------- Own Code ------------------*/

func (while WhileStmt) pretty() string {
	return "while(" + while.cond.pretty() + ") \n" + while.block.pretty()
}
func (block Block) pretty() string {
	return "{ \n" + block.stmt.pretty() + " \n}"
}
func (printStmt PrintStmt) pretty() string {
	return "print(" + printStmt.exp.pretty() + ")"
}

func (ifTehenElse IfThenElse) pretty() string {
	return "if(" + ifTehenElse.cond.pretty() + ") \n" + ifTehenElse.thenStmt.pretty() + "\n else \n" + ifTehenElse.elseStmt.pretty()
}

func (assign Assign) pretty() string {
	return assign.lhs + " = " + assign.rhs.pretty()
}

func (e Not) pretty() string {
	return "!" + e.exp.pretty()
}

func (e Equal) pretty() string {
	return e[0].pretty() + "==" + e[1].pretty()
}
func (e Lesser) pretty() string {
	return e[0].pretty() + "<" + e[1].pretty()
}

func (e Group) pretty() string {
	return " (" + e.exp.pretty() + ") "
}
