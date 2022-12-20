package main

// Simple imperative language

/*
vars       Variable names, start with lower-case letter

prog      ::= block
block     ::= "{" statement "}"
statement ::=  statement ";" statement           -- Command sequence
            |  vars ":=" exp                     -- Variable declaration
            |  vars "=" exp                      -- Variable assignment
            |  "while" exp block                 -- While
            |  "if" exp block "else" block       -- If-then-else
            |  "print" exp                       -- Print

exp ::= 0 | 1 | -1 | ...     -- Integers
     | "true" | "false"      -- Booleans
     | exp "+" exp           -- Addition
     | exp "*" exp           -- Multiplication
     | exp "||" exp          -- Disjunction
     | exp "&&" exp          -- Conjunction
     | "!" exp               -- Negation
     | exp "==" exp          -- Equality test
     | exp "<" exp           -- Lesser test
     | "(" exp ")"           -- Grouping of expressions
     | vars                  -- Variables
*/

// Interface

type Exp interface {
	pretty() string
	eval(s ValState) Val
	infer(t TyState) Type
}

type Stmt interface {
	pretty() string
	eval(s ValState)
	check(t TyState) bool
}

// Statement cases (incomplete)

type Seq [2]Stmt
type Decl struct {
	lhs string
	rhs Exp
}
type IfThenElse struct {
	cond     Exp
	thenStmt Block
	elseStmt Block
}

type Assign struct {
	lhs string
	rhs Exp
}

//-------------- Own Code --------//

type WhileStmt struct {
	cond  Exp
	block Block
}

type Block struct {
	stmt Stmt
}

type PrintStmt struct {
	exp Exp
}

//--------------------------------//

// Expression cases (incomplete)

type Bool bool
type Num int
type Mult [2]Exp
type Plus [2]Exp
type And [2]Exp
type Or [2]Exp
type Var string

// -------------- Own Code ----------//
type Not struct {
	exp Exp
}

type Equal [2]Exp
type Lesser [2]Exp
type Group struct {
	exp Exp
}

//----------------------------------//

// Helper functions to build ASTs by hand

func number(x int) Exp {
	return Num(x)
}

func boolean(x bool) Exp {
	return Bool(x)
}

func plus(x, y Exp) Exp {
	return (Plus)([2]Exp{x, y})

	// The type Plus is defined as the two element array consisting of Exp elements.
	// Plus and [2]Exp are isomorphic but different types.
	// We first build the AST value [2]Exp{x,y}.
	// Then cast this value (of type [2]Exp) into a value of type Plus.

}

func mult(x, y Exp) Exp {
	return (Mult)([2]Exp{x, y})
}

func and(x, y Exp) Exp {
	return (And)([2]Exp{x, y})
}

func or(x, y Exp) Exp {
	return (Or)([2]Exp{x, y})
}

/*--------------- Own Code --------------------*/
//helper functions
func seq(stmt_1, stmt_2 Stmt) Stmt {
	return (Seq)([2]Stmt{stmt_1, stmt_2})
}

func decl(lhs string, rhs Exp) Stmt {
	return Decl{lhs, rhs}
}

func ifThenElse(cond Exp, thenStmt, elseStmt Block) Stmt {
	return IfThenElse{cond, thenStmt, elseStmt}
}

func assign(lhs string, rhs Exp) Stmt {
	return Assign{lhs, rhs}
}

func whileStmt(cond Exp, block Block) Stmt {
	return WhileStmt{cond, block}
}

func block(stmt Stmt) Block {
	return Block{stmt}
}

func printStmt(exp Exp) Stmt {
	return PrintStmt{exp}
}

func not(exp Exp) Exp {
	return Not{exp}
}

func equal(x, y Exp) Exp {
	return Equal{x, y}
}

func lesser(x, y Exp) Exp {
	return Lesser{x, y}
}

func group(exp Exp) Exp {
	return Group{exp}
}

func variable(s string) Exp {
	return Var(s)
}

/*--------------------------------------------------*/

func main() {
	test()
}
