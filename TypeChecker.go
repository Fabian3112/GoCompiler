package main

// Types

type Type int

const (
	TyIllTyped Type = 0
	TyInt      Type = 1
	TyBool     Type = 2
)

func showType(t Type) string {
	var s string
	switch {
	case t == TyInt:
		s = "Int"
	case t == TyBool:
		s = "Bool"
	case t == TyIllTyped:
		s = "Illtyped"
	}
	return s
}

// Type State is a mapping from variable names to types
type TyState map[string]*Type

/////////////////////////
// Stmt instances

func (stmt Seq) check(t TyState) bool {
	if !stmt[0].check(t) {
		return false
	}
	return stmt[1].check(t)
}

func (decl Decl) check(t TyState) bool {
	ty := decl.rhs.infer(t)
	if ty == TyIllTyped {
		return false
	}

	x := (string)(decl.lhs)
	t[x] = &ty
	return true
}

func (a Assign) check(t TyState) bool {
	x := (string)(a.lhs)
	_, ok := t[x]
	if !ok {
		return false
	}
	return *t[x] == a.rhs.infer(t)
}

/////////////////////////
// Exp instances

func (x Var) infer(t TyState) Type {
	y := (string)(x)
	ty, ok := t[y]
	if ok {
		return *ty
	} else {
		return TyIllTyped // variable does not exist yields illtyped
	}

}

func (x Bool) infer(t TyState) Type {
	return TyBool
}

func (x Num) infer(t TyState) Type {
	return TyInt
}

func (e Mult) infer(t TyState) Type {
	t1 := e[0].infer(t)
	t2 := e[1].infer(t)
	if t1 == TyInt && t2 == TyInt {
		return TyInt
	}
	return TyIllTyped
}

func (e Plus) infer(t TyState) Type {
	t1 := e[0].infer(t)
	t2 := e[1].infer(t)
	if t1 == TyInt && t2 == TyInt {
		return TyInt
	}
	return TyIllTyped
}

func (e And) infer(t TyState) Type {
	t1 := e[0].infer(t)
	t2 := e[1].infer(t)
	if t1 == TyBool && t2 == TyBool {
		return TyBool
	}
	return TyIllTyped
}

func (e Or) infer(t TyState) Type {
	t1 := e[0].infer(t)
	t2 := e[1].infer(t)
	if t1 == TyBool && t2 == TyBool {
		return TyBool
	}
	return TyIllTyped
}

//------------------- Own Code -------------
func (i IfThenElse) check(t TyState) bool {
	if i.cond.infer(t) != TyBool {
		return false
	} else if !i.thenStmt.check(t) {
		return false
	} else {
		return i.elseStmt.check(t)
	}
}

func (w WhileStmt) check(t TyState) bool {
	if w.cond.infer(t) != TyBool {
		return false
	}
	return w.block.check(t)
}

func (b Block) check(t TyState) bool {
	t_new := make(map[string]*Type)

	for key, element := range t {
		t_new[key] = element
	}
	return b.stmt.check(t_new)
}

func (p PrintStmt) check(t TyState) bool {
	return p.exp.infer(t) != TyIllTyped
}

func (n Not) infer(t TyState) Type {
	if n.exp.infer(t) == TyBool {
		return TyBool
	} else {
		return TyIllTyped
	}
}

func (e Equal) infer(t TyState) Type {
	tya := e[0].infer(t)
	tyb := e[1].infer(t)
	if tya == tyb && tya != TyIllTyped {
		return TyBool
	} else {
		return TyIllTyped
	}
}

func (l Lesser) infer(t TyState) Type {
	if l[0].infer(t) != TyInt {
		return TyIllTyped
	} else if l[1].infer(t) != TyInt {
		return TyIllTyped
	} else {
		return TyBool
	}
}

func (g Group) infer(t TyState) Type {
	return g.exp.infer(t)
}
