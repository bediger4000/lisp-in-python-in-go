package lispy2
import (
	"math"
	//"fmt"
)

type Environment struct {
	env map[Node]Node
	outer *Environment
	proc Node
}

func (e *Environment) find(sym Node) (Node, bool) {
	val, ok := e.env[sym]
	if ok {
		return val, ok
	}
	if e.outer != nil {
		val, ok = e.outer.find(sym)
	}
	return val, ok
}
func (e *Environment) def(sym Node, val Node) {
	e.env[sym] = val
}


func newEnv(params Node, args Node, outer *Environment) (*Environment) {
	var e Environment
	e.outer = outer
	e.env = make(map[Node]Node)

	p := params.(List)
	a := args.(List)

	for i, sp := range p {
		e.def(sp, a[i])
	}

	return &e
}

func StandardEnv() (*Environment) {
	var env Environment
	env.env = make(map[Node]Node)

	env.env[S("+")] = Fn(addNodes)
	env.env[S("-")] = Fn(subtractNodes)
	env.env[S("*")] = Fn(multiplyNodes)
	env.env[S("/")] = Fn(divideNodes)
	env.env[S("==")] = Fn(equalNodes)
	env.env[S("pi")] = F(math.Pi)
	env.env[S("cdr")] = Fn(cdr)
	env.env[S("car")] = Fn(car)
	env.env[S("cons")] = Fn(cons)

	return &env
}

type Pr struct {
	body Node
	params Node
	env *Environment
}
func (pr Pr) isNode() { }

func (pr Pr) run(n Node, e *Environment) (Node) {
	local := newEnv(pr.params, n.(List)[1:], pr.env)
	return Eval(pr.body, local)
}

func newProcedure(params Node, body Node, env *Environment) (Pr) {
	var r Pr
	r.body = body
	r.params = params
	r.env = env
	return r
}
