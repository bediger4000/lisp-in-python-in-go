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
			return node
		}
	case List:
		l := node.(List)
		switch l[0].(S) {
		case "if":
			cond := Eval(l[1], GlobalEnv)
			if cond.(B) {
				return Eval(l[2], GlobalEnv)
			} else {
				return Eval(l[3], GlobalEnv)
			}
		case "quote":
			return l[1]
		case "define":
			symbol := Eval(l[1], GlobalEnv)
			value  := Eval(l[2], GlobalEnv)
			GlobalEnv[symbol] = value
			return value
		case "begin":
			var answer Node
			for _, n := range l {
				answer = Eval(n, GlobalEnv)
			}
			return answer
		}
		proc := Eval(l[0], GlobalEnv)
	default:
		return node
	}
	return node
}
