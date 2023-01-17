package main

import "fmt"

// Values

type Kind int

const (
	ValueInt  Kind = 0
	ValueBool Kind = 1
	Undefined Kind = 2
)

type Val struct {
	flag Kind
	valI int
	valB bool
}

func mkInt(x int) Val {
	return Val{flag: ValueInt, valI: x}
}
func mkBool(x bool) Val {
	return Val{flag: ValueBool, valB: x}
}
func mkUndefined() Val {
	return Val{flag: Undefined}
}

func showVal(v Val) string {
	var s string
	switch {
	case v.flag == ValueInt:
		s = Num(v.valI).pretty()
	case v.flag == ValueBool:
		s = Bool(v.valB).pretty()
	case v.flag == Undefined:
		s = "Undefined"
	}
	return s
}

// Value State is a mapping from variable names to values
type ValState map[string]*Val

/////////////////////////
// Stmt instances

func (stmt Seq) eval(s ValState) string {
	var result string
	result += stmt[0].eval(s)
	result += stmt[1].eval(s)
	return result
}

func (ite IfThenElse) eval(s ValState) string {
	var result string
	v := ite.cond.eval(s)
	if v.flag == ValueBool {
		switch {
		case v.valB:
			result += ite.thenStmt.eval(s)
		case !v.valB:
			result += ite.elseStmt.eval(s)
		}

	} else {
		fmt.Printf("if-then-else eval fail")
		result += "if-then-else eval fail"
	}
	return result
}

// Maps are represented via points.
// Hence, maps are passed by "reference" and the update is visible for the caller as well.
func (decl Decl) eval(s ValState) string {
	v := decl.rhs.eval(s)
	x := (string)(decl.lhs)
	s[x] = &v
	return ""
}

/////////////////////////
// Exp instances

func (x Bool) eval(s ValState) Val {
	return mkBool((bool)(x))
}

func (x Num) eval(s ValState) Val {
	return mkInt((int)(x))
}

func (e Mult) eval(s ValState) Val {
	n1 := e[0].eval(s)
	n2 := e[1].eval(s)
	if n1.flag == ValueInt && n2.flag == ValueInt {
		return mkInt(n1.valI * n2.valI)
	}
	return mkUndefined()
}

func (e Plus) eval(s ValState) Val {
	n1 := e[0].eval(s)
	n2 := e[1].eval(s)
	if n1.flag == ValueInt && n2.flag == ValueInt {
		return mkInt(n1.valI + n2.valI)
	}
	return mkUndefined()
}

func (e And) eval(s ValState) Val {
	b1 := e[0].eval(s)
	b2 := e[1].eval(s)
	switch {
	case b1.flag == ValueBool && !b1.valB:
		return mkBool(false)
	case b1.flag == ValueBool && b2.flag == ValueBool:
		return mkBool(b1.valB && b2.valB)
	}
	return mkUndefined()
}

func (e Or) eval(s ValState) Val {
	b1 := e[0].eval(s)
	b2 := e[1].eval(s)
	switch {
	case b1.flag == ValueBool && b1.valB:
		return mkBool(true)
	case b1.flag == ValueBool && b2.flag == ValueBool:
		return mkBool(b1.valB || b2.valB)
	}
	return mkUndefined()
}

/*----------------- Own Code -----------------------*/

func (a Assign) eval(s ValState) string {
	v := a.rhs.eval(s)
	x := (string)(a.lhs)
	oldValue, ok := s[x]
	if !ok {
		fmt.Printf("Variable unknown. Var: " + x)
		return fmt.Sprintf("Variable unknown. Var: " + x)
	}
	if v.flag != oldValue.flag {
		fmt.Printf("Variable assignement Failed. Var: " + x)
		return fmt.Sprintf("Variable assignement Failed. Var: " + x)
	}

	oldValue.valB = v.valB
	oldValue.valI = v.valI
	return ""
}

func (w WhileStmt) eval(s ValState) string {
	var result string
	for w.cond.eval(s).valB {
		result += w.block.eval(s)
	}
	return result
}

func (b Block) eval(s ValState) string {
	s_new := make(map[string]*Val)

	for key, element := range s {
		s_new[key] = element
	}

	return b.stmt.eval(s_new)
}

func (p PrintStmt) eval(s ValState) string {
	val := p.exp.eval(s)
	if val.flag == ValueBool {
		fmt.Printf("%t \n", val.valB)
		return fmt.Sprintf("%t", val.valB)
	} else if val.flag == ValueInt {
		fmt.Printf("%d \n", val.valI)
		return fmt.Sprintf("%d", val.valI)
	} else {
		fmt.Printf("Error Evaluating Print Statement Illtyped Value")
		return "Error Evaluating Print Statement Illtyped Value"
	}

}

func (e Not) eval(s ValState) Val {
	val := e.exp.eval(s)
	if val.flag == ValueBool {
		return mkBool(!val.valB)
	} else {
		return mkUndefined()
	}
}

func (e Equal) eval(s ValState) Val {
	b1 := e[0].eval(s)
	b2 := e[1].eval(s)
	switch {
	case b1.flag == ValueBool && b2.flag == ValueBool:
		return mkBool(b1.valB == b2.valB)
	case b1.flag == ValueInt && b2.flag == ValueInt:
		return mkBool(b1.valI == b2.valI)
	}
	return mkUndefined()
}

func (e Lesser) eval(s ValState) Val {
	b1 := e[0].eval(s)
	b2 := e[1].eval(s)
	if b1.flag == ValueInt && b2.flag == ValueInt {
		return mkBool(b1.valI < b2.valI)
	}
	return mkUndefined()
}

func (e Group) eval(s ValState) Val {
	return e.exp.eval(s)
}

func (v Var) eval(s ValState) Val {
	val, ok := s[string(v)]
	if ok {
		return *val
	} else {
		return mkUndefined()
	}

}
