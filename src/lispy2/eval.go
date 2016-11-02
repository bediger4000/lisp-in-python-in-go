package lispy2
import (
	//"fmt"
)

func Eval(node Node, env Environment) (Node){
	switch node.(type) {
	case S:
		val, found := env[node]
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
				cond := Eval(l[1], GlobalEnv)
				var exp Node
				if cond.(B) {
					exp = l[2]
				} else {
					exp = l[3]
				}
				return Eval(exp, GlobalEnv)
			case "quote":
				return l[1]
			case "define":
				symbol := Eval(l[1], GlobalEnv)
				if symbol == nil {
					symbol = l[1]
				}
				value  := Eval(l[2], GlobalEnv)
				GlobalEnv[symbol] = value
				return nil
			case "begin":
				var answer Node
				for _, n := range l {
					answer = Eval(n, GlobalEnv)
				}
				return answer
			}
			proc := Eval(l[0], GlobalEnv)
			switch proc.(type) {
			case Fn:
				return proc.(Fn)(l)
			}
		default:
			return Node(nil)
		}
	default:
		return node
	}
	return node
}
