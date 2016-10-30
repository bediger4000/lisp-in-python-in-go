package lispy2
import (
	// "fmt"
)

func Eval(list Node, env Environment) (Node){
	switch list.(type) {
	case List:
		switch list.(List)[0].(type) {
		case S:
			fn, found := env[list.(List)[0]]
			if found {
				return fn.(Fn)(list)
			}
		case Fn:
			return list.(List)[0].(Fn)(list)
		default:
		}
	case I, F, B:
		return list
	case S:
		n, ok := env[list]
		if ok {
			return n
		} else {
			return list.(S)
		}
	case Fn:
	}
	return nil
}
