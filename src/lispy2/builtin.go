package lispy2
import (
	//"fmt"
)

func addNodes(n Node) (Node) {
	switch n.(type) {
	case List:
		sum := 0.0
		l := n.(List)
		for _, node := range l[1:] {
			sum += float64(Eval(node, GlobalEnv).(F))
		}
		return F(sum)
	default:
	}
	return nil
}
func subtractNodes(n Node) (Node) {
	switch n.(type) {
	case List:
		l := n.(List)
		sum := float64(Eval(l[1], GlobalEnv).(F))
		for _, node := range l[2:] {
			sum -= float64(Eval(node, GlobalEnv).(F))
		}
		return F(sum)
	default:
	}
	return nil
}
func multiplyNodes(n Node) (Node) {
	switch n.(type) {
	case List:
		l := n.(List)
		answer := float64(Eval(l[1], GlobalEnv).(F))
		for _, node := range l[2:] {
			answer *= float64(Eval(node, GlobalEnv).(F))
		}
		return F(answer)
	default:
	}
	return nil
}
func divideNodes(n Node) (Node) {
	l := n.(List)
	answer := float64(Eval(l[1], GlobalEnv).(F))
	answer /= float64(Eval(l[2], GlobalEnv).(F))
	return F(answer)
}
func equalNodes(n Node) (Node) {
	l := n.(List)
	a := l[1]
	switch a.(type) {
	case F:
		b := l[1].(F)
		c := l[2].(F)
		return B(b == c)
	case S:
		b := l[1].(S)
		c := l[2].(S)
		return B(b == c)
	case B:
		b := l[1].(B)
		c := l[2].(B)
		return B(b == c)
	}
	return B(false)
}
