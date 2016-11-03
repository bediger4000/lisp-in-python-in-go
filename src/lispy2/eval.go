package lispy2
import (
	//"fmt"
)

func Eval(node Node, env *Environment) (Node){
	switch node.(type) {
	case S:
		val, found := env.find(node)
		if found {
			return val
		} else {
			return Node(nil)
		}
	case List:
		l := node.(List)
		switch l[0].(type) {
		case S:
			switch l[0].(S) {
			case "if":
				cond := Eval(l[1], env)
				var exp Node
				if cond.(B) {
					exp = l[2]
				} else {
					exp = l[3]
				}
				return Eval(exp, env)
			case "quote":
				return l[1]
			case "define":
				symbol := Eval(l[1], env)
				if symbol == nil {
					symbol = l[1]
				}
				value  := Eval(l[2], env)
				env.def(symbol, value)
				return Node(nil)
			case "begin":
				var answer Node
				for _, n := range l {
					answer = Eval(n, env)
				}
				return answer
			case "lambda":
				pr := newProcedure(l[1], l[2], env)
				return pr
			}
			proc := Eval(l[0], env)
			switch proc.(type) {
			case Fn:
				return proc.(Fn)(l, env)
			case Pr:
				pr := proc.(Pr)
				return pr.run(l, env)
			}
		default:
			return Node(nil)
		}
	default:
		return node
	}
	return node
}
